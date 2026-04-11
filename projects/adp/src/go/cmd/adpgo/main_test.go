package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestListEntitiesCommandPrintsDecodedEntities(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/executeAdpTask" {
			t.Fatalf("path = %q, want /executeAdpTask", r.URL.Path)
		}

		var req map[string]any
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("decode request: %v", err)
		}

		if got := req["taskType"]; got != "List Entities" {
			t.Fatalf("taskType = %#v", got)
		}

		cfg, ok := req["taskConfiguration"].(map[string]any)
		if !ok {
			t.Fatalf("taskConfiguration type = %T", req["taskConfiguration"])
		}
		if got := cfg["adp_listEntities_type"]; got != "singleMindServer" {
			t.Fatalf("adp_listEntities_type = %#v", got)
		}

		_, _ = io.WriteString(w, `{"executionId":"1","taskType":"List Entities","loggingEnabled":"true","progressMax":1,"executionStatus":"success","executionRootDir":"root","contextId":"ctx","executionPersistent":"true","progressCurrent":1,"progressPercentage":1.0,"taskDisplayName":"List entities","executionMetaData":{"adp_entities_output_file_name":"output.json","adp_entities_json_output":"[{\"id\":\"entity-1\",\"displayName\":\"Entity One\"}]"}}`)
	}))
	defer server.Close()

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd := newApp(stdout, stderr)

	err := cmd.Run(context.Background(), []string{
		"adpgo",
		"--host", server.URL,
		"--path", "",
		"--user", "adp",
		"--password", "secret",
		"list-entities",
		"--type", "singleMindServer",
	})
	if err != nil {
		t.Fatalf("Run error: %v", err)
	}

	if stderr.Len() != 0 {
		t.Fatalf("stderr = %q", stderr.String())
	}

	if got, want := strings.TrimSpace(stdout.String()), "[\n  {\n    \"id\": \"entity-1\",\n    \"displayName\": \"Entity One\"\n  }\n]"; got != want {
		t.Fatalf("stdout = %q, want %q", got, want)
	}
}

func TestQueryEngineCommandParsesTaxonomyFlags(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req map[string]any
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("decode request: %v", err)
		}

		cfg, ok := req["taskConfiguration"].(map[string]any)
		if !ok {
			t.Fatalf("taskConfiguration type = %T", req["taskConfiguration"])
		}

		rawTaxonomies, ok := cfg["adp_queryEngine_engineTaxonomies"].([]any)
		if !ok {
			t.Fatalf("engine taxonomies type = %T", cfg["adp_queryEngine_engineTaxonomies"])
		}
		if len(rawTaxonomies) != 2 {
			t.Fatalf("len(engine taxonomies) = %d, want 2", len(rawTaxonomies))
		}

		first, ok := rawTaxonomies[0].(map[string]any)
		if !ok {
			t.Fatalf("first taxonomy type = %T", rawTaxonomies[0])
		}
		second, ok := rawTaxonomies[1].(map[string]any)
		if !ok {
			t.Fatalf("second taxonomy type = %T", rawTaxonomies[1])
		}

		if first["taxonomy"] != "rm_source" || first["query"] != "email" || first["negation"] != false {
			t.Fatalf("first taxonomy = %#v", first)
		}
		if second["taxonomy"] != "rm_mimetype" || second["query"] != "pdf" || second["negation"] != true {
			t.Fatalf("second taxonomy = %#v", second)
		}

		_, _ = io.WriteString(w, `{"executionId":"2","taskType":"Query Engine","loggingEnabled":"true","progressMax":1,"executionStatus":"success","executionRootDir":"root","contextId":"ctx","executionPersistent":"true","progressCurrent":1,"progressPercentage":1.0,"taskDisplayName":"Query engine","executionMetaData":{"adp_query_engine_documents_count":"100","adp_query_engine_aggregated_value":"500"}}`)
	}))
	defer server.Close()

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd := newApp(stdout, stderr)

	err := cmd.Run(context.Background(), []string{
		"adpgo",
		"--host", server.URL,
		"--path", "",
		"--user", "adp",
		"--password", "secret",
		"query-engine",
		"--engineName", "myEngine",
		"--engineTaxonomies", "rm_source=email",
		"--engineTaxonomies", "rm_mimetype!=pdf",
	})
	if err != nil {
		t.Fatalf("Run error: %v", err)
	}

	if stderr.Len() != 0 {
		t.Fatalf("stderr = %q", stderr.String())
	}

	if got, want := strings.TrimSpace(stdout.String()), "{\n  \"documentsCount\": 100,\n  \"aggregatedValue\": \"500\"\n}"; got != want {
		t.Fatalf("stdout = %q, want %q", got, want)
	}
}

func TestCommandFailureFormatsTaskExecutionError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, `{"executionId":"exec-123","taskType":"List Entities","loggingEnabled":"true","progressMax":1,"executionStatus":"failed","executionRootDir":"root","contextId":"ctx","executionPersistent":"true","progressCurrent":0,"progressPercentage":0.0,"taskDisplayName":"List entities","errorMessage":"boom","executionMetaData":null}`)
	}))
	defer server.Close()

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd := newApp(stdout, stderr)

	err := cmd.Run(context.Background(), []string{
		"adpgo",
		"--host", server.URL,
		"--path", "",
		"--user", "adp",
		"--password", "secret",
		"list-entities",
	})
	if err == nil {
		t.Fatal("expected error")
	}

	if stdout.Len() != 0 {
		t.Fatalf("stdout = %q", stdout.String())
	}

	text := stderr.String()
	for _, want := range []string{"Error: boom", "ExecutionID: exec-123", "TaskType: List Entities"} {
		if !strings.Contains(text, want) {
			t.Fatalf("stderr = %q, missing %q", text, want)
		}
	}
}
