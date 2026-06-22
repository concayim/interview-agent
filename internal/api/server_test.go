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
)

func TestHealthAndTokenProtection(t *testing.T) {
	dir := t.TempDir()
	store, err := config.NewStore(dir)
	if err != nil {
		t.Fatal(err)
	}
	evaluator := agent.NewEinoEvaluator(store)
	service := interview.NewService(evaluator)
	handler := New(slog.New(slog.NewTextHandler(io.Discard, nil)), dir, "secret", store, service, evaluator).Handler()

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
}
