package interview

import (
	"context"
	"strings"
	"testing"
	"unicode"

	"interview-agent/internal/agent"
	"interview-agent/internal/domain"
	"interview-agent/internal/knowledge"
	"interview-agent/internal/questions"
	"interview-agent/internal/skills"
)

type stubEvaluator struct{}

func (stubEvaluator) Evaluate(context.Context, agent.EvaluationInput) (domain.Evaluation, error) {
	return domain.Evaluation{Score: 88, Summary: "测试评价", Strengths: []string{"结构清晰"}, Improvements: []string{"补充边界"}, Source: "llm"}, nil
}

func TestInterviewLifecycleIncludesStandardAnswersInReport(t *testing.T) {
	service := NewService(stubEvaluator{})
	session, err := service.Start(StartInput{CandidateName: "小林", Language: "python", Difficulty: "easy", QuestionCount: 2})
	if err != nil {
		t.Fatal(err)
	}
	if session.CurrentQuestion == nil {
		t.Fatal("missing first question")
	}
	for index := 0; index < session.Total; index++ {
		result, err := service.Answer(context.Background(), session.ID, AnswerInput{Answer: "这是一个包含技术细节的完整测试回答。", ElapsedSeconds: 12})
		if err != nil {
			t.Fatal(err)
		}
		if index < session.Total-1 && result.NextQuestion == nil {
			t.Fatal("missing next question")
		}
		if index == session.Total-1 {
			if !result.Completed || result.Report == nil {
				t.Fatal("expected completed report")
			}
			if len(result.Report.Answers) != session.Total {
				t.Fatalf("expected %d report answers", session.Total)
			}
			for _, answer := range result.Report.Answers {
				if answer.Question.StandardAnswer == "" {
					t.Fatal("standard answer was not included")
				}
			}
		}
	}
}

func TestStartPrefersKnowledgeBaseQAOverModelGeneration(t *testing.T) {
	catalog, err := skills.Load()
	if err != nil {
		t.Fatal(err)
	}
	store, err := knowledge.NewStore(t.TempDir(), catalog.DomainSkills())
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.Put("computer-golang", knowledge.PutInput{
		Question:   "知识库里的 goroutine 调度题",
		Answer:     "标准答案说明 GMP 调度模型。",
		KeyPoints:  []string{"GMP", "调度"},
		Tags:       []string{"Go", "并发"},
		Difficulty: "medium",
	})
	if err != nil {
		t.Fatal(err)
	}

	service := NewService(generatingEvaluator{}, WithCatalog(catalog), WithKnowledge(store))
	session, err := service.Start(StartInput{DomainSkillID: "computer-golang", Difficulty: "medium", QuestionCount: 1})
	if err != nil {
		t.Fatal(err)
	}
	if session.QuestionSource != "knowledge" {
		t.Fatalf("expected knowledge question source, got %q", session.QuestionSource)
	}
	if session.CurrentQuestion == nil || session.CurrentQuestion.Prompt != "知识库里的 goroutine 调度题" {
		t.Fatalf("expected stored QA question, got %#v", session.CurrentQuestion)
	}
}

func TestEnglishInterviewDoesNotUseChineseKnowledgeQA(t *testing.T) {
	catalog, err := skills.Load()
	if err != nil {
		t.Fatal(err)
	}
	store, err := knowledge.NewStore(t.TempDir(), catalog.DomainSkills())
	if err != nil {
		t.Fatal(err)
	}
	_, err = store.Put("computer-rust", knowledge.PutInput{
		Question:   "请说明 Rust 所有权。",
		Answer:     "每个值都有唯一所有者。",
		KeyPoints:  []string{"所有权"},
		Tags:       []string{"Rust"},
		Difficulty: "easy",
	})
	if err != nil {
		t.Fatal(err)
	}

	service := NewService(failingEvaluator{}, WithCatalog(catalog), WithKnowledge(store))
	session, err := service.Start(StartInput{DomainSkillID: "computer-rust", SpeechLanguage: "en-US", Difficulty: "easy", QuestionCount: 1})
	if err != nil {
		t.Fatal(err)
	}
	if session.QuestionSource == "knowledge" || session.CurrentQuestion == nil || containsHan(session.CurrentQuestion.Prompt) {
		t.Fatalf("expected English fallback instead of Chinese knowledge QA, got %#v", session)
	}
}

func TestScoreBelow75KeepsQuestionAndScore75Advances(t *testing.T) {
	evaluator := &adaptiveEvaluator{scores: []int{74, 75}}
	service := NewService(evaluator)
	session, err := service.Start(StartInput{Language: "golang", InterviewerSkillID: "vera-challenger", Difficulty: "easy", QuestionCount: 1})
	if err != nil {
		t.Fatal(err)
	}

	firstAnswer := "我认为 goroutine 是一种轻量线程。"
	result, err := service.Answer(context.Background(), session.ID, AnswerInput{Answer: firstAnswer, ElapsedSeconds: 12})
	if err != nil {
		t.Fatal(err)
	}
	if !result.Accepted || !result.RequiresFollowUp || result.Completed {
		t.Fatalf("score 74 should request a follow-up, got %#v", result)
	}
	if result.Current != 0 || result.NextQuestion == nil || result.NextQuestion.Prompt != "请具体说明它由谁调度，以及如何映射到线程。" {
		t.Fatalf("low score advanced or lost follow-up: %#v", result)
	}
	current, err := service.Get(session.ID)
	if err != nil {
		t.Fatal(err)
	}
	if current.CurrentQuestion == nil || current.CurrentQuestion.ID != result.NextQuestion.ID || current.CurrentQuestion.Prompt != result.NextQuestion.Prompt {
		t.Fatalf("session view did not preserve active follow-up: %#v", current)
	}

	secondAnswer := "运行时通过 GMP 调度，把 G 分配给持有 P 的 M 执行。"
	result, err = service.Answer(context.Background(), session.ID, AnswerInput{Answer: secondAnswer, ElapsedSeconds: 9})
	if err != nil {
		t.Fatal(err)
	}
	if result.RequiresFollowUp || !result.Completed || result.Current != 1 || result.Report == nil {
		t.Fatalf("score 75 should complete the one-question interview, got %#v", result)
	}
	if len(evaluator.inputs) != 2 || len(evaluator.inputs[1].PreviousAnswers) != 1 || evaluator.inputs[1].PreviousAnswers[0] != firstAnswer {
		t.Fatalf("follow-up evaluation did not receive prior attempts: %#v", evaluator.inputs)
	}
	if evaluator.inputs[1].Question.Prompt != "请具体说明它由谁调度，以及如何映射到线程。" {
		t.Fatalf("expected follow-up prompt in second evaluation, got %#v", evaluator.inputs[1])
	}
	if evaluator.inputs[0].InterviewerName != "Vera · 压力挑战官" || evaluator.inputs[0].InterviewerPrompt == "" {
		t.Fatalf("interviewer style was not passed to evaluation: %#v", evaluator.inputs[0])
	}
	if !strings.Contains(result.Report.Answers[0].Answer, firstAnswer) || !strings.Contains(result.Report.Answers[0].Answer, secondAnswer) {
		t.Fatalf("report did not preserve cumulative answers: %q", result.Report.Answers[0].Answer)
	}
}

func TestLocalFallbackWhenModelUnavailable(t *testing.T) {
	service := NewService(failingEvaluator{})
	session, err := service.Start(StartInput{Language: "golang", Difficulty: "easy", QuestionCount: 1})
	if err != nil {
		t.Fatal(err)
	}
	result, err := service.Answer(context.Background(), session.ID, AnswerInput{Answer: "goroutine 使用 GMP 调度并复用线程，栈可以动态增长。"})
	if err != nil {
		t.Fatal(err)
	}
	if result.Evaluation == nil || result.Evaluation.Source != "local" {
		t.Fatalf("expected local fallback, got %#v", result.Evaluation)
	}
}

func TestSynchronizedSkillStartsWithGenericLocalQuestions(t *testing.T) {
	service := NewService(failingEvaluator{})
	session, err := service.Start(StartInput{DomainSkillID: "computer-rust", Difficulty: "mixed", QuestionCount: 5})
	if err != nil {
		t.Fatal(err)
	}
	if session.DomainSkillName != "Rust 工程师" || session.Total != 5 || session.QuestionSource != "built-in" {
		t.Fatalf("unexpected synchronized skill session %#v", session)
	}
	if session.CurrentQuestion == nil || session.CurrentQuestion.ID != "rust-generic-easy-1" {
		t.Fatalf("expected generic Rust fallback question, got %#v", session.CurrentQuestion)
	}
}

func TestRustLifetimeAnswerMatchesConcreteLocalKeyPoints(t *testing.T) {
	questions := questions.Generic("rust", "Rust 工程师", []string{"所有权", "生命周期", "并发安全"}, "medium", 1)
	if len(questions) != 1 {
		t.Fatalf("expected one Rust question, got %d", len(questions))
	}
	answer := "生命周期用于描述引用有效期之间的约束，编译器据此避免悬垂引用；函数返回借用数据时需要用生命周期参数表达输入输出引用关系。"
	evaluation := localEvaluate(questions[0], answer, agent.ErrNotConfigured, "echo-coach", "zh-CN")
	if evaluation.Score < 80 {
		t.Fatalf("expected a strong Rust lifetime answer to score at least 80, got %d: %s", evaluation.Score, evaluation.Summary)
	}
}

func TestRustLifetimeGenericWordingDoesNotEarnScenarioPoint(t *testing.T) {
	generated := questions.Generic("rust", "Rust 工程师", []string{"所有权", "生命周期", "并发安全"}, "medium", 1)
	answer := "生命周期是 Rust 的重要机制，它有工作机制、边界条件和实际场景，函数返回时也可能用到。"
	evaluation := localEvaluate(generated[0], answer, agent.ErrNotConfigured, "echo-coach", "zh-CN")
	if evaluation.Score > 40 {
		t.Fatalf("generic wording should not match concrete borrowing concepts, got %d: %s", evaluation.Score, evaluation.Summary)
	}
}

func TestLocalFallbackCanPassExactStandardAnswerWithoutKeyPoints(t *testing.T) {
	question := domain.Question{Prompt: "什么是幂等？", StandardAnswer: "相同请求执行一次或多次产生相同的业务结果。", Tags: []string{"API"}}
	evaluation := localEvaluate(question, question.StandardAnswer, agent.ErrNotConfigured, "echo-coach", "zh-CN")
	if evaluation.Score < 75 {
		t.Fatalf("exact standard answer should pass without key points, got %d", evaluation.Score)
	}
}

func TestLowScoreFeedbackDoesNotRevealKnowledgeAnswer(t *testing.T) {
	service := NewService(leakyEvaluator{})
	session, err := service.Start(StartInput{Language: "golang", Difficulty: "easy", QuestionCount: 1})
	if err != nil {
		t.Fatal(err)
	}
	result, err := service.Answer(context.Background(), session.ID, AnswerInput{Answer: "我只知道它是轻量级的。"})
	if err != nil {
		t.Fatal(err)
	}
	if result.Evaluation == nil || !result.RequiresFollowUp {
		t.Fatalf("expected follow-up evaluation, got %#v", result)
	}
	exposed := strings.Join(append(append([]string{result.Evaluation.Summary}, result.Evaluation.Strengths...), result.Evaluation.Improvements...), " ")
	if strings.Contains(exposed, "GMP") || strings.Contains(exposed, "P 与 M") {
		t.Fatalf("in-progress feedback leaked answer details: %q", exposed)
	}
	if strings.Contains(result.Evaluation.FollowUpQuestion, "GMP") || strings.Contains(result.Evaluation.FollowUpQuestion, "P 与 M") {
		t.Fatalf("follow-up leaked answer details: %q", result.Evaluation.FollowUpQuestion)
	}
}

func TestCandidateCanAskForHintWithoutAdvancing(t *testing.T) {
	service := NewService(stubEvaluator{})
	session, err := service.Start(StartInput{Language: "python", Difficulty: "easy", QuestionCount: 1})
	if err != nil {
		t.Fatal(err)
	}
	result, err := service.Answer(context.Background(), session.ID, AnswerInput{Answer: "提示一下"})
	if err != nil {
		t.Fatal(err)
	}
	if result.Accepted {
		t.Fatal("hint should not be accepted as a final answer")
	}
	if result.Intent != "hint" || result.AssistantReply == "" {
		t.Fatalf("expected hint reply, got %#v", result)
	}
	next, err := service.Get(session.ID)
	if err != nil {
		t.Fatal(err)
	}
	if next.Current != 0 || next.CurrentQuestion == nil || next.CurrentQuestion.ID != session.CurrentQuestion.ID {
		t.Fatalf("question advanced after hint: %#v", next)
	}
}

func TestStartUsesGeneratedQuestionsWhenAvailable(t *testing.T) {
	service := NewService(generatingEvaluator{})
	session, err := service.Start(StartInput{Language: "golang", Difficulty: "mixed", QuestionCount: 2})
	if err != nil {
		t.Fatal(err)
	}
	if session.QuestionSource != "model" {
		t.Fatalf("expected model questions, got %s", session.QuestionSource)
	}
	if session.CurrentQuestion == nil || session.CurrentQuestion.ID != "generated-test-1" {
		t.Fatalf("expected generated current question, got %#v", session.CurrentQuestion)
	}
}

func TestEnglishSessionPassesOutputLanguageToQuestionGenerator(t *testing.T) {
	evaluator := &languageCapturingEvaluator{}
	service := NewService(evaluator)
	_, err := service.Start(StartInput{DomainSkillID: "computer-rust", SpeechLanguage: "en-US", Difficulty: "mixed", QuestionCount: 2})
	if err != nil {
		t.Fatal(err)
	}
	if evaluator.input.OutputLanguage != "en-US" {
		t.Fatalf("expected en-US output language, got %#v", evaluator.input)
	}
}

func TestEnglishSessionUsesEnglishQuestionsWhenModelUnavailable(t *testing.T) {
	service := NewService(failingEvaluator{})
	session, err := service.Start(StartInput{DomainSkillID: "computer-rust", SpeechLanguage: "english", Difficulty: "mixed", QuestionCount: 5, IncludeFoundation: true})
	if err != nil {
		t.Fatal(err)
	}
	if session.SpeechLanguage != "en-US" || session.CurrentQuestion == nil {
		t.Fatalf("unexpected English session %#v", session)
	}
	if strings.IndexFunc(session.CurrentQuestion.Prompt, func(r rune) bool { return unicode.Is(unicode.Han, r) }) >= 0 {
		t.Fatalf("expected English fallback question, got %q", session.CurrentQuestion.Prompt)
	}
}

func TestEnglishSessionUsesEnglishDefaultCandidateName(t *testing.T) {
	service := NewService(failingEvaluator{})
	session, err := service.Start(StartInput{DomainSkillID: "computer-rust", SpeechLanguage: "en-US", Difficulty: "mixed", QuestionCount: 1})
	if err != nil {
		t.Fatal(err)
	}
	if session.CandidateName != "Candidate" {
		t.Fatalf("expected English default candidate name, got %q", session.CandidateName)
	}
}

func TestEnglishSessionRejectsMixedLanguageModelQuestions(t *testing.T) {
	service := NewService(generatingEvaluator{})
	session, err := service.Start(StartInput{DomainSkillID: "computer-rust", SpeechLanguage: "en-US", Difficulty: "mixed", QuestionCount: 2})
	if err != nil {
		t.Fatal(err)
	}
	if session.QuestionSource != "built-in" || session.CurrentQuestion == nil || containsHan(session.CurrentQuestion.Prompt) {
		t.Fatalf("expected English fallback after mixed-language model output, got %#v", session)
	}
}

func TestEnglishRepeatAndHintRepliesStayEnglishWithoutAdvancing(t *testing.T) {
	for _, message := range []string{"Could you repeat the question?", "Can you give me a hint?"} {
		service := NewService(failingEvaluator{})
		session, err := service.Start(StartInput{DomainSkillID: "computer-rust", SpeechLanguage: "en-US", Difficulty: "mixed", QuestionCount: 2})
		if err != nil {
			t.Fatal(err)
		}
		result, err := service.Answer(context.Background(), session.ID, AnswerInput{Answer: message})
		if err != nil {
			t.Fatal(err)
		}
		if result.Accepted || result.AssistantReply == "" || containsHan(result.AssistantReply) {
			t.Fatalf("expected English non-advancing reply for %q, got %#v", message, result)
		}
		current, err := service.Get(session.ID)
		if err != nil || current.Current != 0 {
			t.Fatalf("English request advanced the interview: %#v, %v", current, err)
		}
	}
}

func TestEnglishSkipUsesEnglishLocalEvaluation(t *testing.T) {
	service := NewService(failingEvaluator{})
	session, err := service.Start(StartInput{DomainSkillID: "computer-rust", SpeechLanguage: "en-US", Difficulty: "mixed", QuestionCount: 2})
	if err != nil {
		t.Fatal(err)
	}
	result, err := service.Answer(context.Background(), session.ID, AnswerInput{Answer: "Please skip this one."})
	if err != nil {
		t.Fatal(err)
	}
	if !result.Accepted || result.Intent != "skip" || result.Evaluation == nil || containsHan(result.Evaluation.Summary) {
		t.Fatalf("expected English skip evaluation, got %#v", result)
	}
}

func TestEnglishAnswerUsesEnglishLocalEvaluation(t *testing.T) {
	service := NewService(failingEvaluator{})
	session, err := service.Start(StartInput{DomainSkillID: "computer-rust", SpeechLanguage: "en-US", Difficulty: "mixed", QuestionCount: 1})
	if err != nil {
		t.Fatal(err)
	}
	result, err := service.Answer(context.Background(), session.ID, AnswerInput{Answer: "Ownership gives each value one owner, and borrowing lets references access data without taking ownership."})
	if err != nil {
		t.Fatal(err)
	}
	if result.Evaluation == nil || result.Evaluation.Source != "local" {
		t.Fatalf("expected local evaluation, got %#v", result)
	}
	values := append([]string{result.Evaluation.Summary}, result.Evaluation.Strengths...)
	values = append(values, result.Evaluation.Improvements...)
	for _, value := range values {
		if containsHan(value) {
			t.Fatalf("English evaluation contains Han characters: %q", value)
		}
	}
}

func containsHan(value string) bool {
	return strings.IndexFunc(value, func(r rune) bool { return unicode.Is(unicode.Han, r) }) >= 0
}

func TestCandidateCanSkipQuestion(t *testing.T) {
	service := NewService(stubEvaluator{})
	session, err := service.Start(StartInput{Language: "golang", Difficulty: "easy", QuestionCount: 2})
	if err != nil {
		t.Fatal(err)
	}
	result, err := service.Answer(context.Background(), session.ID, AnswerInput{Answer: "这题先跳过", ElapsedSeconds: 8})
	if err != nil {
		t.Fatal(err)
	}
	if !result.Accepted || result.Intent != "skip" {
		t.Fatalf("expected accepted skip, got %#v", result)
	}
	if result.Evaluation == nil || result.Evaluation.Score != 0 || result.Evaluation.Source != "local" {
		t.Fatalf("expected local zero evaluation for skip, got %#v", result.Evaluation)
	}
	if result.Current != 1 || result.NextQuestion == nil {
		t.Fatalf("skip did not advance to next question: %#v", result)
	}
}

func TestModelIntentResolverHandlesOffTopicWithoutAdvancing(t *testing.T) {
	service := NewService(intentStubEvaluator{result: agent.IntentResult{Intent: "off_topic", Accepted: false, AssistantReply: "我理解你在换个话题，不过我们先回到当前题。"}})
	session, err := service.Start(StartInput{Language: "java", Difficulty: "easy", QuestionCount: 1})
	if err != nil {
		t.Fatal(err)
	}
	result, err := service.Answer(context.Background(), session.ID, AnswerInput{Answer: "今天天气不错，我想聊点别的"})
	if err != nil {
		t.Fatal(err)
	}
	if result.Accepted {
		t.Fatal("off-topic message should not be accepted as an answer")
	}
	if result.Intent != "off_topic" || result.AssistantReply == "" {
		t.Fatalf("expected off-topic reply, got %#v", result)
	}
	next, err := service.Get(session.ID)
	if err != nil {
		t.Fatal(err)
	}
	if next.Current != 0 {
		t.Fatalf("off-topic message advanced question: %#v", next)
	}
}

func TestSkillSelectionAndFoundationMix(t *testing.T) {
	service := NewService(stubEvaluator{})
	session, err := service.Start(StartInput{CandidateName: "小顾", DomainSkillID: "computer-java", InterviewerSkillID: "vera-challenger", IncludeFoundation: true, Difficulty: "mixed", QuestionCount: 5})
	if err != nil {
		t.Fatal(err)
	}
	if session.DomainSkillName != "Java 工程师" || session.InterviewerName != "Vera · 压力挑战官" {
		t.Fatalf("skills not applied: %#v", session)
	}
	stored := service.sessions[session.ID]
	foundationCount := 0
	for _, question := range stored.Questions {
		if question.Language == "foundation" {
			foundationCount++
		}
	}
	if foundationCount != 1 {
		t.Fatalf("expected one foundation question, got %d", foundationCount)
	}
}

type failingEvaluator struct{}

func (failingEvaluator) Evaluate(context.Context, agent.EvaluationInput) (domain.Evaluation, error) {
	return domain.Evaluation{}, agent.ErrNotConfigured
}

type intentStubEvaluator struct {
	result agent.IntentResult
}

func (intentStubEvaluator) Evaluate(context.Context, agent.EvaluationInput) (domain.Evaluation, error) {
	return domain.Evaluation{Score: 88, Summary: "测试评价", Strengths: []string{"结构清晰"}, Improvements: []string{"补充边界"}, Source: "llm"}, nil
}

func (e intentStubEvaluator) ResolveIntent(context.Context, agent.IntentInput) (agent.IntentResult, error) {
	return e.result, nil
}

type generatingEvaluator struct{}

func (generatingEvaluator) Evaluate(context.Context, agent.EvaluationInput) (domain.Evaluation, error) {
	return domain.Evaluation{Score: 90, Summary: "测试评价", Strengths: []string{"结构清晰"}, Improvements: []string{"补充边界"}, Source: "llm"}, nil
}

func (generatingEvaluator) GenerateQuestions(context.Context, agent.QuestionGenerationInput) ([]domain.Question, error) {
	return []domain.Question{
		{ID: "generated-test-1", Language: "golang", Difficulty: "medium", Prompt: "模型生成题 1", StandardAnswer: "标准答案", KeyPoints: []string{"关键点"}, Tags: []string{"Go"}},
		{ID: "generated-test-2", Language: "golang", Difficulty: "hard", Prompt: "模型生成题 2", StandardAnswer: "标准答案", KeyPoints: []string{"关键点"}, Tags: []string{"Go"}},
	}, nil
}

type languageCapturingEvaluator struct {
	input agent.QuestionGenerationInput
}

type adaptiveEvaluator struct {
	scores []int
	inputs []agent.EvaluationInput
}

type leakyEvaluator struct{}

func (leakyEvaluator) Evaluate(context.Context, agent.EvaluationInput) (domain.Evaluation, error) {
	return domain.Evaluation{
		Score:            30,
		Summary:          "标准答案要求 GMP 调度。",
		Strengths:        []string{"提到了 goroutine"},
		Improvements:     []string{"补充 P 与 M 的映射关系"},
		FollowUpQuestion: "GMP 中 P 与 M 如何映射？",
		Source:           "llm",
	}, nil
}

func (e *adaptiveEvaluator) Evaluate(_ context.Context, input agent.EvaluationInput) (domain.Evaluation, error) {
	e.inputs = append(e.inputs, input)
	score := e.scores[len(e.inputs)-1]
	evaluation := domain.Evaluation{Score: score, Summary: "相似度评分", Strengths: []string{"已覆盖部分概念"}, Improvements: []string{"补充调度细节"}, Source: "llm"}
	if score < 75 {
		evaluation.FollowUpQuestion = "请具体说明它由谁调度，以及如何映射到线程。"
	}
	return evaluation, nil
}

func (*languageCapturingEvaluator) Evaluate(context.Context, agent.EvaluationInput) (domain.Evaluation, error) {
	return domain.Evaluation{}, agent.ErrNotConfigured
}

func (e *languageCapturingEvaluator) GenerateQuestions(_ context.Context, input agent.QuestionGenerationInput) ([]domain.Question, error) {
	e.input = input
	return nil, agent.ErrNotConfigured
}
