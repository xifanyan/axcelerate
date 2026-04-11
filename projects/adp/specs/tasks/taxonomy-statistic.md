# Taxonomy Statistic Task Specification

## Overview

| Property | Value |
|----------|-------|
| Task Type | `Taxonomy Statistic` |
| Description | Retrieves category counts for a taxonomy |
| Display Name | Taxonomy statistic |
| Subcommand | `taxonomy-statistic` |

---

## Semantic Inputs

These are the user-facing fields for the request-construction API.

| Field | Type | Default | Required | Description |
|-------|------|---------|----------|-------------|
| engineName | string | — | Conditional | Engine name |
| engineQuery | string | "*" | No | Query string |
| computeCounts | boolean | true | No | Compute entity counts |
| listCategoryProperties | boolean | false | No | List category properties |
| engineTaxonomies | EngineTaxonomyArg[] | [] | No | Engine taxonomies to filter |
| outputTaxonomies | OutputTaxonomiesArg[] | [] | No | Output taxonomies configuration |
| applicationIdentifier | string | "" | Conditional | Application identifier |

> engineName and applicationIdentifier are mutually exclusive selectors. Exactly one must be provided.

### EngineTaxonomyArg

See [api-contract.md](../api-contract.md#enginetaxonomyarg).

### OutputTaxonomiesArg

See [api-contract.md](../api-contract.md#outputtaxonomiesarg).

---

## Raw Default Configuration

> Configuration below shows **all fields with their exact default values** from [API-SPEC.md](../../API-SPEC.md). This is for reference only. Clients must not pre-populate all fields. See [request-construction.md](../request-construction.md).

```json
{
  "taskType": "Taxonomy Statistic",
  "taskConfiguration": {
    "adp_progressTaskTimeout": 0,
    "adp_taxonomyStatistic_outputJsonAbsFilePath": "adp_taxonomy_statistics_json_file_path",
    "adp_taxonomyStatistic_applicationIdentifier": "",
    "adp_taskActive": true,
    "adp_taxonomyStatistic_adp_taxonomyStatistic_mainQueryType": null,
    "adp_executionPersistent": true,
    "adp_taxonomyStatistic_engineUserName": "{adp_user}",
    "adp_abortWfOnFailure": true,
    "adp_taxonomyStatistic_applicationType": "",
    "adp_taxonomyStatistic_computeCounts": "true",
    "adp_loggingEnabled": true,
    "adp_taxonomyStatistic_outputJsonFilePath": null,
    "adp_taxonomyStatistic_engineTaxonomies": [],
    "adp_taxonomyStatistic_engineUserPassword": "",
    "adp_taxonomyStatistic_outputXmlAbsFilePath": "adp_taxonomy_statistics_xml_file_path",
    "adp_taxonomyStatistic_engineQuery": "*",
    "adp_taxonomyStatistic_listCategoryProperties": "false",
    "adp_taxonomyStatistic_outputTaxonomies": [],
    "adp_taxonomyStatistic_outputJson": "adp_taxonomy_statistics_json_output",
    "adp_taxonomyStatistic_engineType": "true",
    "adp_cleanUpHistory": false,
    "adp_taxonomyStatistic_outputXmlFilePath": null,
    "adp_taxonomyStatistic_outputFields": [],
    "adp_taxonomyStatistic_engineGlobalSearch": "",
    "adp_taxonomyStatistic_listDocuments": "false",
    "adp_taskTimeout": 0,
    "adp_taxonomyStatistic_engineName": "{adp_engine_name}"
  },
  "taskDescription": "Retrieves category counts for a taxonomy",
  "taskDisplayName": "Taxonomy statistic"
}
```

---

## Field Reference (with Defaults)

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| adp_progressTaskTimeout | integer | 0 | Progress task timeout |
| adp_taxonomyStatistic_outputJsonAbsFilePath | string | "adp_taxonomy_statistics_json_file_path" | Output JSON absolute file path |
| adp_taxonomyStatistic_applicationIdentifier | string | "" | Application identifier |
| adp_taskActive | boolean | true | Whether task is active |
| adp_taxonomyStatistic_adp_taxonomyStatistic_mainQueryType | null | null | Main query type |
| adp_executionPersistent | boolean | true | Persist execution |
| adp_taxonomyStatistic_engineUserName | string | "{adp_user}" | Engine username |
| adp_abortWfOnFailure | boolean | true | Abort workflow on failure |
| adp_taxonomyStatistic_applicationType | string | "" | Application type |
| adp_taxonomyStatistic_computeCounts | string | "true" | Compute entity counts |
| adp_loggingEnabled | boolean | true | Enable logging |
| adp_taxonomyStatistic_outputJsonFilePath | null | null | Output JSON file path |
| adp_taxonomyStatistic_engineTaxonomies | array | [] | Engine taxonomies |
| adp_taxonomyStatistic_engineUserPassword | string | "" | Engine password |
| adp_taxonomyStatistic_outputXmlAbsFilePath | string | "adp_taxonomy_statistics_xml_file_path" | Output XML absolute file path |
| adp_taxonomyStatistic_engineQuery | string | "*" | Query string |
| adp_taxonomyStatistic_listCategoryProperties | string | "false" | List category properties |
| adp_taxonomyStatistic_outputTaxonomies | array | [] | Output taxonomies |
| adp_taxonomyStatistic_outputJson | string | "adp_taxonomy_statistics_json_output" | Output JSON variable |
| adp_taxonomyStatistic_engineType | string | "true" | Engine type |
| adp_cleanUpHistory | boolean | false | Clean up history |
| adp_taxonomyStatistic_outputXmlFilePath | null | null | Output XML file path |
| adp_taxonomyStatistic_outputFields | array | [] | Output fields |
| adp_taxonomyStatistic_engineGlobalSearch | string | "" | Global search |
| adp_taxonomyStatistic_listDocuments | string | "false" | List documents |
| adp_taskTimeout | integer | 0 | Task timeout |
| adp_taxonomyStatistic_engineName | string | "{adp_engine_name}" | Engine name |

---

## Raw Example Request

> Example below matches **exactly** the default configuration from [API-SPEC.md](../../API-SPEC.md).

```json
{
  "taskType": "Taxonomy Statistic",
  "taskConfiguration": {
    "adp_progressTaskTimeout": 0,
    "adp_taxonomyStatistic_outputJsonAbsFilePath": "adp_taxonomy_statistics_json_file_path",
    "adp_taxonomyStatistic_applicationIdentifier": "",
    "adp_taskActive": true,
    "adp_taxonomyStatistic_adp_taxonomyStatistic_mainQueryType": null,
    "adp_executionPersistent": true,
    "adp_taxonomyStatistic_engineUserName": "{adp_user}",
    "adp_abortWfOnFailure": true,
    "adp_taxonomyStatistic_applicationType": "",
    "adp_taxonomyStatistic_computeCounts": "true",
    "adp_loggingEnabled": true,
    "adp_taxonomyStatistic_outputJsonFilePath": null,
    "adp_taxonomyStatistic_engineTaxonomies": [],
    "adp_taxonomyStatistic_engineUserPassword": "",
    "adp_taxonomyStatistic_outputXmlAbsFilePath": "adp_taxonomy_statistics_xml_file_path",
    "adp_taxonomyStatistic_engineQuery": "*",
    "adp_taxonomyStatistic_listCategoryProperties": "false",
    "adp_taxonomyStatistic_outputTaxonomies": [],
    "adp_taxonomyStatistic_outputJson": "adp_taxonomy_statistics_json_output",
    "adp_taxonomyStatistic_engineType": "true",
    "adp_cleanUpHistory": false,
    "adp_taxonomyStatistic_outputXmlFilePath": null,
    "adp_taxonomyStatistic_outputFields": [],
    "adp_taxonomyStatistic_engineGlobalSearch": "",
    "adp_taxonomyStatistic_listDocuments": "false",
    "adp_taskTimeout": 0,
    "adp_taxonomyStatistic_engineName": "{adp_engine_name}"
  },
  "taskDescription": "Retrieves category counts for a taxonomy",
  "taskDisplayName": "Taxonomy statistic"
}
```

---

## CLI Arguments

See [cli.md](../cli.md) for global flags and naming conventions.

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--engineName` | string | — | Engine name |
| `--engineQuery` | string | "*" | Query string |
| `--computeCounts` | boolean | true | Compute entity counts |
| `--listCategoryProperties` | boolean | false | List category properties |
| `--engineTaxonomies` | string | — | Engine taxonomies filter (repeatable, see format below) |
| `--outputTaxonomies` | string | — | Output taxonomies: comma-separated list or JSON array |
| `--applicationIdentifier` | string | "" | Application identifier |

> engineName and applicationIdentifier are mutually exclusive selectors. Exactly one must be provided.

### EngineTaxonomiesArg CLI Format

Shorthand format: `Taxonomy=Query` or `Taxonomy!=Query`

| Format | Description | Example |
|--------|-------------|---------|
| `Taxonomy=Query` | Equals (negation=false) | `rm_mimetype=pdf` |
| `Taxonomy!=Query` | Not equals (negation=true) | `rm_source!=email` |

Repeatable flag.

### CLI Examples

```bash
# With outputTaxonomies (comma-separated taxonomy names)
adpgo taxonomy-statistic --engineName "myEngine" --outputTaxonomies "rm_source,meta_documentcharacteristics"

# Using application identifier
adpgo taxonomy-statistic --applicationIdentifier "my-app-id" --outputTaxonomies "rm_source"

# Single taxonomy equals
adpgo taxonomy-statistic --engineName "myEngine" --engineTaxonomies "rm_mimetype=pdf"

# Multiple taxonomies
adpgo taxonomy-statistic --engineName "myEngine" --engineTaxonomies "rm_source=email" --engineTaxonomies "rm_mimetype=pdf"

# Negation (not equal)
adpgo taxonomy-statistic --engineName "myEngine" --engineTaxonomies "rm_source!=email"
```

---

## Raw Example Response (without category properties)

```json
{
  "executionId": "40d068f9-ea14-4316-a890-9b8b2b889f18",
  "taskType": "Taxonomy Statistic",
  "loggingEnabled": "true",
  "progressMax": 1,
  "executionStatus": "success",
  "executionRootDir": "E:\\MindServer\\Projects\\adp.adp\\adpRootDir",
  "contextId": "37178a68-90be-4dc0-a447-43b23784b6ed",
  "executionPersistent": "true",
  "progressCurrent": 1,
  "progressPercentage": 1.0,
  "taskDisplayName": "",
  "executionMetaData": {
    "adp_taxonomy_statistics_json_file_path": "E:\\MindServer\\Projects\\adp.adp\\adpRootDir\\taxonomy_stats.json",
    "adp_taxonomy_statistics_json_output": "{\"date\":\"Wed Mar 18 02:56:03 EDT 2026\",\"searchParameter\":[{\"key\":\"rm_main\",\"value\":\"[*]\"},{\"key\":\"rm_pagesize\",\"value\":\"[-1]\"}],\"statistics\":{\"taxonomy\":[{\"id\":\"rm_source\",\"category\":[{\"id\":\"file_demo_04\",\"displayName\":\"file_demo_04\",\"count\":761},{\"id\":\"new_demo_02\",\"displayName\":\"new_demo_02\",\"count\":2},{\"id\":\"file_dmoe_03\",\"displayName\":\"file_dmoe_03\",\"count\":1}]}]}}"
  }
}
```

### Decoded (without category properties)

```json
{
  "date": "Wed Mar 18 02:56:03 EDT 2026",
  "searchParameter": [
    { "key": "rm_main", "value": "[*]" },
    { "key": "rm_pagesize", "value": "[-1]" }
  ],
  "statistics": {
    "taxonomy": [
      {
        "id": "rm_source",
        "category": [
          { "id": "file_demo_04", "displayName": "file_demo_04", "count": 761 },
          { "id": "new_demo_02", "displayName": "new_demo_02", "count": 2 },
          { "id": "file_dmoe_03", "displayName": "file_dmoe_03", "count": 1 }
        ]
      }
    ]
  }
}
```

---

## Raw Example Response (with category properties)

```json
{
  "executionId": "78ad9874-b009-4102-bf2e-3f88c3246dfe",
  "taskType": "Taxonomy Statistic",
  "loggingEnabled": "true",
  "progressMax": 1,
  "executionStatus": "success",
  "executionRootDir": "E:\\MindServer\\Projects\\adp.adp\\adpRootDir",
  "contextId": "53fc1124-b35e-4f23-8fab-6e1515543e14",
  "executionPersistent": "true",
  "progressCurrent": 1,
  "progressPercentage": 1.0,
  "taskDisplayName": "",
  "executionMetaData": {
    "adp_taxonomy_statistics_json_file_path": "E:\\MindServer\\Projects\\adp.adp\\adpRootDir\\taxonomy_stats.json",
    "adp_taxonomy_statistics_json_output": "{\"date\":\"Wed Mar 18 03:03:06 EDT 2026\",\"searchParameter\":[{\"key\":\"rm_main\",\"value\":\"[*]\"},{\"key\":\"rm_pagesize\",\"value\":\"[-1]\"}],\"statistics\":{\"taxonomy\":[{\"id\":\"rm_source\",\"category\":[{\"id\":\"file_demo_04\",\"displayName\":\"file_demo_04\",\"count\":761,\"properties\":{\"rm_prop_editDate\":[\"1761795365031\"],\"rm_prop_creationDate\":[\"1761795365031\"],\"rm_prop_creator\":[\"system\"]}},{\"id\":\"new_demo_02\",\"displayName\":\"new_demo_02\",\"count\":2,\"properties\":{\"rm_prop_editDate\":[\"1761794623213\"],\"rm_prop_creationDate\":[\"1761794623213\"],\"rm_prop_creator\":[\"system\"]}},{\"id\":\"file_dmoe_03\",\"displayName\":\"file_dmoe_03\",\"count\":1,\"properties\":{\"rm_prop_editDate\":[\"1761620568340\"],\"rm_prop_creationDate\":[\"1761620568340\"],\"rm_prop_creator\":[\"system\"]}}]}]}}"
  }
}
```

---

## Decoded Result

### Result Type

```
TaxonomyStatisticResult {
    outputFile: string
    statistics: StatisticsDocument
}
```

### StatisticsDocument

```
StatisticsDocument {
    date: string
    searchParameter: SearchParameter[]
    statistics: Statistics
}

SearchParameter {
    key: string
    value: string
}

Statistics {
    taxonomy: TaxonomyEntry[]
}

TaxonomyEntry {
    id: string
    category: Category[]
}

Category {
    id: string
    displayName: string
    count: integer
    properties: Record<string, string[]> | absent  # present only when listCategoryProperties is enabled
}
```

### Decoding Rules

1. Map `executionMetaData.adp_taxonomy_statistics_json_file_path` to `outputFile`
2. Parse `executionMetaData.adp_taxonomy_statistics_json_output` as a JSON string into `StatisticsDocument`
3. `count` fields are integers in the parsed JSON
4. `properties` field is absent by default; when `listCategoryProperties` is enabled, it is a `Record<string, string[]>`

---

## executionMetaData Contract

| Field | Type | Description |
|-------|------|-------------|
| adp_taxonomy_statistics_json_file_path | string | Output file path |
| adp_taxonomy_statistics_json_output | string | JSON string containing taxonomy statistics — must be parsed |

### JSON String Fields

| Field | Parse As |
|-------|----------|
| adp_taxonomy_statistics_json_output | `StatisticsDocument` |

---

## Failure Response

On `executionStatus: "failed"`:

```json
{
  "executionId": "f9463001-dc1f-486a-a8a0-efaca8dd29cb",
  "taskType": "Taxonomy Statistic",
  "executionStatus": "failed",
  "errorMessage": "Error message details",
  "executionMetaData": null
}
```
