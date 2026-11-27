package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name          string
		authorization string
		wantKey       string
		wantErr       bool
	}{
		{
			name:          "valid API key",
			authorization: "ApiKey test-key",
			wantKey:       "test-key",
		},
		{
			name:    "missing authorization header",
			wantErr: true,
		},
		{
			name:          "malformed authorization header",
			authorization: "ApiKey",
			wantErr:       true,
		},
		{
			name:          "unsupported authorization scheme",
			authorization: "Bearer test-key",
			wantErr:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := make(http.Header)
			if tt.authorization != "" {
				headers.Set("Authorization", tt.authorization)
			}

			got, err := GetAPIKey(headers)
			if (err != nil) != tt.wantErr {
				t.Fatalf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil && got != tt.wantKey {
				t.Errorf("GetAPIKey() = %q, want %q", got, tt.wantKey)
			}
		})
	}
}
