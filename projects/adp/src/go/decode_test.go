package adp

import "testing"

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
