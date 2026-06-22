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
		HTTPClient:          &http.Client{Timeout: 45 * time.Second},
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
