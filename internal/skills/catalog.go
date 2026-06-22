package skills

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"sort"
	"strings"
)

//go:embed manifests/interviewer/*.json manifests/domain/*/*.json
var manifestFS embed.FS

type Skill struct {
	ID              string   `json:"id"`
	Kind            string   `json:"kind"`
	Name            string   `json:"name"`
	ShortLabel      string   `json:"shortLabel"`
	Description     string   `json:"description"`
	Accent          string   `json:"accent"`
	Avatar          string   `json:"avatar,omitempty"`
	Prompt          string   `json:"prompt,omitempty"`
	OpeningLine     string   `json:"openingLine,omitempty"`
	EvaluationFocus []string `json:"evaluationFocus,omitempty"`
	FeedbackTone    string   `json:"feedbackTone,omitempty"`
	Industry        string   `json:"industry,omitempty"`
	IndustryName    string   `json:"industryName,omitempty"`
	Domain          string   `json:"domain,omitempty"`
	Language        string   `json:"language,omitempty"`
	KnowledgeBaseID string   `json:"knowledgeBaseId,omitempty"`
	Topics          []string `json:"topics,omitempty"`
}

type Industry struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type PublicCatalog struct {
	Interviewers []Skill    `json:"interviewers"`
	Domains      []Skill    `json:"domains"`
	Industries   []Industry `json:"industries"`
}

type Catalog struct {
	all          map[string]Skill
	interviewers []Skill
	domains      []Skill
}

func Load() (*Catalog, error) {
	catalog := &Catalog{all: make(map[string]Skill)}
	err := fs.WalkDir(manifestFS, "manifests", func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() || !strings.HasSuffix(path, ".json") {
			return nil
		}
		data, err := manifestFS.ReadFile(path)
		if err != nil {
			return err
		}
		var skill Skill
		if err := json.Unmarshal(data, &skill); err != nil {
			return fmt.Errorf("parse skill %s: %w", path, err)
		}
		if skill.ID == "" || skill.Kind == "" || skill.Name == "" {
			return fmt.Errorf("skill %s is missing required fields", path)
		}
		if _, exists := catalog.all[skill.ID]; exists {
			return fmt.Errorf("duplicate skill id %s", skill.ID)
		}
		catalog.all[skill.ID] = skill
		switch skill.Kind {
		case "interviewer":
			catalog.interviewers = append(catalog.interviewers, skill)
		case "domain":
			catalog.domains = append(catalog.domains, skill)
		default:
			return fmt.Errorf("unknown skill kind %s", skill.Kind)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	sort.Slice(catalog.interviewers, func(i, j int) bool { return catalog.interviewers[i].ID < catalog.interviewers[j].ID })
	sort.Slice(catalog.domains, func(i, j int) bool {
		if catalog.domains[i].Language == "foundation" {
			return true
		}
		if catalog.domains[j].Language == "foundation" {
			return false
		}
		return catalog.domains[i].ID < catalog.domains[j].ID
	})
	return catalog, nil
}

func (c *Catalog) Get(id string) (Skill, bool) { skill, ok := c.all[id]; return skill, ok }

func (c *Catalog) Interviewer(id string) (Skill, bool) {
	skill, ok := c.Get(id)
	return skill, ok && skill.Kind == "interviewer"
}

func (c *Catalog) Domain(id string) (Skill, bool) {
	skill, ok := c.Get(id)
	return skill, ok && skill.Kind == "domain"
}

func (c *Catalog) DomainByLanguage(language string) (Skill, bool) {
	for _, skill := range c.domains {
		if skill.Language == language {
			return skill, true
		}
	}
	return Skill{}, false
}

func (c *Catalog) Public() PublicCatalog {
	industries := make(map[string]string)
	for _, skill := range c.domains {
		if skill.Industry != "" {
			industries[skill.Industry] = skill.IndustryName
		}
	}
	result := PublicCatalog{Interviewers: append([]Skill(nil), c.interviewers...), Domains: append([]Skill(nil), c.domains...)}
	for i := range result.Interviewers {
		result.Interviewers[i].Prompt = ""
	}
	for i := range result.Domains {
		result.Domains[i].Prompt = ""
	}
	for id, name := range industries {
		result.Industries = append(result.Industries, Industry{ID: id, Name: name})
	}
	sort.Slice(result.Industries, func(i, j int) bool { return result.Industries[i].ID < result.Industries[j].ID })
	return result
}

func (c *Catalog) DomainSkills() []Skill { return append([]Skill(nil), c.domains...) }
