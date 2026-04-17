package searchwebapi

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSearchRecords(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("query") != "hello" {
			t.Fatalf("query = %q", r.URL.Query().Get("query"))
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":{"successful":true,"backendStatus":"ok","httpStatus":200,"errorMessage":""},"numberResults":1,"results":[{"rank":1,"relevance":1,"id":"r1","uniqueField":"u1","fields":[],"folderSets":[],"body":"","page":0,"pageCount":0}]}`))
	}))
	defer ts.Close()

	client, err := NewClient(Config{BaseURL: ts.URL})
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	result, err := client.SearchRecords(context.Background(), "p1", "default", SearchRecordsOptions{Query: "hello"})
	if err != nil {
		t.Fatalf("SearchRecords() error = %v", err)
	}
	if len(result.Results) != 1 || result.Results[0].ID != "r1" {
		t.Fatalf("results = %+v", result.Results)
	}
}

func TestSearchRecordsStream(t *testing.T) {
	body := strings.Join([]string{
		`{"status":{"successful":true,"backendStatus":"ok","httpStatus":200,"errorMessage":""},"numberResults":1,"results":[]}`,
		`{"rank":1,"relevance":1,"id":"r1","uniqueField":"u1","fields":[],"folderSets":[],"body":"","page":0,"pageCount":0}`,
	}, "\n")

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/x-ndjson")
		_, _ = w.Write([]byte(body))
	}))
	defer ts.Close()

	client, err := NewClient(Config{BaseURL: ts.URL})
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	stream, err := client.SearchRecordsStream(context.Background(), "p1", "default", SearchRecordsOptions{Query: "hello"})
	if err != nil {
		t.Fatalf("SearchRecordsStream() error = %v", err)
	}
	defer stream.Close()

	if stream.Meta.NumberResults != 1 {
		t.Fatalf("meta = %+v", stream.Meta)
	}
	record, err := stream.Next()
	if err != nil {
		t.Fatalf("Next() error = %v", err)
	}
	if record.ID != "r1" {
		t.Fatalf("record = %+v", record)
	}
	_, err = stream.Next()
	if err != io.EOF {
		t.Fatalf("expected EOF, got %v", err)
	}
}
