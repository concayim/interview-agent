package api

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

const (
	volcSpeechWebSocketURL      = "wss://openspeech.bytedance.com/api/v3/sauc/bigmodel_async"
	defaultVolcSpeechResourceID = "volc.bigasr.sauc.duration"
	maxRealtimeAudioFrameSize   = 256 << 10
)

var realtimeSpeechUpgrader = websocket.Upgrader{
	ReadBufferSize:  16 << 10,
	WriteBufferSize: 16 << 10,
	CheckOrigin:     func(*http.Request) bool { return true },
}

type realtimeClientEvent struct {
	audio []byte
	stop  bool
	err   error
}

type volcSpeechResult struct {
	text    string
	final   bool
	err     error
	payload []byte
}

func (s *Server) realtimeTranscribeSpeech(w http.ResponseWriter, r *http.Request) {
	cfg := s.config.Get()
	if strings.TrimSpace(cfg.SpeechAPIKey) == "" {
		writeError(w, http.StatusBadRequest, "请先在模型设置中填写火山语音 API Key")
		return
	}

	clientConn, err := realtimeSpeechUpgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer clientConn.Close()
	clientConn.SetReadLimit(maxRealtimeAudioFrameSize)

	resourceID := strings.TrimSpace(cfg.SpeechResourceID)
	if resourceID == "" {
		resourceID = defaultVolcSpeechResourceID
	}
	connectID := uuid.NewString()
	headers := volcSpeechAuthHeaders(cfg.SpeechAppID, cfg.SpeechAPIKey)
	headers.Set("X-Api-Resource-Id", resourceID)
	headers.Set("X-Api-Connect-Id", connectID)
	headers.Set("X-Api-Request-Id", connectID)
	headers.Set("X-Api-Sequence", "-1")

	ctx, cancel := context.WithTimeout(r.Context(), 30*time.Minute)
	defer cancel()
	upstreamConn, response, err := websocket.DefaultDialer.DialContext(ctx, volcSpeechWebSocketURL, headers)
	if err != nil {
		message := "连接火山实时语音识别失败"
		if response != nil {
			message = fmt.Sprintf("连接火山实时语音识别失败（HTTP %d）", response.StatusCode)
		}
		_ = clientConn.WriteJSON(map[string]string{"type": "error", "error": message})
		return
	}
	defer upstreamConn.Close()

	language := normalizeSpeechLanguage(r.URL.Query().Get("language"))
	startPayload, err := json.Marshal(map[string]any{
		"user": map[string]any{"uid": "interview-copilot"},
		"audio": map[string]any{
			"format": "pcm", "codec": "raw", "rate": 16000, "bits": 16, "channel": 1, "language": language,
		},
		"request": map[string]any{
			"reqid": connectID, "model_name": "bigmodel", "enable_itn": true, "enable_punc": true,
			"show_utterances": true, "result_type": "full", "enable_nonstream": true, "end_window_size": 600,
		},
	})
	if err != nil {
		_ = clientConn.WriteJSON(map[string]string{"type": "error", "error": err.Error()})
		return
	}
	if err := upstreamConn.WriteMessage(websocket.BinaryMessage, buildVolcSpeechFrame(0x1, 0x0, 0x1, startPayload)); err != nil {
		_ = clientConn.WriteJSON(map[string]string{"type": "error", "error": "初始化火山实时语音识别失败：" + err.Error()})
		return
	}
	if err := clientConn.WriteJSON(map[string]string{"type": "ready"}); err != nil {
		return
	}

	clientEvents := make(chan realtimeClientEvent, 16)
	upstreamEvents := make(chan volcSpeechResult, 16)
	go readRealtimeClient(clientConn, clientEvents)
	go readVolcSpeech(upstreamConn, upstreamEvents)

	var stopTimer <-chan time.Time
	stopping := false
	lastText := ""
	for {
		select {
		case event := <-clientEvents:
			if event.err != nil {
				return
			}
			if event.stop {
				if stopping {
					continue
				}
				stopping = true
				if err := upstreamConn.WriteMessage(websocket.BinaryMessage, buildVolcSpeechFrame(0x2, 0x2, 0x0, nil)); err != nil {
					_ = clientConn.WriteJSON(map[string]string{"type": "error", "error": "结束语音识别失败：" + err.Error()})
					return
				}
				stopTimer = time.After(8 * time.Second)
				continue
			}
			if stopping || len(event.audio) == 0 {
				continue
			}
			if err := upstreamConn.WriteMessage(websocket.BinaryMessage, buildVolcSpeechFrame(0x2, 0x0, 0x0, event.audio)); err != nil {
				_ = clientConn.WriteJSON(map[string]string{"type": "error", "error": "发送语音数据失败：" + err.Error()})
				return
			}
		case result := <-upstreamEvents:
			if result.err != nil {
				_ = clientConn.WriteJSON(map[string]string{"type": "error", "error": "火山语音识别失败：" + result.err.Error()})
				return
			}
			if result.text != "" {
				lastText = result.text
				eventType := "partial"
				if stopping && result.final {
					eventType = "final"
				}
				if err := clientConn.WriteJSON(map[string]string{"type": eventType, "text": result.text}); err != nil {
					return
				}
			}
			if stopping && result.final {
				return
			}
		case <-stopTimer:
			_ = clientConn.WriteJSON(map[string]string{"type": "final", "text": lastText})
			return
		case <-ctx.Done():
			_ = clientConn.WriteJSON(map[string]string{"type": "error", "error": "实时语音识别连接已超时"})
			return
		}
	}
}

func volcSpeechAuthHeaders(appID, key string) http.Header {
	headers := http.Header{}
	if strings.TrimSpace(appID) != "" {
		headers.Set("X-Api-App-Key", strings.TrimSpace(appID))
		headers.Set("X-Api-Access-Key", strings.TrimSpace(key))
	} else {
		headers.Set("X-Api-Key", strings.TrimSpace(key))
	}
	return headers
}

func readRealtimeClient(conn *websocket.Conn, events chan<- realtimeClientEvent) {
	for {
		messageType, payload, err := conn.ReadMessage()
		if err != nil {
			events <- realtimeClientEvent{err: err}
			return
		}
		switch messageType {
		case websocket.BinaryMessage:
			events <- realtimeClientEvent{audio: payload}
		case websocket.TextMessage:
			var message struct {
				Type string `json:"type"`
			}
			if json.Unmarshal(payload, &message) == nil && message.Type == "stop" {
				events <- realtimeClientEvent{stop: true}
			}
		}
	}
}

func readVolcSpeech(conn *websocket.Conn, events chan<- volcSpeechResult) {
	for {
		messageType, payload, err := conn.ReadMessage()
		if err != nil {
			events <- volcSpeechResult{err: err}
			return
		}
		if messageType != websocket.BinaryMessage {
			events <- volcSpeechResult{err: errors.New(strings.TrimSpace(string(payload)))}
			return
		}
		result := parseVolcSpeechFrame(payload)
		events <- result
		if result.err != nil || result.final {
			return
		}
	}
}

func buildVolcSpeechFrame(messageType, flags, serialization byte, payload []byte) []byte {
	frame := bytes.NewBuffer(make([]byte, 0, 8+len(payload)))
	frame.WriteByte(0x11)
	frame.WriteByte(messageType<<4 | flags)
	frame.WriteByte(serialization << 4)
	frame.WriteByte(0x00)
	_ = binary.Write(frame, binary.BigEndian, uint32(len(payload)))
	_, _ = frame.Write(payload)
	return frame.Bytes()
}

func parseVolcSpeechFrame(frame []byte) volcSpeechResult {
	if len(frame) < 8 {
		return volcSpeechResult{err: errors.New("火山返回了不完整的数据帧")}
	}
	headerSize := int(frame[0]&0x0f) * 4
	messageType := frame[1] >> 4
	flags := frame[1] & 0x0f
	offset := headerSize
	if flags == 0x1 || flags == 0x3 {
		if len(frame) < offset+4 {
			return volcSpeechResult{err: errors.New("火山返回的数据帧缺少序号")}
		}
		offset += 4
	}
	if messageType == 0xf {
		if len(frame) < offset+8 {
			return volcSpeechResult{err: errors.New("火山返回了不完整的错误帧")}
		}
		code := binary.BigEndian.Uint32(frame[offset : offset+4])
		offset += 4
		size := int(binary.BigEndian.Uint32(frame[offset : offset+4]))
		offset += 4
		if size < 0 || len(frame) < offset+size {
			return volcSpeechResult{err: fmt.Errorf("火山语音错误 %d", code)}
		}
		return volcSpeechResult{err: fmt.Errorf("%d：%s", code, strings.TrimSpace(string(frame[offset:offset+size])))}
	}
	if len(frame) < offset+4 {
		return volcSpeechResult{err: errors.New("火山返回的数据帧缺少正文长度")}
	}
	size := int(binary.BigEndian.Uint32(frame[offset : offset+4]))
	offset += 4
	if size < 0 || len(frame) < offset+size {
		return volcSpeechResult{err: errors.New("火山返回的数据帧正文不完整")}
	}
	payload := frame[offset : offset+size]
	var response struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Result  struct {
			Text       string `json:"text"`
			Utterances []struct {
				Definite bool `json:"definite"`
			} `json:"utterances"`
		} `json:"result"`
	}
	if err := json.Unmarshal(payload, &response); err != nil {
		return volcSpeechResult{err: fmt.Errorf("无法解析火山语音响应：%w", err), payload: payload}
	}
	if response.Code != 0 && response.Code != 1000 {
		return volcSpeechResult{err: fmt.Errorf("%d：%s", response.Code, response.Message), payload: payload}
	}
	final := flags == 0x2 || flags == 0x3
	return volcSpeechResult{text: strings.TrimSpace(response.Result.Text), final: final, payload: payload}
}

func normalizeSpeechLanguage(value string) string {
	if strings.HasPrefix(strings.ToLower(strings.TrimSpace(value)), "en") {
		return "en-US"
	}
	return "zh-CN"
}
