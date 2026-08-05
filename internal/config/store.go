package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"sync"
)

type ModelConfig struct {
	APIKey           string `json:"apiKey"`
	BaseURL          string `json:"baseUrl"`
	Model            string `json:"model"`
	Enabled          bool   `json:"enabled"`
	SpeechAPIKey     string `json:"speechApiKey,omitempty"`
	SpeechAppID      string `json:"speechAppId,omitempty"`
	SpeechResourceID string `json:"speechResourceId,omitempty"`
	TTSAPIKey        string `json:"ttsApiKey,omitempty"`
	TTSAppID         string `json:"ttsAppId,omitempty"`
	TTSResourceID    string `json:"ttsResourceId,omitempty"`
	TTSSpeaker       string `json:"ttsSpeaker,omitempty"`
	TTSEnabled       bool   `json:"ttsEnabled,omitempty"`
}

func (c ModelConfig) Ready() bool {
	return c.Enabled && c.APIKey != "" && c.Model != ""
}

type PublicModelConfig struct {
	BaseURL          string `json:"baseUrl"`
	Model            string `json:"model"`
	Enabled          bool   `json:"enabled"`
	HasAPIKey        bool   `json:"hasApiKey"`
	SpeechAppID      string `json:"speechAppId"`
	SpeechResourceID string `json:"speechResourceId"`
	HasSpeechAPIKey  bool   `json:"hasSpeechApiKey"`
	TTSAppID         string `json:"ttsAppId"`
	TTSResourceID    string `json:"ttsResourceId"`
	TTSSpeaker       string `json:"ttsSpeaker"`
	TTSEnabled       bool   `json:"ttsEnabled"`
	HasTTSAPIKey     bool   `json:"hasTtsApiKey"`
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
	var persisted map[string]json.RawMessage
	if json.Unmarshal(data, &persisted) == nil {
		if _, legacySpeechModel := persisted["speechModel"]; legacySpeechModel {
			delete(persisted, "speechModel")
			cleaned, marshalErr := json.MarshalIndent(persisted, "", "  ")
			if marshalErr != nil {
				return nil, marshalErr
			}
			if writeErr := os.WriteFile(s.path, cleaned, 0o600); writeErr != nil {
				return nil, writeErr
			}
		}
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
	return PublicModelConfig{
		BaseURL: cfg.BaseURL, Model: cfg.Model, Enabled: cfg.Enabled, HasAPIKey: cfg.APIKey != "",
		SpeechAppID: cfg.SpeechAppID, SpeechResourceID: cfg.SpeechResourceID, HasSpeechAPIKey: cfg.SpeechAPIKey != "",
		TTSAppID: cfg.TTSAppID, TTSResourceID: cfg.TTSResourceID, TTSSpeaker: cfg.TTSSpeaker, TTSEnabled: cfg.TTSEnabled, HasTTSAPIKey: cfg.TTSAPIKey != "",
	}
}

func (s *Store) Save(next ModelConfig, preserveAPIKey bool, preserveSecrets ...bool) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if preserveAPIKey && next.APIKey == "" {
		next.APIKey = s.cfg.APIKey
	}
	if len(preserveSecrets) > 0 && preserveSecrets[0] && next.SpeechAPIKey == "" {
		next.SpeechAPIKey = s.cfg.SpeechAPIKey
	}
	if len(preserveSecrets) > 1 && preserveSecrets[1] && next.TTSAPIKey == "" {
		next.TTSAPIKey = s.cfg.TTSAPIKey
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
