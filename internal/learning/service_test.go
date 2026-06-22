package learning

import "testing"

func TestParseAtomFeed(t *testing.T) {
	source := feedSource{Name: "Official", URL: "https://example.com/feed", DomainSkillID: "computer-golang"}
	resources, err := parseFeed([]byte(`<feed xmlns="http://www.w3.org/2005/Atom"><entry><title>Go &amp; Tools</title><link rel="alternate" href="https://example.com/go"/><summary>&lt;b&gt;Useful&lt;/b&gt; article</summary><updated>2026-06-20T00:00:00Z</updated></entry></feed>`), source)
	if err != nil {
		t.Fatal(err)
	}
	if len(resources) != 1 || resources[0].Summary != "Useful article" || resources[0].Kind != "article" {
		t.Fatalf("unexpected resources %#v", resources)
	}
}

func TestSelectionPersistence(t *testing.T) {
	dir := t.TempDir()
	service, err := NewService(dir)
	if err != nil {
		t.Fatal(err)
	}
	if err := service.SetSelected("docs-go-tour", true); err != nil {
		t.Fatal(err)
	}
	reloaded, err := NewService(dir)
	if err != nil {
		t.Fatal(err)
	}
	result := reloaded.List("computer-golang", "all", true)
	if len(result.Resources) != 1 || result.Resources[0].ID != "docs-go-tour" {
		t.Fatalf("unexpected selected resources %#v", result.Resources)
	}
}
