# Taxonomy Statistic Task Specification

## Overview

| Property | Value |
|----------|-------|
| Task Type | `Taxonomy Statistic` |
| Description | Retrieves category counts for a taxonomy |
| Display Name | Taxonomy statistic |

---

## Default Configuration

> Configuration below shows **all fields with their exact default values** from [[API-SPEC.md]]

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

## Example Request

> Example below matches **exactly** the default configuration from API-SPEC.md

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

## Example Response

```json
{
  "ExecutionID": "f9463001-dc1f-486a-a8a0-efaca8dd29cb",
  "TaskType": "Taxonomy Statistic",
  "LoggingEnabled": false,
  "ProgressMax": 1,
  "ExecutionStatus": "success",
  "ExecutionRootDir": "E:\\MindServer\\Projects\\adp.adp\\adpRootDir",
  "ContextID": "2e5a47e4-d9c8-4547-aaba-45c0a3774d47",
  "ExecutionPersistent": false,
  "ProgressCurrent": 1,
  "ProgressPercentage": 1,
  "TaskDisplayName": "",
  "ExecutionMetaData": {
    "adp_taxonomy_statistics_json_file_path": "E:\\MindServer\\Projects\\adp.adp\\adpRootDir\\taxonomy_stats.json",
    "adp_taxonomy_statistics_json_output": "[{\"category\":\"/documents/reports\",\"count\":150},{\"category\":\"/documents/invoices\",\"count\":75}]"
  }
}
```

---

## Response Fields

All responses include the common fields. Taxonomy Statistic-specific `ExecutionMetaData` fields:

| Field | Type | Description |
|-------|------|-------------|
| adp_taxonomy_statistics_json_file_path | string | Output file path |
| adp_taxonomy_statistics_json_output | string | JSON string containing taxonomy statistics |
