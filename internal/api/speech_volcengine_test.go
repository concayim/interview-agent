package api

import (
	"encoding/json"
	"testing"
)

func TestBuildVolcSpeechFrame(t *testing.T) {
	payload := []byte(`{"audio":{"format":"pcm"}}`)
	frame := buildVolcSpeechFrame(0x1, 0x0, 0x1, payload)
	if len(frame) != len(payload)+8 {
		t.Fatalf("unexpected frame size: %d", len(frame))
	}
	if frame[0] != 0x11 || frame[1] != 0x10 || frame[2] != 0x10 {
		t.Fatalf("unexpected protocol header: %x", frame[:4])
	}
}

func TestParseVolcSpeechFrame(t *testing.T) {
	payload, err := json.Marshal(map[string]any{
		"code":   1000,
		"result": map[string]any{"text": "你好，世界"},
	})
	if err != nil {
		t.Fatal(err)
	}
	result := parseVolcSpeechFrame(buildVolcSpeechFrame(0x9, 0x2, 0x1, payload))
	if result.err != nil {
		t.Fatal(result.err)
	}
	if result.text != "你好，世界" || !result.final {
		t.Fatalf("unexpected result: %#v", result)
	}
}

func TestParseVolcSpeechErrorFrame(t *testing.T) {
	payload := []byte("invalid credentials")
	frame := buildVolcSpeechFrame(0xf, 0x0, 0x1, payload)
	frame = append(frame[:4], append([]byte{0, 0, 17, 112}, frame[4:]...)...)
	result := parseVolcSpeechFrame(frame)
	if result.err == nil {
		t.Fatal("expected provider error")
	}
}

func TestNormalizeSpeechLanguage(t *testing.T) {
	if normalizeSpeechLanguage("en-US") != "en-US" || normalizeSpeechLanguage("zh") != "zh-CN" {
		t.Fatal("unexpected language normalization")
	}
}
