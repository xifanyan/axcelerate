package adp

import "testing"

func TestParseEngineTaxonomyArgEquals(t *testing.T) {
	got, err := parseEngineTaxonomyArg("rm_mimetype=pdf")
	if err != nil {
		t.Fatalf("parseEngineTaxonomyArg error: %v", err)
	}
	if got.Taxonomy != "rm_mimetype" || got.Negation || got.Query != "pdf" {
		t.Fatalf("got = %#v", got)
	}
}

func TestParseEngineTaxonomyArgNotEquals(t *testing.T) {
	got, err := parseEngineTaxonomyArg("rm_source!=email")
	if err != nil {
		t.Fatalf("parseEngineTaxonomyArg error: %v", err)
	}
	if got.Taxonomy != "rm_source" || !got.Negation || got.Query != "email" {
		t.Fatalf("got = %#v", got)
	}
}

func TestParseEngineTaxonomyArgTrimsWhitespace(t *testing.T) {
	got, err := parseEngineTaxonomyArg(" rm_source != email ")
	if err != nil {
		t.Fatalf("parseEngineTaxonomyArg error: %v", err)
	}
	if got.Taxonomy != "rm_source" || !got.Negation || got.Query != "email" {
		t.Fatalf("got = %#v", got)
	}
}

func TestParseEngineTaxonomyArgInvalid(t *testing.T) {
	_, err := parseEngineTaxonomyArg("rm_source")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestParseEngineTaxonomyArgRejectsTrimmedEmptyValues(t *testing.T) {
	_, err := parseEngineTaxonomyArg("  != value ")
	if err == nil {
		t.Fatal("expected error")
	}

	_, err = parseEngineTaxonomyArg(" taxonomy =   ")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestParseOutputTaxonomiesEmpty(t *testing.T) {
	got, err := parseOutputTaxonomies("")
	if err != nil {
		t.Fatalf("parseOutputTaxonomies error: %v", err)
	}
	if got != nil {
		t.Fatalf("got = %#v", got)
	}
}

func TestParseOutputTaxonomiesCSV(t *testing.T) {
	got, err := parseOutputTaxonomies("rm_source,meta_documentcharacteristics")
	if err != nil {
		t.Fatalf("parseOutputTaxonomies error: %v", err)
	}
	if len(got) != 2 || got[0].Taxonomy != "rm_source" || got[0].Mode != "Category counts" || got[0].MaximumNumberOfCategories != 10 || got[1].Taxonomy != "meta_documentcharacteristics" || got[1].Mode != "Category counts" || got[1].MaximumNumberOfCategories != 10 {
		t.Fatalf("got = %#v", got)
	}
}

func TestParseOutputTaxonomiesJSON(t *testing.T) {
	got, err := parseOutputTaxonomies(`[{"taxonomy":"rm_source","mode":"Top categories","maximumNumberOfCategories":3}]`)
	if err != nil {
		t.Fatalf("parseOutputTaxonomies error: %v", err)
	}
	if len(got) != 1 || got[0].Taxonomy != "rm_source" || got[0].Mode != "Top categories" || got[0].MaximumNumberOfCategories != 3 {
		t.Fatalf("got = %#v", got)
	}
}

func TestParseConfigArgsJSON(t *testing.T) {
	got, err := parseConfigArgs(`[{"Configuration ID":"dataSource.file_demo_01","Dynamic Component Names":"x","Field list":"name,value","Name value list":"crawlLocationClassifierRules","Application type":"","Entity type":""}]`)
	if err != nil {
		t.Fatalf("parseConfigArgs error: %v", err)
	}
	if len(got) != 1 || got[0].ConfigurationID != "dataSource.file_demo_01" || got[0].DynamicComponentNames != "x" || got[0].FieldList != "name,value" || got[0].NameValueList != "crawlLocationClassifierRules" || got[0].ApplicationType != "" || got[0].EntityType != "" {
		t.Fatalf("got = %#v", got)
	}
}

func TestParseConfigArgsEmpty(t *testing.T) {
	got, err := parseConfigArgs("")
	if err != nil {
		t.Fatalf("parseConfigArgs error: %v", err)
	}
	if got != nil {
		t.Fatalf("got = %#v", got)
	}
}

func TestParseBatchScriptParametersJSON(t *testing.T) {
	got, err := parseBatchScriptParameters(`[{"Parameter":"-File"},{"Parameter":"c:/demo/test.ps1"}]`)
	if err != nil {
		t.Fatalf("parseBatchScriptParameters error: %v", err)
	}
	if len(got) != 2 || got[0].Parameter != "-File" || got[1].Parameter != "c:/demo/test.ps1" {
		t.Fatalf("got = %#v", got)
	}
}

func TestParseBatchScriptParametersEmpty(t *testing.T) {
	got, err := parseBatchScriptParameters("")
	if err != nil {
		t.Fatalf("parseBatchScriptParameters error: %v", err)
	}
	if got != nil {
		t.Fatalf("got = %#v", got)
	}
}
