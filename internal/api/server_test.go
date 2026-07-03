package api

import (
	"bytes"
	"io"
	"log/slog"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"interview-agent/internal/agent"
	"interview-agent/internal/config"
	"interview-agent/internal/interview"
	"interview-agent/internal/knowledge"
	"interview-agent/internal/learning"
	"interview-agent/internal/questions"
	"interview-agent/internal/skills"
)

func TestHealthAndTokenProtection(t *testing.T) {
	dir := t.TempDir()
	store, err := config.NewStore(dir)
	if err != nil {
		t.Fatal(err)
	}
	evaluator := agent.NewEinoEvaluator(store)
	catalog, err := skills.Load()
	if err != nil {
		t.Fatal(err)
	}
	knowledgeStore, err := knowledge.NewStore(dir, catalog.DomainSkills())
	if err != nil {
		t.Fatal(err)
	}
	if err := knowledgeStore.SeedQuestions(questions.All()); err != nil {
		t.Fatal(err)
	}
	learningService, err := learning.NewService(dir)
	if err != nil {
		t.Fatal(err)
	}
	service := interview.NewService(evaluator, interview.WithCatalog(catalog), interview.WithKnowledge(knowledgeStore))
	handler := New(slog.New(slog.NewTextHandler(io.Discard, nil)), dir, "secret", store, service, evaluator, catalog, knowledgeStore, learningService).Handler()

	health := httptest.NewRequest(http.MethodGet, "/api/v1/health", nil)
	healthResponse := httptest.NewRecorder()
	handler.ServeHTTP(healthResponse, health)
	if healthResponse.Code != http.StatusOK {
		t.Fatalf("health status: %d", healthResponse.Code)
	}

	unauthorized := httptest.NewRequest(http.MethodGet, "/api/v1/config/model", nil)
	unauthorizedResponse := httptest.NewRecorder()
	handler.ServeHTTP(unauthorizedResponse, unauthorized)
	if unauthorizedResponse.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %d", unauthorizedResponse.Code)
	}

	publicSkills := httptest.NewRequest(http.MethodGet, "/api/v1/skills", nil)
	publicSkillsResponse := httptest.NewRecorder()
	handler.ServeHTTP(publicSkillsResponse, publicSkills)
	if publicSkillsResponse.Code != http.StatusOK {
		t.Fatalf("expected public skills 200, got %d", publicSkillsResponse.Code)
	}

	authorized := httptest.NewRequest(http.MethodGet, "/api/v1/config/model", nil)
	authorized.Header.Set("X-Interview-Agent-Token", "secret")
	authorizedResponse := httptest.NewRecorder()
	handler.ServeHTTP(authorizedResponse, authorized)
	if authorizedResponse.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", authorizedResponse.Code)
	}

	for _, path := range []string{"/api/v1/skills", "/api/v1/knowledge/bases", "/api/v1/learning/resources"} {
		request := httptest.NewRequest(http.MethodGet, path, nil)
		request.Header.Set("X-Interview-Agent-Token", "secret")
		response := httptest.NewRecorder()
		handler.ServeHTTP(response, request)
		if response.Code != http.StatusOK {
			t.Fatalf("%s: expected 200, got %d", path, response.Code)
		}
	}
}

func TestSpeechTranscriptionRequiresSpeechModel(t *testing.T) {
	dir := t.TempDir()
	store, err := config.NewStore(dir)
	if err != nil {
		t.Fatal(err)
	}
	if err := store.Save(config.ModelConfig{APIKey: "sk-test", BaseURL: "https://api.example.com/v1", Model: "chat-model", Enabled: true}, false); err != nil {
		t.Fatal(err)
	}
	handler := newTestHandler(t, dir, "secret", store)

	body, contentType := speechRequestBody(t)
	request := httptest.NewRequest(http.MethodPost, "/api/v1/speech/transcriptions", body)
	request.Header.Set("Content-Type", contentType)
	request.Header.Set("X-Interview-Agent-Token", "secret")
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d: %s", response.Code, response.Body.String())
	}
	if !strings.Contains(response.Body.String(), "Speech Model") {
		t.Fatalf("expected speech model guidance, got %s", response.Body.String())
	}
}

func TestSpeechTranscriptionUsesSpeechModel(t *testing.T) {
	var receivedModel string
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/audio/transcriptions" {
			t.Fatalf("unexpected upstream path: %s", r.URL.Path)
		}
		if got := r.Header.Get("Authorization"); got != "Bearer sk-test" {
			t.Fatalf("unexpected authorization header: %s", got)
		}
		if err := r.ParseMultipartForm(12 << 20); err != nil {
			t.Fatal(err)
		}
		receivedModel = r.FormValue("model")
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"text":"你好，我是候选人"}`))
	}))
	defer upstream.Close()

	dir := t.TempDir()
	store, err := config.NewStore(dir)
	if err != nil {
		t.Fatal(err)
	}
	if err := store.Save(config.ModelConfig{APIKey: "sk-test", BaseURL: upstream.URL + "/v1", Model: "chat-model", SpeechModel: "whisper-1", Enabled: true}, false); err != nil {
		t.Fatal(err)
	}
	handler := newTestHandler(t, dir, "secret", store)

	body, contentType := speechRequestBody(t)
	request := httptest.NewRequest(http.MethodPost, "/api/v1/speech/transcriptions", body)
	request.Header.Set("Content-Type", contentType)
	request.Header.Set("X-Interview-Agent-Token", "secret")
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", response.Code, response.Body.String())
	}
	if receivedModel != "whisper-1" {
		t.Fatalf("expected speech model whisper-1, got %q", receivedModel)
	}
	if !strings.Contains(response.Body.String(), "候选人") {
		t.Fatalf("expected transcript text, got %s", response.Body.String())
	}
}

func TestSpeechTranscriptionStreamForwardsDeltas(t *testing.T) {
	var receivedStream string
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/audio/transcriptions" {
			t.Fatalf("unexpected upstream path: %s", r.URL.Path)
		}
		if err := r.ParseMultipartForm(12 << 20); err != nil {
			t.Fatal(err)
		}
		receivedStream = r.FormValue("stream")
		w.Header().Set("Content-Type", "text/event-stream")
		_, _ = w.Write([]byte("event: transcript.text.delta\ndata: {\"type\":\"transcript.text.delta\",\"delta\":\"你好\"}\n\n"))
		_, _ = w.Write([]byte("event: transcript.text.delta\ndata: {\"type\":\"transcript.text.delta\",\"delta\":\"，世界\"}\n\n"))
		_, _ = w.Write([]byte("event: transcript.text.done\ndata: {\"type\":\"transcript.text.done\",\"text\":\"你好，世界\"}\n\n"))
	}))
	defer upstream.Close()

	dir := t.TempDir()
	store, err := config.NewStore(dir)
	if err != nil {
		t.Fatal(err)
	}
	if err := store.Save(config.ModelConfig{APIKey: "sk-test", BaseURL: upstream.URL + "/v1", Model: "chat-model", SpeechModel: "whisper-1", Enabled: true}, false); err != nil {
		t.Fatal(err)
	}
	handler := newTestHandler(t, dir, "secret", store)

	body, contentType := speechRequestBody(t)
	request := httptest.NewRequest(http.MethodPost, "/api/v1/speech/transcriptions/stream", body)
	request.Header.Set("Content-Type", contentType)
	request.Header.Set("X-Interview-Agent-Token", "secret")
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", response.Code, response.Body.String())
	}
	if receivedStream != "true" {
		t.Fatalf("expected stream=true, got %q", receivedStream)
	}
	bodyText := response.Body.String()
	if !strings.Contains(bodyText, "event: delta") || !strings.Contains(bodyText, "你好") || !strings.Contains(bodyText, "世界") {
		t.Fatalf("expected streamed transcript deltas, got %s", bodyText)
	}
}

func TestSpeechTranscriptionStreamHandlesNestedDeltasAndCRLF(t *testing.T) {
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(12 << 20); err != nil {
			t.Fatal(err)
		}
		w.Header().Set("Content-Type", "text/event-stream")
		_, _ = w.Write([]byte("data: {\"choices\":[{\"delta\":{\"content\":\"hello\"}}]}\r\n\r\n"))
		_, _ = w.Write([]byte("data: {\"choices\":[{\"delta\":{\"content\":\" world\"},\"finish_reason\":\"stop\"}]}\r\n\r\n"))
		_, _ = w.Write([]byte("data: [DONE]\r\n\r\n"))
	}))
	defer upstream.Close()

	dir := t.TempDir()
	store, err := config.NewStore(dir)
	if err != nil {
		t.Fatal(err)
	}
	if err := store.Save(config.ModelConfig{APIKey: "sk-test", BaseURL: upstream.URL + "/v1", Model: "chat-model", SpeechModel: "whisper-1", Enabled: true}, false); err != nil {
		t.Fatal(err)
	}
	handler := newTestHandler(t, dir, "secret", store)

	body, contentType := speechRequestBody(t)
	request := httptest.NewRequest(http.MethodPost, "/api/v1/speech/transcriptions/stream", body)
	request.Header.Set("Content-Type", contentType)
	request.Header.Set("X-Interview-Agent-Token", "secret")
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", response.Code, response.Body.String())
	}
	bodyText := response.Body.String()
	if !strings.Contains(bodyText, "hello") || !strings.Contains(bodyText, "world") {
		t.Fatalf("expected nested streamed deltas, got %s", bodyText)
	}
	if strings.Count(bodyText, "event: done") != 1 {
		t.Fatalf("expected one done event, got %s", bodyText)
	}
}

func TestSpeechTranscriptionStreamFallsBackWhenProviderRejectsStream(t *testing.T) {
	var calls int
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(12 << 20); err != nil {
			t.Fatal(err)
		}
		calls++
		if r.FormValue("stream") == "true" {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(`{"error":{"message":"unsupported parameter: stream"}}`))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"text":"fallback transcript"}`))
	}))
	defer upstream.Close()

	dir := t.TempDir()
	store, err := config.NewStore(dir)
	if err != nil {
		t.Fatal(err)
	}
	if err := store.Save(config.ModelConfig{APIKey: "sk-test", BaseURL: upstream.URL + "/v1", Model: "chat-model", SpeechModel: "whisper-1", Enabled: true}, false); err != nil {
		t.Fatal(err)
	}
	handler := newTestHandler(t, dir, "secret", store)

	body, contentType := speechRequestBody(t)
	request := httptest.NewRequest(http.MethodPost, "/api/v1/speech/transcriptions/stream", body)
	request.Header.Set("Content-Type", contentType)
	request.Header.Set("X-Interview-Agent-Token", "secret")
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", response.Code, response.Body.String())
	}
	if calls != 2 {
		t.Fatalf("expected stream request and fallback request, got %d calls", calls)
	}
	if !strings.Contains(response.Body.String(), "fallback transcript") {
		t.Fatalf("expected fallback transcript, got %s", response.Body.String())
	}
}

func speechRequestBody(t *testing.T) (*bytes.Buffer, string) {
	t.Helper()
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	part, err := writer.CreateFormFile("audio", "answer.webm")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := part.Write([]byte("fake audio")); err != nil {
		t.Fatal(err)
	}
	if err := writer.WriteField("language", "zh-CN"); err != nil {
		t.Fatal(err)
	}
	if err := writer.Close(); err != nil {
		t.Fatal(err)
	}
	return &body, writer.FormDataContentType()
}

func newTestHandler(t *testing.T, dir string, token string, store *config.Store) http.Handler {
	t.Helper()
	evaluator := agent.NewEinoEvaluator(store)
	catalog, err := skills.Load()
	if err != nil {
		t.Fatal(err)
	}
	knowledgeStore, err := knowledge.NewStore(dir, catalog.DomainSkills())
	if err != nil {
		t.Fatal(err)
	}
	if err := knowledgeStore.SeedQuestions(questions.All()); err != nil {
		t.Fatal(err)
	}
	learningService, err := learning.NewService(dir)
	if err != nil {
		t.Fatal(err)
	}
	service := interview.NewService(evaluator, interview.WithCatalog(catalog), interview.WithKnowledge(knowledgeStore))
	return New(slog.New(slog.NewTextHandler(io.Discard, nil)), dir, token, store, service, evaluator, catalog, knowledgeStore, learningService).Handler()
}
