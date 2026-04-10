package adp

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
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

func TestExecuteUsesPUTJSONAndAuthHeaders(t *testing.T) {
	t.Helper()

	var gotMethod string
	var gotContentType string
	var gotUser string
	var gotPassword string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotContentType = r.Header.Get("Content-Type")
		gotUser = r.Header.Get("Auth-Username")
		gotPassword = r.Header.Get("Auth-Password")
		io.WriteString(w, `{"executionId":"1","taskType":"List Entities","loggingEnabled":"true","progressMax":1,"executionStatus":"success","executionRootDir":"root","contextId":"ctx","executionPersistent":"true","progressCurrent":1,"progressPercentage":1.0,"taskDisplayName":"List entities","executionMetaData":{}}`)
	}))
	defer server.Close()

	client, err := NewClient(ClientConfig{BaseURL: server.URL, Username: "adp", Password: "secret"})
	if err != nil {
		t.Fatalf("NewClient error: %v", err)
	}

	_, err = client.execute(context.Background(), "/executeAdpTask", rawTaskRequest{
		TaskType:          "List Entities",
		TaskConfiguration: map[string]any{"adp_loggingEnabled": true},
		TaskDescription:   "Writes a list of entities ot an output variable",
		TaskDisplayName:   "List entities",
	})
	if err != nil {
		t.Fatalf("execute returned error: %v", err)
	}

	if gotMethod != http.MethodPut {
		t.Fatalf("method = %s, want PUT", gotMethod)
	}
	if gotContentType != "application/json" {
		t.Fatalf("content-type = %q", gotContentType)
	}
	if gotUser != "adp" || gotPassword != "secret" {
		t.Fatalf("unexpected auth headers %q / %q", gotUser, gotPassword)
	}
}

func TestExecuteReturnsTaskExecutionErrorOnFailedStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `{"executionId":"2","taskType":"List Entities","loggingEnabled":"true","progressMax":1,"executionStatus":"failed","executionRootDir":"root","contextId":"ctx","executionPersistent":"true","progressCurrent":0,"progressPercentage":0.0,"taskDisplayName":"List entities","errorMessage":"boom","executionMetaData":null}`)
	}))
	defer server.Close()

	client, err := NewClient(ClientConfig{BaseURL: server.URL, Username: "adp", Password: "secret"})
	if err != nil {
		t.Fatalf("NewClient error: %v", err)
	}

	_, err = client.execute(context.Background(), "/executeAdpTask", rawTaskRequest{TaskType: "List Entities"})
	if err == nil {
		t.Fatal("expected error")
	}

	var execErr *TaskExecutionError
	if !errors.As(err, &execErr) {
		t.Fatalf("expected TaskExecutionError, got %T", err)
	}
	if execErr.ExecutionID != "2" || execErr.ErrorMessage != "boom" {
		t.Fatalf("unexpected execution error: %+v", execErr)
	}
}
