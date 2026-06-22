package questions

import "testing"

func TestSelectMixedQuestionsAndResumeRelevance(t *testing.T) {
	selected, err := Select("golang", "mixed", 5, []string{"性能优化"})
	if err != nil {
		t.Fatal(err)
	}
	if len(selected) != 5 {
		t.Fatalf("expected 5 questions, got %d", len(selected))
	}
	if selected[0].ID != "go-04" {
		t.Fatalf("expected resume-relevant question first, got %s", selected[0].ID)
	}
	for _, question := range selected {
		if question.StandardAnswer == "" || len(question.KeyPoints) == 0 {
			t.Fatalf("question %s is missing review material", question.ID)
		}
	}
}

func TestEveryLanguageHasQuestions(t *testing.T) {
	for _, language := range Languages() {
		selected, err := Select(language, "mixed", 6, nil)
		if err != nil {
			t.Fatalf("%s: %v", language, err)
		}
		if len(selected) < 5 {
			t.Fatalf("%s has only %d questions", language, len(selected))
		}
	}
}
