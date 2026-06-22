package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"sync"
)

type ModelConfig struct {
	APIKey  string `json:"apiKey"`
	BaseURL string `json:"baseUrl"`
	Model   string `json:"model"`
	Enabled bool   `json:"enabled"`
}

func (c ModelConfig) Ready() bool {
	return c.Enabled && c.APIKey != "" && c.Model != ""
}

type PublicModelConfig struct {
	BaseURL   string `json:"baseUrl"`
	Model     string `json:"model"`
	Enabled   bool   `json:"enabled"`
	HasAPIKey bool   `json:"hasApiKey"`
}

type Store struct {
	mu   sync.RWMutex
	path string
	cfg  ModelConfig
}

func NewStore(dataDir string) (*Store, error) {
	if err := os.MkdirAll(dataDir, 0o700); err != nil {
		return nil, err
	}
	s := &Store{path: filepath.Join(dataDir, "model-config.json")}
	data, err := os.ReadFile(s.path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return s, nil
		}
		return nil, err
	}
	if err := json.Unmarshal(data, &s.cfg); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Store) Get() ModelConfig {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.cfg
}

func (s *Store) Public() PublicModelConfig {
	cfg := s.Get()
	return PublicModelConfig{BaseURL: cfg.BaseURL, Model: cfg.Model, Enabled: cfg.Enabled, HasAPIKey: cfg.APIKey != ""}
}

func (s *Store) Save(next ModelConfig, preserveAPIKey bool) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if preserveAPIKey && next.APIKey == "" {
		next.APIKey = s.cfg.APIKey
	}
	data, err := json.MarshalIndent(next, "", "  ")
	if err != nil {
		return err
	}
	tmp := s.path + ".tmp"
	if err := os.WriteFile(tmp, data, 0o600); err != nil {
		return err
	}
	if err := os.Rename(tmp, s.path); err != nil {
		return err
	}
	s.cfg = next
	return nil
}
