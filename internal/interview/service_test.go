package interview

import (
	"context"
	"testing"

	"interview-agent/internal/agent"
	"interview-agent/internal/domain"
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
	totalTurns := session.Total * (session.FollowUpTotal + 1)
	for index := 0; index < totalTurns; index++ {
		result, err := service.Answer(context.Background(), session.ID, AnswerInput{Answer: "这是一个包含技术细节的完整测试回答。", ElapsedSeconds: 12})
		if err != nil {
			t.Fatal(err)
		}
		if index < totalTurns-1 && result.NextQuestion == nil {
			t.Fatal("missing next question")
		}
		if index == totalTurns-1 {
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
				if len(answer.FollowUps) != session.FollowUpTotal {
					t.Fatalf("expected %d follow-ups, got %d", session.FollowUpTotal, len(answer.FollowUps))
				}
				if answer.AverageScore != 88 {
					t.Fatalf("unexpected average score %d", answer.AverageScore)
				}
			}
		}
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
	if result.Evaluation.Source != "local" {
		t.Fatalf("expected local fallback, got %s", result.Evaluation.Source)
	}
	followUpResult, err := service.Answer(context.Background(), session.ID, AnswerInput{Answer: "线上会通过限流和降级控制风险，并用 p99 指标和压测验证容量边界。"})
	if err != nil {
		t.Fatal(err)
	}
	if followUpResult.Evaluation.Source != "local" || followUpResult.Evaluation.Score < 65 {
		t.Fatalf("expected focused local follow-up evaluation, got %#v", followUpResult.Evaluation)
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

func TestFollowUpProgressionUsesStyleRounds(t *testing.T) {
	service := NewService(stubEvaluator{})
	session, err := service.Start(StartInput{DomainSkillID: "computer-golang", InterviewerSkillID: "atlas-architect", Difficulty: "mixed", QuestionCount: 1})
	if err != nil {
		t.Fatal(err)
	}
	if session.FollowUpTotal != 3 {
		t.Fatalf("expected 3 Atlas follow-ups, got %d", session.FollowUpTotal)
	}
	seenPrompts := map[string]bool{}
	for round := 1; round <= 3; round++ {
		result, err := service.Answer(context.Background(), session.ID, AnswerInput{Answer: "包含一些关键点的回答。"})
		if err != nil {
			t.Fatal(err)
		}
		if result.NextQuestion == nil || !result.NextQuestion.FollowUp || result.NextQuestion.Round != round {
			t.Fatalf("round %d did not produce expected follow-up: %#v", round, result.NextQuestion)
		}
		if seenPrompts[result.NextQuestion.Prompt] {
			t.Fatalf("follow-up repeated at round %d: %s", round, result.NextQuestion.Prompt)
		}
		seenPrompts[result.NextQuestion.Prompt] = true
	}
	result, err := service.Answer(context.Background(), session.ID, AnswerInput{Answer: "最后一轮补充回答。"})
	if err != nil {
		t.Fatal(err)
	}
	if !result.QuestionCompleted || !result.Completed {
		t.Fatalf("main question did not complete after follow-ups: %#v", result)
	}
}

func TestCustomMainQuestionCount(t *testing.T) {
	service := NewService(stubEvaluator{})
	session, err := service.Start(StartInput{DomainSkillID: "computer-python", InterviewerSkillID: "echo-coach", IncludeFoundation: true, Difficulty: "mixed", QuestionCount: 11})
	if err != nil {
		t.Fatal(err)
	}
	if session.Total != 11 {
		t.Fatalf("expected 11 custom main questions, got %d", session.Total)
	}
	if _, err := service.Start(StartInput{DomainSkillID: "computer-python", InterviewerSkillID: "echo-coach", Difficulty: "mixed", QuestionCount: 21}); err == nil {
		t.Fatal("expected question count validation error")
	}
}

type failingEvaluator struct{}

func (failingEvaluator) Evaluate(context.Context, agent.EvaluationInput) (domain.Evaluation, error) {
	return domain.Evaluation{}, agent.ErrNotConfigured
}
