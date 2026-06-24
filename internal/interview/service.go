package interview

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"

	"interview-agent/internal/agent"
	"interview-agent/internal/domain"
	"interview-agent/internal/knowledge"
	"interview-agent/internal/parser"
	"interview-agent/internal/questions"
	"interview-agent/internal/skills"
)

var (
	ErrNotFound  = errors.New("面试不存在")
	ErrCompleted = errors.New("面试已经结束")
)

type StartInput struct {
	CandidateName      string `json:"candidateName"`
	ResumeID           string `json:"resumeId"`
	Language           string `json:"language,omitempty"`
	DomainSkillID      string `json:"domainSkillId"`
	InterviewerSkillID string `json:"interviewerSkillId"`
	IncludeFoundation  bool   `json:"includeFoundation"`
	Difficulty         string `json:"difficulty"`
	QuestionCount      int    `json:"questionCount"`
}

type AnswerInput struct {
	Answer         string `json:"answer"`
	ElapsedSeconds int    `json:"elapsedSeconds"`
}

type SessionView struct {
	ID                 string                 `json:"id"`
	CandidateName      string                 `json:"candidateName"`
	Language           string                 `json:"language"`
	Difficulty         string                 `json:"difficulty"`
	Industry           string                 `json:"industry"`
	DomainSkillID      string                 `json:"domainSkillId"`
	DomainSkillName    string                 `json:"domainSkillName"`
	InterviewerSkillID string                 `json:"interviewerSkillId"`
	InterviewerName    string                 `json:"interviewerName"`
	InterviewerOpening string                 `json:"interviewerOpening"`
	IncludeFoundation  bool                   `json:"includeFoundation"`
	Status             string                 `json:"status"`
	Phase              string                 `json:"phase"`
	Current            int                    `json:"current"`
	Total              int                    `json:"total"`
	FollowUpRound      int                    `json:"followUpRound"`
	FollowUpTotal      int                    `json:"followUpTotal"`
	CurrentQuestion    *domain.PublicQuestion `json:"currentQuestion,omitempty"`
	StartedAt          time.Time              `json:"startedAt"`
}

type AnswerResult struct {
	Evaluation        domain.Evaluation      `json:"evaluation"`
	QuestionCompleted bool                   `json:"questionCompleted"`
	Completed         bool                   `json:"completed"`
	NextQuestion      *domain.PublicQuestion `json:"nextQuestion,omitempty"`
	Current           int                    `json:"current"`
	Total             int                    `json:"total"`
	FollowUpRound     int                    `json:"followUpRound"`
	FollowUpTotal     int                    `json:"followUpTotal"`
	Phase             string                 `json:"phase"`
	Report            *domain.Report         `json:"report,omitempty"`
}

type Service struct {
	mu        sync.RWMutex
	sessions  map[string]*domain.Session
	resumes   map[string]domain.Resume
	evaluator agent.Evaluator
	catalog   *skills.Catalog
	knowledge *knowledge.Store
}

type Option func(*Service)

func WithCatalog(catalog *skills.Catalog) Option {
	return func(service *Service) { service.catalog = catalog }
}
func WithKnowledge(store *knowledge.Store) Option {
	return func(service *Service) { service.knowledge = store }
}

func NewService(evaluator agent.Evaluator, options ...Option) *Service {
	catalog, _ := skills.Load()
	service := &Service{sessions: make(map[string]*domain.Session), resumes: make(map[string]domain.Resume), evaluator: evaluator, catalog: catalog}
	for _, option := range options {
		option(service)
	}
	return service
}

func (s *Service) AddResume(resume domain.Resume) domain.Resume {
	s.mu.Lock()
	defer s.mu.Unlock()
	resume.ID = newID("resume")
	resume.CreatedAt = time.Now()
	s.resumes[resume.ID] = resume
	return resume
}

func (s *Service) Start(input StartInput) (SessionView, error) {
	input.CandidateName = strings.TrimSpace(input.CandidateName)
	input.Language = strings.ToLower(strings.TrimSpace(input.Language))
	input.DomainSkillID = strings.TrimSpace(input.DomainSkillID)
	input.InterviewerSkillID = strings.TrimSpace(input.InterviewerSkillID)
	input.Difficulty = strings.ToLower(strings.TrimSpace(input.Difficulty))
	if input.CandidateName == "" {
		input.CandidateName = "候选人"
	}
	if input.Difficulty == "" {
		input.Difficulty = "mixed"
	}
	if input.QuestionCount < 1 {
		return SessionView{}, fmt.Errorf("主问题数量至少为 1")
	}
	if input.QuestionCount > 20 {
		return SessionView{}, fmt.Errorf("主问题数量不能超过 20")
	}
	if s.catalog == nil {
		return SessionView{}, fmt.Errorf("Skill 目录不可用")
	}
	var domainSkill skills.Skill
	var ok bool
	if input.DomainSkillID != "" {
		domainSkill, ok = s.catalog.Domain(input.DomainSkillID)
	} else {
		domainSkill, ok = s.catalog.DomainByLanguage(input.Language)
	}
	if !ok || domainSkill.Language == "foundation" {
		return SessionView{}, fmt.Errorf("请选择有效的行业/语言 Skill")
	}
	interviewerSkillID := input.InterviewerSkillID
	if interviewerSkillID == "" {
		interviewerSkillID = "echo-coach"
	}
	interviewerSkill, ok := s.catalog.Interviewer(interviewerSkillID)
	if !ok {
		return SessionView{}, fmt.Errorf("请选择有效的面试官 Skill")
	}
	if interviewerSkill.FollowUpRounds < 2 {
		interviewerSkill.FollowUpRounds = 2
	}
	if interviewerSkill.FollowUpRounds > 3 {
		interviewerSkill.FollowUpRounds = 3
	}
	input.Language = domainSkill.Language
	var keywords []string
	s.mu.RLock()
	if input.ResumeID != "" {
		if resume, ok := s.resumes[input.ResumeID]; ok {
			keywords = resume.Keywords
		} else {
			s.mu.RUnlock()
			return SessionView{}, fmt.Errorf("简历不存在")
		}
	}
	s.mu.RUnlock()
	selected, err := questions.SelectWithFoundation(input.Language, input.Difficulty, input.QuestionCount, keywords, input.IncludeFoundation)
	if err != nil {
		return SessionView{}, err
	}
	if s.knowledge != nil {
		if err := s.knowledge.RecordIssued(selected); err != nil {
			return SessionView{}, fmt.Errorf("记录出题 QA 失败: %w", err)
		}
	}
	session := &domain.Session{ID: newID("session"), CandidateName: input.CandidateName, ResumeID: input.ResumeID, Language: input.Language, Difficulty: input.Difficulty, Industry: domainSkill.Industry, DomainSkillID: domainSkill.ID, DomainSkillName: domainSkill.Name, InterviewerSkillID: interviewerSkill.ID, InterviewerName: interviewerSkill.Name, InterviewerOpening: interviewerSkill.OpeningLine, InterviewerPrompt: interviewerSkill.Prompt, InterviewerIntro: interviewerSkill.IntroductionPrompt, EvaluationFocus: interviewerSkill.EvaluationFocus, FeedbackTone: interviewerSkill.FeedbackTone, IncludeFoundation: input.IncludeFoundation, FollowUpTotal: interviewerSkill.FollowUpRounds, Status: "active", Phase: "introduction", Questions: selected, ResumeKeywords: append([]string(nil), keywords...), StartedAt: time.Now()}
	s.mu.Lock()
	s.sessions[session.ID] = session
	s.mu.Unlock()
	return view(session), nil
}

func (s *Service) Get(id string) (SessionView, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	session, ok := s.sessions[id]
	if !ok {
		return SessionView{}, ErrNotFound
	}
	return view(session), nil
}

func (s *Service) Answer(ctx context.Context, id string, input AnswerInput) (AnswerResult, error) {
	input.Answer = strings.TrimSpace(input.Answer)
	if input.Answer == "" {
		return AnswerResult{}, fmt.Errorf("回答不能为空")
	}
	if len([]rune(input.Answer)) > 8000 {
		return AnswerResult{}, fmt.Errorf("回答不能超过 8000 个字符")
	}
	if input.ElapsedSeconds < 0 {
		input.ElapsedSeconds = 0
	}
	if input.ElapsedSeconds > 7200 {
		input.ElapsedSeconds = 7200
	}

	s.mu.RLock()
	session, ok := s.sessions[id]
	if !ok {
		s.mu.RUnlock()
		return AnswerResult{}, ErrNotFound
	}
	if session.Status == "completed" {
		s.mu.RUnlock()
		return AnswerResult{}, ErrCompleted
	}
	if session.Phase == "introduction" {
		s.mu.RUnlock()
		return s.answerIntroduction(id, input)
	}
	if session.Current >= len(session.Questions) {
		s.mu.RUnlock()
		return AnswerResult{}, ErrCompleted
	}
	question := session.Questions[session.Current]
	currentIndex := session.Current
	currentRound := session.FollowUpRound
	currentPrompt := question.Prompt
	if currentRound > 0 {
		currentPrompt = session.CurrentPrompt
	}
	interviewerName := session.InterviewerName
	interviewerPrompt := session.InterviewerPrompt
	evaluationFocus := append([]string(nil), session.EvaluationFocus...)
	feedbackTone := session.FeedbackTone
	interviewerSkillID := session.InterviewerSkillID
	followUpTotal := session.FollowUpTotal
	previousAnswers := currentAnswers(session)
	s.mu.RUnlock()

	evaluationQuestion := question
	evaluationQuestion.Prompt = currentPrompt
	evaluation, err := s.evaluator.Evaluate(ctx, agent.EvaluationInput{Question: evaluationQuestion, CandidateAnswer: input.Answer, InterviewerName: interviewerName, InterviewerPrompt: interviewerPrompt, EvaluationFocus: evaluationFocus, FeedbackTone: feedbackTone})
	if err != nil {
		if currentRound > 0 {
			evaluation = localFollowUpEvaluate(input.Answer, err, interviewerSkillID)
		} else {
			evaluation = localEvaluate(evaluationQuestion, input.Answer, err, interviewerSkillID)
		}
	}

	nextRound := currentRound + 1
	nextPrompt := ""
	if nextRound <= followUpTotal {
		answers := append(previousAnswers, input.Answer)
		if generator, ok := s.evaluator.(agent.FollowUpGenerator); ok {
			nextPrompt, err = generator.GenerateFollowUp(ctx, agent.FollowUpInput{Question: question, CandidateAnswers: answers, Round: nextRound, Total: followUpTotal, InterviewerName: interviewerName, InterviewerPrompt: interviewerPrompt, EvaluationFocus: evaluationFocus})
		}
		if nextPrompt == "" || err != nil {
			nextPrompt = fallbackFollowUp(question, answers, nextRound, followUpTotal, interviewerSkillID)
		}
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	session, ok = s.sessions[id]
	if !ok {
		return AnswerResult{}, ErrNotFound
	}
	if session.Status == "completed" || session.Current >= len(session.Questions) {
		return AnswerResult{}, ErrCompleted
	}
	// Prevent a slow duplicate request from recording the same interview round twice.
	if session.Current != currentIndex || session.FollowUpRound != currentRound || session.Questions[session.Current].ID != question.ID {
		return AnswerResult{}, fmt.Errorf("该轮回答已经提交，请继续当前面试")
	}
	if currentRound == 0 {
		session.Answers = append(session.Answers, domain.AnswerRecord{Question: question, Answer: input.Answer, ElapsedSeconds: input.ElapsedSeconds, Evaluation: evaluation, AverageScore: evaluation.Score})
	} else {
		if len(session.Answers) == 0 || session.Answers[len(session.Answers)-1].Question.ID != question.ID {
			return AnswerResult{}, fmt.Errorf("追问上下文不存在")
		}
		record := &session.Answers[len(session.Answers)-1]
		record.FollowUps = append(record.FollowUps, domain.FollowUpRecord{Round: currentRound, Prompt: currentPrompt, Answer: input.Answer, ElapsedSeconds: input.ElapsedSeconds, Evaluation: evaluation})
		record.AverageScore = answerAverage(*record)
	}
	result := AnswerResult{Evaluation: evaluation, Current: session.Current, Total: len(session.Questions), FollowUpRound: currentRound, FollowUpTotal: followUpTotal, Phase: "technical"}
	if nextRound <= followUpTotal {
		session.FollowUpRound = nextRound
		session.CurrentPrompt = nextPrompt
		next := followUpPublic(question, nextPrompt, nextRound, followUpTotal)
		result.NextQuestion = &next
		result.FollowUpRound = nextRound
		return result, nil
	}

	result.QuestionCompleted = true
	session.Current++
	session.FollowUpRound = 0
	session.CurrentPrompt = ""
	result.Current = session.Current
	result.FollowUpRound = 0
	if session.Current >= len(session.Questions) {
		now := time.Now()
		session.Status = "completed"
		session.Phase = "completed"
		result.Phase = "completed"
		session.CompletedAt = &now
		result.Completed = true
		report := buildReport(session)
		result.Report = &report
		return result, nil
	}
	next := technicalPublic(session, session.Questions[session.Current])
	result.NextQuestion = &next
	return result, nil
}

func (s *Service) Report(id string) (domain.Report, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	session, ok := s.sessions[id]
	if !ok {
		return domain.Report{}, ErrNotFound
	}
	if session.Status != "completed" {
		return domain.Report{}, fmt.Errorf("面试尚未结束")
	}
	return buildReport(session), nil
}

func (s *Service) answerIntroduction(id string, input AnswerInput) (AnswerResult, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	session, ok := s.sessions[id]
	if !ok {
		return AnswerResult{}, ErrNotFound
	}
	if session.Status == "completed" {
		return AnswerResult{}, ErrCompleted
	}
	if session.Phase != "introduction" {
		return AnswerResult{}, fmt.Errorf("自我介绍已经提交，请继续当前面试")
	}
	prompt := introductionPrompt(session)
	session.Introduction = &domain.IntroductionRecord{Prompt: prompt, Answer: input.Answer, ElapsedSeconds: input.ElapsedSeconds}
	session.IntroductionWords = parser.ExtractKeywords(input.Answer, 12)
	session.Questions = questions.Rank(session.Questions, mergeKeywords(session.ResumeKeywords, session.IntroductionWords))
	session.Phase = "technical"
	next := technicalPublic(session, session.Questions[0])
	evaluation := domain.Evaluation{Summary: "自我介绍已记录，接下来的问题会结合你的简历与刚才提到的经历。", Strengths: []string{"完成面试开场"}, Improvements: []string{}, Source: "local"}
	return AnswerResult{Evaluation: evaluation, QuestionCompleted: true, NextQuestion: &next, Current: 0, Total: len(session.Questions), FollowUpTotal: session.FollowUpTotal, Phase: "technical"}, nil
}

func introductionPublic(session *domain.Session) domain.PublicQuestion {
	return domain.PublicQuestion{ID: "introduction", PromptID: session.ID + "-introduction", Stage: "introduction", Language: session.Language, Difficulty: session.Difficulty, Prompt: introductionPrompt(session), Tags: []string{"自我介绍"}}
}

func introductionPrompt(session *domain.Session) string {
	if prompt := strings.TrimSpace(session.InterviewerIntro); prompt != "" {
		return strings.ReplaceAll(prompt, "{candidate}", session.CandidateName)
	}
	return fmt.Sprintf("%s，欢迎你。正式开始前，请先用 1–2 分钟介绍一下自己，可以说说你的技术方向、最近的项目经历，以及这次最希望展示的能力。", session.CandidateName)
}

func technicalPublic(session *domain.Session, question domain.Question) domain.PublicQuestion {
	result := question.Public()
	result.FollowUpTotal = session.FollowUpTotal
	result.LeadIn = questionLeadIn(session, question)
	return result
}

func questionLeadIn(session *domain.Session, question domain.Question) string {
	if keyword := matchingKeyword(question, session.ResumeKeywords); keyword != "" {
		if session.Current == 0 {
			return fmt.Sprintf("谢谢你的介绍。我看到你的简历里提到了「%s」，我们就从这段经历展开。", keyword)
		}
		return fmt.Sprintf("接下来结合你简历里的「%s」，我想再深入问一个问题。", keyword)
	}
	if keyword := matchingKeyword(question, session.IntroductionWords); keyword != "" {
		if session.Current == 0 {
			return fmt.Sprintf("你刚才提到了「%s」，我们顺着这个方向继续。", keyword)
		}
		return fmt.Sprintf("回到你自我介绍中提到的「%s」，再看一个相关问题。", keyword)
	}
	if session.Current == 0 {
		if session.ResumeID != "" {
			return "谢谢你的介绍。我会结合简历中的技术经历继续提问，下面进入第一个问题。"
		}
		return fmt.Sprintf("谢谢你的介绍。下面我们进入「%s」方向的技术问题。", session.DomainSkillName)
	}
	return ""
}

func matchingKeyword(question domain.Question, keywords []string) string {
	haystack := strings.ToLower(strings.Join(append(append([]string{question.Prompt}, question.Tags...), question.KeyPoints...), " "))
	for _, keyword := range keywords {
		if keyword = strings.TrimSpace(keyword); keyword != "" && strings.Contains(haystack, strings.ToLower(keyword)) {
			return keyword
		}
	}
	return ""
}

func mergeKeywords(groups ...[]string) []string {
	seen := make(map[string]bool)
	result := make([]string, 0)
	for _, group := range groups {
		for _, keyword := range group {
			key := strings.ToLower(strings.TrimSpace(keyword))
			if key != "" && !seen[key] {
				seen[key] = true
				result = append(result, keyword)
			}
		}
	}
	return result
}

func view(session *domain.Session) SessionView {
	result := SessionView{ID: session.ID, CandidateName: session.CandidateName, Language: session.Language, Difficulty: session.Difficulty, Industry: session.Industry, DomainSkillID: session.DomainSkillID, DomainSkillName: session.DomainSkillName, InterviewerSkillID: session.InterviewerSkillID, InterviewerName: session.InterviewerName, InterviewerOpening: session.InterviewerOpening, IncludeFoundation: session.IncludeFoundation, Status: session.Status, Phase: session.Phase, Current: session.Current, Total: len(session.Questions), FollowUpRound: session.FollowUpRound, FollowUpTotal: session.FollowUpTotal, StartedAt: session.StartedAt}
	if session.Status == "active" && session.Current < len(session.Questions) {
		if session.Phase == "introduction" {
			q := introductionPublic(session)
			result.CurrentQuestion = &q
			return result
		}
		question := session.Questions[session.Current]
		q := technicalPublic(session, question)
		if session.FollowUpRound > 0 {
			q = followUpPublic(question, session.CurrentPrompt, session.FollowUpRound, session.FollowUpTotal)
		}
		result.CurrentQuestion = &q
	}
	return result
}

func followUpPublic(question domain.Question, prompt string, round, total int) domain.PublicQuestion {
	return domain.PublicQuestion{ID: question.ID, PromptID: fmt.Sprintf("%s-followup-%d", question.ID, round), Stage: "follow_up", Language: question.Language, Difficulty: question.Difficulty, Prompt: prompt, Tags: question.Tags, FollowUp: true, Round: round, FollowUpTotal: total}
}

func currentAnswers(session *domain.Session) []string {
	if session.FollowUpRound == 0 || len(session.Answers) == 0 {
		return nil
	}
	record := session.Answers[len(session.Answers)-1]
	if session.Current >= len(session.Questions) || record.Question.ID != session.Questions[session.Current].ID {
		return nil
	}
	answers := []string{record.Answer}
	for _, followUp := range record.FollowUps {
		answers = append(answers, followUp.Answer)
	}
	return answers
}

func fallbackFollowUp(question domain.Question, answers []string, round, total int, interviewerSkillID string) string {
	combined := strings.ToLower(strings.Join(answers, " "))
	missingPoints := make([]string, 0)
	for _, group := range question.KeyPoints {
		found := false
		for _, term := range strings.Split(group, "/") {
			if term = strings.TrimSpace(strings.ToLower(term)); term != "" && strings.Contains(combined, term) {
				found = true
				break
			}
		}
		if !found {
			missingPoints = append(missingPoints, strings.Split(group, "/")[0])
		}
	}
	missing := ""
	if len(missingPoints) > 0 {
		missing = missingPoints[min(round-1, len(missingPoints)-1)]
	}
	if missing != "" {
		switch interviewerSkillID {
		case "atlas-architect":
			if round == 1 {
				return fmt.Sprintf("追问 %d/%d：如果把「%s」放进真实系统，你会如何定义约束和容量？", round, total, missing)
			}
			if round == 2 {
				return fmt.Sprintf("追问 %d/%d：围绕「%s」，最可能的失败模式是什么？你会怎样降级？", round, total, missing)
			}
			return fmt.Sprintf("追问 %d/%d：针对「%s」的方案，你会怎样验证效果，并说明最终取舍？", round, total, missing)
		case "vera-challenger":
			if round == 1 {
				return fmt.Sprintf("追问 %d/%d：你还没有说明「%s」的准确机制，请直接补充。", round, total, missing)
			}
			if round == 2 {
				return fmt.Sprintf("追问 %d/%d：给我一个「%s」容易被误用的反例，以及后果。", round, total, missing)
			}
			return fmt.Sprintf("追问 %d/%d：线上因「%s」出现故障时，你先看什么证据，如何止损？", round, total, missing)
		case "socrates-guide":
			if round == 1 {
				return fmt.Sprintf("追问 %d/%d：你认为「%s」与刚才的结论是什么关系？沿着假设推导一下。", round, total, missing)
			}
			return fmt.Sprintf("追问 %d/%d：如果「%s」的前提不成立，你会怎样修正刚才的答案？", round, total, missing)
		default:
			if round == 1 {
				return fmt.Sprintf("追问 %d/%d：能进一步解释「%s」的原理吗？", round, total, missing)
			}
			return fmt.Sprintf("追问 %d/%d：请结合一个实际场景说明「%s」如何落地，以及要注意什么。", round, total, missing)
		}
	}
	switch round {
	case 1:
		return fmt.Sprintf("追问 %d/%d：如果条件发生变化，你刚才的结论在哪些边界下不再成立？", round, total)
	case 2:
		return fmt.Sprintf("追问 %d/%d：请用一个真实项目或故障案例说明你会怎样应用这个判断。", round, total)
	default:
		return fmt.Sprintf("追问 %d/%d：如果让你重新设计一次，你会做出什么取舍，如何验证结果？", round, total)
	}
}

func answerAverage(record domain.AnswerRecord) int {
	total := record.Evaluation.Score
	for _, followUp := range record.FollowUps {
		total += followUp.Evaluation.Score
	}
	return total / (len(record.FollowUps) + 1)
}

func localEvaluate(question domain.Question, answer string, modelErr error, interviewerSkillID string) domain.Evaluation {
	lower := strings.ToLower(answer)
	hits := 0
	matched := make([]string, 0)
	missing := make([]string, 0)
	for _, group := range question.KeyPoints {
		found := false
		for _, term := range strings.Split(group, "/") {
			term = strings.TrimSpace(strings.ToLower(term))
			if term != "" && strings.Contains(lower, term) {
				found = true
				break
			}
		}
		label := strings.Split(group, "/")[0]
		if found {
			hits++
			matched = append(matched, label)
		} else {
			missing = append(missing, label)
		}
	}
	score := 20
	if len(question.KeyPoints) > 0 {
		score += hits * 80 / len(question.KeyPoints)
	}
	if len([]rune(answer)) < 20 {
		score = min(score, 25)
	}
	if score > 100 {
		score = 100
	}
	strengths := []string{"回答已覆盖部分核心概念"}
	if len(matched) > 0 {
		strengths = []string{"覆盖了关键点：" + strings.Join(matched, "、")}
	}
	improvements := []string{"建议结合标准答案补全技术细节，并给出工程场景或例子"}
	if len(missing) > 0 {
		prefix := "还可补充："
		switch interviewerSkillID {
		case "vera-challenger":
			prefix = "关键缺失，必须补充："
		case "atlas-architect":
			prefix = "请补齐约束与边界："
		case "socrates-guide":
			prefix = "继续追问自己这些概念如何关联："
		}
		improvements = []string{prefix + strings.Join(missing[:min(3, len(missing))], "、")}
	}
	summary := fmt.Sprintf("本地规则识别到 %d/%d 个关键点。", hits, len(question.KeyPoints))
	if modelErr != nil && !errors.Is(modelErr, agent.ErrNotConfigured) {
		summary += " 大模型暂不可用，本题已自动降级评分。"
	}
	return domain.Evaluation{Score: score, Summary: summary, Strengths: strengths, Improvements: improvements, Source: "local"}
}

func localFollowUpEvaluate(answer string, modelErr error, interviewerSkillID string) domain.Evaluation {
	lower := strings.ToLower(answer)
	signalGroups := []struct {
		name  string
		terms []string
	}{
		{"原理机制", []string{"因为", "原理", "机制", "runtime", "gmp", "内存", "调度", "线程", "进程"}},
		{"边界取舍", []string{"边界", "取舍", "限制", "容量", "失败", "风险", "反例", "前提"}},
		{"实践验证", []string{"项目", "线上", "压测", "监控", "指标", "验证", "故障", "p99", "pprof"}},
		{"解决方案", []string{"通过", "使用", "设计", "降级", "限流", "取消", "队列", "发布"}},
	}
	covered := make([]string, 0)
	missing := make([]string, 0)
	for _, group := range signalGroups {
		found := false
		for _, term := range group.terms {
			if strings.Contains(lower, term) {
				found = true
				break
			}
		}
		if found {
			covered = append(covered, group.name)
		} else {
			missing = append(missing, group.name)
		}
	}
	score := 35 + len(covered)*15
	if len([]rune(answer)) < 20 {
		score = min(score, 35)
	}
	if score > 95 {
		score = 95
	}
	strengths := []string{"正面回应了本轮追问"}
	if len(covered) > 0 {
		strengths = []string{"追问回答体现了：" + strings.Join(covered, "、")}
	}
	prefix := "还可补充："
	if interviewerSkillID == "vera-challenger" {
		prefix = "回答仍缺少："
	} else if interviewerSkillID == "atlas-architect" {
		prefix = "架构回答还应补齐："
	}
	improvements := []string{"补充一个具体边界或项目例子"}
	if len(missing) > 0 {
		improvements = []string{prefix + strings.Join(missing[:min(2, len(missing))], "、")}
	}
	summary := fmt.Sprintf("本地追问评分识别到 %d/%d 类深度信号。", len(covered), len(signalGroups))
	if modelErr != nil && !errors.Is(modelErr, agent.ErrNotConfigured) {
		summary += " 大模型暂不可用，已自动降级评分。"
	}
	return domain.Evaluation{Score: score, Summary: summary, Strengths: strengths, Improvements: improvements, Source: "local"}
}

func buildReport(session *domain.Session) domain.Report {
	completed := time.Now()
	if session.CompletedAt != nil {
		completed = *session.CompletedAt
	}
	total := 0
	var highlights, focus []string
	for _, answer := range session.Answers {
		score := answer.AverageScore
		if score == 0 {
			score = answerAverage(answer)
		}
		total += score
		highlights = append(highlights, answer.Evaluation.Strengths...)
		focus = append(focus, answer.Evaluation.Improvements...)
		for _, followUp := range answer.FollowUps {
			highlights = append(highlights, followUp.Evaluation.Strengths...)
			focus = append(focus, followUp.Evaluation.Improvements...)
		}
	}
	score := 0
	if len(session.Answers) > 0 {
		score = total / len(session.Answers)
	}
	var introduction *domain.IntroductionRecord
	if session.Introduction != nil {
		copy := *session.Introduction
		introduction = &copy
	}
	return domain.Report{SessionID: session.ID, CandidateName: session.CandidateName, Language: session.Language, Difficulty: session.Difficulty, DomainSkillName: session.DomainSkillName, InterviewerName: session.InterviewerName, Score: score, Answered: len(session.Answers), Duration: int(completed.Sub(session.StartedAt).Seconds()), Highlights: topUnique(highlights, 3), FocusAreas: topUnique(focus, 3), Introduction: introduction, Answers: append([]domain.AnswerRecord(nil), session.Answers...), StartedAt: session.StartedAt, CompletedAt: completed}
}

func topUnique(values []string, limit int) []string {
	counts := map[string]int{}
	for _, value := range values {
		if value = strings.TrimSpace(value); value != "" {
			counts[value]++
		}
	}
	type item struct {
		value string
		count int
	}
	items := make([]item, 0, len(counts))
	for value, count := range counts {
		items = append(items, item{value, count})
	}
	sort.Slice(items, func(i, j int) bool { return items[i].count > items[j].count })
	if len(items) > limit {
		items = items[:limit]
	}
	result := make([]string, len(items))
	for i := range items {
		result[i] = items[i].value
	}
	return result
}

func newID(prefix string) string {
	buf := make([]byte, 8)
	if _, err := rand.Read(buf); err != nil {
		return fmt.Sprintf("%s-%d", prefix, time.Now().UnixNano())
	}
	return prefix + "-" + hex.EncodeToString(buf)
}
