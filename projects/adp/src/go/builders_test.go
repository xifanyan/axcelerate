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
