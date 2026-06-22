package knowledge

import (
	"testing"

	"interview-agent/internal/domain"
	"interview-agent/internal/skills"
)

func TestSeedRecordAndSearch(t *testing.T) {
	catalog, err := skills.Load()
	if err != nil {
		t.Fatal(err)
	}
	store, err := NewStore(t.TempDir(), catalog.DomainSkills())
	if err != nil {
		t.Fatal(err)
	}
	question := domain.Question{ID: "test-go", Language: "golang", Difficulty: "easy", Prompt: "goroutine 如何调度？", StandardAnswer: "使用 GMP 调度。", KeyPoints: []string{"GMP"}, Tags: []string{"Go", "并发"}}
	if err := store.SeedQuestions([]domain.Question{question}); err != nil {
		t.Fatal(err)
	}
	if err := store.RecordIssued([]domain.Question{question}); err != nil {
		t.Fatal(err)
	}
	items, err := store.Search("computer-golang", "GMP", 10)
	if err != nil {
		t.Fatal(err)
	}
	if len(items) != 1 || items[0].IssuedCount != 1 {
		t.Fatalf("unexpected results %#v", items)
	}
	bases := store.Bases()
	if len(bases) < 5 {
		t.Fatalf("expected per-language bases, got %d", len(bases))
	}
}

func TestPutManualQA(t *testing.T) {
	catalog, _ := skills.Load()
	store, _ := NewStore(t.TempDir(), catalog.DomainSkills())
	item, err := store.Put("computer-foundation", PutInput{Question: "什么是时间复杂度？", Answer: "描述输入规模增长时算法耗时的渐近趋势。", Tags: []string{"算法"}})
	if err != nil {
		t.Fatal(err)
	}
	if item.Source != "manual" {
		t.Fatalf("unexpected source %s", item.Source)
	}
}
