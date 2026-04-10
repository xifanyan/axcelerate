package adp

import (
	"strings"
	"testing"
)

func TestMetaObjectReturnsMap(t *testing.T) {
	meta, err := metaObject(map[string]any{"x": 1})
	if err != nil {
		t.Fatalf("metaObject error: %v", err)
	}
	if meta["x"].(int) != 1 {
		t.Fatalf("meta[x] = %#v", meta["x"])
	}
}

func TestMetaObjectAcceptsEmptyArrayForNoResultTasks(t *testing.T) {
	meta, err := metaObject([]any{})
	if err != nil {
		t.Fatalf("metaObject error: %v", err)
	}
	if len(meta) != 0 {
		t.Fatalf("len(meta) = %d, want 0", len(meta))
	}
}

func TestMetaObjectRejectsNonEmptyArray(t *testing.T) {
	_, err := metaObject([]any{"x"})
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestStringFieldReturnsMissingKeyError(t *testing.T) {
	_, err := stringField(map[string]any{}, "missing")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestStringFieldRejectsNonStringValue(t *testing.T) {
	_, err := stringField(map[string]any{"key": 1}, "key")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestJSONStringFieldDecodesIntoTarget(t *testing.T) {
	var entities []Entity
	err := jsonStringField(map[string]any{"adp_entities_json_output": `[{"id":"a"}]`}, "adp_entities_json_output", &entities)
	if err != nil {
		t.Fatalf("jsonStringField error: %v", err)
	}
	if len(entities) != 1 || entities[0].ID != "a" {
		t.Fatalf("entities = %#v", entities)
	}
}

func TestIntStringFieldParsesStringValue(t *testing.T) {
	got, err := intStringField(map[string]any{"adp_query_engine_documents_count": "100"}, "adp_query_engine_documents_count")
	if err != nil {
		t.Fatalf("intStringField error: %v", err)
	}
	if got != 100 {
		t.Fatalf("got = %d, want 100", got)
	}
}

func TestIntStringFieldRejectsInvalidInteger(t *testing.T) {
	_, err := intStringField(map[string]any{"adp_query_engine_documents_count": "abc"}, "adp_query_engine_documents_count")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestJSONStringFieldRejectsMalformedJSON(t *testing.T) {
	var entities []Entity
	err := jsonStringField(map[string]any{"adp_entities_json_output": `[{`}, "adp_entities_json_output", &entities)
	if err == nil {
		t.Fatal("expected error")
	}
	if !strings.Contains(err.Error(), "adp_entities_json_output") {
		t.Fatalf("error = %v", err)
	}
}

type builderCommonOwner struct {
	taskActive          bool
	taskTimeout         int
	executionPersistent bool
	abortWfOnFailure    bool
	loggingEnabled      bool
	cleanUpHistory      bool
}

func requireBuilderCommonOwner(common builderCommon[builderCommonOwner]) builderCommon[builderCommonOwner] {
	return common
}

func TestBuilderCommonSettersReturnPointerOwnerAndApplyValues(t *testing.T) {
	owner := &builderCommonOwner{}
	common := requireBuilderCommonOwner(newBuilderCommon(owner))
	if got := common.TaskActive(false); got != owner {
		t.Fatalf("TaskActive returned %p, want %p", got, owner)
	}
	common.TaskTimeout(12)
	common.ExecutionPersistent(false)
	common.AbortWfOnFailure(false)
	common.LoggingEnabled(false)
	common.CleanUpHistory(true)

	config := map[string]any{}
	common.apply(config)

	if config["adp_taskActive"] != false {
		t.Fatalf("adp_taskActive = %#v", config["adp_taskActive"])
	}
	if config["adp_taskTimeout"] != 12 {
		t.Fatalf("adp_taskTimeout = %#v", config["adp_taskTimeout"])
	}
	if config["adp_executionPersistent"] != false {
		t.Fatalf("adp_executionPersistent = %#v", config["adp_executionPersistent"])
	}
	if config["adp_abortWfOnFailure"] != false {
		t.Fatalf("adp_abortWfOnFailure = %#v", config["adp_abortWfOnFailure"])
	}
	if config["adp_loggingEnabled"] != false {
		t.Fatalf("adp_loggingEnabled = %#v", config["adp_loggingEnabled"])
	}
	if config["adp_cleanUpHistory"] != true {
		t.Fatalf("adp_cleanUpHistory = %#v", config["adp_cleanUpHistory"])
	}
}
