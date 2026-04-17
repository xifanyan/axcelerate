package searchwebapi

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetBinaryByRecordID(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("field") != "native" {
			t.Fatalf("field = %q", r.URL.Query().Get("field"))
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		_, _ = w.Write([]byte("hello"))
	}))
	defer ts.Close()

	client, err := NewClient(Config{BaseURL: ts.URL})
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	result, err := client.GetBinaryByRecordID(context.Background(), "p1", "default", "r1", "native")
	if err != nil {
		t.Fatalf("GetBinaryByRecordID() error = %v", err)
	}
	defer result.Body.Close()
	data, err := io.ReadAll(result.Body)
	if err != nil {
		t.Fatalf("ReadAll() error = %v", err)
	}
	if string(data) != "hello" {
		t.Fatalf("body = %q", string(data))
	}
}
