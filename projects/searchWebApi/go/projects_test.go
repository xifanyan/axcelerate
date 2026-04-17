package searchwebapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListProjects(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/searchWebApi/projects" {
			t.Fatalf("path = %q", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":{"successful":true,"backendStatus":"ok","httpStatus":200,"errorMessage":""},"numberResults":1,"results":[{"id":"p1"}]}`))
	}))
	defer ts.Close()

	client, err := NewClient(Config{BaseURL: ts.URL})
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	result, err := client.ListProjects(context.Background())
	if err != nil {
		t.Fatalf("ListProjects() error = %v", err)
	}
	if len(result.Results) != 1 || result.Results[0].ID != "p1" {
		t.Fatalf("results = %+v", result.Results)
	}
}

func TestListCollections(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/searchWebApi/projects/project-a/collections" {
			t.Fatalf("path = %q", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":{"successful":true,"backendStatus":"ok","httpStatus":200,"errorMessage":""},"numberResults":1,"results":[{"id":"default","displayName":"Documents"}]}`))
	}))
	defer ts.Close()

	client, err := NewClient(Config{BaseURL: ts.URL})
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	result, err := client.ListCollections(context.Background(), "project-a")
	if err != nil {
		t.Fatalf("ListCollections() error = %v", err)
	}
	if len(result.Results) != 1 || result.Results[0].DisplayName != "Documents" {
		t.Fatalf("results = %+v", result.Results)
	}
}
