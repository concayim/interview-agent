package questions

import (
	"strings"
	"testing"
)

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

func TestGenericQuestionsRespectDifficulty(t *testing.T) {
	selected := Generic("rust", "Rust 工程师", []string{"Rust", "所有权", "并发安全"}, "hard", 2)
	if len(selected) != 2 {
		t.Fatalf("expected 2 generic questions, got %d", len(selected))
	}
	for _, question := range selected {
		if question.Language != "rust" || question.Difficulty != "hard" || question.StandardAnswer == "" {
			t.Fatalf("unexpected generic question %#v", question)
		}
	}
}

func TestGenericTopicQuestionsUseConcreteKnowledge(t *testing.T) {
	cases := []struct {
		language string
		name     string
		topics   []string
	}{
		{"assembly", "Assembly 工程师", []string{"指令集", "寄存器", "内存模型"}},
		{"c", "C 工程师", []string{"指针", "内存管理", "编译链接"}},
		{"csharp", "C# 工程师", []string{"C#", ".NET", "异步编程"}},
		{"dart", "Dart 工程师", []string{"Dart", "异步编程", "Flutter"}},
		{"javascript", "JavaScript 工程师", []string{"JavaScript", "事件循环", "异步编程"}},
		{"julia", "Julia 工程师", []string{"Julia", "多重派发", "科学计算"}},
		{"kotlin", "Kotlin 工程师", []string{"Kotlin", "协程", "空安全"}},
		{"lua", "Lua 工程师", []string{"Lua", "协程", "元表"}},
		{"perl", "Perl 工程师", []string{"Perl", "正则表达式", "文本处理"}},
		{"php", "PHP 工程师", []string{"PHP", "Web 开发", "类型系统"}},
		{"powershell", "PowerShell 工程师", []string{"PowerShell", "管道", "自动化"}},
		{"r", "R 工程师", []string{"R", "数据分析", "统计计算"}},
		{"ruby", "Ruby 工程师", []string{"Ruby", "对象模型", "元编程"}},
		{"rust", "Rust 工程师", []string{"所有权", "生命周期", "并发安全"}},
		{"scala", "Scala 工程师", []string{"Scala", "函数式编程", "类型系统"}},
		{"sql", "SQL 工程师", []string{"SQL", "查询优化", "事务"}},
		{"swift", "Swift 工程师", []string{"Swift", "值语义", "并发"}},
		{"typescript", "TypeScript 工程师", []string{"TypeScript", "类型系统", "泛型"}},
		{"vbscript", "VBScript 工程师", []string{"VBScript", "脚本自动化", "COM"}},
		{"verilog", "Verilog 工程师", []string{"RTL", "时序逻辑", "仿真"}},
		{"zig", "Zig 工程师", []string{"内存管理", "编译期", "系统编程"}},
	}
	placeholders := map[string]bool{"概念定义": true, "工作机制": true, "适用场景": true, "边界条件": true, "语言定位": true, "核心特性": true, "模块边界": true, "性能分析": true, "对比维度": true}
	for _, item := range cases {
		questions := Generic(item.language, item.name, item.topics, "mixed", 6)
		if len(questions) != 6 {
			t.Fatalf("%s: expected six mixed questions, got %d", item.language, len(questions))
		}
		topicQuestion := questions[2]
		if !strings.Contains(topicQuestion.StandardAnswer, item.topics[1]) {
			t.Errorf("%s: standard answer is not topic-specific: %s", item.language, topicQuestion.StandardAnswer)
		}
		for _, question := range questions {
			mentionsTopic := false
			for _, topic := range item.topics {
				mentionsTopic = mentionsTopic || strings.Contains(question.StandardAnswer, topic)
			}
			if !mentionsTopic {
				t.Errorf("%s/%s: standard answer does not mention a language topic: %s", item.language, question.ID, question.StandardAnswer)
			}
			for _, keyPoint := range question.KeyPoints {
				if placeholders[keyPoint] {
					t.Errorf("%s/%s: placeholder key point %q", item.language, question.ID, keyPoint)
				}
			}
		}
	}
	for _, item := range cases {
		for _, question := range Generic(item.language, item.name, item.topics, "mixed", 6) {
			for _, group := range question.KeyPoints {
				if strings.Contains(group, "/") {
					t.Errorf("%s/%s: generated key point %q still uses legacy slash separator", item.language, question.ID, group)
				}
			}
		}
	}
}
