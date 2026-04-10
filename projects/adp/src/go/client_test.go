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
