package skills

import "testing"

func TestLoadCatalog(t *testing.T) {
	catalog, err := Load()
	if err != nil {
		t.Fatal(err)
	}
	public := catalog.Public()
	if len(public.Interviewers) < 4 {
		t.Fatalf("expected interviewer skills, got %d", len(public.Interviewers))
	}
	if len(public.Domains) < 5 {
		t.Fatalf("expected domain skills, got %d", len(public.Domains))
	}
	for _, skill := range public.Domains {
		if skill.KnowledgeBaseID == "" {
			t.Fatalf("domain skill %s has no knowledge base", skill.ID)
		}
	}
	vera, ok := catalog.Interviewer("vera-challenger")
	if !ok || vera.Prompt == "" || vera.IntroductionPrompt == "" || vera.FollowUpRounds != 3 {
		t.Fatal("private interviewer prompt was not loaded")
	}
	for _, skill := range public.Interviewers {
		if skill.Prompt != "" || skill.IntroductionPrompt != "" {
			t.Fatalf("public catalog leaked prompt for %s", skill.ID)
		}
	}
}
