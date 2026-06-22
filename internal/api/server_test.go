package api

import (
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
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
