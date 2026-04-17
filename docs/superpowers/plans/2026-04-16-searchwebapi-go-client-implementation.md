# searchWebApi Go Client Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Generate a working Go SDK and small example CLI for the full `searchWebApi` contract at `C:\Users\pyan\ai-generated\searchWebApi\go`.

**Architecture:** Build one Go module with a flat root `Client`, shared transport/auth helpers, shared wire-shaped schema types, and resource-focused method files. Keep request construction and response decoding close to the raw API while making binary, multipart, and NDJSON behavior explicit.

**Tech Stack:** Go, `net/http`, `encoding/json`, `mime/multipart`, `bufio`, `httptest`

---

## File Structure

- Create: `C:\Users\pyan\ai-generated\searchWebApi\go\go.mod`
- Create: `C:\Users\pyan\ai-generated\searchWebApi\go\README.md`
- Create: `C:\Users\pyan\ai-generated\searchWebApi\go\client.go`
- Create: `C:\Users\pyan\ai-generated\searchWebApi\go\auth.go`
- Create: `C:\Users\pyan\ai-generated\searchWebApi\go\transport.go`
- Create: `C:\Users\pyan\ai-generated\searchWebApi\go\types.go`
- Create: `C:\Users\pyan\ai-generated\searchWebApi\go\projects.go`
- Create: `C:\Users\pyan\ai-generated\searchWebApi\go\collections.go`
- Create: `C:\Users\pyan\ai-generated\searchWebApi\go\records.go`
- Create: `C:\Users\pyan\ai-generated\searchWebApi\go\binary.go`
- Create: `C:\Users\pyan\ai-generated\searchWebApi\go\measures.go`
- Create: `C:\Users\pyan\ai-generated\searchWebApi\go\cached_searches.go`
- Create: `C:\Users\pyan\ai-generated\searchWebApi\go\change_queue.go`
- Create: `C:\Users\pyan\ai-generated\searchWebApi\go\insert_remove.go`
- Create: `C:\Users\pyan\ai-generated\searchWebApi\go\session.go`
- Create: `C:\Users\pyan\ai-generated\searchWebApi\go\client_test.go`
- Create: `C:\Users\pyan\ai-generated\searchWebApi\go\projects_test.go`
- Create: `C:\Users\pyan\ai-generated\searchWebApi\go\collections_test.go`
- Create: `C:\Users\pyan\ai-generated\searchWebApi\go\records_test.go`
- Create: `C:\Users\pyan\ai-generated\searchWebApi\go\binary_test.go`
- Create: `C:\Users\pyan\ai-generated\searchWebApi\go\measures_test.go`
- Create: `C:\Users\pyan\ai-generated\searchWebApi\go\insert_remove_test.go`
- Create: `C:\Users\pyan\ai-generated\searchWebApi\go\cmd\searchwebapi-example\main.go`

---

### Task 1: Bootstrap The Module And Root Client

**Files:**
- Create: `C:\Users\pyan\ai-generated\searchWebApi\go\go.mod`
- Create: `C:\Users\pyan\ai-generated\searchWebApi\go\client.go`
- Create: `C:\Users\pyan\ai-generated\searchWebApi\go\auth.go`
- Test: `C:\Users\pyan\ai-generated\searchWebApi\go\client_test.go`

- [ ] **Step 1: Write the failing root-client tests**

```go
package searchwebapi

import (
	"net/http"
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

func TestNewClientUsesProvidedHTTPClient(t *testing.T) {
	h := &http.Client{}
	client, err := NewClient(Config{BaseURL: "https://example.com", HTTPClient: h})
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}
	if client.httpClient != h {
		t.Fatal("expected custom http client to be used")
	}
}
```

- [ ] **Step 2: Run the test to verify it fails**

Run: `go test ./...`
Expected: FAIL with undefined `NewClient`, `Config`, and session methods.

- [ ] **Step 3: Write the minimal root client implementation**

```go
package searchwebapi

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

type Config struct {
	BaseURL     string
	Username    string
	Password    string
	BearerToken string
	SessionID   string
	SessionType string
	MDCToken    string
	MDCMethod   string
	HTTPClient  *http.Client
}

type Client struct {
	baseURL     *url.URL
	httpClient  *http.Client
	auth        authConfig
	sessionType string
	mdcToken    string
	mdcMethod   string

	mu        sync.RWMutex
	sessionID string
}

func NewClient(cfg Config) (*Client, error) {
	if strings.TrimSpace(cfg.BaseURL) == "" {
		return nil, errors.New("base URL is required")
	}

	parsed, err := url.Parse(strings.TrimSpace(cfg.BaseURL))
	if err != nil {
		return nil, err
	}

	parsed.Path = strings.TrimRight(parsed.Path, "/") + "/searchWebApi"
	parsed.RawQuery = ""
	parsed.Fragment = ""

	httpClient := cfg.HTTPClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	client := &Client{
		baseURL:    parsed,
		httpClient: httpClient,
		auth: authConfig{
			username:    cfg.Username,
			password:    cfg.Password,
			bearerToken: cfg.BearerToken,
		},
		sessionType: cfg.SessionType,
		mdcToken:    cfg.MDCToken,
		mdcMethod:   cfg.MDCMethod,
		sessionID:   cfg.SessionID,
	}

	return client, nil
}

func (c *Client) SessionID() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.sessionID
}

func (c *Client) SetSessionID(sessionID string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.sessionID = sessionID
}

func (c *Client) ClearSessionID() {
	c.SetSessionID("")
}
```

```go
package searchwebapi

type authConfig struct {
	username    string
	password    string
	bearerToken string
}
```

- [ ] **Step 4: Run the tests to verify the root client passes**

Run: `go test ./...`
Expected: PASS for `client_test.go`, remaining package failures still possible.

---

### Task 2: Add Shared Transport, Errors, And Wire Types

**Files:**
- Create: `C:\Users\pyan\ai-generated\searchWebApi\go\transport.go`
- Create: `C:\Users\pyan\ai-generated\searchWebApi\go\types.go`
- Modify: `C:\Users\pyan\ai-generated\searchWebApi\go\client_test.go`

- [ ] **Step 1: Write failing transport and header tests**

```go
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

	client, _ := NewClient(Config{BaseURL: ts.URL, Username: "user", Password: "pass", SessionID: "start"})
	var out LoginResult
	if err := client.doJSON(context.Background(), http.MethodPost, "/login", nil, nil, nil, &out); err != nil {
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
```

- [ ] **Step 2: Write minimal transport helpers and core types**

```go
package searchwebapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"path"
	"strings"
)

type ResponseError struct {
	StatusCode int
	Body       string
}

func (e *ResponseError) Error() string {
	return fmt.Sprintf("unexpected HTTP status %d", e.StatusCode)
}

func (c *Client) newURL(relPath string, query url.Values) string {
	u := *c.baseURL
	u.Path = path.Join(c.baseURL.Path, relPath)
	if query != nil {
		u.RawQuery = query.Encode()
	}
	return u.String()
}

func (c *Client) doJSON(ctx context.Context, method, relPath string, query url.Values, headers http.Header, body any, out any) error {
	var reader io.Reader
	if body != nil {
		buf := &bytes.Buffer{}
		if err := json.NewEncoder(buf).Encode(body); err != nil {
			return err
		}
		reader = buf
	}

	req, err := http.NewRequestWithContext(ctx, method, c.newURL(relPath, query), reader)
	if err != nil {
		return err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	c.applyHeaders(req, headers)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	c.captureSession(resp)

	if resp.StatusCode >= 400 {
		data, _ := io.ReadAll(resp.Body)
		return &ResponseError{StatusCode: resp.StatusCode, Body: string(data)}
	}
	if out == nil {
		_, _ = io.Copy(io.Discard, resp.Body)
		return nil
	}
	return json.NewDecoder(resp.Body).Decode(out)
}

func (c *Client) applyHeaders(req *http.Request, headers http.Header) {
	for key, values := range headers {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}
	if c.auth.username != "" || c.auth.password != "" {
		req.SetBasicAuth(c.auth.username, c.auth.password)
	}
	if c.auth.bearerToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.auth.bearerToken)
	}
	if sessionID := c.SessionID(); sessionID != "" {
		req.Header.Set("SWA-SESSION", sessionID)
	}
	if c.sessionType != "" {
		req.Header.Set("SWA-SESSION-TYPE", c.sessionType)
	}
	if c.mdcToken != "" {
		req.Header.Set("SWA-MDC-TOKEN", c.mdcToken)
	}
	if c.mdcMethod != "" {
		req.Header.Set("SWA-MDC-METHOD", c.mdcMethod)
	}
}

func (c *Client) captureSession(resp *http.Response) {
	if sessionID := strings.TrimSpace(resp.Header.Get("SWA-SESSION")); sessionID != "" {
		c.SetSessionID(sessionID)
	}
}

func writeMultipartJSONAndFiles(fieldName string, requestBody any, files [][]byte) (*bytes.Buffer, string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormField(fieldName)
	if err != nil {
		return nil, "", err
	}
	if err := json.NewEncoder(part).Encode(requestBody); err != nil {
		return nil, "", err
	}
	for i, data := range files {
		filePart, err := writer.CreateFormFile("binaries", fmt.Sprintf("binary-%d.bin", i))
		if err != nil {
			return nil, "", err
		}
		if _, err := filePart.Write(data); err != nil {
			return nil, "", err
		}
	}
	if err := writer.Close(); err != nil {
		return nil, "", err
	}
	return body, writer.FormDataContentType(), nil
}
```

```go
package searchwebapi

type StatusObject struct {
	Successful    bool   `json:"successful"`
	BackendStatus string `json:"backendStatus"`
	HTTPStatus    int    `json:"httpStatus"`
	ErrorMessage  string `json:"errorMessage"`
}

type LoginResult struct {
	Status StatusObject `json:"status"`
}

type LogoutResult struct {
	Status StatusObject `json:"status"`
}

type Project struct {
	ID string `json:"id"`
}

type ProjectsResult struct {
	Status        StatusObject `json:"status"`
	NumberResults int64        `json:"numberResults"`
	Results       []Project    `json:"results"`
}
```

- [ ] **Step 3: Run the tests to verify shared transport passes**

Run: `go test ./...`
Expected: PASS for header/session tests, remaining endpoint tests still pending.

---

### Task 3: Implement Session, Project, And Collection Operations

**Files:**
- Create: `C:\Users\pyan\ai-generated\searchWebApi\go\session.go`
- Create: `C:\Users\pyan\ai-generated\searchWebApi\go\projects.go`
- Create: `C:\Users\pyan\ai-generated\searchWebApi\go\collections.go`
- Create: `C:\Users\pyan\ai-generated\searchWebApi\go\projects_test.go`
- Create: `C:\Users\pyan\ai-generated\searchWebApi\go\collections_test.go`

- [ ] **Step 1: Write failing endpoint tests for session, projects, and collections**

```go
func TestListProjects(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/searchWebApi/projects" {
			t.Fatalf("path = %q", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":{"successful":true,"backendStatus":"ok","httpStatus":200,"errorMessage":""},"numberResults":1,"results":[{"id":"p1"}]}`))
	}))
	defer ts.Close()

	client, _ := NewClient(Config{BaseURL: ts.URL})
	result, err := client.ListProjects(context.Background())
	if err != nil {
		t.Fatalf("ListProjects() error = %v", err)
	}
	if len(result.Results) != 1 || result.Results[0].ID != "p1" {
		t.Fatalf("results = %+v", result.Results)
	}
}
```

- [ ] **Step 2: Add the minimal shared types and methods for these resources**

```go
package searchwebapi

import "context"

func (c *Client) Login(ctx context.Context) (*LoginResult, error) {
	var result LoginResult
	if err := c.doJSON(ctx, "POST", "/login", nil, nil, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) Logout(ctx context.Context) (*LogoutResult, error) {
	var result LogoutResult
	if err := c.doJSON(ctx, "DELETE", "/logout", nil, nil, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
```

```go
package searchwebapi

import "context"

func (c *Client) ListProjects(ctx context.Context) (*ProjectsResult, error) {
	var result ProjectsResult
	if err := c.doJSON(ctx, "GET", "/projects", nil, nil, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
```

```go
package searchwebapi

type ProjectResource struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

type ProjectResourcesResult struct {
	Status        StatusObject      `json:"status"`
	NumberResults int64             `json:"numberResults"`
	Results       []ProjectResource `json:"results"`
}

type Collection struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName,omitempty"`
}

type CollectionsResult struct {
	Status        StatusObject `json:"status"`
	NumberResults int64        `json:"numberResults"`
	Results       []Collection `json:"results"`
}
```

- [ ] **Step 3: Expand collections coverage with option structs and methods**

```go
type FolderValuesOptions struct {
	Query                  string
	Language               string
	JoinRestriction        string
	Prefix                 string
	RestrictFoldersByQuery string
	ReturnEmptyFolders     *bool
	Order                  string
	Offset                 *int
	Limit                  *int
	SearchCacheControl     string
}

func (c *Client) ListCollections(ctx context.Context, projectID string) (*CollectionsResult, error) {
	var result CollectionsResult
	if err := c.doJSON(ctx, "GET", "/projects/"+projectID+"/collections", nil, nil, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
```

- [ ] **Step 4: Run the targeted tests**

Run: `go test ./...`
Expected: PASS for session, projects, and collection operations.

---

### Task 4: Implement Record Search, Tokens, Fetch, Changes, And NDJSON Streaming

**Files:**
- Create: `C:\Users\pyan\ai-generated\searchWebApi\go\records.go`
- Create: `C:\Users\pyan\ai-generated\searchWebApi\go\records_test.go`

- [ ] **Step 1: Write failing tests for record search and stream decoding**

```go
func TestSearchRecords(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("query") != "hello" {
			t.Fatalf("query = %q", r.URL.Query().Get("query"))
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":{"successful":true,"backendStatus":"ok","httpStatus":200,"errorMessage":""},"numberResults":1,"results":[{"rank":1,"relevance":1,"id":"r1","uniqueField":"u1","fields":[],"folderSets":[],"body":"","page":0,"pageCount":0}]}`))
	}))
	defer ts.Close()

	client, _ := NewClient(Config{BaseURL: ts.URL})
	result, err := client.SearchRecords(context.Background(), "p1", "default", SearchRecordsOptions{Query: "hello"})
	if err != nil {
		t.Fatalf("SearchRecords() error = %v", err)
	}
	if len(result.Results) != 1 || result.Results[0].ID != "r1" {
		t.Fatalf("results = %+v", result.Results)
	}
}
```

```go
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

	client, _ := NewClient(Config{BaseURL: ts.URL})
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
}
```

- [ ] **Step 2: Add the shared record types and search methods**

```go
type Field struct {
	ID          string `json:"id"`
	Value       string `json:"value"`
	ValueObject any    `json:"valueObject,omitempty"`
}

type Folder struct {
	ID          string  `json:"id"`
	DisplayName string  `json:"displayName"`
	Properties  []Field `json:"properties,omitempty"`
}

type FolderSet struct {
	ID    string   `json:"id"`
	Value []Folder `json:"value"`
}

type Record struct {
	Rank       int64       `json:"rank"`
	Relevance  float32     `json:"relevance"`
	ID         string      `json:"id"`
	UniqueField string     `json:"uniqueField"`
	Fields     []Field     `json:"fields"`
	FolderSets []FolderSet `json:"folderSets"`
	Body       string      `json:"body"`
	Page       int         `json:"page"`
	PageCount  int         `json:"pageCount"`
}

type SearchResult struct {
	Status              StatusObject               `json:"status"`
	NumberResults       int64                      `json:"numberResults"`
	Results             []Record                   `json:"results"`
	SponsoredLinks      []SponsoredLink            `json:"sponsoredLinks,omitempty"`
	SpellingSuggestions *SpellingSuggestionResult  `json:"spellingSuggestions,omitempty"`
}
```

- [ ] **Step 3: Add token, fetch, change, and in-document methods**

```go
func (c *Client) FetchRecordContent(ctx context.Context, projectID, collectionID, recordID string, opts FetchRecordContentOptions) (*Record, error) {
	var result Record
	if err := c.doJSON(ctx, "GET", "/projects/"+projectID+"/collections/"+collectionID+"/records/"+recordID+"/content", opts.values(), nil, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) ChangeRecordContent(ctx context.Context, projectID, collectionID, recordID string, requests []ChangeRequest) (*ChangeResult, error) {
	var result ChangeResult
	if err := c.doJSON(ctx, "PUT", "/projects/"+projectID+"/collections/"+collectionID+"/records/"+recordID+"/content", nil, nil, requests, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
```

- [ ] **Step 4: Run the record tests**

Run: `go test ./...`
Expected: PASS for JSON search, NDJSON streaming, token lifecycle, and fetch/change flows.

---

### Task 5: Implement Binary, Measures, Cached Searches, Change Queue, And Insert/Remove

**Files:**
- Create: `C:\Users\pyan\ai-generated\searchWebApi\go\binary.go`
- Create: `C:\Users\pyan\ai-generated\searchWebApi\go\measures.go`
- Create: `C:\Users\pyan\ai-generated\searchWebApi\go\cached_searches.go`
- Create: `C:\Users\pyan\ai-generated\searchWebApi\go\change_queue.go`
- Create: `C:\Users\pyan\ai-generated\searchWebApi\go\insert_remove.go`
- Create: `C:\Users\pyan\ai-generated\searchWebApi\go\binary_test.go`
- Create: `C:\Users\pyan\ai-generated\searchWebApi\go\measures_test.go`
- Create: `C:\Users\pyan\ai-generated\searchWebApi\go\insert_remove_test.go`

- [ ] **Step 1: Write failing binary and multipart tests**

```go
func TestGetBinaryByRecordID(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/octet-stream")
		_, _ = w.Write([]byte("hello"))
	}))
	defer ts.Close()

	client, _ := NewClient(Config{BaseURL: ts.URL})
	result, err := client.GetBinaryByRecordID(context.Background(), "p1", "default", "r1", "native")
	if err != nil {
		t.Fatalf("GetBinaryByRecordID() error = %v", err)
	}
	defer result.Body.Close()
	data, _ := io.ReadAll(result.Body)
	if string(data) != "hello" {
		t.Fatalf("body = %q", string(data))
	}
}
```

- [ ] **Step 2: Add binary response wrapper and remaining resource methods**

```go
type BinaryResponse struct {
	Header      http.Header
	ContentType string
	Body        io.ReadCloser
}

func (c *Client) GetBinaryByRecordID(ctx context.Context, projectID, collectionID, recordID, field string) (*BinaryResponse, error) {
	query := url.Values{}
	query.Set("field", field)
	return c.doBinary(ctx, "/projects/"+projectID+"/collections/"+collectionID+"/binary/"+recordID+"/content", query)
}
```

- [ ] **Step 3: Add measure, cache, queue, and insert/remove types and methods**

```go
type ChangeResult struct {
	Status StatusObject `json:"status"`
}

type InsertRemoveResult struct {
	Status StatusObject `json:"status"`
}

type WaitForPendingChangesResult struct {
	Success bool `json:"success"`
}

func (c *Client) WaitForPendingChanges(ctx context.Context, projectID, collectionID string, opts WaitForPendingChangesOptions) (*WaitForPendingChangesResult, error) {
	var result WaitForPendingChangesResult
	if err := c.doJSON(ctx, "GET", "/projects/"+projectID+"/collections/"+collectionID+"/changes/queue", opts.values(), nil, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
```

- [ ] **Step 4: Run the resource tests**

Run: `go test ./...`
Expected: PASS for binary download, measure, queue, cache, and insert/remove coverage.

---

### Task 6: Add The Example CLI And README

**Files:**
- Create: `C:\Users\pyan\ai-generated\searchWebApi\go\cmd\searchwebapi-example\main.go`
- Create: `C:\Users\pyan\ai-generated\searchWebApi\go\README.md`

- [ ] **Step 1: Write a failing CLI build check**

Run: `go build ./cmd/searchwebapi-example`
Expected: FAIL because the CLI entrypoint does not exist yet.

- [ ] **Step 2: Add the small example CLI**

```go
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	searchwebapi "searchwebapi"
)

func main() {
	baseURL := flag.String("base-url", envOr("SEARCHWEBAPI_BASE_URL", ""), "Base URL")
	username := flag.String("username", envOr("SEARCHWEBAPI_USERNAME", ""), "Username")
	password := flag.String("password", envOr("SEARCHWEBAPI_PASSWORD", ""), "Password")
	project := flag.String("project", "", "Project ID")
	collection := flag.String("collection", "default", "Collection ID")
	query := flag.String("query", "*", "Search query")
	command := flag.String("command", "projects-list", "projects-list or records-search")
	flag.Parse()

	client, err := searchwebapi.NewClient(searchwebapi.Config{
		BaseURL:  *baseURL,
		Username: *username,
		Password: *password,
	})
	if err != nil {
		fatal(err)
	}

	ctx := context.Background()
	switch *command {
	case "projects-list":
		result, err := client.ListProjects(ctx)
		fatalIf(err)
		printJSON(result)
	case "records-search":
		result, err := client.SearchRecords(ctx, *project, *collection, searchwebapi.SearchRecordsOptions{Query: *query})
		fatalIf(err)
		printJSON(result)
	default:
		fatal(fmt.Errorf("unknown command %q", *command))
	}
}
```

- [ ] **Step 3: Add the README**

```md
# searchwebapi

Generated Go client for the `searchWebApi` contract.

## Build

```bash
go test ./...
go build ./...
```

## Example CLI

```bash
go run ./cmd/searchwebapi-example --base-url https://example.test --command projects-list
go run ./cmd/searchwebapi-example --base-url https://example.test --command records-search --project my-project --collection default --query hello
```
```

- [ ] **Step 4: Run the CLI and package build checks**

Run: `go build ./...`
Expected: PASS.

---

### Task 7: Final Verification

**Files:**
- Verify: `C:\Users\pyan\ai-generated\searchWebApi\go\*.go`
- Verify: `C:\Users\pyan\ai-generated\searchWebApi\go\cmd\searchwebapi-example\main.go`

- [ ] **Step 1: Run formatting**

Run: `gofmt -w *.go cmd/searchwebapi-example/main.go`
Expected: files are formatted in place.

- [ ] **Step 2: Run the full test suite**

Run: `go test ./...`
Expected: PASS.

- [ ] **Step 3: Run the full build**

Run: `go build ./...`
Expected: PASS.

- [ ] **Step 4: Inspect the module tree**

Run: `go list ./...`
Expected: at least the root package and `cmd/searchwebapi-example`.

- [ ] **Step 5: Commit if and only if explicitly requested by the user**

```bash
git add docs/superpowers/plans/2026-04-16-searchwebapi-go-client-implementation.md
git commit -m "docs: add searchWebApi Go client plan"
```
