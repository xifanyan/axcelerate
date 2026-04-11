package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
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

func TestListEntitiesCommandReadsGlobalConfigFile(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/executeAdpTask" {
			t.Fatalf("path = %q, want /executeAdpTask", r.URL.Path)
		}

		if got, want := r.Header.Get("Auth-Username"), "adp"; got != want {
			t.Fatalf("Auth-Username = %q, want %q", got, want)
		}
		if got, want := r.Header.Get("Auth-Password"), "secret"; got != want {
			t.Fatalf("Auth-Password = %q, want %q", got, want)
		}

		var req map[string]any
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("decode request: %v", err)
		}

		if got := req["taskType"]; got != "List Entities" {
			t.Fatalf("taskType = %#v", got)
		}

		_, _ = io.WriteString(w, `{"executionId":"1","taskType":"List Entities","loggingEnabled":"true","progressMax":1,"executionStatus":"success","executionRootDir":"root","contextId":"ctx","executionPersistent":"true","progressCurrent":1,"progressPercentage":1.0,"taskDisplayName":"List entities","executionMetaData":{"adp_entities_output_file_name":"output.json","adp_entities_json_output":"[]"}}`)
	}))
	defer server.Close()

	dir := t.TempDir()
	writeADPConfigFile(t, dir, map[string]any{
		"host":     server.URL,
		"user":     "adp",
		"password": "secret",
		"path":     "",
	})

	withWorkingDir(t, dir)

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd := newApp(stdout, stderr)

	err := cmd.Run(context.Background(), []string{
		"adpgo",
		"list-entities",
	})
	if err != nil {
		t.Fatalf("Run error: %v", err)
	}

	if stderr.Len() != 0 {
		t.Fatalf("stderr = %q", stderr.String())
	}

	if got, want := strings.TrimSpace(stdout.String()), "[]"; got != want {
		t.Fatalf("stdout = %q, want %q", got, want)
	}
}

func TestListEntitiesCommandCLIFlagsOverrideConfigFile(t *testing.T) {
	configServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Fatalf("unexpected request to config-backed server: %s", r.URL.Path)
	}))
	defer configServer.Close()

	cliServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/executeAdpTask" {
			t.Fatalf("path = %q, want /executeAdpTask", r.URL.Path)
		}

		if got, want := r.Header.Get("Auth-Username"), "cli-user"; got != want {
			t.Fatalf("Auth-Username = %q, want %q", got, want)
		}
		if got, want := r.Header.Get("Auth-Password"), "cli-secret"; got != want {
			t.Fatalf("Auth-Password = %q, want %q", got, want)
		}

		_, _ = io.WriteString(w, `{"executionId":"1","taskType":"List Entities","loggingEnabled":"true","progressMax":1,"executionStatus":"success","executionRootDir":"root","contextId":"ctx","executionPersistent":"true","progressCurrent":1,"progressPercentage":1.0,"taskDisplayName":"List entities","executionMetaData":{"adp_entities_output_file_name":"output.json","adp_entities_json_output":"[]"}}`)
	}))
	defer cliServer.Close()

	dir := t.TempDir()
	writeADPConfigFile(t, dir, map[string]any{
		"host":     configServer.URL,
		"user":     "config-user",
		"password": "config-secret",
		"path":     "/wrong-base",
	})

	withWorkingDir(t, dir)

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd := newApp(stdout, stderr)

	err := cmd.Run(context.Background(), []string{
		"adpgo",
		"--host", cliServer.URL,
		"--path", "",
		"--user", "cli-user",
		"--password", "cli-secret",
		"list-entities",
	})
	if err != nil {
		t.Fatalf("Run error: %v", err)
	}

	if stderr.Len() != 0 {
		t.Fatalf("stderr = %q", stderr.String())
	}

	if got, want := strings.TrimSpace(stdout.String()), "[]"; got != want {
		t.Fatalf("stdout = %q, want %q", got, want)
	}
}

func TestRunReportsInvalidConfigFile(t *testing.T) {
	dir := t.TempDir()
	configPath := filepath.Join(dir, "adp_config.json")
	if err := os.WriteFile(configPath, []byte(`{"host":`), 0o644); err != nil {
		t.Fatalf("write config: %v", err)
	}

	withWorkingDir(t, dir)

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	exitCode := run(stdout, stderr, []string{
		"adpgo",
		"list-entities",
	})

	if exitCode != 1 {
		t.Fatalf("exitCode = %d, want 1", exitCode)
	}
	if stdout.Len() != 0 {
		t.Fatalf("stdout = %q", stdout.String())
	}
	if got := stderr.String(); !strings.Contains(got, "adp_config.json") || !strings.Contains(strings.ToLower(got), "invalid") {
		t.Fatalf("stderr = %q", got)
	}
}

func TestRunReportsMissingRequiredGlobalsAfterConfigResolution(t *testing.T) {
	dir := t.TempDir()
	writeADPConfigFile(t, dir, map[string]any{
		"host": "https://example.test",
		"path": "",
	})

	withWorkingDir(t, dir)

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	exitCode := run(stdout, stderr, []string{
		"adpgo",
		"list-entities",
	})

	if exitCode != 1 {
		t.Fatalf("exitCode = %d, want 1", exitCode)
	}
	if stdout.Len() != 0 {
		t.Fatalf("stdout = %q", stdout.String())
	}
	text := stderr.String()
	for _, want := range []string{"user", "password"} {
		if !strings.Contains(text, want) {
			t.Fatalf("stderr = %q, missing %q", text, want)
		}
	}
	if strings.Contains(text, "host") {
		t.Fatalf("stderr = %q, unexpectedly mentions host", text)
	}
}

func TestRunReportsMalformedHostFromConfig(t *testing.T) {
	dir := t.TempDir()
	writeADPConfigFile(t, dir, map[string]any{
		"host":     "http://[::1",
		"user":     "adp",
		"password": "secret",
	})

	withWorkingDir(t, dir)

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	exitCode := run(stdout, stderr, []string{
		"adpgo",
		"list-entities",
	})

	if exitCode != 1 {
		t.Fatalf("exitCode = %d, want 1", exitCode)
	}
	if stdout.Len() != 0 {
		t.Fatalf("stdout = %q", stdout.String())
	}
	if got := strings.ToLower(stderr.String()); !strings.Contains(got, "host") {
		t.Fatalf("stderr = %q", stderr.String())
	}
}

func TestListEntitiesCommandCLIPathOverridesConfigFile(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/executeAdpTask" {
			t.Fatalf("path = %q, want /executeAdpTask", r.URL.Path)
		}

		_, _ = io.WriteString(w, `{"executionId":"1","taskType":"List Entities","loggingEnabled":"true","progressMax":1,"executionStatus":"success","executionRootDir":"root","contextId":"ctx","executionPersistent":"true","progressCurrent":1,"progressPercentage":1.0,"taskDisplayName":"List entities","executionMetaData":{"adp_entities_output_file_name":"output.json","adp_entities_json_output":"[]"}}`)
	}))
	defer server.Close()

	dir := t.TempDir()
	writeADPConfigFile(t, dir, map[string]any{
		"host":     server.URL,
		"user":     "adp",
		"password": "secret",
		"path":     "/wrong-base",
	})

	withWorkingDir(t, dir)

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd := newApp(stdout, stderr)

	err := cmd.Run(context.Background(), []string{
		"adpgo",
		"--path", "",
		"list-entities",
	})
	if err != nil {
		t.Fatalf("Run error: %v", err)
	}

	if stderr.Len() != 0 {
		t.Fatalf("stderr = %q", stderr.String())
	}

	if got, want := strings.TrimSpace(stdout.String()), "[]"; got != want {
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

func TestQueryEngineCommandRequiresExactlyOneSelector(t *testing.T) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd := newApp(stdout, stderr)

	err := cmd.Run(context.Background(), []string{
		"adpgo",
		"--host", "https://example.test",
		"--user", "adp",
		"--password", "secret",
		"query-engine",
	})
	if err == nil {
		t.Fatal("expected error")
	}
	if !strings.Contains(err.Error(), "exactly one of engineName or applicationIdentifier is required") {
		t.Fatalf("error = %q", err)
	}
}

func TestQueryEngineCommandRejectsBothSelectors(t *testing.T) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd := newApp(stdout, stderr)

	err := cmd.Run(context.Background(), []string{
		"adpgo",
		"--host", "https://example.test",
		"--user", "adp",
		"--password", "secret",
		"query-engine",
		"--engineName", "engineA",
		"--applicationIdentifier", "appA",
	})
	if err == nil {
		t.Fatal("expected error")
	}
	if !strings.Contains(err.Error(), "engineName and applicationIdentifier are mutually exclusive") {
		t.Fatalf("error = %q", err)
	}
}

func TestQueryEngineCommandAllowsApplicationIdentifier(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg := decodeTaskConfiguration(t, r)
		assertApplicationSelectorOnly(t, cfg, "adp_queryEngine_applicationIdentifier", "adp_queryEngine_engineName")

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
		"--applicationIdentifier", "appA",
	})
	if err != nil {
		t.Fatalf("Run error: %v", err)
	}
	if stderr.Len() != 0 {
		t.Fatalf("stderr = %q", stderr.String())
	}
	if !strings.Contains(stdout.String(), `"documentsCount": 100`) {
		t.Fatalf("stdout = %q", stdout.String())
	}
	if strings.Contains(stdout.String(), "engineName") {
		t.Fatalf("stdout = %q", stdout.String())
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

func TestCSVMergeCommandParsesFieldMappingsJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req map[string]any
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("decode request: %v", err)
		}

		cfg, ok := req["taskConfiguration"].(map[string]any)
		if !ok {
			t.Fatalf("taskConfiguration type = %T", req["taskConfiguration"])
		}

		fieldMappings, ok := cfg["adp_csvMerge_fieldMappings"].([]any)
		if !ok {
			t.Fatalf("fieldMappings type = %T", cfg["adp_csvMerge_fieldMappings"])
		}
		if len(fieldMappings) != 1 {
			t.Fatalf("len(fieldMappings) = %d, want 1", len(fieldMappings))
		}

		mapping, ok := fieldMappings[0].(map[string]any)
		if !ok {
			t.Fatalf("mapping type = %T", fieldMappings[0])
		}
		if mapping["csvField"] != "id" || mapping["targetField"] != "rm_id" {
			t.Fatalf("mapping = %#v", mapping)
		}

		_, _ = io.WriteString(w, `{"executionId":"3","taskType":"CSV Merge","loggingEnabled":"true","progressMax":1,"executionStatus":"success","executionRootDir":"root","contextId":"ctx","executionPersistent":"true","progressCurrent":1,"progressPercentage":1.0,"taskDisplayName":"Csv merge task","executionMetaData":[]}`)
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
		"csv-merge",
		"--csvFile", "input.csv",
		"--fieldMappings", `[{"csvField":"id","targetField":"rm_id"}]`,
	})
	if err != nil {
		t.Fatalf("Run error: %v", err)
	}

	if stderr.Len() != 0 {
		t.Fatalf("stderr = %q", stderr.String())
	}

	if got, want := strings.TrimSpace(stdout.String()), "{}"; got != want {
		t.Fatalf("stdout = %q, want %q", got, want)
	}
}

func TestTaxonomyStatisticCommandRequiresExactlyOneSelector(t *testing.T) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd := newApp(stdout, stderr)

	err := cmd.Run(context.Background(), []string{
		"adpgo",
		"--host", "https://example.test",
		"--user", "adp",
		"--password", "secret",
		"taxonomy-statistic",
	})
	if err == nil {
		t.Fatal("expected error")
	}
	if !strings.Contains(err.Error(), "exactly one of engineName or applicationIdentifier is required") {
		t.Fatalf("error = %q", err)
	}
}

func TestTaxonomyStatisticCommandRejectsBothSelectors(t *testing.T) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd := newApp(stdout, stderr)

	err := cmd.Run(context.Background(), []string{
		"adpgo",
		"--host", "https://example.test",
		"--user", "adp",
		"--password", "secret",
		"taxonomy-statistic",
		"--engineName", "engineA",
		"--applicationIdentifier", "appA",
	})
	if err == nil {
		t.Fatal("expected error")
	}
	if !strings.Contains(err.Error(), "engineName and applicationIdentifier are mutually exclusive") {
		t.Fatalf("error = %q", err)
	}
}

func TestTaxonomyStatisticCommandAllowsApplicationIdentifier(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg := decodeTaskConfiguration(t, r)
		assertApplicationSelectorOnly(t, cfg, "adp_taxonomyStatistic_applicationIdentifier", "adp_taxonomyStatistic_engineName")

		_, _ = io.WriteString(w, `{"executionId":"7","taskType":"Taxonomy Statistic","loggingEnabled":"true","progressMax":1,"executionStatus":"success","executionRootDir":"root","contextId":"ctx","executionPersistent":"true","progressCurrent":1,"progressPercentage":1.0,"taskDisplayName":"Taxonomy statistic","executionMetaData":{"adp_taxonomy_statistics_json_file_path":"taxonomy_stats.json","adp_taxonomy_statistics_json_output":"{\"date\":\"Wed\",\"searchParameter\":[],\"statistics\":{\"taxonomy\":[{\"id\":\"rm_source\",\"category\":[{\"id\":\"file_demo_04\",\"displayName\":\"file_demo_04\",\"count\":761}]}]}}"}}`)
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
		"taxonomy-statistic",
		"--applicationIdentifier", "appA",
	})
	if err != nil {
		t.Fatalf("Run error: %v", err)
	}
	if stderr.Len() != 0 {
		t.Fatalf("stderr = %q", stderr.String())
	}
	if !strings.Contains(stdout.String(), `"rm_source"`) {
		t.Fatalf("stdout = %q", stdout.String())
	}
}

func TestQueryEngineCommandRequiresEngineName(t *testing.T) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd := newApp(stdout, stderr)

	err := cmd.Run(context.Background(), []string{
		"adpgo",
		"--host", "https://example.test",
		"--user", "adp",
		"--password", "secret",
		"query-engine",
	})
	if err == nil {
		t.Fatal("expected error")
	}

	if !strings.Contains(err.Error(), "Required flag \"engineName\" not set") {
		t.Fatalf("error = %q", err)
	}
}

func TestTaxonomyStatisticCommandRequiresEngineName(t *testing.T) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd := newApp(stdout, stderr)

	err := cmd.Run(context.Background(), []string{
		"adpgo",
		"--host", "https://example.test",
		"--user", "adp",
		"--password", "secret",
		"taxonomy-statistic",
	})
	if err == nil {
		t.Fatal("expected error")
	}

	if !strings.Contains(err.Error(), "Required flag \"engineName\" not set") {
		t.Fatalf("error = %q", err)
	}
}

func TestRunPrintsParserErrorsToStderr(t *testing.T) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	exitCode := run(stdout, stderr, []string{
		"adpgo",
		"--host", "https://example.test",
		"--user", "adp",
		"--password", "secret",
		"query-engine",
	})

	if exitCode != 1 {
		t.Fatalf("exitCode = %d, want 1", exitCode)
	}
	if !strings.Contains(stderr.String(), "Required flag \"engineName\" not set") {
		t.Fatalf("stderr = %q", stderr.String())
	}
}

func TestRunPrintsTaxonomyStatisticParserErrorsToStderr(t *testing.T) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	exitCode := run(stdout, stderr, []string{
		"adpgo",
		"--host", "https://example.test",
		"--user", "adp",
		"--password", "secret",
		"taxonomy-statistic",
	})

	if exitCode != 1 {
		t.Fatalf("exitCode = %d, want 1", exitCode)
	}
	if !strings.Contains(stderr.String(), "Required flag \"engineName\" not set") {
		t.Fatalf("stderr = %q", stderr.String())
	}
}

func TestRunDoesNotDuplicateTaskExecutionErrors(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, `{"executionId":"exec-123","taskType":"List Entities","loggingEnabled":"true","progressMax":1,"executionStatus":"failed","executionRootDir":"root","contextId":"ctx","executionPersistent":"true","progressCurrent":0,"progressPercentage":0.0,"taskDisplayName":"List entities","errorMessage":"boom","executionMetaData":null}`)
	}))
	defer server.Close()

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	exitCode := run(stdout, stderr, []string{
		"adpgo",
		"--host", server.URL,
		"--path", "",
		"--user", "adp",
		"--password", "secret",
		"list-entities",
	})

	if exitCode != 1 {
		t.Fatalf("exitCode = %d, want 1", exitCode)
	}
	if stdout.Len() != 0 {
		t.Fatalf("stdout = %q", stdout.String())
	}

	got := stderr.String()
	const want = "Error: boom\nExecutionID: exec-123\nTaskType: List Entities\n"
	if got != want {
		t.Fatalf("stderr = %q, want %q", got, want)
	}
}

func TestCreateOcrJobCommandStartsWithoutWaitingByDefault(t *testing.T) {
	requests := make([]string, 0, 2)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests = append(requests, r.URL.Path)
		_, _ = io.WriteString(w, `{"executionId":"ocr-123","taskType":"Create OCR Job","loggingEnabled":"true","progressMax":0,"executionStatus":"","executionRootDir":"root","contextId":"ctx","executionPersistent":"true","progressCurrent":0,"progressPercentage":0.0,"taskDisplayName":"Create OCR Job","executionMetaData":[]}`)
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
		"create-ocr-job",
		"--engineName", "engineA",
	})
	if err != nil {
		t.Fatalf("Run error: %v", err)
	}
	if stderr.Len() != 0 {
		t.Fatalf("stderr = %q", stderr.String())
	}
	if len(requests) != 1 || requests[0] != "/executeAdpTaskAsync" {
		t.Fatalf("requests = %#v", requests)
	}
	if !strings.Contains(stdout.String(), `"executionId": "ocr-123"`) {
		t.Fatalf("stdout = %q", stdout.String())
	}
}

func TestCreateOcrJobCommandWaitsWhenRequested(t *testing.T) {
	requests := make([]string, 0, 2)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests = append(requests, r.URL.Path)
		if len(requests) == 1 {
			_, _ = io.WriteString(w, `{"executionId":"ocr-123","taskType":"Create OCR Job","loggingEnabled":"true","progressMax":0,"executionStatus":"","executionRootDir":"root","contextId":"ctx","executionPersistent":"true","progressCurrent":0,"progressPercentage":0.0,"taskDisplayName":"Create OCR Job","executionMetaData":[]}`)
			return
		}
		_, _ = io.WriteString(w, `{"executionId":"ocr-123","taskType":"Create OCR Job","loggingEnabled":"true","progressMax":1,"executionStatus":"success","executionRootDir":"root","contextId":"ctx","executionPersistent":"true","progressCurrent":1,"progressPercentage":1.0,"taskDisplayName":"Create OCR Job","executionMetaData":[]}`)
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
		"create-ocr-job",
		"--engineName", "engineA",
		"--wait",
	})
	if err != nil {
		t.Fatalf("Run error: %v", err)
	}
	if stderr.Len() != 0 {
		t.Fatalf("stderr = %q", stderr.String())
	}
	if len(requests) != 2 || requests[0] != "/executeAdpTaskAsync" || requests[1] != "/statusAndProgress" {
		t.Fatalf("requests = %#v", requests)
	}
	if got, want := strings.TrimSpace(stdout.String()), "{}"; got != want {
		t.Fatalf("stdout = %q, want %q", got, want)
	}
}

func TestCreateOcrJobCommandRequiresExactlyOneSelector(t *testing.T) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd := newApp(stdout, stderr)

	err := cmd.Run(context.Background(), []string{
		"adpgo",
		"--host", "https://example.test",
		"--user", "adp",
		"--password", "secret",
		"create-ocr-job",
	})
	if err == nil {
		t.Fatal("expected error")
	}
	if !strings.Contains(err.Error(), "exactly one of engineName or applicationIdentifier is required") {
		t.Fatalf("error = %q", err)
	}
}

func TestCreateOcrJobCommandRejectsBothSelectors(t *testing.T) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd := newApp(stdout, stderr)

	err := cmd.Run(context.Background(), []string{
		"adpgo",
		"--host", "https://example.test",
		"--user", "adp",
		"--password", "secret",
		"create-ocr-job",
		"--engineName", "engineA",
		"--applicationIdentifier", "appA",
	})
	if err == nil {
		t.Fatal("expected error")
	}
	if !strings.Contains(err.Error(), "engineName and applicationIdentifier are mutually exclusive") {
		t.Fatalf("error = %q", err)
	}
}

func TestCreateOcrJobCommandAllowsApplicationIdentifierWithoutEngineName(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/executeAdpTaskAsync" {
			t.Fatalf("path = %q, want /executeAdpTaskAsync", r.URL.Path)
		}
		cfg := decodeTaskConfiguration(t, r)
		assertApplicationSelectorOnly(t, cfg, "adp_createOcrJob_applicationIdentifier", "adp_createOcrJob_engineName")

		_, _ = io.WriteString(w, `{"executionId":"ocr-123","taskType":"Create OCR Job","loggingEnabled":"true","progressMax":0,"executionStatus":"","executionRootDir":"root","contextId":"ctx","executionPersistent":"true","progressCurrent":0,"progressPercentage":0.0,"taskDisplayName":"Create OCR Job","executionMetaData":[]}`)
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
		"create-ocr-job",
		"--applicationIdentifier", "appA",
	})
	if err != nil {
		t.Fatalf("Run error: %v", err)
	}
	if stderr.Len() != 0 {
		t.Fatalf("stderr = %q", stderr.String())
	}
	if !strings.Contains(stdout.String(), `"executionId": "ocr-123"`) {
		t.Fatalf("stdout = %q", stdout.String())
	}
}

func TestCSVMergeCommandRequiresCSVFile(t *testing.T) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd := newApp(stdout, stderr)

	err := cmd.Run(context.Background(), []string{
		"adpgo",
		"--host", "https://example.test",
		"--user", "adp",
		"--password", "secret",
		"csv-merge",
	})
	if err == nil {
		t.Fatal("expected error")
	}
	if !strings.Contains(err.Error(), "Required flag \"csvFile\" not set") {
		t.Fatalf("error = %q", err)
	}
}

func TestCSVMergeCommandRequiresExactlyOneSelector(t *testing.T) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd := newApp(stdout, stderr)

	err := cmd.Run(context.Background(), []string{
		"adpgo",
		"--host", "https://example.test",
		"--user", "adp",
		"--password", "secret",
		"csv-merge",
		"--csvFile", "input.csv",
	})
	if err == nil {
		t.Fatal("expected error")
	}
	if !strings.Contains(err.Error(), "exactly one of engineName or applicationIdentifier is required") {
		t.Fatalf("error = %q", err)
	}
}

func TestCSVMergeCommandRejectsBothSelectors(t *testing.T) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd := newApp(stdout, stderr)

	err := cmd.Run(context.Background(), []string{
		"adpgo",
		"--host", "https://example.test",
		"--user", "adp",
		"--password", "secret",
		"csv-merge",
		"--csvFile", "input.csv",
		"--engineName", "engineA",
		"--applicationIdentifier", "appA",
	})
	if err == nil {
		t.Fatal("expected error")
	}
	if !strings.Contains(err.Error(), "engineName and applicationIdentifier are mutually exclusive") {
		t.Fatalf("error = %q", err)
	}
}

func TestCSVMergeCommandAllowsApplicationIdentifierWithoutEngineName(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg := decodeTaskConfiguration(t, r)
		assertApplicationSelectorOnly(t, cfg, "adp_csvMerge_applicationIdentifier", "adp_csvMerge_engineName")

		_, _ = io.WriteString(w, `{"executionId":"3","taskType":"CSV Merge","loggingEnabled":"true","progressMax":1,"executionStatus":"success","executionRootDir":"root","contextId":"ctx","executionPersistent":"true","progressCurrent":1,"progressPercentage":1.0,"taskDisplayName":"Csv merge task","executionMetaData":[]}`)
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
		"csv-merge",
		"--csvFile", "input.csv",
		"--applicationIdentifier", "appA",
	})
	if err != nil {
		t.Fatalf("Run error: %v", err)
	}
	if stderr.Len() != 0 {
		t.Fatalf("stderr = %q", stderr.String())
	}
	if got, want := strings.TrimSpace(stdout.String()), "{}"; got != want {
		t.Fatalf("stdout = %q, want %q", got, want)
	}
}

func TestCLICommandRequiresBatchScriptPath(t *testing.T) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd := newApp(stdout, stderr)

	err := cmd.Run(context.Background(), []string{
		"adpgo",
		"--host", "https://example.test",
		"--user", "adp",
		"--password", "secret",
		"cli",
	})
	if err == nil {
		t.Fatal("expected error")
	}
	if !strings.Contains(err.Error(), "Required flag \"batchScriptPath\" not set") {
		t.Fatalf("error = %q", err)
	}
}

func decodeTaskConfiguration(t *testing.T, r *http.Request) map[string]any {
	t.Helper()

	var req map[string]any
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		t.Fatalf("decode request: %v", err)
	}

	cfg, ok := req["taskConfiguration"].(map[string]any)
	if !ok {
		t.Fatalf("taskConfiguration type = %T", req["taskConfiguration"])
	}

	return cfg
}

func assertApplicationSelectorOnly(t *testing.T, cfg map[string]any, applicationKey, engineKey string) {
	t.Helper()

	if got := cfg[applicationKey]; got != "appA" {
		t.Fatalf("%s = %#v", applicationKey, got)
	}
	if _, ok := cfg[engineKey]; ok {
		t.Fatalf("taskConfiguration should omit %s: %#v", engineKey, cfg)
	}
}

func writeADPConfigFile(t *testing.T, dir string, config map[string]any) {
	t.Helper()

	body, err := json.Marshal(config)
	if err != nil {
		t.Fatalf("marshal config: %v", err)
	}

	configPath := filepath.Join(dir, "adp_config.json")
	if err := os.WriteFile(configPath, body, 0o644); err != nil {
		t.Fatalf("write config: %v", err)
	}
}

func withWorkingDir(t *testing.T, dir string) {
	t.Helper()

	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatalf("chdir %q: %v", dir, err)
	}
	t.Cleanup(func() {
		if err := os.Chdir(wd); err != nil {
			t.Fatalf("restore working directory: %v", err)
		}
	})
}
