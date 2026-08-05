package api

import (
	"bytes"
	"encoding/base64"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBuildDoubaoTTSRequest(t *testing.T) {
	body, err := buildDoubaoTTSRequest("请介绍一下你自己", "zh_male_beijingxiaoye_moon_bigtts")
	if err != nil {
		t.Fatal(err)
	}
	want := `{"user":{"uid":"interview-copilot"},"req_params":{"text":"请介绍一下你自己","speaker":"zh_male_beijingxiaoye_moon_bigtts","audio_params":{"format":"mp3","sample_rate":24000,"bit_rate":64000}}}`
	if string(body) != want {
		t.Fatalf("unexpected request body:\n%s\nwant:\n%s", body, want)
	}
}

func TestDecodeDoubaoTTSStream(t *testing.T) {
	first := base64.StdEncoding.EncodeToString([]byte("audio-one"))
	second := base64.StdEncoding.EncodeToString([]byte("audio-two"))
	stream := bytes.NewBufferString(`{"code":0,"data":"` + first + `"}` + "\n" + `{"code":0,"data":"` + second + `"}` + "\n" + `{"code":20000000,"done":true}` + "\n")
	var output bytes.Buffer
	if err := decodeDoubaoTTSStream(stream, &output); err != nil {
		t.Fatal(err)
	}
	if output.String() != "audio-oneaudio-two" {
		t.Fatalf("unexpected audio: %q", output.String())
	}
}

func TestDecodeDoubaoTTSStreamReturnsProviderError(t *testing.T) {
	stream := bytes.NewBufferString(`{"code":45000001,"message":"speaker is invalid"}` + "\n")
	if err := decodeDoubaoTTSStream(stream, io.Discard); err == nil {
		t.Fatal("expected provider error")
	}
}

func TestDoubaoTTSHeadersSupportBothConsoleVersions(t *testing.T) {
	legacy := doubaoTTSHeaders("app-id", "access-token", "seed-tts-1.0")
	if legacy.Get("X-Api-App-Id") != "app-id" || legacy.Get("X-Api-Access-Key") != "access-token" || legacy.Get("X-Api-Key") != "" {
		t.Fatalf("unexpected legacy headers: %#v", legacy)
	}
	if legacy.Get("X-Api-App-Key") != "" || legacy.Get("X-Api-Connect-Id") == "" {
		t.Fatalf("legacy headers use an invalid app header or omit connect id: %#v", legacy)
	}
	modern := doubaoTTSHeaders("", "api-key", "seed-tts-2.0")
	if modern.Get("X-Api-Key") != "api-key" || modern.Get("X-Api-Resource-Id") != "seed-tts-2.0" || modern.Get("X-Api-Connect-Id") == "" {
		t.Fatalf("unexpected modern headers: %#v", modern)
	}
}

func TestSynthesizeSpeechStreamsMP3(t *testing.T) {
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/tts/unidirectional" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if r.Header.Get("X-Api-Key") != "tts-key" || r.Header.Get("X-Api-Resource-Id") != "seed-tts-2.0" {
			t.Fatalf("unexpected headers: %#v", r.Header)
		}
		encoded := base64.StdEncoding.EncodeToString([]byte("fake-mp3"))
		_, _ = w.Write([]byte(`{"code":0,"data":"` + encoded + `"}` + "\n" + `{"code":20000000,"done":true}` + "\n"))
	}))
	defer upstream.Close()

	server := &Server{ttsBaseURL: upstream.URL}
	request := httptest.NewRequest(http.MethodPost, "/api/v1/speech/synthesis", bytes.NewBufferString(`{"text":"你好"}`))
	response := httptest.NewRecorder()
	server.synthesizeSpeechWithConfig(response, request, ttsConfig{APIKey: "tts-key", ResourceID: "seed-tts-2.0", Speaker: "voice"})
	if response.Code != http.StatusOK || response.Header().Get("Content-Type") != "audio/mpeg" || response.Body.String() != "fake-mp3" {
		t.Fatalf("unexpected response: status=%d type=%q body=%q", response.Code, response.Header().Get("Content-Type"), response.Body.String())
	}
}
