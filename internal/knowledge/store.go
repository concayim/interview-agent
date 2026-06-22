package knowledge

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"
	"unicode"

	"interview-agent/internal/domain"
	"interview-agent/internal/skills"
)

type Base struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Industry     string   `json:"industry"`
	IndustryName string   `json:"industryName"`
	Language     string   `json:"language"`
	Accent       string   `json:"accent"`
	Topics       []string `json:"topics"`
	ItemCount    int      `json:"itemCount"`
	IssuedCount  int      `json:"issuedCount"`
}

type QAItem struct {
	ID          string    `json:"id"`
	BaseID      string    `json:"baseId"`
	Question    string    `json:"question"`
	Answer      string    `json:"answer"`
	KeyPoints   []string  `json:"keyPoints"`
	Tags        []string  `json:"tags"`
	Language    string    `json:"language"`
	Difficulty  string    `json:"difficulty"`
	Source      string    `json:"source"`
	IssuedCount int       `json:"issuedCount"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type PutInput struct {
	Question   string   `json:"question"`
	Answer     string   `json:"answer"`
	KeyPoints  []string `json:"keyPoints"`
	Tags       []string `json:"tags"`
	Difficulty string   `json:"difficulty"`
}

type collection struct {
	Base  Base              `json:"base"`
	Items map[string]QAItem `json:"items"`
}

type Store struct {
	mu          sync.RWMutex
	dir         string
	collections map[string]*collection
}

func NewStore(dataDir string, domainSkills []skills.Skill) (*Store, error) {
	dir := filepath.Join(dataDir, "knowledge")
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return nil, err
	}
	store := &Store{dir: dir, collections: make(map[string]*collection)}
	for _, skill := range domainSkills {
		if skill.KnowledgeBaseID == "" {
			continue
		}
		base := Base{ID: skill.KnowledgeBaseID, Name: skill.Name, Description: skill.Description, Industry: skill.Industry, IndustryName: skill.IndustryName, Language: skill.Language, Accent: skill.Accent, Topics: append([]string(nil), skill.Topics...)}
		store.collections[base.ID] = &collection{Base: base, Items: make(map[string]QAItem)}
	}
	for id, target := range store.collections {
		data, err := os.ReadFile(store.path(id))
		if errors.Is(err, os.ErrNotExist) {
			continue
		}
		if err != nil {
			return nil, err
		}
		var saved collection
		if err := json.Unmarshal(data, &saved); err != nil {
			return nil, fmt.Errorf("load knowledge base %s: %w", id, err)
		}
		if saved.Items != nil {
			target.Items = saved.Items
		}
	}
	return store, nil
}

func (s *Store) SeedQuestions(questions []domain.Question) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	touched := map[string]bool{}
	now := time.Now()
	for _, question := range questions {
		baseID := baseIDForLanguage(question.Language)
		collection, ok := s.collections[baseID]
		if !ok {
			continue
		}
		id := "qa-" + question.ID
		existing, exists := collection.Items[id]
		if exists {
			existing.Question = question.Prompt
			existing.Answer = question.StandardAnswer
			existing.KeyPoints = append([]string(nil), question.KeyPoints...)
			existing.Tags = append([]string(nil), question.Tags...)
			existing.Difficulty = question.Difficulty
			existing.UpdatedAt = now
			collection.Items[id] = existing
		} else {
			collection.Items[id] = QAItem{ID: id, BaseID: baseID, Question: question.Prompt, Answer: question.StandardAnswer, KeyPoints: append([]string(nil), question.KeyPoints...), Tags: append([]string(nil), question.Tags...), Language: question.Language, Difficulty: question.Difficulty, Source: "built-in", CreatedAt: now, UpdatedAt: now}
		}
		touched[baseID] = true
	}
	for id := range touched {
		if err := s.persistLocked(id); err != nil {
			return err
		}
	}
	return nil
}

func (s *Store) RecordIssued(questions []domain.Question) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	touched := map[string]bool{}
	now := time.Now()
	for _, question := range questions {
		baseID := baseIDForLanguage(question.Language)
		collection, ok := s.collections[baseID]
		if !ok {
			continue
		}
		id := "qa-" + question.ID
		item, exists := collection.Items[id]
		if !exists {
			item = QAItem{ID: id, BaseID: baseID, Question: question.Prompt, Answer: question.StandardAnswer, KeyPoints: append([]string(nil), question.KeyPoints...), Tags: append([]string(nil), question.Tags...), Language: question.Language, Difficulty: question.Difficulty, Source: "interview", CreatedAt: now}
		}
		item.IssuedCount++
		item.UpdatedAt = now
		collection.Items[id] = item
		touched[baseID] = true
	}
	for id := range touched {
		if err := s.persistLocked(id); err != nil {
			return err
		}
	}
	return nil
}

func (s *Store) Put(baseID string, input PutInput) (QAItem, error) {
	input.Question = strings.TrimSpace(input.Question)
	input.Answer = strings.TrimSpace(input.Answer)
	if input.Question == "" || input.Answer == "" {
		return QAItem{}, fmt.Errorf("问题和答案不能为空")
	}
	if len([]rune(input.Question)) > 2000 || len([]rune(input.Answer)) > 12000 {
		return QAItem{}, fmt.Errorf("QA 内容过长")
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	collection, ok := s.collections[baseID]
	if !ok {
		return QAItem{}, fmt.Errorf("知识库不存在")
	}
	now := time.Now()
	id := "qa-custom-" + shortHash(input.Question)
	item := QAItem{ID: id, BaseID: baseID, Question: input.Question, Answer: input.Answer, KeyPoints: unique(input.KeyPoints), Tags: unique(input.Tags), Language: collection.Base.Language, Difficulty: input.Difficulty, Source: "manual", CreatedAt: now, UpdatedAt: now}
	if previous, exists := collection.Items[id]; exists {
		item.CreatedAt = previous.CreatedAt
		item.IssuedCount = previous.IssuedCount
	}
	collection.Items[id] = item
	if err := s.persistLocked(baseID); err != nil {
		return QAItem{}, err
	}
	return item, nil
}

func (s *Store) Bases() []Base {
	s.mu.RLock()
	defer s.mu.RUnlock()
	result := make([]Base, 0, len(s.collections))
	for _, collection := range s.collections {
		base := collection.Base
		base.ItemCount = len(collection.Items)
		for _, item := range collection.Items {
			base.IssuedCount += item.IssuedCount
		}
		result = append(result, base)
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i].Language == "foundation" {
			return true
		}
		if result[j].Language == "foundation" {
			return false
		}
		return result[i].ID < result[j].ID
	})
	return result
}

func (s *Store) Search(baseID, query string, limit int) ([]QAItem, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	collection, ok := s.collections[baseID]
	if !ok {
		return nil, fmt.Errorf("知识库不存在")
	}
	if limit < 1 {
		limit = 20
	}
	if limit > 100 {
		limit = 100
	}
	type ranked struct {
		item  QAItem
		score int
	}
	rankedItems := make([]ranked, 0, len(collection.Items))
	for _, item := range collection.Items {
		score := searchScore(item, query)
		if strings.TrimSpace(query) != "" && score == 0 {
			continue
		}
		rankedItems = append(rankedItems, ranked{item: item, score: score})
	}
	sort.SliceStable(rankedItems, func(i, j int) bool {
		if rankedItems[i].score != rankedItems[j].score {
			return rankedItems[i].score > rankedItems[j].score
		}
		if rankedItems[i].item.IssuedCount != rankedItems[j].item.IssuedCount {
			return rankedItems[i].item.IssuedCount > rankedItems[j].item.IssuedCount
		}
		return rankedItems[i].item.UpdatedAt.After(rankedItems[j].item.UpdatedAt)
	})
	if len(rankedItems) > limit {
		rankedItems = rankedItems[:limit]
	}
	result := make([]QAItem, len(rankedItems))
	for i := range rankedItems {
		result[i] = rankedItems[i].item
	}
	return result, nil
}

func (s *Store) path(id string) string { return filepath.Join(s.dir, id+".json") }

func (s *Store) persistLocked(id string) error {
	collection, ok := s.collections[id]
	if !ok {
		return fmt.Errorf("knowledge base %s not found", id)
	}
	data, err := json.MarshalIndent(collection, "", "  ")
	if err != nil {
		return err
	}
	tmp := s.path(id) + ".tmp"
	if err := os.WriteFile(tmp, data, 0o600); err != nil {
		return err
	}
	return os.Rename(tmp, s.path(id))
}

func baseIDForLanguage(language string) string {
	if language == "foundation" {
		return "computer-foundation"
	}
	return "computer-" + language
}

func searchScore(item QAItem, query string) int {
	query = strings.ToLower(strings.TrimSpace(query))
	if query == "" {
		return item.IssuedCount
	}
	question := strings.ToLower(item.Question)
	answer := strings.ToLower(item.Answer)
	tags := strings.ToLower(strings.Join(item.Tags, " "))
	score := 0
	if strings.Contains(question, query) {
		score += 20
	}
	if strings.Contains(tags, query) {
		score += 16
	}
	if strings.Contains(answer, query) {
		score += 8
	}
	for _, token := range tokens(query) {
		if strings.Contains(question, token) {
			score += 5
		}
		if strings.Contains(tags, token) {
			score += 4
		}
		if strings.Contains(answer, token) {
			score += 2
		}
	}
	return score
}

func tokens(value string) []string {
	return strings.FieldsFunc(value, func(r rune) bool { return unicode.IsSpace(r) || unicode.IsPunct(r) || unicode.IsSymbol(r) })
}

func shortHash(value string) string {
	sum := sha256.Sum256([]byte(value))
	return hex.EncodeToString(sum[:8])
}

func unique(values []string) []string {
	seen := map[string]bool{}
	result := make([]string, 0, len(values))
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value != "" && !seen[value] {
			seen[value] = true
			result = append(result, value)
		}
	}
	return result
}
