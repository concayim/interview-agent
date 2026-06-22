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
