package adp

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
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

	if client.debugOut != io.Discard {
		t.Fatalf("debugOut = %#v, want io.Discard", client.debugOut)
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

func TestExecuteReturnsHTTPErrorForNon2xxResponses(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadGateway)
		io.WriteString(w, "upstream failed")
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

	if got := err.Error(); got != "unexpected HTTP status 502 Bad Gateway: upstream failed" {
		t.Fatalf("error = %q", got)
	}
}

func TestExecuteReturnsDecodeErrorForMalformedResponseBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "not-json")
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

	if got := err.Error(); !strings.HasPrefix(got, "decode response:") {
		t.Fatalf("error = %q", got)
	}
}

func TestPollSendsExecutionID(t *testing.T) {
	var gotBody map[string]any

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewDecoder(r.Body).Decode(&gotBody); err != nil {
			t.Fatalf("decode request body: %v", err)
		}
		io.WriteString(w, `{"executionId":"3","taskType":"Create OCR Job","loggingEnabled":"true","progressMax":1,"executionStatus":"running","executionRootDir":"root","contextId":"ctx","executionPersistent":"true","progressCurrent":0,"progressPercentage":0.0,"taskDisplayName":"Create OCR Job","executionMetaData":[]}`)
	}))
	defer server.Close()

	client, err := NewClient(ClientConfig{BaseURL: server.URL, Username: "adp", Password: "secret"})
	if err != nil {
		t.Fatalf("NewClient error: %v", err)
	}

	_, err = client.Poll(context.Background(), "abc-123")
	if err != nil {
		t.Fatalf("Poll error: %v", err)
	}

	if gotBody["executionId"] != "abc-123" {
		t.Fatalf("executionId body = %#v", gotBody["executionId"])
	}
}

func TestWaitPollsUntilTerminalSuccess(t *testing.T) {
	calls := 0

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		if calls == 1 {
			io.WriteString(w, `{"executionId":"4","taskType":"Create OCR Job","loggingEnabled":"true","progressMax":1,"executionStatus":"running","executionRootDir":"root","contextId":"ctx","executionPersistent":"true","progressCurrent":0,"progressPercentage":0.0,"taskDisplayName":"Create OCR Job","executionMetaData":[]}`)
			return
		}
		io.WriteString(w, `{"executionId":"4","taskType":"Create OCR Job","loggingEnabled":"true","progressMax":1,"executionStatus":"success","executionRootDir":"root","contextId":"ctx","executionPersistent":"true","progressCurrent":1,"progressPercentage":1.0,"taskDisplayName":"Create OCR Job","executionMetaData":[]}`)
	}))
	defer server.Close()

	client, err := NewClient(ClientConfig{BaseURL: server.URL, Username: "adp", Password: "secret"})
	if err != nil {
		t.Fatalf("NewClient error: %v", err)
	}

	resp, err := client.Wait(context.Background(), "4", time.Millisecond)
	if err != nil {
		t.Fatalf("Wait error: %v", err)
	}
	if resp.ExecutionStatus != "success" {
		t.Fatalf("status = %q, want success", resp.ExecutionStatus)
	}
	if calls != 2 {
		t.Fatalf("calls = %d, want 2", calls)
	}
}

func TestPollReturnsProtocolErrorForEmptyExecutionStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `{"executionId":"5","taskType":"Create OCR Job","loggingEnabled":"true","progressMax":1,"executionStatus":"","executionRootDir":"root","contextId":"ctx","executionPersistent":"true","progressCurrent":0,"progressPercentage":0.0,"taskDisplayName":"Create OCR Job","executionMetaData":[]}`)
	}))
	defer server.Close()

	client, err := NewClient(ClientConfig{BaseURL: server.URL, Username: "adp", Password: "secret"})
	if err != nil {
		t.Fatalf("NewClient error: %v", err)
	}

	_, err = client.Poll(context.Background(), "abc-123")
	if err == nil {
		t.Fatal("expected error")
	}

	if got := err.Error(); got != "invalid polling response: missing executionStatus" {
		t.Fatalf("error = %q", got)
	}
}
