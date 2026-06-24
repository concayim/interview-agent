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

func TestCustomQuestionCountUsesFoundationCapacity(t *testing.T) {
	selected, err := SelectWithFoundation("golang", "mixed", 11, nil, true)
	if err != nil {
		t.Fatal(err)
	}
	if len(selected) != 11 {
		t.Fatalf("expected 11 custom questions, got %d", len(selected))
	}
	seen := map[string]bool{}
	for _, question := range selected {
		if seen[question.ID] {
			t.Fatalf("duplicate question %s", question.ID)
		}
		seen[question.ID] = true
	}
	withoutFoundation, err := SelectWithFoundation("golang", "mixed", 6, nil, false)
	if err != nil || len(withoutFoundation) != 6 {
		t.Fatalf("expected 6 language questions, got %d (%v)", len(withoutFoundation), err)
	}
}

func TestSingleQuestionKeepsSelectedLanguage(t *testing.T) {
	selected, err := SelectWithFoundation("golang", "mixed", 1, nil, true)
	if err != nil {
		t.Fatal(err)
	}
	if len(selected) != 1 || selected[0].Language != "golang" {
		t.Fatalf("expected one Golang question, got %#v", selected)
	}
}
