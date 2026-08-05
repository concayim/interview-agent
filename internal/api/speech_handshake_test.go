package api

import (
	"net/http"
	"testing"
)

func TestVolcSpeechAuthHeaders(t *testing.T) {
	tests := []struct {
		name   string
		appID  string
		key    string
		assert func(*testing.T, http.Header)
	}{
		{
			name:  "legacy console",
			appID: "123456789",
			key:   "access-token",
			assert: func(t *testing.T, headers http.Header) {
				if headers.Get("X-Api-App-Key") != "123456789" || headers.Get("X-Api-Access-Key") != "access-token" || headers.Get("X-Api-Key") != "" {
					t.Fatalf("unexpected legacy headers: %#v", headers)
				}
			},
		},
		{
			name: "new console",
			key:  "api-key",
			assert: func(t *testing.T, headers http.Header) {
				if headers.Get("X-Api-Key") != "api-key" || headers.Get("X-Api-App-Key") != "" || headers.Get("X-Api-Access-Key") != "" {
					t.Fatalf("unexpected new headers: %#v", headers)
				}
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			headers := volcSpeechAuthHeaders(test.appID, test.key)
			test.assert(t, headers)
		})
	}
}
