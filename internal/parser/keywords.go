package parser

import (
	"sort"
	"strings"
	"unicode"
)

type keywordDefinition struct {
	Name    string
	Aliases []string
}

var keywordDictionary = []keywordDefinition{
	{"Go", []string{"golang", "go language", "go语言", "go"}}, {"Java", []string{"java"}},
	{"Python", []string{"python"}}, {"C++", []string{"c++", "cpp"}}, {"JavaScript", []string{"javascript", "js"}},
	{"TypeScript", []string{"typescript"}}, {"React", []string{"react"}}, {"Vue", []string{"vue"}},
	{"Spring Boot", []string{"spring boot", "springboot"}}, {"Gin", []string{"gin"}}, {"Django", []string{"django"}},
	{"MySQL", []string{"mysql"}}, {"PostgreSQL", []string{"postgresql", "postgres"}}, {"Redis", []string{"redis"}},
	{"MongoDB", []string{"mongodb", "mongo"}}, {"Kafka", []string{"kafka"}}, {"RabbitMQ", []string{"rabbitmq"}},
	{"Docker", []string{"docker"}}, {"Kubernetes", []string{"kubernetes", "k8s"}}, {"Linux", []string{"linux"}},
	{"AWS", []string{"aws"}}, {"阿里云", []string{"阿里云"}}, {"微服务", []string{"微服务", "microservice"}},
	{"分布式系统", []string{"分布式", "distributed system"}}, {"高并发", []string{"高并发", "high concurrency"}},
	{"REST API", []string{"rest api", "restful"}}, {"gRPC", []string{"grpc"}}, {"GraphQL", []string{"graphql"}},
	{"Git", []string{"git"}}, {"CI/CD", []string{"ci/cd", "cicd"}}, {"DevOps", []string{"devops"}},
	{"机器学习", []string{"机器学习", "machine learning"}}, {"深度学习", []string{"深度学习", "deep learning"}},
	{"大模型", []string{"大模型", "llm", "langchain", "eino"}}, {"数据结构", []string{"数据结构"}},
	{"算法", []string{"算法", "algorithm"}}, {"性能优化", []string{"性能优化", "performance tuning"}},
}

func ExtractKeywords(text string, limit int) []string {
	lower := strings.ToLower(text)
	type hit struct {
		name  string
		count int
	}
	hits := make([]hit, 0)
	for _, def := range keywordDictionary {
		count := 0
		for _, alias := range def.Aliases {
			count += countTerm(lower, strings.ToLower(alias))
		}
		if count > 0 {
			hits = append(hits, hit{def.Name, count})
		}
	}
	sort.SliceStable(hits, func(i, j int) bool { return hits[i].count > hits[j].count })
	if limit <= 0 || limit > len(hits) {
		limit = len(hits)
	}
	result := make([]string, limit)
	for i := 0; i < limit; i++ {
		result[i] = hits[i].name
	}
	return result
}

func countTerm(text, term string) int {
	if term == "" {
		return 0
	}
	count, offset := 0, 0
	for {
		idx := strings.Index(text[offset:], term)
		if idx < 0 {
			return count
		}
		start := offset + idx
		end := start + len(term)
		if validBoundary(text, start, end, term) {
			count++
		}
		offset = end
		if offset >= len(text) {
			return count
		}
	}
}

func validBoundary(text string, start, end int, term string) bool {
	first, _ := utf8Rune(term)
	if unicode.Is(unicode.Han, first) || strings.ContainsAny(term, "+/#") || strings.Contains(term, " ") {
		return true
	}
	if start > 0 {
		before, _ := utf8LastRune(text[:start])
		if unicode.IsLetter(before) || unicode.IsDigit(before) {
			return false
		}
	}
	if end < len(text) {
		after, _ := utf8Rune(text[end:])
		if unicode.IsLetter(after) || unicode.IsDigit(after) {
			return false
		}
	}
	return true
}

func utf8Rune(value string) (rune, int) {
	for _, r := range value {
		return r, len(string(r))
	}
	return 0, 0
}

func utf8LastRune(value string) (rune, int) {
	runes := []rune(value)
	if len(runes) == 0 {
		return 0, 0
	}
	r := runes[len(runes)-1]
	return r, len(string(r))
}
