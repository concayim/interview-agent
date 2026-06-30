package agent

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"interview-agent/internal/config"
	"interview-agent/internal/domain"

	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/schema"
)

var ErrNotConfigured = errors.New("大模型尚未配置")

type EvaluationInput struct {
	Question          domain.Question
	CandidateAnswer   string
	InterviewerName   string
	InterviewerPrompt string
	EvaluationFocus   []string
	FeedbackTone      string
}

type IntentInput struct {
	Question          domain.Question
	CandidateMessage  string
	InterviewerName   string
	InterviewerPrompt string
	FeedbackTone      string
}

type IntentResult struct {
	Intent         string
	Accepted       bool
	AssistantReply string
}

type QuestionGenerationInput struct {
	Language          string
	Difficulty        string
	QuestionCount     int
	CandidateKeywords []string
	DomainSkillName   string
	DomainPrompt      string
	Topics            []string
	IncludeFoundation bool
}

type Evaluator interface {
	Evaluate(context.Context, EvaluationInput) (domain.Evaluation, error)
}

type IntentResolver interface {
	ResolveIntent(context.Context, IntentInput) (IntentResult, error)
}

type QuestionGenerator interface {
	GenerateQuestions(context.Context, QuestionGenerationInput) ([]domain.Question, error)
}

type EinoEvaluator struct{ config *config.Store }

func NewEinoEvaluator(store *config.Store) *EinoEvaluator { return &EinoEvaluator{config: store} }

func (e *EinoEvaluator) Evaluate(ctx context.Context, input EvaluationInput) (domain.Evaluation, error) {
	cfg := e.config.Get()
	if !cfg.Ready() {
		return domain.Evaluation{}, ErrNotConfigured
	}
	temperature := float32(0.2)
	maxTokens := 700
	model, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
		APIKey:              cfg.APIKey,
		BaseURL:             strings.TrimRight(cfg.BaseURL, "/"),
		Model:               cfg.Model,
		Temperature:         &temperature,
		MaxCompletionTokens: &maxTokens,
		HTTPClient:          &http.Client{Timeout: 14 * time.Second},
	})
	if err != nil {
		return domain.Evaluation{}, fmt.Errorf("创建 Eino ChatModel 失败: %w", err)
	}
	prompt := fmt.Sprintf(`你是面试官 Skill「%s」。风格指令：%s
评价重点：%s。反馈语气：%s。
根据题目、标准答案和关键点评价候选人的回答。
只返回合法 JSON，不要使用 Markdown。结构必须是：
{"score":0到100的整数,"summary":"两句以内的中文总结","strengths":["具体优点"],"improvements":["可执行的改进建议"]}
不要因为措辞与标准答案不同而扣分，重点判断技术事实、推理和工程意识。回答为空时得 0 分。

题目：%s
标准答案：%s
关键点：%s
候选人回答：%s`, input.InterviewerName, input.InterviewerPrompt, strings.Join(input.EvaluationFocus, "、"), input.FeedbackTone, input.Question.Prompt, input.Question.StandardAnswer, strings.Join(input.Question.KeyPoints, "、"), input.CandidateAnswer)
	response, err := model.Generate(ctx, []*schema.Message{
		{Role: schema.System, Content: "你是 Interview Copilot 的答案评估 Agent，必须输出可解析的 JSON。"},
		{Role: schema.User, Content: prompt},
	})
	if err != nil {
		return domain.Evaluation{}, fmt.Errorf("调用大模型失败: %w", err)
	}
	var result struct {
		Score        int      `json:"score"`
		Summary      string   `json:"summary"`
		Strengths    []string `json:"strengths"`
		Improvements []string `json:"improvements"`
	}
	if err := json.Unmarshal([]byte(stripCodeFence(response.Content)), &result); err != nil {
		return domain.Evaluation{}, fmt.Errorf("解析模型评价失败: %w", err)
	}
	if result.Score < 0 {
		result.Score = 0
	}
	if result.Score > 100 {
		result.Score = 100
	}
	return domain.Evaluation{Score: result.Score, Summary: result.Summary, Strengths: result.Strengths, Improvements: result.Improvements, Source: "llm"}, nil
}

func (e *EinoEvaluator) ResolveIntent(ctx context.Context, input IntentInput) (IntentResult, error) {
	cfg := e.config.Get()
	if !cfg.Ready() {
		return IntentResult{}, ErrNotConfigured
	}
	temperature := float32(0.1)
	maxTokens := 360
	model, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
		APIKey:              cfg.APIKey,
		BaseURL:             strings.TrimRight(cfg.BaseURL, "/"),
		Model:               cfg.Model,
		Temperature:         &temperature,
		MaxCompletionTokens: &maxTokens,
		HTTPClient:          &http.Client{Timeout: 8 * time.Second},
	})
	if err != nil {
		return IntentResult{}, fmt.Errorf("创建 Eino ChatModel 失败: %w", err)
	}
	prompt := fmt.Sprintf(`你是面试官 Skill「%s」。风格指令：%s。反馈语气：%s。
你正在进行技术面试，需要判断候选人的这句话在当前题目下的真实意图。
只返回合法 JSON，不要使用 Markdown。结构必须是：
{"intent":"answer|hint|clarify|repeat|skip|off_topic|smalltalk","accepted":true或false,"assistantReply":"中文回复，最多两句"}

判定规则：
- answer：候选人正在尝试回答题目，即使不完整、口语化或包含错误，也应 accepted=true，并留空 assistantReply。
- skip：候选人明确想跳过/下一题，accepted=true。
- hint/clarify/repeat：候选人请求提示、解释题意或重复题目，accepted=false，并给出可帮助继续作答的回复。
- off_topic：候选人输入与当前题无关、抱怨题目不相关、转移话题、谈状态或说不知道怎么关联，accepted=false；要简短理解其处境，再把话题拉回当前题。
- smalltalk：问候、感谢、闲聊或非面试内容，accepted=false；自然回应后拉回当前题。
- 不要把明显偏题内容硬判成 answer。
- 不要泄露标准答案或关键点清单。

当前题目：%s
题目标签：%s
候选人输入：%s`, input.InterviewerName, input.InterviewerPrompt, input.FeedbackTone, input.Question.Prompt, strings.Join(input.Question.Tags, "、"), input.CandidateMessage)
	response, err := model.Generate(ctx, []*schema.Message{
		{Role: schema.System, Content: "你是 Interview Copilot 的意图识别 Agent，必须输出可解析的 JSON。"},
		{Role: schema.User, Content: prompt},
	})
	if err != nil {
		return IntentResult{}, fmt.Errorf("调用大模型失败: %w", err)
	}
	var result IntentResult
	if err := json.Unmarshal([]byte(stripCodeFence(response.Content)), &result); err != nil {
		return IntentResult{}, fmt.Errorf("解析模型意图失败: %w", err)
	}
	result.Intent = strings.TrimSpace(strings.ToLower(result.Intent))
	switch result.Intent {
	case "answer", "skip":
		result.Accepted = true
		result.AssistantReply = ""
	case "hint", "clarify", "repeat", "off_topic", "smalltalk":
		result.Accepted = false
	default:
		result.Intent = "answer"
		result.Accepted = true
		result.AssistantReply = ""
	}
	if !result.Accepted && strings.TrimSpace(result.AssistantReply) == "" {
		result.AssistantReply = "我先把你拉回当前题：可以先讲一个粗略结论，再补充原因和边界。"
	}
	return result, nil
}

func (e *EinoEvaluator) GenerateQuestions(ctx context.Context, input QuestionGenerationInput) ([]domain.Question, error) {
	cfg := e.config.Get()
	if !cfg.Ready() {
		return nil, ErrNotConfigured
	}
	if input.QuestionCount < 1 {
		input.QuestionCount = 5
	}
	if input.QuestionCount > 10 {
		input.QuestionCount = 10
	}
	temperature := float32(0.55)
	maxTokens := 1800
	model, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
		APIKey:              cfg.APIKey,
		BaseURL:             strings.TrimRight(cfg.BaseURL, "/"),
		Model:               cfg.Model,
		Temperature:         &temperature,
		MaxCompletionTokens: &maxTokens,
		HTTPClient:          &http.Client{Timeout: 14 * time.Second},
	})
	if err != nil {
		return nil, fmt.Errorf("创建 Eino ChatModel 失败: %w", err)
	}
	prompt := fmt.Sprintf(`你是 Interview Copilot 的面试出题 Agent。
请为「%s」生成 %d 道技术面试题。只返回合法 JSON，不要使用 Markdown。结构必须是：
{"questions":[{"difficulty":"easy|medium|hard","prompt":"题目","standardAnswer":"复盘用标准答案","keyPoints":["关键点1","关键点2"],"tags":["标签1","标签2"]}]}

要求：
- 题目语言/方向：%s。
- 难度：%s；mixed 表示基础、进阶、挑战均衡。
- 领域 Skill 指令：%s
- 主题：%s
- 候选人关键词：%s
- 是否混入计算机基础公共题：%t
- 每题必须可独立作答，避免重复，避免泄露“这是模型生成”的措辞。
- standardAnswer 用于面试结束后的复盘，可以更完整；prompt 不要包含答案。
- keyPoints 每题 4 到 7 个，支持中文或中英混写。
- tags 每题 2 到 4 个。`, input.DomainSkillName, input.QuestionCount, input.Language, input.Difficulty, input.DomainPrompt, strings.Join(input.Topics, "、"), strings.Join(input.CandidateKeywords, "、"), input.IncludeFoundation)
	response, err := model.Generate(ctx, []*schema.Message{
		{Role: schema.System, Content: "你是 Interview Copilot 的出题 Agent，必须输出可解析 JSON。"},
		{Role: schema.User, Content: prompt},
	})
	if err != nil {
		return nil, fmt.Errorf("调用大模型失败: %w", err)
	}
	var result struct {
		Questions []struct {
			Difficulty     string   `json:"difficulty"`
			Prompt         string   `json:"prompt"`
			StandardAnswer string   `json:"standardAnswer"`
			KeyPoints      []string `json:"keyPoints"`
			Tags           []string `json:"tags"`
		} `json:"questions"`
	}
	if err := json.Unmarshal([]byte(stripCodeFence(response.Content)), &result); err != nil {
		return nil, fmt.Errorf("解析模型题目失败: %w", err)
	}
	questions := make([]domain.Question, 0, len(result.Questions))
	for index, item := range result.Questions {
		prompt := strings.TrimSpace(item.Prompt)
		standardAnswer := strings.TrimSpace(item.StandardAnswer)
		if prompt == "" || standardAnswer == "" || len(item.KeyPoints) == 0 {
			continue
		}
		difficulty := strings.TrimSpace(strings.ToLower(item.Difficulty))
		if difficulty != "easy" && difficulty != "medium" && difficulty != "hard" {
			difficulty = input.Difficulty
			if difficulty == "" || difficulty == "mixed" {
				difficulty = "medium"
			}
		}
		questions = append(questions, domain.Question{
			ID:             fmt.Sprintf("generated-%d-%d", time.Now().UnixNano(), index),
			Language:       input.Language,
			Difficulty:     difficulty,
			Prompt:         prompt,
			StandardAnswer: standardAnswer,
			KeyPoints:      cleanStrings(item.KeyPoints, 7),
			Tags:           cleanStrings(item.Tags, 4),
		})
	}
	if len(questions) == 0 {
		return nil, fmt.Errorf("模型没有生成可用题目")
	}
	if len(questions) > input.QuestionCount {
		questions = questions[:input.QuestionCount]
	}
	return questions, nil
}

func (e *EinoEvaluator) Test(ctx context.Context) error {
	cfg := e.config.Get()
	if !cfg.Ready() {
		return ErrNotConfigured
	}
	temperature := float32(0)
	maxTokens := 32
	model, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
		APIKey: cfg.APIKey, BaseURL: strings.TrimRight(cfg.BaseURL, "/"), Model: cfg.Model,
		Temperature: &temperature, MaxCompletionTokens: &maxTokens,
		HTTPClient: &http.Client{Timeout: 20 * time.Second},
	})
	if err != nil {
		return err
	}
	_, err = model.Generate(ctx, []*schema.Message{{Role: schema.User, Content: "请只回复 OK"}})
	return err
}

func stripCodeFence(value string) string {
	value = strings.TrimSpace(value)
	if strings.HasPrefix(value, "```") {
		value = strings.TrimPrefix(value, "```json")
		value = strings.TrimPrefix(value, "```")
		value = strings.TrimSuffix(value, "```")
	}
	return strings.TrimSpace(value)
}

func cleanStrings(values []string, limit int) []string {
	result := make([]string, 0, min(len(values), limit))
	seen := map[string]bool{}
	for _, value := range values {
		value = strings.TrimSpace(value)
		key := strings.ToLower(value)
		if value == "" || seen[key] {
			continue
		}
		seen[key] = true
		result = append(result, value)
		if len(result) >= limit {
			break
		}
	}
	return result
}
