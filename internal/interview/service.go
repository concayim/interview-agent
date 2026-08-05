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
	VideoEnabled       bool   `json:"videoEnabled"`
	SpeechLanguage     string `json:"speechLanguage"`
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
	VideoEnabled       bool                   `json:"videoEnabled"`
	SpeechLanguage     string                 `json:"speechLanguage"`
	QuestionSource     string                 `json:"questionSource"`
	Status             string                 `json:"status"`
	Current            int                    `json:"current"`
	Total              int                    `json:"total"`
	CurrentQuestion    *domain.PublicQuestion `json:"currentQuestion,omitempty"`
	StartedAt          time.Time              `json:"startedAt"`
}

type AnswerResult struct {
	Evaluation     *domain.Evaluation     `json:"evaluation,omitempty"`
	Completed      bool                   `json:"completed"`
	NextQuestion   *domain.PublicQuestion `json:"nextQuestion,omitempty"`
	Current        int                    `json:"current"`
	Total          int                    `json:"total"`
	Report         *domain.Report         `json:"report,omitempty"`
	Accepted       bool                   `json:"accepted"`
	Intent         string                 `json:"intent"`
	AssistantReply string                 `json:"assistantReply,omitempty"`
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
	input.SpeechLanguage = normalizeSpeechLanguage(input.SpeechLanguage)
	if input.CandidateName == "" {
		input.CandidateName = "候选人"
	}
	if input.Difficulty == "" {
		input.Difficulty = "mixed"
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
	selected, questionSource, err := s.selectQuestions(input, domainSkill, keywords)
	if err != nil {
		return SessionView{}, err
	}
	if s.knowledge != nil {
		if err := s.knowledge.RecordIssued(selected); err != nil {
			return SessionView{}, fmt.Errorf("记录出题 QA 失败: %w", err)
		}
	}
	session := &domain.Session{ID: newID("session"), CandidateName: input.CandidateName, ResumeID: input.ResumeID, Language: input.Language, Difficulty: input.Difficulty, Industry: domainSkill.Industry, DomainSkillID: domainSkill.ID, DomainSkillName: domainSkill.Name, InterviewerSkillID: interviewerSkill.ID, InterviewerName: interviewerSkill.Name, InterviewerOpening: interviewerSkill.OpeningLine, InterviewerPrompt: interviewerSkill.Prompt, EvaluationFocus: interviewerSkill.EvaluationFocus, FeedbackTone: interviewerSkill.FeedbackTone, IncludeFoundation: input.IncludeFoundation, VideoEnabled: input.VideoEnabled, SpeechLanguage: input.SpeechLanguage, QuestionSource: questionSource, Status: "active", Questions: selected, StartedAt: time.Now()}
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
	question := session.Questions[session.Current]
	current := session.Current
	total := len(session.Questions)
	interviewerName := session.InterviewerName
	interviewerPrompt := session.InterviewerPrompt
	evaluationFocus := append([]string(nil), session.EvaluationFocus...)
	feedbackTone := session.FeedbackTone
	interviewerSkillID := session.InterviewerSkillID
	s.mu.RUnlock()

	intentResult := s.resolveCandidateIntent(ctx, agent.IntentInput{Question: question, CandidateMessage: input.Answer, InterviewerName: interviewerName, InterviewerPrompt: interviewerPrompt, FeedbackTone: feedbackTone})
	if !intentResult.Accepted {
		return AnswerResult{Accepted: false, Intent: intentResult.Intent, AssistantReply: intentResult.AssistantReply, Current: current, Total: total}, nil
	}
	var evaluation domain.Evaluation
	if intentResult.Intent == "skip" {
		evaluation = skippedEvaluation()
	} else {
		var err error
		evaluation, err = s.evaluator.Evaluate(ctx, agent.EvaluationInput{Question: question, CandidateAnswer: input.Answer, InterviewerName: interviewerName, InterviewerPrompt: interviewerPrompt, EvaluationFocus: evaluationFocus, FeedbackTone: feedbackTone})
		if err != nil {
			evaluation = localEvaluate(question, input.Answer, err, interviewerSkillID)
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
	// Prevent a slow duplicate request from recording the same question twice.
	if session.Questions[session.Current].ID != question.ID {
		return AnswerResult{}, fmt.Errorf("该题已经提交，请继续下一题")
	}
	session.Answers = append(session.Answers, domain.AnswerRecord{Question: question, Answer: input.Answer, ElapsedSeconds: input.ElapsedSeconds, Evaluation: evaluation})
	session.Current++
	result := AnswerResult{Evaluation: &evaluation, Accepted: true, Intent: intentResult.Intent, Current: session.Current, Total: len(session.Questions)}
	if session.Current >= len(session.Questions) {
		now := time.Now()
		session.Status = "completed"
		session.CompletedAt = &now
		result.Completed = true
		report := buildReport(session)
		result.Report = &report
		return result, nil
	}
	next := session.Questions[session.Current].Public()
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

func (s *Service) selectQuestions(input StartInput, domainSkill skills.Skill, keywords []string) ([]domain.Question, string, error) {
	if generator, ok := s.evaluator.(agent.QuestionGenerator); ok {
		ctx, cancel := context.WithTimeout(context.Background(), 16*time.Second)
		defer cancel()
		generated, err := generator.GenerateQuestions(ctx, agent.QuestionGenerationInput{
			Language:          input.Language,
			Difficulty:        input.Difficulty,
			QuestionCount:     input.QuestionCount,
			CandidateKeywords: keywords,
			DomainSkillName:   domainSkill.Name,
			DomainPrompt:      domainSkill.Prompt,
			Topics:            domainSkill.Topics,
			IncludeFoundation: input.IncludeFoundation,
		})
		if err == nil && len(generated) > 0 {
			return generated, "model", nil
		}
	}
	selected, err := questions.SelectWithFoundation(input.Language, input.Difficulty, input.QuestionCount, keywords, input.IncludeFoundation)
	if err != nil {
		selected = questions.Generic(input.Language, domainSkill.Name, domainSkill.Topics, input.Difficulty, input.QuestionCount)
		if len(selected) > 0 {
			return selected, "built-in", nil
		}
	}
	return selected, "built-in", err
}

func normalizeSpeechLanguage(value string) string {
	value = strings.ToLower(strings.TrimSpace(value))
	switch value {
	case "en", "en-us", "english":
		return "en-US"
	case "zh", "zh-cn", "chinese", "中文":
		return "zh-CN"
	default:
		return "zh-CN"
	}
}

func view(session *domain.Session) SessionView {
	result := SessionView{ID: session.ID, CandidateName: session.CandidateName, Language: session.Language, Difficulty: session.Difficulty, Industry: session.Industry, DomainSkillID: session.DomainSkillID, DomainSkillName: session.DomainSkillName, InterviewerSkillID: session.InterviewerSkillID, InterviewerName: session.InterviewerName, InterviewerOpening: session.InterviewerOpening, IncludeFoundation: session.IncludeFoundation, VideoEnabled: session.VideoEnabled, SpeechLanguage: session.SpeechLanguage, QuestionSource: session.QuestionSource, Status: session.Status, Current: session.Current, Total: len(session.Questions), StartedAt: session.StartedAt}
	if session.Status == "active" && session.Current < len(session.Questions) {
		q := session.Questions[session.Current].Public()
		result.CurrentQuestion = &q
	}
	return result
}

func (s *Service) resolveCandidateIntent(ctx context.Context, input agent.IntentInput) agent.IntentResult {
	if result, handled := detectLocalCandidateIntent(input.CandidateMessage, input.Question); handled {
		return result
	}
	if resolver, ok := s.evaluator.(agent.IntentResolver); ok {
		if result, err := resolver.ResolveIntent(ctx, input); err == nil {
			return normalizeIntentResult(result)
		}
	}
	return agent.IntentResult{Intent: "answer", Accepted: true}
}

func normalizeIntentResult(result agent.IntentResult) agent.IntentResult {
	result.Intent = strings.TrimSpace(strings.ToLower(result.Intent))
	switch result.Intent {
	case "answer", "skip":
		result.Accepted = true
		result.AssistantReply = ""
	case "hint", "clarify", "repeat", "off_topic", "smalltalk":
		result.Accepted = false
	default:
		result.Intent = "answer"
		result.Accepted = true
		result.AssistantReply = ""
	}
	if !result.Accepted && strings.TrimSpace(result.AssistantReply) == "" {
		result.AssistantReply = "我理解你的意思。我们先回到当前题，你可以从一句结论开始，后面再补充原因和边界。"
	}
	return result
}

func detectLocalCandidateIntent(answer string, question domain.Question) (agent.IntentResult, bool) {
	compact := strings.ToLower(strings.TrimSpace(answer))
	if compact == "" {
		return agent.IntentResult{}, false
	}
	if containsAny(compact, "跳过", "下一题", "不会答了", "放弃这题", "过吧", "pass", "skip") {
		return agent.IntentResult{Intent: "skip", Accepted: true}, true
	}
	if len([]rune(compact)) > 80 {
		return agent.IntentResult{}, false
	}
	if containsAny(compact, "重复", "再说一遍", "再发", "重新发", "上一题", "题目是什么") {
		return agent.IntentResult{Intent: "repeat", Accepted: false, AssistantReply: fmt.Sprintf("当然。当前题目是：%s", question.Prompt)}, true
	}
	if containsAny(compact, "什么意思", "没懂", "看不懂", "解释一下", "换个说法", "换种说法", "展开一下", "题目意思") {
		return agent.IntentResult{Intent: "clarify", Accepted: false, AssistantReply: clarifyQuestionReply(question)}, true
	}
	if containsAny(compact, "提示", "hint", "怎么答", "思路", "不会", "不知道", "没思路", "帮我") {
		return agent.IntentResult{Intent: "hint", Accepted: false, AssistantReply: hintQuestionReply(question)}, true
	}
	return agent.IntentResult{}, false
}

func containsAny(value string, needles ...string) bool {
	for _, needle := range needles {
		if strings.Contains(value, needle) {
			return true
		}
	}
	return false
}

func clarifyQuestionReply(question domain.Question) string {
	topic := strings.Join(question.Tags, "、")
	if topic == "" {
		topic = "这道题"
	}
	return fmt.Sprintf("我换个说法：这题想看你能不能围绕「%s」讲清概念、机制和边界。你可以先用一句话给结论，再解释为什么，最后补一个适用场景或容易踩的坑。", topic)
}

func hintQuestionReply(question domain.Question) string {
	topic := strings.Join(question.Tags, "、")
	if topic == "" {
		topic = "题目里的核心概念"
	}
	return fmt.Sprintf("可以按这个顺序组织：先定义「%s」，再说它解决了什么问题，然后补充限制、反例或项目里的使用经验。先答一个粗版本也可以，我会根据你的回答继续点评。", topic)
}

func skippedEvaluation() domain.Evaluation {
	return domain.Evaluation{
		Score:        0,
		Summary:      "本题已按你的要求跳过。",
		Strengths:    []string{"能主动识别当前卡点"},
		Improvements: []string{"复盘时先对照标准答案补齐核心概念，再用一句话重述"},
		Source:       "local",
	}
}

func localEvaluate(question domain.Question, answer string, modelErr error, interviewerSkillID string) domain.Evaluation {
	lower := strings.ToLower(answer)
	hits := 0
	matched := make([]string, 0)
	missing := make([]string, 0)
	for _, group := range question.KeyPoints {
		found := false
		separator := "/"
		if strings.Contains(group, "|") {
			separator = "|"
		}
		for _, term := range strings.Split(group, separator) {
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

func buildReport(session *domain.Session) domain.Report {
	completed := time.Now()
	if session.CompletedAt != nil {
		completed = *session.CompletedAt
	}
	total := 0
	var highlights, focus []string
	for _, answer := range session.Answers {
		total += answer.Evaluation.Score
		highlights = append(highlights, answer.Evaluation.Strengths...)
		focus = append(focus, answer.Evaluation.Improvements...)
	}
	score := 0
	if len(session.Answers) > 0 {
		score = total / len(session.Answers)
	}
	return domain.Report{SessionID: session.ID, CandidateName: session.CandidateName, Language: session.Language, Difficulty: session.Difficulty, DomainSkillName: session.DomainSkillName, InterviewerName: session.InterviewerName, Score: score, Answered: len(session.Answers), Duration: int(completed.Sub(session.StartedAt).Seconds()), Highlights: topUnique(highlights, 3), FocusAreas: topUnique(focus, 3), Answers: append([]domain.AnswerRecord(nil), session.Answers...), StartedAt: session.StartedAt, CompletedAt: completed}
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
