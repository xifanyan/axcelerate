package adp

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testClientForBuilder(t *testing.T, handler http.HandlerFunc) *Client {
	t.Helper()
	server := httptest.NewServer(handler)
	t.Cleanup(server.Close)

	client, err := NewClient(ClientConfig{BaseURL: server.URL, Username: "adp", Password: "secret"})
	if err != nil {
		t.Fatalf("NewClient error: %v", err)
	}
	return client
}

func TestListEntitiesBuildRequestIsSparse(t *testing.T) {
	client := testClientForBuilder(t, func(w http.ResponseWriter, r *http.Request) {})

	req, err := NewListEntitiesBuilder(client).
		Type("singleMindServer").
		WhiteList("id,displayName,processStatus").
		buildRequest()
	if err != nil {
		t.Fatalf("buildRequest error: %v", err)
	}

	if req.TaskType != "List Entities" {
		t.Fatalf("taskType = %q", req.TaskType)
	}
	if len(req.TaskConfiguration) != 2 {
		t.Fatalf("taskConfiguration length = %d, want 2", len(req.TaskConfiguration))
	}
	if req.TaskConfiguration["adp_listEntities_type"] != "singleMindServer" {
		t.Fatalf("missing type mapping: %#v", req.TaskConfiguration)
	}
	if req.TaskConfiguration["adp_listEntities_whiteList"] != "id,displayName,processStatus" {
		t.Fatalf("missing whitelist mapping: %#v", req.TaskConfiguration)
	}
}

func TestListEntitiesBuildRequestIncludesCommonBuilderFields(t *testing.T) {
	client := testClientForBuilder(t, func(w http.ResponseWriter, r *http.Request) {})

	req, err := NewListEntitiesBuilder(client).
		Type("singleMindServer").
		LoggingEnabled(false).
		TaskTimeout(45).
		buildRequest()
	if err != nil {
		t.Fatalf("buildRequest error: %v", err)
	}

	if req.TaskConfiguration["adp_loggingEnabled"] != false {
		t.Fatalf("missing common logging mapping: %#v", req.TaskConfiguration)
	}
	if req.TaskConfiguration["adp_taskTimeout"] != 45 {
		t.Fatalf("missing common timeout mapping: %#v", req.TaskConfiguration)
	}
}

func TestListEntitiesExecuteDecodesEntities(t *testing.T) {
	client := testClientForBuilder(t, func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `{"executionId":"5","taskType":"List Entities","loggingEnabled":"true","progressMax":1,"executionStatus":"success","executionRootDir":"root","contextId":"ctx","executionPersistent":"true","progressCurrent":1,"progressPercentage":1.0,"taskDisplayName":"List entities","executionMetaData":{"adp_entities_output_file_name":"output.json","adp_entities_json_output":"[{\"id\":\"a\",\"displayName\":\"A\"}]"}}`)
	})

	got, err := NewListEntitiesBuilder(client).Type("singleMindServer").Execute(context.Background())
	if err != nil {
		t.Fatalf("Execute error: %v", err)
	}
	if got.OutputFile != "output.json" || len(got.Entities) != 1 || got.Entities[0].ID != "a" {
		t.Fatalf("got = %#v", got)
	}
}

func TestQueryEngineRequiresEngineName(t *testing.T) {
	client := testClientForBuilder(t, func(w http.ResponseWriter, r *http.Request) {})

	_, err := NewQueryEngineBuilder(client).Execute(context.Background())
	if err == nil || err.Error() != "engineName is required" {
		t.Fatalf("err = %v", err)
	}
}

func TestQueryEngineRejectsEmptyEngineName(t *testing.T) {
	client := testClientForBuilder(t, func(w http.ResponseWriter, r *http.Request) {})

	_, err := NewQueryEngineBuilder(client).EngineName("").Execute(context.Background())
	if err == nil || err.Error() != "engineName is required" {
		t.Fatalf("err = %v", err)
	}
}

func TestQueryEngineBuildsTaxonomyFiltersAndDecodesResult(t *testing.T) {
	client := testClientForBuilder(t, func(w http.ResponseWriter, r *http.Request) {
		var req rawTaskRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Decode request error: %v", err)
		}
		if req.TaskConfiguration["adp_queryEngine_engineName"] != "engineA" {
			t.Fatalf("engineName = %#v", req.TaskConfiguration["adp_queryEngine_engineName"])
		}

		taxonomies, ok := req.TaskConfiguration["adp_queryEngine_engineTaxonomies"].([]any)
		if !ok || len(taxonomies) != 1 {
			t.Fatalf("engineTaxonomies = %#v", req.TaskConfiguration["adp_queryEngine_engineTaxonomies"])
		}

		taxonomy, ok := taxonomies[0].(map[string]any)
		if !ok {
			t.Fatalf("taxonomy = %#v", taxonomies[0])
		}
		if taxonomy["taxonomy"] != "rm_mimetype" || taxonomy["query"] != "pdf" || taxonomy["negation"] != false {
			t.Fatalf("taxonomy = %#v", taxonomy)
		}

		io.WriteString(w, `{"executionId":"6","taskType":"Query Engine","loggingEnabled":"true","progressMax":1,"executionStatus":"success","executionRootDir":"root","contextId":"ctx","executionPersistent":"true","progressCurrent":1,"progressPercentage":1.0,"taskDisplayName":"Query engine","executionMetaData":{"adp_query_engine_documents_count":"100","adp_query_engine_aggregated_value":"500"}}`)
	})

	got, err := NewQueryEngineBuilder(client).
		EngineName("engineA").
		EngineTaxonomies([]EngineTaxonomyArg{{Taxonomy: "rm_mimetype", Query: "pdf"}}).
		Execute(context.Background())
	if err != nil {
		t.Fatalf("Execute error: %v", err)
	}
	if got.DocumentsCount != 100 || got.AggregatedValue != "500" {
		t.Fatalf("got = %#v", got)
	}
}

func TestDecodeQueryEngineRequiresExpectedMetadataFields(t *testing.T) {
	t.Run("documents count must be string", func(t *testing.T) {
		_, err := decodeQueryEngine(map[string]any{
			"adp_query_engine_documents_count":  100,
			"adp_query_engine_aggregated_value": "500",
		})
		if err == nil || err.Error() != "adp_query_engine_documents_count must be a string" {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("aggregated value is required", func(t *testing.T) {
		_, err := decodeQueryEngine(map[string]any{
			"adp_query_engine_documents_count": "100",
		})
		if err == nil || err.Error() != "missing adp_query_engine_aggregated_value" {
			t.Fatalf("err = %v", err)
		}
	})
}

func TestTaxonomyStatisticRequiresEngineName(t *testing.T) {
	client := testClientForBuilder(t, func(w http.ResponseWriter, r *http.Request) {})
	_, err := NewTaxonomyStatisticBuilder(client).Execute(context.Background())
	if err == nil || err.Error() != "engineName is required" {
		t.Fatalf("err = %v", err)
	}
}

func TestTaxonomyStatisticBuildRequestSerializesBoolFlagsAsStringsAndOmitsUnsetFields(t *testing.T) {
	client := testClientForBuilder(t, func(w http.ResponseWriter, r *http.Request) {})

	req, err := NewTaxonomyStatisticBuilder(client).
		EngineName("engineA").
		ComputeCounts(true).
		ListCategoryProperties(false).
		buildRequest()
	if err != nil {
		t.Fatalf("buildRequest error: %v", err)
	}

	if got := req.TaskConfiguration["adp_taxonomyStatistic_computeCounts"]; got != "true" {
		t.Fatalf("computeCounts = %#v", got)
	}
	if got := req.TaskConfiguration["adp_taxonomyStatistic_listCategoryProperties"]; got != "false" {
		t.Fatalf("listCategoryProperties = %#v", got)
	}
	if _, ok := req.TaskConfiguration["adp_taxonomyStatistic_engineQuery"]; ok {
		t.Fatalf("engineQuery should be omitted: %#v", req.TaskConfiguration)
	}
	if _, ok := req.TaskConfiguration["adp_taxonomyStatistic_engineTaxonomies"]; ok {
		t.Fatalf("engineTaxonomies should be omitted: %#v", req.TaskConfiguration)
	}
	if _, ok := req.TaskConfiguration["adp_taxonomyStatistic_outputTaxonomies"]; ok {
		t.Fatalf("outputTaxonomies should be omitted: %#v", req.TaskConfiguration)
	}
	if _, ok := req.TaskConfiguration["adp_taxonomyStatistic_applicationIdentifier"]; ok {
		t.Fatalf("applicationIdentifier should be omitted: %#v", req.TaskConfiguration)
	}
}

func TestTaxonomyStatisticDecodesStatisticsDocument(t *testing.T) {
	client := testClientForBuilder(t, func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `{"executionId":"7","taskType":"Taxonomy Statistic","loggingEnabled":"true","progressMax":1,"executionStatus":"success","executionRootDir":"root","contextId":"ctx","executionPersistent":"true","progressCurrent":1,"progressPercentage":1.0,"taskDisplayName":"Taxonomy statistic","executionMetaData":{"adp_taxonomy_statistics_json_file_path":"taxonomy_stats.json","adp_taxonomy_statistics_json_output":"{\"date\":\"Wed\",\"searchParameter\":[],\"statistics\":{\"taxonomy\":[{\"id\":\"rm_source\",\"category\":[{\"id\":\"file_demo_04\",\"displayName\":\"file_demo_04\",\"count\":761}]}]}}"}}`)
	})

	got, err := NewTaxonomyStatisticBuilder(client).
		EngineName("engineA").
		Execute(context.Background())
	if err != nil {
		t.Fatalf("Execute error: %v", err)
	}
	if got.OutputFile != "taxonomy_stats.json" || len(got.Statistics.Statistics.Taxonomy) != 1 {
		t.Fatalf("got = %#v", got)
	}
}

func TestStartApplicationDecodesURL(t *testing.T) {
	type requestCapture struct {
		req rawTaskRequest
		err error
	}
	requestCh := make(chan requestCapture, 1)

	client := testClientForBuilder(t, func(w http.ResponseWriter, r *http.Request) {
		var req rawTaskRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		requestCh <- requestCapture{req: req, err: err}
		io.WriteString(w, `{"executionId":"8","taskType":"Start Application","loggingEnabled":"true","progressMax":1,"executionStatus":"success","executionRootDir":"root","contextId":"ctx","executionPersistent":"true","progressCurrent":1,"progressPercentage":1.0,"taskDisplayName":"Start application","executionMetaData":{"adp_started_application_url":"https://example/app"}}`)
	})

	got, err := NewStartApplicationBuilder(client).ApplicationIdentifier("app").Execute(context.Background())
	if err != nil {
		t.Fatalf("Execute error: %v", err)
	}
	captured := <-requestCh
	if captured.err != nil {
		t.Fatalf("Decode request error: %v", captured.err)
	}
	if captured.req.TaskType != "Start Application" {
		t.Fatalf("taskType = %q", captured.req.TaskType)
	}
	if captured.req.TaskConfiguration["adp_startApplication_applicationIdentifier"] != "app" {
		t.Fatalf("taskConfiguration = %#v", captured.req.TaskConfiguration)
	}
	if got.ApplicationURL != "https://example/app" {
		t.Fatalf("got = %#v", got)
	}
}

func TestExportDocumentsDecodesCounts(t *testing.T) {
	type requestCapture struct {
		req rawTaskRequest
		err error
	}
	requestCh := make(chan requestCapture, 1)

	client := testClientForBuilder(t, func(w http.ResponseWriter, r *http.Request) {
		var req rawTaskRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		requestCh <- requestCapture{req: req, err: err}
		io.WriteString(w, `{"executionId":"9","taskType":"Export Documents","loggingEnabled":"true","progressMax":1,"executionStatus":"success","executionRootDir":"root","contextId":"ctx","executionPersistent":"true","progressCurrent":1,"progressPercentage":1.0,"taskDisplayName":"Export documents task","executionMetaData":{"adp_exportDocuments_exportFileName":"export.csv","adp_exportDocuments_exportPath":"/tmp/export.csv","adp_exportDocuments_searchResultSize":"1000"}}`)
	})

	got, err := NewExportDocumentsBuilder(client).Query("*").Execute(context.Background())
	if err != nil {
		t.Fatalf("Execute error: %v", err)
	}
	captured := <-requestCh
	if captured.err != nil {
		t.Fatalf("Decode request error: %v", captured.err)
	}
	if captured.req.TaskType != "Export Documents" {
		t.Fatalf("taskType = %q", captured.req.TaskType)
	}
	if captured.req.TaskConfiguration["adp_exportDocuments_query"] != "*" {
		t.Fatalf("taskConfiguration = %#v", captured.req.TaskConfiguration)
	}
	if got.SearchResultSize != 1000 {
		t.Fatalf("got = %#v", got)
	}
}

func TestCSVMergeRequiresCSVFileAndAcceptsEmptyArrayMetadata(t *testing.T) {
	client := testClientForBuilder(t, func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `{"executionId":"10","taskType":"CSV Merge","loggingEnabled":"true","progressMax":2,"executionStatus":"success","executionRootDir":"root","contextId":"ctx","executionPersistent":"true","progressCurrent":2,"progressPercentage":1.0,"taskDisplayName":"Csv merge task","executionMetaData":[]}`)
	})

	_, err := NewCSVMergeBuilder(client).Execute(context.Background())
	if err == nil || err.Error() != "csvFile is required" {
		t.Fatalf("err = %v", err)
	}

	got, err := NewCSVMergeBuilder(client).CSVFile("/tmp/data.csv").Execute(context.Background())
	if err != nil {
		t.Fatalf("Execute error: %v", err)
	}
	if got != (CSVMergeResult{}) {
		t.Fatalf("got = %#v", got)
	}
}

func TestCSVMergeBuildRequestSerializesFieldMappingsAsArrayAndOmitsEmptyValue(t *testing.T) {
	client := testClientForBuilder(t, func(w http.ResponseWriter, r *http.Request) {})

	req, err := NewCSVMergeBuilder(client).
		CSVFile("/tmp/data.csv").
		FieldMappings([]map[string]any{{"csvField": "source", "adpField": "target"}}).
		buildRequest()
	if err != nil {
		t.Fatalf("buildRequest error: %v", err)
	}

	mappings, ok := req.TaskConfiguration["adp_csvMerge_fieldMappings"].([]map[string]any)
	if !ok || len(mappings) != 1 {
		t.Fatalf("fieldMappings = %#v", req.TaskConfiguration["adp_csvMerge_fieldMappings"])
	}
	if mappings[0]["csvField"] != "source" || mappings[0]["adpField"] != "target" {
		t.Fatalf("fieldMappings = %#v", mappings)
	}

	req, err = NewCSVMergeBuilder(client).
		CSVFile("/tmp/data.csv").
		FieldMappings([]map[string]any{}).
		buildRequest()
	if err != nil {
		t.Fatalf("buildRequest error: %v", err)
	}
	if _, ok := req.TaskConfiguration["adp_csvMerge_fieldMappings"]; ok {
		t.Fatalf("fieldMappings should be omitted: %#v", req.TaskConfiguration)
	}
}
