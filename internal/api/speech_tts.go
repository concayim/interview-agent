package api

import (
	"bufio"
	"bytes"
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const (
	defaultDoubaoTTSBaseURL = "https://openspeech.bytedance.com"
	maxTTSInputRunes        = 3000
)

type ttsConfig struct {
	APIKey     string
	AppID      string
	ResourceID string
	Speaker    string
}

func (s *Server) synthesizeSpeech(w http.ResponseWriter, r *http.Request) {
	cfg := s.config.Get()
	s.synthesizeSpeechWithConfig(w, r, ttsConfig{
		APIKey: cfg.TTSAPIKey, AppID: cfg.TTSAppID, ResourceID: cfg.TTSResourceID, Speaker: cfg.TTSSpeaker,
	})
}

func (s *Server) synthesizeSpeechWithConfig(w http.ResponseWriter, r *http.Request, cfg ttsConfig) {
	if strings.TrimSpace(cfg.APIKey) == "" || strings.TrimSpace(cfg.Speaker) == "" {
		writeError(w, http.StatusBadRequest, "请先配置豆包 TTS 凭据和音色 ID")
		return
	}
	var input struct {
		Text string `json:"text"`
	}
	if err := readJSON(r, &input); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	input.Text = strings.TrimSpace(input.Text)
	if input.Text == "" {
		writeError(w, http.StatusBadRequest, "朗读文本不能为空")
		return
	}
	if len([]rune(input.Text)) > maxTTSInputRunes {
		writeError(w, http.StatusBadRequest, "单次朗读文本不能超过 3000 字")
		return
	}
	body, err := buildDoubaoTTSRequest(input.Text, cfg.Speaker)
	if err != nil {
		s.internalError(w, err)
		return
	}
	baseURL := s.ttsBaseURL
	if baseURL == "" {
		baseURL = defaultDoubaoTTSBaseURL
	}
	ctx, cancel := context.WithTimeout(r.Context(), 60*time.Second)
	defer cancel()
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, strings.TrimRight(baseURL, "/")+"/api/v3/tts/unidirectional", bytes.NewReader(body))
	if err != nil {
		s.internalError(w, err)
		return
	}
	resourceID := strings.TrimSpace(cfg.ResourceID)
	if resourceID == "" {
		resourceID = "seed-tts-2.0"
	}
	request.Header = doubaoTTSHeaders(cfg.AppID, cfg.APIKey, resourceID)
	request.Header.Set("Content-Type", "application/json")
	response, err := (&http.Client{Timeout: 60 * time.Second}).Do(request)
	if err != nil {
		writeError(w, http.StatusBadGateway, "豆包 TTS 连接失败："+err.Error())
		return
	}
	defer response.Body.Close()
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		payload, _ := io.ReadAll(io.LimitReader(response.Body, 1<<20))
		writeError(w, http.StatusBadGateway, "豆包 TTS 合成失败："+extractProviderError(payload))
		return
	}
	var audio bytes.Buffer
	if err := decodeDoubaoTTSStream(response.Body, &audio); err != nil {
		writeError(w, http.StatusBadGateway, "豆包 TTS 合成失败："+err.Error())
		return
	}
	if audio.Len() == 0 {
		writeError(w, http.StatusBadGateway, "豆包 TTS 没有返回音频")
		return
	}
	w.Header().Set("Content-Type", "audio/mpeg")
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", audio.Len()))
	w.WriteHeader(http.StatusOK)
	_, _ = io.Copy(w, &audio)
}

func buildDoubaoTTSRequest(text, speaker string) ([]byte, error) {
	return json.Marshal(struct {
		User struct {
			UID string `json:"uid"`
		} `json:"user"`
		ReqParams struct {
			Text        string `json:"text"`
			Speaker     string `json:"speaker"`
			AudioParams struct {
				Format     string `json:"format"`
				SampleRate int    `json:"sample_rate"`
				BitRate    int    `json:"bit_rate"`
			} `json:"audio_params"`
		} `json:"req_params"`
	}{
		User: struct {
			UID string `json:"uid"`
		}{UID: "interview-copilot"},
		ReqParams: struct {
			Text        string `json:"text"`
			Speaker     string `json:"speaker"`
			AudioParams struct {
				Format     string `json:"format"`
				SampleRate int    `json:"sample_rate"`
				BitRate    int    `json:"bit_rate"`
			} `json:"audio_params"`
		}{Text: text, Speaker: speaker, AudioParams: struct {
			Format     string `json:"format"`
			SampleRate int    `json:"sample_rate"`
			BitRate    int    `json:"bit_rate"`
		}{Format: "mp3", SampleRate: 24000, BitRate: 64000}},
	})
}

func doubaoTTSHeaders(appID, apiKey, resourceID string) http.Header {
	headers := http.Header{}
	if strings.TrimSpace(appID) != "" {
		headers.Set("X-Api-App-Id", strings.TrimSpace(appID))
		headers.Set("X-Api-Access-Key", strings.TrimSpace(apiKey))
	} else {
		headers.Set("X-Api-Key", strings.TrimSpace(apiKey))
	}
	headers.Set("X-Api-Resource-Id", strings.TrimSpace(resourceID))
	headers.Set("X-Api-Connect-Id", rand.Text())
	return headers
}

func decodeDoubaoTTSStream(reader io.Reader, writer io.Writer) error {
	scanner := bufio.NewScanner(reader)
	scanner.Buffer(make([]byte, 64<<10), 2<<20)
	seenDone := false
	for scanner.Scan() {
		line := bytes.TrimSpace(scanner.Bytes())
		if len(line) == 0 {
			continue
		}
		var frame struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Done    bool   `json:"done"`
			Data    string `json:"data"`
		}
		if err := json.Unmarshal(line, &frame); err != nil {
			return fmt.Errorf("无法解析音频流：%w", err)
		}
		if frame.Code != 0 && frame.Code != 20000000 {
			return fmt.Errorf("%d：%s", frame.Code, strings.TrimSpace(frame.Message))
		}
		if frame.Data != "" {
			audio, err := base64.StdEncoding.DecodeString(frame.Data)
			if err != nil {
				return fmt.Errorf("无法解码音频：%w", err)
			}
			if _, err := writer.Write(audio); err != nil {
				return err
			}
		}
		if frame.Done || frame.Code == 20000000 {
			seenDone = true
			break
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	if !seenDone {
		return errors.New("音频流提前结束")
	}
	return nil
}
