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
	for _, language := range []string{"assembly", "c", "cpp", "csharp", "dart", "golang", "java", "javascript", "julia", "kotlin", "lua", "perl", "php", "powershell", "python", "r", "ruby", "rust", "scala", "sql", "swift", "typescript", "vbscript", "verilog", "zig"} {
		if _, ok := catalog.DomainByLanguage(language); !ok {
			t.Errorf("missing synchronized language skill %s", language)
		}
	}
	vera, ok := catalog.Interviewer("vera-challenger")
	if !ok || vera.Prompt == "" {
		t.Fatal("private interviewer prompt was not loaded")
	}
	for _, skill := range public.Interviewers {
		if skill.Prompt != "" {
			t.Fatalf("public catalog leaked prompt for %s", skill.ID)
		}
	}
}
