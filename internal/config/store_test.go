package config

import "testing"

func TestSavePreservesIndependentSpeechAndTTSSecrets(t *testing.T) {
	store, err := NewStore(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	initial := ModelConfig{
		APIKey: "chat-key", Model: "chat-model", Enabled: true,
		SpeechAPIKey: "asr-token", SpeechAppID: "asr-app",
		TTSAPIKey: "tts-key", TTSAppID: "tts-app", TTSResourceID: "seed-tts-2.0", TTSSpeaker: "voice", TTSEnabled: true,
	}
	if err := store.Save(initial, false, false, false); err != nil {
		t.Fatal(err)
	}
	next := initial
	next.APIKey = ""
	next.SpeechAPIKey = ""
	next.TTSAPIKey = ""
	if err := store.Save(next, true, true, true); err != nil {
		t.Fatal(err)
	}
	got := store.Get()
	if got.APIKey != "chat-key" || got.SpeechAPIKey != "asr-token" || got.TTSAPIKey != "tts-key" {
		t.Fatalf("secrets were not preserved independently: %#v", got)
	}
	public := store.Public()
	if !public.HasAPIKey || !public.HasSpeechAPIKey || !public.HasTTSAPIKey || !public.TTSEnabled || public.TTSSpeaker != "voice" {
		t.Fatalf("unexpected public config: %#v", public)
	}
}
