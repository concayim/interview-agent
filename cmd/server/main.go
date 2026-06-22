package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"interview-agent/internal/agent"
	"interview-agent/internal/api"
	"interview-agent/internal/config"
	"interview-agent/internal/interview"
)

func main() {
	port := flag.Int("port", envInt("INTERVIEW_AGENT_PORT", 46831), "local HTTP port")
	flag.Parse()
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	dataDir, err := resolveDataDir()
	if err != nil {
		logger.Error("resolve data directory", "error", err)
		os.Exit(1)
	}
	configStore, err := config.NewStore(dataDir)
	if err != nil {
		logger.Error("load model config", "error", err)
		os.Exit(1)
	}
	evaluator := agent.NewEinoEvaluator(configStore)
	service := interview.NewService(evaluator)
	handler := api.New(logger, dataDir, os.Getenv("INTERVIEW_AGENT_TOKEN"), configStore, service, evaluator).Handler()
	listener, err := net.Listen("tcp4", fmt.Sprintf("127.0.0.1:%d", *port))
	if err != nil {
		logger.Error("listen", "error", err)
		os.Exit(1)
	}
	server := &http.Server{Handler: handler, ReadHeaderTimeout: 5 * time.Second, ReadTimeout: 70 * time.Second, WriteTimeout: 70 * time.Second, IdleTimeout: 90 * time.Second}
	go func() {
		logger.Info("Interview Copilot backend ready", "address", listener.Addr(), "dataDir", dataDir)
		if err := server.Serve(listener); err != nil && err != http.ErrServerClosed {
			logger.Error("serve", "error", err)
			os.Exit(1)
		}
	}()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Error("shutdown", "error", err)
	}
}

func resolveDataDir() (string, error) {
	if value := os.Getenv("INTERVIEW_AGENT_DATA_DIR"); value != "" {
		if err := os.MkdirAll(value, 0o700); err != nil {
			return "", err
		}
		return value, nil
	}
	root, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	dir := filepath.Join(root, "interview-copilot")
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return "", err
	}
	return dir, nil
}

func envInt(name string, fallback int) int {
	var value int
	if _, err := fmt.Sscanf(os.Getenv(name), "%d", &value); err == nil && value > 0 {
		return value
	}
	return fallback
}
