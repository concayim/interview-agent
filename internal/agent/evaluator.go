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

type Evaluator interface {
	Evaluate(context.Context, EvaluationInput) (domain.Evaluation, error)
}

type FollowUpInput struct {
	Question          domain.Question
	CandidateAnswers  []string
	Round             int
	Total             int
	InterviewerName   string
	InterviewerPrompt string
	EvaluationFocus   []string
}

type FollowUpGenerator interface {
	GenerateFollowUp(context.Context, FollowUpInput) (string, error)
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
		HTTPClient:          &http.Client{Timeout: 40 * time.Second},
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

func (e *EinoEvaluator) GenerateFollowUp(ctx context.Context, input FollowUpInput) (string, error) {
	cfg := e.config.Get()
	if !cfg.Ready() {
		return "", ErrNotConfigured
	}
	temperature := float32(0.35)
	maxTokens := 180
	model, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
		APIKey:              cfg.APIKey,
		BaseURL:             strings.TrimRight(cfg.BaseURL, "/"),
		Model:               cfg.Model,
		Temperature:         &temperature,
		MaxCompletionTokens: &maxTokens,
		HTTPClient:          &http.Client{Timeout: 30 * time.Second},
	})
	if err != nil {
		return "", fmt.Errorf("创建 Eino ChatModel 失败: %w", err)
	}
	prompt := fmt.Sprintf(`你是面试官 Skill「%s」。风格指令：%s
这是主问题后的第 %d/%d 轮追问。请只输出一个自然、具体的中文追问，不要输出答案、解释、编号或引号，控制在 80 字以内。
追问应基于候选人刚才的回答继续深入，优先检查未说明的边界、原理、取舍、反例或真实项目经验；不要重复主问题。

主问题：%s
标准答案关键点：%s
评价重点：%s
候选人截至目前的回答：
%s`, input.InterviewerName, input.InterviewerPrompt, input.Round, input.Total, input.Question.Prompt, strings.Join(input.Question.KeyPoints, "、"), strings.Join(input.EvaluationFocus, "、"), strings.Join(input.CandidateAnswers, "\n---\n"))
	response, err := model.Generate(ctx, []*schema.Message{
		{Role: schema.System, Content: "你正在进行连续技术面试。只输出下一句追问，不泄露参考答案。"},
		{Role: schema.User, Content: prompt},
	})
	if err != nil {
		return "", fmt.Errorf("生成追问失败: %w", err)
	}
	question := strings.Trim(strings.TrimSpace(response.Content), "\"'“”")
	if question == "" {
		return "", fmt.Errorf("模型返回了空追问")
	}
	runes := []rune(question)
	if len(runes) > 120 {
		question = string(runes[:120])
	}
	return question, nil
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
