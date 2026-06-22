package learning

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"
)

type Resource struct {
	ID             string    `json:"id"`
	Title          string    `json:"title"`
	URL            string    `json:"url"`
	Kind           string    `json:"kind"`
	Source         string    `json:"source"`
	Authority      string    `json:"authority"`
	Summary        string    `json:"summary"`
	DomainSkillIDs []string  `json:"domainSkillIds"`
	PublishedAt    time.Time `json:"publishedAt,omitempty"`
	Selected       bool      `json:"selected"`
	Live           bool      `json:"live"`
}

type ListResult struct {
	Resources   []Resource `json:"resources"`
	RefreshedAt time.Time  `json:"refreshedAt,omitempty"`
	Warnings    []string   `json:"warnings,omitempty"`
}

type feedSource struct{ ID, Name, URL, DomainSkillID string }

var feeds = []feedSource{
	{ID: "go-blog", Name: "The Go Blog", URL: "https://go.dev/blog/feed.atom", DomainSkillID: "computer-golang"},
	{ID: "inside-java", Name: "Inside Java", URL: "https://inside.java/feed.xml", DomainSkillID: "computer-java"},
	{ID: "spring-blog", Name: "Spring Blog", URL: "https://spring.io/blog.atom", DomainSkillID: "computer-java"},
	{ID: "python-blog", Name: "Python Insider", URL: "https://blog.python.org/feeds/posts/default", DomainSkillID: "computer-python"},
	{ID: "cpp-blog", Name: "Microsoft C++ Team Blog", URL: "https://devblogs.microsoft.com/cppblog/feed/", DomainSkillID: "computer-cpp"},
}

var curated = []Resource{
	{ID: "course-mit-algorithms", Title: "MIT 6.006 · Introduction to Algorithms", URL: "https://ocw.mit.edu/courses/6-006-introduction-to-algorithms-fall-2011/", Kind: "course", Source: "MIT OpenCourseWare", Authority: "MIT", Summary: "包含完整讲义与课程视频的算法基础课程。", DomainSkillIDs: []string{"computer-foundation"}},
	{ID: "docs-go-tour", Title: "A Tour of Go", URL: "https://go.dev/tour/", Kind: "docs", Source: "Go Documentation", Authority: "Go Team", Summary: "Go 官方交互式语言导览，适合系统补齐语法与并发基础。", DomainSkillIDs: []string{"computer-golang"}},
	{ID: "video-go-official", Title: "Go 官方视频频道", URL: "https://www.youtube.com/@golang", Kind: "video", Source: "Go YouTube", Authority: "Go Team", Summary: "Go 团队发布的大会演讲、版本解析和工程实践视频。", DomainSkillIDs: []string{"computer-golang"}},
	{ID: "docs-java-learn", Title: "Learn Java", URL: "https://dev.java/learn/", Kind: "docs", Source: "dev.java", Authority: "Oracle Java Team", Summary: "Java 官方学习路径，覆盖语言基础、集合、并发和现代特性。", DomainSkillIDs: []string{"computer-java"}},
	{ID: "video-java-official", Title: "Java 官方视频频道", URL: "https://www.youtube.com/@java", Kind: "video", Source: "Java YouTube", Authority: "Oracle Java Team", Summary: "Java 平台更新、JVM 和开发实践的官方内容。", DomainSkillIDs: []string{"computer-java"}},
	{ID: "docs-python-tutorial", Title: "The Python Tutorial", URL: "https://docs.python.org/3/tutorial/", Kind: "docs", Source: "Python Documentation", Authority: "Python Software Foundation", Summary: "Python 官方教程，覆盖语言模型、模块、异常、类与标准库。", DomainSkillIDs: []string{"computer-python"}},
	{ID: "video-pycon-us", Title: "PyCon US 视频频道", URL: "https://www.youtube.com/c/pyconus", Kind: "video", Source: "PyCon US", Authority: "Python Software Foundation Community", Summary: "Python 社区权威大会的视频合集，覆盖语言、性能、数据与工程。", DomainSkillIDs: []string{"computer-python"}},
	{ID: "docs-cpp-guidelines", Title: "C++ Core Guidelines", URL: "https://isocpp.github.io/CppCoreGuidelines/CppCoreGuidelines", Kind: "docs", Source: "C++ Core Guidelines", Authority: "Standard C++ Foundation", Summary: "现代 C++ 设计、资源管理、类型安全和并发实践指南。", DomainSkillIDs: []string{"computer-cpp"}},
	{ID: "video-cppcon", Title: "CppCon 视频频道", URL: "https://www.youtube.com/@CppCon", Kind: "video", Source: "CppCon", Authority: "C++ Community", Summary: "C++ 领域最具代表性的技术大会演讲合集。", DomainSkillIDs: []string{"computer-cpp"}},
}

type Service struct {
	mu            sync.RWMutex
	selectionPath string
	selected      map[string]bool
	live          map[string]Resource
	refreshedAt   time.Time
	warnings      []string
	client        *http.Client
}

func NewService(dataDir string) (*Service, error) {
	service := &Service{selectionPath: filepath.Join(dataDir, "learning-selections.json"), selected: make(map[string]bool), live: make(map[string]Resource), client: &http.Client{Timeout: 12 * time.Second}}
	data, err := os.ReadFile(service.selectionPath)
	if err == nil {
		if err := json.Unmarshal(data, &service.selected); err != nil {
			return nil, err
		}
	} else if !os.IsNotExist(err) {
		return nil, err
	}
	return service, nil
}

func (s *Service) List(domainSkillID, kind string, selectedOnly bool) ListResult {
	s.mu.RLock()
	defer s.mu.RUnlock()
	resources := make([]Resource, 0, len(curated)+len(s.live))
	appendIfMatch := func(resource Resource) {
		resource.Selected = s.selected[resource.ID]
		if selectedOnly && !resource.Selected {
			return
		}
		if kind != "" && kind != "all" && resource.Kind != kind {
			return
		}
		if domainSkillID != "" && domainSkillID != "all" && !contains(resource.DomainSkillIDs, domainSkillID) {
			return
		}
		resources = append(resources, resource)
	}
	for _, resource := range curated {
		appendIfMatch(resource)
	}
	for _, resource := range s.live {
		appendIfMatch(resource)
	}
	sort.SliceStable(resources, func(i, j int) bool {
		if resources[i].Selected != resources[j].Selected {
			return resources[i].Selected
		}
		if resources[i].PublishedAt.IsZero() != resources[j].PublishedAt.IsZero() {
			return !resources[i].PublishedAt.IsZero()
		}
		return resources[i].PublishedAt.After(resources[j].PublishedAt)
	})
	return ListResult{Resources: resources, RefreshedAt: s.refreshedAt, Warnings: append([]string(nil), s.warnings...)}
}

func (s *Service) Refresh(ctx context.Context) ListResult {
	type result struct {
		source    feedSource
		resources []Resource
		err       error
	}
	results := make(chan result, len(feeds))
	var wg sync.WaitGroup
	for _, source := range feeds {
		wg.Add(1)
		go func(source feedSource) {
			defer wg.Done()
			resources, err := s.fetchFeed(ctx, source)
			results <- result{source: source, resources: resources, err: err}
		}(source)
	}
	go func() { wg.Wait(); close(results) }()
	s.mu.RLock()
	next := make(map[string]Resource, len(s.live))
	for id, resource := range s.live {
		next[id] = resource
	}
	s.mu.RUnlock()
	warnings := make([]string, 0)
	for result := range results {
		if result.err != nil {
			warnings = append(warnings, result.source.Name+"："+result.err.Error())
			continue
		}
		for id, resource := range next {
			if resource.Source == result.source.Name {
				delete(next, id)
			}
		}
		for _, resource := range result.resources {
			next[resource.ID] = resource
		}
	}
	s.mu.Lock()
	s.live = next
	s.refreshedAt = time.Now()
	s.warnings = warnings
	s.mu.Unlock()
	return s.List("all", "all", false)
}

func (s *Service) SetSelected(id string, selected bool) error {
	if strings.TrimSpace(id) == "" {
		return fmt.Errorf("资源 ID 不能为空")
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if selected {
		s.selected[id] = true
	} else {
		delete(s.selected, id)
	}
	data, err := json.MarshalIndent(s.selected, "", "  ")
	if err != nil {
		return err
	}
	tmp := s.selectionPath + ".tmp"
	if err := os.WriteFile(tmp, data, 0o600); err != nil {
		return err
	}
	return os.Rename(tmp, s.selectionPath)
}

func (s *Service) fetchFeed(ctx context.Context, source feedSource) ([]Resource, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, source.URL, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("User-Agent", "Interview-Copilot/0.2 (+local learning aggregator)")
	response, err := s.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP %d", response.StatusCode)
	}
	data, err := io.ReadAll(io.LimitReader(response.Body, 4<<20))
	if err != nil {
		return nil, err
	}
	return parseFeed(data, source)
}

type feedDocument struct {
	Channel struct {
		Items []struct {
			Title       string `xml:"title"`
			Link        string `xml:"link"`
			Description string `xml:"description"`
			PubDate     string `xml:"pubDate"`
		} `xml:"item"`
	} `xml:"channel"`
	Entries []struct {
		Title string `xml:"title"`
		Links []struct {
			Href string `xml:"href,attr"`
			Rel  string `xml:"rel,attr"`
		} `xml:"link"`
		Summary   string `xml:"summary"`
		Content   string `xml:"content"`
		Updated   string `xml:"updated"`
		Published string `xml:"published"`
	} `xml:"entry"`
}

func parseFeed(data []byte, source feedSource) ([]Resource, error) {
	var document feedDocument
	if err := xml.Unmarshal(data, &document); err != nil {
		return nil, err
	}
	resources := make([]Resource, 0, 12)
	for _, item := range document.Channel.Items {
		if len(resources) >= 12 {
			break
		}
		link, ok := safeFeedURL(item.Link, source.URL)
		if strings.TrimSpace(item.Title) == "" || !ok {
			continue
		}
		resources = append(resources, Resource{ID: "feed-" + shortHash(link), Title: clean(item.Title, 180), URL: link, Kind: "article", Source: source.Name, Authority: source.Name, Summary: clean(item.Description, 240), DomainSkillIDs: []string{source.DomainSkillID}, PublishedAt: parseDate(item.PubDate), Live: true})
	}
	for _, entry := range document.Entries {
		if len(resources) >= 12 {
			break
		}
		link := ""
		for _, candidate := range entry.Links {
			if candidate.Rel == "" || candidate.Rel == "alternate" {
				link = candidate.Href
				break
			}
		}
		link, ok := safeFeedURL(link, source.URL)
		if strings.TrimSpace(entry.Title) == "" || !ok {
			continue
		}
		summary := entry.Summary
		if summary == "" {
			summary = entry.Content
		}
		date := entry.Published
		if date == "" {
			date = entry.Updated
		}
		resources = append(resources, Resource{ID: "feed-" + shortHash(link), Title: clean(entry.Title, 180), URL: link, Kind: "article", Source: source.Name, Authority: source.Name, Summary: clean(summary, 240), DomainSkillIDs: []string{source.DomainSkillID}, PublishedAt: parseDate(date), Live: true})
	}
	if len(resources) == 0 {
		return nil, fmt.Errorf("没有可读取的条目")
	}
	return resources, nil
}

var tagPattern = regexp.MustCompile(`<[^>]+>`)

func clean(value string, limit int) string {
	value = html.UnescapeString(tagPattern.ReplaceAllString(value, " "))
	value = strings.Join(strings.Fields(value), " ")
	runes := []rune(value)
	if len(runes) > limit {
		value = string(runes[:limit]) + "…"
	}
	return value
}
func parseDate(value string) time.Time {
	for _, layout := range []string{time.RFC3339, time.RFC1123Z, time.RFC1123, time.RFC822Z, "Mon, 02 Jan 2006 15:04:05 -0700"} {
		if parsed, err := time.Parse(layout, strings.TrimSpace(value)); err == nil {
			return parsed
		}
	}
	return time.Time{}
}
func shortHash(value string) string {
	sum := sha256.Sum256([]byte(value))
	return hex.EncodeToString(sum[:8])
}
func contains(values []string, target string) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}

func safeFeedURL(value, feedURL string) (string, bool) {
	parsed, err := url.Parse(strings.TrimSpace(value))
	if err != nil || parsed.Scheme != "https" || parsed.Hostname() == "" {
		return "", false
	}
	feed, err := url.Parse(feedURL)
	if err != nil || !strings.EqualFold(parsed.Hostname(), feed.Hostname()) {
		return "", false
	}
	return parsed.String(), true
}
