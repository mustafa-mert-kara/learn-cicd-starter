package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	tests := []struct {
		name      string
		headers   http.Header
		wantToken string
		wantErr   bool
	}{
		{
			name:      "Empty Header",
			headers:   http.Header{},
			wantToken: "",
			wantErr:   true,
		},
		{
			name: "Wrong syntax",
			headers: http.Header{
				"Authorization": []string{"Deneme"},
			},
			wantToken: "",
			wantErr:   true,
		},
		{
			name: "Valid Token",
			headers: http.Header{
				"Authorization": []string{"ApiKey 123"},
			},
			wantToken: "123",
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotToken, err := GetAPIKey(tt.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotToken != tt.wantToken {
				t.Errorf("GetAPIKey() gotToken = %v, want %v", gotToken, tt.wantToken)
			}
		})
	}
}
