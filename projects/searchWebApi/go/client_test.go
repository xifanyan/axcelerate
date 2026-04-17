package searchwebapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewClientNormalizesBaseURL(t *testing.T) {
	client, err := NewClient(Config{BaseURL: "https://example.com/root/"})
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	if got := client.baseURL.String(); got != "https://example.com/root/searchWebApi" {
		t.Fatalf("baseURL = %q", got)
	}
}

func TestNewClientRejectsMissingBaseURL(t *testing.T) {
	_, err := NewClient(Config{})
	if err == nil {
		t.Fatal("expected error for missing base URL")
	}
}

func TestClientSessionAccessors(t *testing.T) {
	client, err := NewClient(Config{BaseURL: "https://example.com"})
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	client.SetSessionID("abc")
	if got := client.SessionID(); got != "abc" {
		t.Fatalf("SessionID() = %q", got)
	}

	client.ClearSessionID()
	if got := client.SessionID(); got != "" {
		t.Fatalf("SessionID() after clear = %q", got)
	}
}

func TestDoJSONAppliesSessionAndAuthHeaders(t *testing.T) {
	var gotSession string
	var gotBasicUser, gotBasicPass string

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotSession = r.Header.Get("SWA-SESSION")
		gotBasicUser, gotBasicPass, _ = r.BasicAuth()
		w.Header().Set("SWA-SESSION", "next-session")
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":{"successful":true,"backendStatus":"ok","httpStatus":200,"errorMessage":""}}`))
	}))
	defer ts.Close()

	client, err := NewClient(Config{BaseURL: ts.URL, Username: "user", Password: "pass", SessionID: "start"})
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	var out LoginResult
	if err := client.doJSON(context.Background(), http.MethodPost, []string{"login"}, nil, nil, nil, &out); err != nil {
		t.Fatalf("doJSON() error = %v", err)
	}

	if gotSession != "start" {
		t.Fatalf("SWA-SESSION header = %q", gotSession)
	}
	if gotBasicUser != "user" || gotBasicPass != "pass" {
		t.Fatalf("basic auth = %q/%q", gotBasicUser, gotBasicPass)
	}
	if client.SessionID() != "next-session" {
		t.Fatalf("stored session = %q", client.SessionID())
	}
}
