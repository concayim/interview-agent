package api

import (
	"context"
	"crypto/subtle"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
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
	next := config.ModelConfig{APIKey: input.APIKey, BaseURL: strings.TrimRight(input.BaseURL, "/"), Model: input.Model, Enabled: input.Enabled}
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

func contextWithTimeout(r *http.Request, duration time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(r.Context(), duration)
}
