package api

import (
	"bufio"
	"bytes"
	"context"
	"crypto/subtle"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"interview-agent/internal/agent"
	"interview-agent/internal/config"
	"interview-agent/internal/interview"
	"interview-agent/internal/knowledge"
	"interview-agent/internal/learning"
	resumeparser "interview-agent/internal/parser"
	"interview-agent/internal/skills"
)

const maxUploadSize = 10 << 20
const maxSpeechUploadSize = 12 << 20

type Server struct {
	logger     *slog.Logger
	dataDir    string
	token      string
	config     *config.Store
	interviews *interview.Service
	evaluator  *agent.EinoEvaluator
	skills     *skills.Catalog
	knowledge  *knowledge.Store
	learning   *learning.Service
}

func New(logger *slog.Logger, dataDir, token string, configStore *config.Store, service *interview.Service, evaluator *agent.EinoEvaluator, skillCatalog *skills.Catalog, knowledgeStore *knowledge.Store, learningService *learning.Service) *Server {
	return &Server{logger: logger, dataDir: dataDir, token: token, config: configStore, interviews: service, evaluator: evaluator, skills: skillCatalog, knowledge: knowledgeStore, learning: learningService}
}

func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/v1/health", s.health)
	mux.HandleFunc("POST /api/v1/resumes", s.uploadResume)
	mux.HandleFunc("POST /api/v1/interviews", s.startInterview)
	mux.HandleFunc("GET /api/v1/interviews/{id}", s.getInterview)
	mux.HandleFunc("POST /api/v1/interviews/{id}/answers", s.answer)
	mux.HandleFunc("GET /api/v1/interviews/{id}/report", s.report)
	mux.HandleFunc("GET /api/v1/config/model", s.getModelConfig)
	mux.HandleFunc("PUT /api/v1/config/model", s.putModelConfig)
	mux.HandleFunc("POST /api/v1/config/model/test", s.testModelConfig)
	mux.HandleFunc("POST /api/v1/speech/transcriptions", s.transcribeSpeech)
	mux.HandleFunc("POST /api/v1/speech/transcriptions/stream", s.streamTranscribeSpeech)
	mux.HandleFunc("GET /api/v1/skills", s.listSkills)
	mux.HandleFunc("GET /api/v1/knowledge/bases", s.listKnowledgeBases)
	mux.HandleFunc("GET /api/v1/knowledge/bases/{id}/qa", s.searchKnowledge)
	mux.HandleFunc("POST /api/v1/knowledge/bases/{id}/qa", s.putKnowledge)
	mux.HandleFunc("GET /api/v1/learning/resources", s.listLearningResources)
	mux.HandleFunc("POST /api/v1/learning/refresh", s.refreshLearningResources)
	mux.HandleFunc("PUT /api/v1/learning/resources/{id}/selection", s.selectLearningResource)
	return s.recover(s.logging(s.cors(s.authorize(mux))))
}

func (s *Server) health(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]any{"status": "ok", "time": time.Now().UTC(), "modelConfigured": s.config.Get().Ready()})
}

func (s *Server) uploadResume(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize+(1<<20))
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		writeError(w, http.StatusBadRequest, "附件不能超过 10 MB")
		return
	}
	file, header, err := r.FormFile("file")
	if err != nil {
		writeError(w, http.StatusBadRequest, "请选择要上传的简历")
		return
	}
	defer file.Close()
	name := filepath.Base(header.Filename)
	ext := strings.ToLower(filepath.Ext(name))
	if ext != ".pdf" && ext != ".docx" {
		writeError(w, http.StatusUnsupportedMediaType, "仅支持 PDF 和 DOCX 文件")
		return
	}
	tmp, err := os.CreateTemp(s.dataDir, "resume-*"+ext)
	if err != nil {
		s.internalError(w, err)
		return
	}
	tmpPath := tmp.Name()
	defer os.Remove(tmpPath)
	defer tmp.Close()
	if _, err := io.Copy(tmp, io.LimitReader(file, maxUploadSize+1)); err != nil {
		s.internalError(w, err)
		return
	}
	if info, err := tmp.Stat(); err != nil || info.Size() > maxUploadSize {
		writeError(w, http.StatusBadRequest, "附件不能超过 10 MB")
		return
	}
	if err := tmp.Close(); err != nil {
		s.internalError(w, err)
		return
	}
	resume, err := resumeparser.Parse(tmpPath, name, header.Header.Get("Content-Type"))
	if err != nil {
		writeError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	resume = s.interviews.AddResume(resume)
	writeJSON(w, http.StatusCreated, resume)
}

func (s *Server) startInterview(w http.ResponseWriter, r *http.Request) {
	var input interview.StartInput
	if err := readJSON(r, &input); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	view, err := s.interviews.Start(input)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, view)
}

func (s *Server) getInterview(w http.ResponseWriter, r *http.Request) {
	view, err := s.interviews.Get(r.PathValue("id"))
	if errors.Is(err, interview.ErrNotFound) {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		s.internalError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, view)
}

func (s *Server) answer(w http.ResponseWriter, r *http.Request) {
	var input interview.AnswerInput
	if err := readJSON(r, &input); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	ctx, cancel := contextWithTimeout(r, 20*time.Second)
	defer cancel()
	result, err := s.interviews.Answer(ctx, r.PathValue("id"), input)
	if errors.Is(err, interview.ErrNotFound) {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}
	if errors.Is(err, interview.ErrCompleted) {
		writeError(w, http.StatusConflict, err.Error())
		return
	}
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, result)
}

func (s *Server) report(w http.ResponseWriter, r *http.Request) {
	report, err := s.interviews.Report(r.PathValue("id"))
	if errors.Is(err, interview.ErrNotFound) {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		writeError(w, http.StatusConflict, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, report)
}

func (s *Server) getModelConfig(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, s.config.Public())
}

func (s *Server) putModelConfig(w http.ResponseWriter, r *http.Request) {
	var input struct {
		APIKey      string `json:"apiKey"`
		BaseURL     string `json:"baseUrl"`
		Model       string `json:"model"`
		SpeechModel string `json:"speechModel"`
		Enabled     bool   `json:"enabled"`
		ClearAPIKey bool   `json:"clearApiKey"`
	}
	if err := readJSON(r, &input); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	input.APIKey = strings.TrimSpace(input.APIKey)
	input.BaseURL = strings.TrimSpace(input.BaseURL)
	input.Model = strings.TrimSpace(input.Model)
	input.SpeechModel = strings.TrimSpace(input.SpeechModel)
	if input.BaseURL != "" {
		parsed, err := url.Parse(input.BaseURL)
		if err != nil || (parsed.Scheme != "http" && parsed.Scheme != "https") || parsed.Host == "" {
			writeError(w, http.StatusBadRequest, "Base URL 必须是有效的 http(s) 地址")
			return
		}
	}
	current := s.config.Get()
	apiKeyReady := input.APIKey != "" || (current.APIKey != "" && !input.ClearAPIKey)
	if input.Enabled && (!apiKeyReady || input.Model == "") {
		writeError(w, http.StatusBadRequest, "启用模型前请填写 API Key 和 Model")
		return
	}
	next := config.ModelConfig{APIKey: input.APIKey, BaseURL: strings.TrimRight(input.BaseURL, "/"), Model: input.Model, SpeechModel: input.SpeechModel, Enabled: input.Enabled}
	if err := s.config.Save(next, !input.ClearAPIKey); err != nil {
		s.internalError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, s.config.Public())
}

func (s *Server) testModelConfig(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := contextWithTimeout(r, 25*time.Second)
	defer cancel()
	if err := s.evaluator.Test(ctx); err != nil {
		writeError(w, http.StatusBadGateway, "模型连接失败："+err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"ok": true, "message": "模型连接成功"})
}

func (s *Server) transcribeSpeech(w http.ResponseWriter, r *http.Request) {
	speechPayload, apiErr := s.parseSpeechTranscriptionPayload(w, r)
	if apiErr != nil {
		writeError(w, apiErr.status, apiErr.message)
		return
	}
	text, err := transcribeSpeechWithoutStream(speechPayload)
	if err != nil {
		writeError(w, http.StatusBadGateway, "语音转写失败："+err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"text": text})
}

func (s *Server) streamTranscribeSpeech(w http.ResponseWriter, r *http.Request) {
	speechPayload, apiErr := s.parseSpeechTranscriptionPayload(w, r)
	if apiErr != nil {
		writeError(w, apiErr.status, apiErr.message)
		return
	}
	upstream, apiErr := newSpeechTranscriptionRequest(speechPayload, true)
	if apiErr != nil {
		writeError(w, apiErr.status, apiErr.message)
		return
	}
	defer upstream.cancel()

	response, err := (&http.Client{Timeout: 30 * time.Second}).Do(upstream.request)
	if err != nil {
		writeError(w, http.StatusBadGateway, "语音转写失败："+err.Error())
		return
	}
	defer response.Body.Close()
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		errorPayload, err := io.ReadAll(io.LimitReader(response.Body, 1<<20))
		if err != nil {
			s.internalError(w, err)
			return
		}
		if shouldRetrySpeechWithoutStream(response.StatusCode, errorPayload) {
			text, retryErr := transcribeSpeechWithoutStream(speechPayload)
			if retryErr == nil {
				writeSpeechTextAsStream(w, text)
				return
			}
		}
		writeError(w, http.StatusBadGateway, "语音转写失败："+extractProviderError(errorPayload))
		return
	}
	flusher, ok := w.(http.Flusher)
	if !ok {
		writeError(w, http.StatusInternalServerError, "当前环境不支持流式输出")
		return
	}

	w.Header().Set("Content-Type", "text/event-stream; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("X-Accel-Buffering", "no")
	if strings.Contains(strings.ToLower(response.Header.Get("Content-Type")), "text/event-stream") {
		forwardSpeechEventStream(w, flusher, response.Body)
		return
	}
	payload, err := io.ReadAll(io.LimitReader(response.Body, 1<<20))
	if err != nil {
		writeSpeechStreamEvent(w, flusher, "error", map[string]string{"error": "读取语音转写响应失败：" + err.Error()})
		return
	}
	text := extractTranscriptText(payload)
	if strings.TrimSpace(text) == "" {
		writeSpeechStreamEvent(w, flusher, "error", map[string]string{"error": "语音转写响应无法解析"})
		return
	}
	writeSpeechStreamEvent(w, flusher, "delta", map[string]string{"text": strings.TrimSpace(text)})
	writeSpeechStreamEvent(w, flusher, "done", map[string]string{"text": strings.TrimSpace(text)})
}

type apiError struct {
	status  int
	message string
}

type speechTranscriptionRequest struct {
	request *http.Request
	cancel  context.CancelFunc
}

type speechTranscriptionPayload struct {
	ctx      context.Context
	apiKey   string
	baseURL  string
	model    string
	language string
	fileName string
	audio    []byte
}

func (s *Server) parseSpeechTranscriptionPayload(w http.ResponseWriter, r *http.Request) (*speechTranscriptionPayload, *apiError) {
	cfg := s.config.Get()
	if !cfg.Ready() {
		return nil, &apiError{status: http.StatusBadRequest, message: "请先配置并启用支持音频转写的模型"}
	}
	speechModel := strings.TrimSpace(cfg.SpeechModel)
	if speechModel == "" {
		return nil, &apiError{status: http.StatusBadRequest, message: "请在模型设置里填写 Speech Model，例如 whisper-1 或服务商提供的音频转写模型"}
	}
	r.Body = http.MaxBytesReader(w, r.Body, maxSpeechUploadSize+(1<<20))
	if err := r.ParseMultipartForm(maxSpeechUploadSize); err != nil {
		return nil, &apiError{status: http.StatusBadRequest, message: "语音片段不能超过 12 MB"}
	}
	file, header, err := r.FormFile("audio")
	if err != nil {
		return nil, &apiError{status: http.StatusBadRequest, message: "请提供语音片段"}
	}
	defer file.Close()

	var audio bytes.Buffer
	if _, err := io.Copy(&audio, io.LimitReader(file, maxSpeechUploadSize+1)); err != nil {
		return nil, &apiError{status: http.StatusInternalServerError, message: err.Error()}
	}
	if audio.Len() > maxSpeechUploadSize {
		return nil, &apiError{status: http.StatusBadRequest, message: "语音片段不能超过 12 MB"}
	}
	baseURL := strings.TrimRight(cfg.BaseURL, "/")
	if baseURL == "" {
		baseURL = "https://api.openai.com/v1"
	}
	return &speechTranscriptionPayload{
		ctx:      r.Context(),
		apiKey:   cfg.APIKey,
		baseURL:  baseURL,
		model:    speechModel,
		language: strings.TrimSpace(r.FormValue("language")),
		fileName: filepath.Base(header.Filename),
		audio:    audio.Bytes(),
	}, nil
}

func newSpeechTranscriptionRequest(payload *speechTranscriptionPayload, stream bool) (*speechTranscriptionRequest, *apiError) {
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	part, err := writer.CreateFormFile("file", payload.fileName)
	if err != nil {
		return nil, &apiError{status: http.StatusInternalServerError, message: err.Error()}
	}
	if _, err := io.Copy(part, bytes.NewReader(payload.audio)); err != nil {
		return nil, &apiError{status: http.StatusInternalServerError, message: err.Error()}
	}
	_ = writer.WriteField("model", payload.model)
	if stream {
		_ = writer.WriteField("stream", "true")
	}
	if payload.language != "" {
		_ = writer.WriteField("language", payload.language)
	}
	if err := writer.Close(); err != nil {
		return nil, &apiError{status: http.StatusInternalServerError, message: err.Error()}
	}

	ctx, cancel := context.WithTimeout(payload.ctx, 30*time.Second)
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, payload.baseURL+"/audio/transcriptions", &body)
	if err != nil {
		cancel()
		return nil, &apiError{status: http.StatusInternalServerError, message: err.Error()}
	}
	request.Header.Set("Authorization", "Bearer "+payload.apiKey)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	return &speechTranscriptionRequest{request: request, cancel: cancel}, nil
}

func transcribeSpeechWithoutStream(payload *speechTranscriptionPayload) (string, error) {
	upstream, apiErr := newSpeechTranscriptionRequest(payload, false)
	if apiErr != nil {
		return "", errors.New(apiErr.message)
	}
	defer upstream.cancel()
	response, err := (&http.Client{Timeout: 30 * time.Second}).Do(upstream.request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(io.LimitReader(response.Body, 1<<20))
	if err != nil {
		return "", err
	}
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return "", errors.New(extractProviderError(body))
	}
	text := extractTranscriptText(body)
	if strings.TrimSpace(text) == "" {
		return "", errors.New("语音转写响应无法解析")
	}
	return strings.TrimSpace(text), nil
}

func shouldRetrySpeechWithoutStream(status int, payload []byte) bool {
	if status != http.StatusBadRequest && status != http.StatusNotFound && status != http.StatusUnprocessableEntity {
		return false
	}
	message := strings.ToLower(extractProviderError(payload))
	return strings.Contains(message, "stream") ||
		strings.Contains(message, "unknown parameter") ||
		strings.Contains(message, "unsupported parameter") ||
		strings.Contains(message, "invalid parameter") ||
		strings.Contains(message, "unrecognized") ||
		strings.Contains(message, "not supported") ||
		strings.Contains(message, "不支持")
}

func writeSpeechTextAsStream(w http.ResponseWriter, text string) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		writeJSON(w, http.StatusOK, map[string]string{"text": text})
		return
	}
	w.Header().Set("Content-Type", "text/event-stream; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("X-Accel-Buffering", "no")
	writeSpeechStreamEvent(w, flusher, "delta", map[string]string{"text": text})
	writeSpeechStreamEvent(w, flusher, "done", map[string]string{"text": text})
}

func forwardSpeechEventStream(w http.ResponseWriter, flusher http.Flusher, body io.Reader) {
	scanner := bufio.NewScanner(body)
	scanner.Buffer(make([]byte, 1024), 1024*1024)
	eventName := ""
	dataLines := []string{}
	sentDelta := false
	sentDone := false
	flushEvent := func() {
		if len(dataLines) == 0 {
			eventName = ""
			return
		}
		data := strings.TrimSpace(strings.Join(dataLines, "\n"))
		event := strings.ToLower(strings.TrimSpace(eventName))
		eventName = ""
		dataLines = nil
		if data == "" {
			return
		}
		if data == "[DONE]" {
			if !sentDone {
				sentDone = true
				writeSpeechStreamEvent(w, flusher, "done", map[string]string{})
			}
			return
		}
		text, done := extractTranscriptDelta([]byte(data))
		if text != "" && (!done || !sentDelta) {
			sentDelta = true
			writeSpeechStreamEvent(w, flusher, "delta", map[string]string{"text": text})
		}
		if done || strings.Contains(event, "done") || strings.Contains(event, "completed") {
			if !sentDone {
				sentDone = true
				writeSpeechStreamEvent(w, flusher, "done", map[string]string{"text": text})
			}
		}
	}
	for scanner.Scan() {
		line := strings.TrimRight(scanner.Text(), "\r")
		if line == "" {
			flushEvent()
			continue
		}
		if strings.HasPrefix(line, "event:") {
			eventName = strings.TrimSpace(strings.TrimPrefix(line, "event:"))
			continue
		}
		if strings.HasPrefix(line, "data:") {
			dataLines = append(dataLines, strings.TrimSpace(strings.TrimPrefix(line, "data:")))
		}
	}
	flushEvent()
	if err := scanner.Err(); err != nil {
		writeSpeechStreamEvent(w, flusher, "error", map[string]string{"error": "读取语音转写流失败：" + err.Error()})
		return
	}
	if !sentDone {
		writeSpeechStreamEvent(w, flusher, "done", map[string]string{})
	}
}

func writeSpeechStreamEvent(w http.ResponseWriter, flusher http.Flusher, event string, payload any) {
	data, err := json.Marshal(payload)
	if err != nil {
		data = []byte(`{}`)
	}
	_, _ = fmt.Fprintf(w, "event: %s\ndata: %s\n\n", event, data)
	flusher.Flush()
}

func extractTranscriptText(payload []byte) string {
	var result struct {
		Text string `json:"text"`
	}
	if err := json.Unmarshal(payload, &result); err == nil {
		return strings.TrimSpace(result.Text)
	}
	return strings.TrimSpace(string(payload))
}

func extractTranscriptDelta(payload []byte) (string, bool) {
	var raw any
	if err := json.Unmarshal(payload, &raw); err != nil {
		return strings.TrimSpace(string(payload)), false
	}
	done := false
	if typed, ok := raw.(map[string]any); ok {
		if value, ok := typed["type"].(string); ok {
			lower := strings.ToLower(value)
			done = strings.Contains(lower, "done") || strings.Contains(lower, "completed") || strings.Contains(lower, "final")
		}
		if value, ok := typed["finish_reason"].(string); ok && strings.TrimSpace(value) != "" {
			done = true
		}
		for _, key := range []string{"done", "completed", "is_final", "final"} {
			if value, ok := typed[key].(bool); ok && value {
				done = true
			}
		}
	}
	return strings.TrimSpace(findTranscriptString(raw)), done
}

func findTranscriptString(value any) string {
	switch typed := value.(type) {
	case string:
		return typed
	case map[string]any:
		for _, key := range []string{"delta", "text", "transcript", "content", "partial", "output_text"} {
			if text, ok := typed[key].(string); ok && strings.TrimSpace(text) != "" {
				return text
			}
		}
		for _, key := range []string{"delta", "text", "transcript", "content", "message", "choice", "choices", "data", "result", "output", "alternatives"} {
			if text := findTranscriptString(typed[key]); text != "" {
				return text
			}
		}
	case []any:
		for _, item := range typed {
			if text := findTranscriptString(item); text != "" {
				return text
			}
		}
	}
	return ""
}

func (s *Server) listSkills(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, s.skills.Public())
}

func (s *Server) listKnowledgeBases(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]any{"bases": s.knowledge.Bases()})
}

func (s *Server) searchKnowledge(w http.ResponseWriter, r *http.Request) {
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	items, err := s.knowledge.Search(r.PathValue("id"), r.URL.Query().Get("query"), limit)
	if err != nil {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"items": items})
}

func (s *Server) putKnowledge(w http.ResponseWriter, r *http.Request) {
	var input knowledge.PutInput
	if err := readJSON(r, &input); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	item, err := s.knowledge.Put(r.PathValue("id"), input)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, item)
}

func (s *Server) listLearningResources(w http.ResponseWriter, r *http.Request) {
	selectedOnly, _ := strconv.ParseBool(r.URL.Query().Get("selectedOnly"))
	writeJSON(w, http.StatusOK, s.learning.List(r.URL.Query().Get("domainSkillId"), r.URL.Query().Get("kind"), selectedOnly))
}

func (s *Server) refreshLearningResources(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := contextWithTimeout(r, 20*time.Second)
	defer cancel()
	writeJSON(w, http.StatusOK, s.learning.Refresh(ctx))
}

func (s *Server) selectLearningResource(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Selected bool `json:"selected"`
	}
	if err := readJSON(r, &input); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := s.learning.SetSelected(r.PathValue("id"), input.Selected); err != nil {
		s.internalError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"id": r.PathValue("id"), "selected": input.Selected})
}

func (s *Server) authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions || isPublicPath(r.URL.Path) || s.token == "" {
			next.ServeHTTP(w, r)
			return
		}
		got := r.Header.Get("X-Interview-Agent-Token")
		if len(got) != len(s.token) || subtle.ConstantTimeCompare([]byte(got), []byte(s.token)) != 1 {
			writeError(w, http.StatusUnauthorized, "无效的本地访问令牌")
			return
		}
		next.ServeHTTP(w, r)
	})
}

func isPublicPath(path string) bool {
	return path == "/api/v1/health" || path == "/api/v1/skills"
}

func (s *Server) cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Interview-Agent-Token")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (s *Server) logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		s.logger.Debug("http request", "method", r.Method, "path", r.URL.Path, "duration", time.Since(start))
	})
}

func (s *Server) recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if value := recover(); value != nil {
				s.logger.Error("panic recovered", "error", value)
				writeError(w, http.StatusInternalServerError, "服务暂时不可用")
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func (s *Server) internalError(w http.ResponseWriter, err error) {
	s.logger.Error("internal error", "error", err)
	writeError(w, http.StatusInternalServerError, "服务内部错误")
}

func readJSON(r *http.Request, target any) error {
	decoder := json.NewDecoder(io.LimitReader(r.Body, 1<<20))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(target); err != nil {
		return fmt.Errorf("请求格式错误: %w", err)
	}
	return nil
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}

func extractProviderError(payload []byte) string {
	var parsed struct {
		Error   any    `json:"error"`
		Message string `json:"message"`
	}
	if err := json.Unmarshal(payload, &parsed); err == nil {
		switch value := parsed.Error.(type) {
		case string:
			if strings.TrimSpace(value) != "" {
				return value
			}
		case map[string]any:
			if message, ok := value["message"].(string); ok && strings.TrimSpace(message) != "" {
				return message
			}
		}
		if strings.TrimSpace(parsed.Message) != "" {
			return parsed.Message
		}
	}
	text := strings.TrimSpace(string(payload))
	if len([]rune(text)) > 300 {
		return string([]rune(text)[:300])
	}
	if text == "" {
		return "服务商没有返回错误详情"
	}
	return text
}

func contextWithTimeout(r *http.Request, duration time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(r.Context(), duration)
}
