package adp

import (
	"testing"
	"time"
)

func TestNewClientNormalizesBaseURLAndTimeout(t *testing.T) {
	client, err := NewClient(ClientConfig{
		BaseURL:  "https://example.com/adp/rest/api/task/",
		Username: "adp",
		Password: "secret",
	})
	if err != nil {
		t.Fatalf("NewClient returned error: %v", err)
	}

	if client.baseURL != "https://example.com/adp/rest/api/task" {
		t.Fatalf("baseURL = %q, want trimmed path", client.baseURL)
	}

	if got, want := client.httpClient.Timeout, 30*time.Second; got != want {
		t.Fatalf("timeout = %v, want %v", got, want)
	}

	if client.debugOut == nil {
		t.Fatal("debugOut should default to a non-nil writer")
	}
}

func TestNewClientRequiresFields(t *testing.T) {
	tests := []struct {
		name string
		cfg  ClientConfig
		want string
	}{
		{
			name: "base url",
			cfg: ClientConfig{
				Username: "adp",
				Password: "secret",
			},
			want: "base URL is required",
		},
		{
			name: "username",
			cfg: ClientConfig{
				BaseURL:  "https://example.com",
				Password: "secret",
			},
			want: "username is required",
		},
		{
			name: "password",
			cfg: ClientConfig{
				BaseURL:  "https://example.com",
				Username: "adp",
			},
			want: "password is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewClient(tt.cfg)
			if err == nil {
				t.Fatal("NewClient returned nil error")
			}

			if err.Error() != tt.want {
				t.Fatalf("error = %q, want %q", err.Error(), tt.want)
			}
		})
	}
}

func TestNewClientTrimsWhitespaceAroundBaseURL(t *testing.T) {
	client, err := NewClient(ClientConfig{
		BaseURL:  "  https://example.com/adp/rest/api/task/  ",
		Username: "adp",
		Password: "secret",
	})
	if err != nil {
		t.Fatalf("NewClient returned error: %v", err)
	}

	if client.baseURL != "https://example.com/adp/rest/api/task" {
		t.Fatalf("baseURL = %q, want fully trimmed path", client.baseURL)
	}
}
