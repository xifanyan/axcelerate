package adp

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
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

func TestQueryEngineRequiresExactlyOneSelector(t *testing.T) {
	client := testClientForBuilder(t, func(w http.ResponseWriter, r *http.Request) {})

	_, err := NewQueryEngineBuilder(client).buildRequest()
	if err == nil || err.Error() != "exactly one of engineName or applicationIdentifier is required" {
		t.Fatalf("err = %v", err)
	}
}

func TestQueryEngineRejectsBothSelectors(t *testing.T) {
	client := testClientForBuilder(t, func(w http.ResponseWriter, r *http.Request) {})

	_, err := NewQueryEngineBuilder(client).
		EngineName("engineA").
		ApplicationIdentifier("appA").
		buildRequest()
	if err == nil || err.Error() != "engineName and applicationIdentifier are mutually exclusive" {
		t.Fatalf("err = %v", err)
	}
}

func TestQueryEngineAllowsApplicationIdentifierWithoutEngineName(t *testing.T) {
	client := testClientForBuilder(t, func(w http.ResponseWriter, r *http.Request) {})

	req, err := NewQueryEngineBuilder(client).
		ApplicationIdentifier("appA").
		buildRequest()
	if err != nil {
		t.Fatalf("buildRequest error: %v", err)
	}
	if req.TaskConfiguration["adp_queryEngine_applicationIdentifier"] != "appA" {
		t.Fatalf("taskConfiguration = %#v", req.TaskConfiguration)
	}
	if _, ok := req.TaskConfiguration["adp_queryEngine_engineName"]; ok {
		t.Fatalf("taskConfiguration should omit engineName: %#v", req.TaskConfiguration)
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

func TestTaxonomyStatisticRequiresExactlyOneSelector(t *testing.T) {
	client := testClientForBuilder(t, func(w http.ResponseWriter, r *http.Request) {})

	_, err := NewTaxonomyStatisticBuilder(client).buildRequest()
	if err == nil || err.Error() != "exactly one of engineName or applicationIdentifier is required" {
		t.Fatalf("err = %v", err)
	}
}

func TestTaxonomyStatisticRejectsBothSelectors(t *testing.T) {
	client := testClientForBuilder(t, func(w http.ResponseWriter, r *http.Request) {})

	_, err := NewTaxonomyStatisticBuilder(client).
		EngineName("engineA").
		ApplicationIdentifier("appA").
		buildRequest()
	if err == nil || err.Error() != "engineName and applicationIdentifier are mutually exclusive" {
		t.Fatalf("err = %v", err)
	}
}

func TestTaxonomyStatisticAllowsApplicationIdentifierWithoutEngineName(t *testing.T) {
	client := testClientForBuilder(t, func(w http.ResponseWriter, r *http.Request) {})

	req, err := NewTaxonomyStatisticBuilder(client).
		ApplicationIdentifier("appA").
		buildRequest()
	if err != nil {
		t.Fatalf("buildRequest error: %v", err)
	}
	if req.TaskConfiguration["adp_taxonomyStatistic_applicationIdentifier"] != "appA" {
		t.Fatalf("taskConfiguration = %#v", req.TaskConfiguration)
	}
	if _, ok := req.TaskConfiguration["adp_taxonomyStatistic_engineName"]; ok {
		t.Fatalf("taskConfiguration should omit engineName: %#v", req.TaskConfiguration)
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
	if len(captured.req.TaskConfiguration) != 1 {
		t.Fatalf("taskConfiguration length = %d, want 1: %#v", len(captured.req.TaskConfiguration), captured.req.TaskConfiguration)
	}
	if captured.req.TaskConfiguration["adp_startApplication_applicationIdentifier"] != "app" {
		t.Fatalf("taskConfiguration = %#v", captured.req.TaskConfiguration)
	}
	if _, ok := captured.req.TaskConfiguration["adp_startApplication_useHttps"]; ok {
		t.Fatalf("taskConfiguration should omit useHttps: %#v", captured.req.TaskConfiguration)
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
	if len(captured.req.TaskConfiguration) != 1 {
		t.Fatalf("taskConfiguration length = %d, want 1: %#v", len(captured.req.TaskConfiguration), captured.req.TaskConfiguration)
	}
	if captured.req.TaskConfiguration["adp_exportDocuments_query"] != "*" {
		t.Fatalf("taskConfiguration = %#v", captured.req.TaskConfiguration)
	}
	if _, ok := captured.req.TaskConfiguration["adp_exportDocuments_waitForExport"]; ok {
		t.Fatalf("taskConfiguration should omit waitForExport: %#v", captured.req.TaskConfiguration)
	}
	if _, ok := captured.req.TaskConfiguration["adp_exportDocuments_exportFields"]; ok {
		t.Fatalf("taskConfiguration should omit exportFields: %#v", captured.req.TaskConfiguration)
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

func TestCSVMergeRequiresExactlyOneSelector(t *testing.T) {
	client := testClientForBuilder(t, func(w http.ResponseWriter, r *http.Request) {})

	_, err := NewCSVMergeBuilder(client).
		CSVFile("/tmp/data.csv").
		buildRequest()
	if err == nil || err.Error() != "exactly one of engineName or applicationIdentifier is required" {
		t.Fatalf("err = %v", err)
	}
}

func TestCSVMergeRejectsBothSelectors(t *testing.T) {
	client := testClientForBuilder(t, func(w http.ResponseWriter, r *http.Request) {})

	_, err := NewCSVMergeBuilder(client).
		CSVFile("/tmp/data.csv").
		EngineName("engineA").
		ApplicationIdentifier("appA").
		buildRequest()
	if err == nil || err.Error() != "engineName and applicationIdentifier are mutually exclusive" {
		t.Fatalf("err = %v", err)
	}
}

func TestCSVMergeAllowsApplicationIdentifierWithoutEngineName(t *testing.T) {
	client := testClientForBuilder(t, func(w http.ResponseWriter, r *http.Request) {})

	req, err := NewCSVMergeBuilder(client).
		CSVFile("/tmp/data.csv").
		ApplicationIdentifier("appA").
		buildRequest()
	if err != nil {
		t.Fatalf("buildRequest error: %v", err)
	}
	if req.TaskConfiguration["adp_csvMerge_applicationIdentifier"] != "appA" {
		t.Fatalf("taskConfiguration = %#v", req.TaskConfiguration)
	}
	if _, ok := req.TaskConfiguration["adp_csvMerge_engineName"]; ok {
		t.Fatalf("taskConfiguration should omit engineName: %#v", req.TaskConfiguration)
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

func TestReadConfigurationBuildsConfigObjectsAndDecodesResult(t *testing.T) {
	type requestCapture struct {
		req rawTaskRequest
		err error
	}
	requestCh := make(chan requestCapture, 1)

	client := testClientForBuilder(t, func(w http.ResponseWriter, r *http.Request) {
		var req rawTaskRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		requestCh <- requestCapture{req: req, err: err}
		io.WriteString(w, `{"executionId":"11","taskType":"Read Configuration","loggingEnabled":"true","progressMax":1,"executionStatus":"success","executionRootDir":"root","contextId":"ctx","executionPersistent":"true","progressCurrent":1,"progressPercentage":1.0,"taskDisplayName":"Read Configuration","executionMetaData":{"adp_entities_output_file_name":"output.json","adp_entities_json_output":"{\"dataSource.file_demo_01\":{\"DynamicComponents\":{\"compA\":{\"Enabled\":true}},\"Global\":{\"Static\":{\"Parameters\":[{\"cells\":[[{\"name\":\"column\",\"value\":\"cell-value\"}]],\"name\":\"param\",\"value\":\"x\"}]}}}}"}}`)
	})

	got, err := NewReadConfigurationBuilder(client).
		ConfigsToRead([]ConfigArg{{
			ConfigurationID:       "dataSource.file_demo_01",
			DynamicComponentNames: "compA,compB",
			FieldList:             "name,value",
			NameValueList:         "k1,v1",
			ApplicationType:       "appType",
		}}).
		Execute(context.Background())
	if err != nil {
		t.Fatalf("Execute error: %v", err)
	}
	captured := <-requestCh
	if captured.err != nil {
		t.Fatalf("Decode request error: %v", captured.err)
	}
	configs, ok := captured.req.TaskConfiguration["adp_readConfiguration_configsToRead"].([]any)
	if !ok || len(configs) != 1 {
		t.Fatalf("configsToRead = %#v", captured.req.TaskConfiguration["adp_readConfiguration_configsToRead"])
	}
	config, ok := configs[0].(map[string]any)
	if !ok {
		t.Fatalf("config = %#v", configs[0])
	}
	if len(config) != 5 {
		t.Fatalf("config = %#v", config)
	}
	if config["Configuration ID"] != "dataSource.file_demo_01" ||
		config["Dynamic Component Names"] != "compA,compB" ||
		config["Field list"] != "name,value" ||
		config["Name value list"] != "k1,v1" ||
		config["Application type"] != "appType" {
		t.Fatalf("config = %#v", config)
	}
	if _, ok := config["Entity type"]; ok {
		t.Fatalf("config should omit empty fields: %#v", config)
	}
	info, ok := got.Configuration["dataSource.file_demo_01"]
	if got.OutputFile != "output.json" || len(got.Configuration) != 1 || !ok {
		t.Fatalf("got = %#v", got)
	}
	if len(info.DynamicComponents) != 1 || len(info.Global.Static.Parameters) != 1 {
		t.Fatalf("info = %#v", info)
	}
	if _, ok := info.DynamicComponents["compA"]; !ok {
		t.Fatalf("info = %#v", info)
	}
	if info.Global.Static.Parameters[0].Name != "param" || info.Global.Static.Parameters[0].Value != "x" {
		t.Fatalf("info = %#v", info)
	}
	if len(info.Global.Static.Parameters[0].Cells) != 1 || len(info.Global.Static.Parameters[0].Cells[0]) != 1 {
		t.Fatalf("info = %#v", info)
	}
	if info.Global.Static.Parameters[0].Cells[0][0].Name != "column" || info.Global.Static.Parameters[0].Cells[0][0].Value != "cell-value" {
		t.Fatalf("info = %#v", info)
	}
}

func TestCLITaskRequiresBatchScriptPathAndDecodesJSONOutput(t *testing.T) {
	client := testClientForBuilder(t, func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `{"executionId":"12","taskType":"CLI","loggingEnabled":"true","progressMax":1,"executionStatus":"success","executionRootDir":"root","contextId":"ctx","executionPersistent":"true","progressCurrent":1,"progressPercentage":1.0,"taskDisplayName":"Command Line Task","executionMetaData":{"cli_result":0,"json_output":"{\"stdout\":\"ok\",\"errout\":\"\"}","cli_error_path":"err.log","cli_result_path":"out.log"}}`)
	})

	_, err := NewCLITaskBuilder(client).Execute(context.Background())
	if err == nil || err.Error() != "batchScriptPath is required" {
		t.Fatalf("err = %v", err)
	}

	got, err := NewCLITaskBuilder(client).BatchScriptPath("c:/demo/script.ps1").Execute(context.Background())
	if err != nil {
		t.Fatalf("Execute error: %v", err)
	}
	if got.Result != 0 || got.JSONOutput["stdout"] != "ok" {
		t.Fatalf("got = %#v", got)
	}
}

func TestDecodeCLITaskRejectsFractionalResult(t *testing.T) {
	_, err := decodeCLITask(map[string]any{
		"cli_result":      0.5,
		"cli_error_path":  "err.log",
		"cli_result_path": "out.log",
	})
	if err == nil || err.Error() != "cli_result must be an integer" {
		t.Fatalf("err = %v", err)
	}
}

func TestDecodeCLITaskRejectsMalformedJSONOutput(t *testing.T) {
	_, err := decodeCLITask(map[string]any{
		"cli_result":      0.0,
		"json_output":     "{",
		"cli_error_path":  "err.log",
		"cli_result_path": "out.log",
	})
	if err == nil || err.Error() == "" {
		t.Fatalf("err = %v", err)
	}
}

func TestDecodeCLITaskAcceptsJSONNumber(t *testing.T) {
	got, err := decodeCLITask(map[string]any{
		"cli_result":      json.Number("7"),
		"cli_error_path":  "err.log",
		"cli_result_path": "out.log",
	})
	if err != nil {
		t.Fatalf("err = %v", err)
	}
	if got.Result != 7 {
		t.Fatalf("got = %#v", got)
	}
}

func TestCLITaskExecutePreservesLargeIntegerMetadataPrecision(t *testing.T) {
	const cliResultText = "9007199254740993"

	client := testClientForBuilder(t, func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `{"executionId":"13","taskType":"CLI","loggingEnabled":"true","progressMax":1,"executionStatus":"success","executionRootDir":"root","contextId":"ctx","executionPersistent":"true","progressCurrent":1,"progressPercentage":1.0,"taskDisplayName":"Command Line Task","executionMetaData":{"cli_result":`+cliResultText+`,"cli_error_path":"err.log","cli_result_path":"out.log"}}`)
	})

	got, err := NewCLITaskBuilder(client).BatchScriptPath("c:/demo/script.ps1").Execute(context.Background())
	if strconv.IntSize == 32 {
		if err == nil || err.Error() != "cli_result out of range" {
			t.Fatalf("err = %v", err)
		}
		return
	}

	if err != nil {
		t.Fatalf("Execute error: %v", err)
	}
	expected, parseErr := strconv.ParseInt(cliResultText, 10, 64)
	if parseErr != nil {
		t.Fatalf("ParseInt error: %v", parseErr)
	}
	if int64(got.Result) != expected {
		t.Fatalf("got = %#v", got)
	}
}

func TestCLITaskExecuteRejectsExponentMetadataResult(t *testing.T) {
	client := testClientForBuilder(t, func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `{"executionId":"14","taskType":"CLI","loggingEnabled":"true","progressMax":1,"executionStatus":"success","executionRootDir":"root","contextId":"ctx","executionPersistent":"true","progressCurrent":1,"progressPercentage":1.0,"taskDisplayName":"Command Line Task","executionMetaData":{"cli_result":1e3,"cli_error_path":"err.log","cli_result_path":"out.log"}}`)
	})

	_, err := NewCLITaskBuilder(client).BatchScriptPath("c:/demo/script.ps1").Execute(context.Background())
	if err == nil || err.Error() != "cli_result must be an integer" {
		t.Fatalf("err = %v", err)
	}
}

func TestCreateOcrJobStartBuildsRestrictions(t *testing.T) {
	var gotBody map[string]any
	var gotPath string

	client := testClientForBuilder(t, func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		if err := json.NewDecoder(r.Body).Decode(&gotBody); err != nil {
			t.Fatalf("decode request: %v", err)
		}
		io.WriteString(w, `{"executionId":"13","taskType":"Create OCR Job","loggingEnabled":"true","progressMax":0,"executionStatus":"","executionRootDir":"root","contextId":"ctx","executionPersistent":"true","progressCurrent":0,"progressPercentage":0.0,"taskDisplayName":"Create OCR Job","executionMetaData":[]}`)
	})

	_, err := NewCreateOcrJobBuilder(client).
		EngineName("singleMindServer.demo00001").
		JobName("demo_ocr").
		Restrictions([]EngineTaxonomyArg{{Taxonomy: "rm_source", Negation: false, Query: "file_demo_04"}}).
		AdvancedRestrictions([]EngineTaxonomyArg{{Taxonomy: "rm_mimetype", Negation: false, Query: "application%2Fpdf"}}).
		Start(context.Background())
	if err != nil {
		t.Fatalf("Start error: %v", err)
	}
	if gotPath != "/executeAdpTaskAsync" {
		t.Fatalf("path = %q", gotPath)
	}

	cfg := gotBody["taskConfiguration"].(map[string]any)
	restrictions := cfg["adp_createOcrJob_restrictions"].([]any)
	if len(restrictions) != 1 {
		t.Fatalf("restrictions = %#v", restrictions)
	}
	advancedRestrictions := cfg["adp_createOcrJob_AdvancedRestrictions"].([]any)
	if len(advancedRestrictions) != 1 {
		t.Fatalf("advancedRestrictions = %#v", advancedRestrictions)
	}
}

func TestCreateOcrJobWaitReturnsEmptyResult(t *testing.T) {
	calls := 0
	paths := make([]string, 0, 2)
	client := testClientForBuilder(t, func(w http.ResponseWriter, r *http.Request) {
		calls++
		paths = append(paths, r.URL.Path)
		if calls == 1 {
			io.WriteString(w, `{"executionId":"13","taskType":"Create OCR Job","loggingEnabled":"true","progressMax":0,"executionStatus":"","executionRootDir":"root","contextId":"ctx","executionPersistent":"true","progressCurrent":0,"progressPercentage":0.0,"taskDisplayName":"Create OCR Job","executionMetaData":[]}`)
			return
		}
		io.WriteString(w, `{"executionId":"13","taskType":"Create OCR Job","loggingEnabled":"true","progressMax":1,"executionStatus":"success","executionRootDir":"root","contextId":"ctx","executionPersistent":"true","progressCurrent":1,"progressPercentage":1.0,"taskDisplayName":"Create OCR Job","executionMetaData":[]}`)
	})

	got, err := NewCreateOcrJobBuilder(client).EngineName("engineA").Wait(context.Background(), time.Millisecond)
	if err != nil {
		t.Fatalf("Wait error: %v", err)
	}
	if len(paths) == 0 || paths[0] != "/executeAdpTaskAsync" {
		t.Fatalf("paths = %#v", paths)
	}
	if got != (CreateOcrJobResult{}) {
		t.Fatalf("got = %#v", got)
	}
}

func TestCreateOcrJobBuildRequestOmitsEmptyRestrictionSlices(t *testing.T) {
	client := testClientForBuilder(t, func(w http.ResponseWriter, r *http.Request) {})

	req, err := NewCreateOcrJobBuilder(client).
		Restrictions([]EngineTaxonomyArg{}).
		AdvancedRestrictions([]EngineTaxonomyArg{}).
		buildRequest()
	if err != nil {
		t.Fatalf("buildRequest error: %v", err)
	}

	if _, ok := req.TaskConfiguration["adp_createOcrJob_restrictions"]; ok {
		t.Fatalf("restrictions should be omitted: %#v", req.TaskConfiguration)
	}
	if _, ok := req.TaskConfiguration["adp_createOcrJob_AdvancedRestrictions"]; ok {
		t.Fatalf("advancedRestrictions should be omitted: %#v", req.TaskConfiguration)
	}
}

func TestCreateOcrJobRequiresExactlyOneSelector(t *testing.T) {
	client := testClientForBuilder(t, func(w http.ResponseWriter, r *http.Request) {})

	_, err := NewCreateOcrJobBuilder(client).buildRequest()
	if err == nil || err.Error() != "exactly one of engineName or applicationIdentifier is required" {
		t.Fatalf("err = %v", err)
	}
}

func TestCreateOcrJobRejectsBothSelectors(t *testing.T) {
	client := testClientForBuilder(t, func(w http.ResponseWriter, r *http.Request) {})

	_, err := NewCreateOcrJobBuilder(client).
		EngineName("engineA").
		ApplicationIdentifier("appA").
		buildRequest()
	if err == nil || err.Error() != "engineName and applicationIdentifier are mutually exclusive" {
		t.Fatalf("err = %v", err)
	}
}

func TestCreateOcrJobAllowsApplicationIdentifierWithoutEngineName(t *testing.T) {
	client := testClientForBuilder(t, func(w http.ResponseWriter, r *http.Request) {})

	req, err := NewCreateOcrJobBuilder(client).
		ApplicationIdentifier("appA").
		buildRequest()
	if err != nil {
		t.Fatalf("buildRequest error: %v", err)
	}
	if req.TaskConfiguration["adp_createOcrJob_applicationIdentifier"] != "appA" {
		t.Fatalf("taskConfiguration = %#v", req.TaskConfiguration)
	}
	if _, ok := req.TaskConfiguration["adp_createOcrJob_engineName"]; ok {
		t.Fatalf("taskConfiguration should omit engineName: %#v", req.TaskConfiguration)
	}
}
