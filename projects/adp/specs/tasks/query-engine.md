# Query Engine Task Specification

## Overview

| Property | Value |
|----------|-------|
| Task Type | `Query Engine` |
| Description | Queries an engine |
| Display Name | Query engine |

---

## Default Configuration

> Configuration below shows **all fields with their exact default values** from [API-SPEC.md](../../API-SPEC.md)

```json
{
  "taskType": "Query Engine",
  "taskConfiguration": {
    "adp_progressTaskTimeout": 0,
    "adp_queryEngine_fieldName": "virtual_filesize",
    "adp_queryEngine_enableSiblingExpansion": "false",
    "adp_queryEngine_engineName": "{adp_engineName}",
    "adp_taskActive": true,
    "adp_executionPersistent": true,
    "adp_queryEngine_engineUserPassword": "",
    "adp_abortWfOnFailure": true,
    "adp_loggingEnabled": true,
    "adp_queryEngine_engineTaxonomies": [],
    "adp_queryEngine_engineUserName": "{adp_user}",
    "adp_queryEngine_engineType": "true",
    "adp_queryEngine_saveVariable": "{engine_save}",
    "adp_queryEngine_categoryToDelete": "",
    "adp_queryEngine_activateCategoryDeletion": false,
    "adp_queryEngine_applicationIdentifier": "",
    "adp_queryEngine_taxonomyToDelete": "",
    "adp_queryEngine_successIfCountIs": "{adp_expectedDsDoccount}",
    "adp_queryEngine_category": "",
    "adp_queryEngine_activateTagging": false,
    "adp_queryEngine_globalSearchId": "",
    "adp_queryEngine_aggregatedValue": "adp_query_engine_aggregated_value",
    "adp_queryEngine_AdvancedRestrictions": [],
    "adp_queryEngine_taxonomy": "",
    "adp_queryEngine_globalSearchJson": "",
    "adp_queryEngine_saveCompareString": "true",
    "adp_cleanUpHistory": false,
    "adp_queryEngine_numberOfDocuments": "adp_query_engine_documents_count",
    "adp_queryEngine_siblingFields": "rm_attachmentroot",
    "adp_queryEngine_engineQuery": "*",
    "adp_queryEngine_mainQueryType": null,
    "adp_queryEngine_waitForResult": false,
    "adp_queryEngine_categoryDisplayName": "",
    "adp_queryEngine_waitWhileCountIs": "{adp_oldDsDoccount}",
    "adp_taskTimeout": 0,
    "adp_queryEngine_applicationType": "",
    "adp_queryEngine_exitOnValueChanged": true
  },
  "taskDescription": "Queries an engine",
  "taskDisplayName": "Query engine"
}
```

---

## Field Reference (with Defaults)

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| adp_progressTaskTimeout | integer | 0 | Progress task timeout |
| adp_queryEngine_fieldName | string | "virtual_filesize" | Field name |
| adp_queryEngine_enableSiblingExpansion | string | "false" | Enable sibling expansion |
| adp_queryEngine_engineName | string | "{adp_engineName}" | Engine name |
| adp_taskActive | boolean | true | Whether task is active |
| adp_executionPersistent | boolean | true | Persist execution |
| adp_queryEngine_engineUserPassword | string | "" | Engine password |
| adp_abortWfOnFailure | boolean | true | Abort workflow on failure |
| adp_loggingEnabled | boolean | true | Enable logging |
| adp_queryEngine_engineTaxonomies | array | [] | Engine taxonomies |
| adp_queryEngine_engineUserName | string | "{adp_user}" | Engine username |
| adp_queryEngine_engineType | string | "true" | Engine type |
| adp_queryEngine_saveVariable | string | "{engine_save}" | Variable to save results |
| adp_queryEngine_categoryToDelete | string | "" | Category to delete |
| adp_queryEngine_activateCategoryDeletion | boolean | false | Activate category deletion |
| adp_queryEngine_applicationIdentifier | string | "" | Application identifier |
| adp_queryEngine_taxonomyToDelete | string | "" | Taxonomy to delete |
| adp_queryEngine_successIfCountIs | string | "{adp_expectedDsDoccount}" | Success if count matches |
| adp_queryEngine_category | string | "" | Query category |
| adp_queryEngine_activateTagging | boolean | false | Enable tagging |
| adp_queryEngine_globalSearchId | string | "" | Global search ID |
| adp_queryEngine_aggregatedValue | string | "adp_query_engine_aggregated_value" | Aggregated value output |
| adp_queryEngine_AdvancedRestrictions | array | [] | Advanced restrictions |
| adp_queryEngine_taxonomy | string | "" | Taxonomy filter |
| adp_queryEngine_globalSearchJson | string | "" | Global search JSON |
| adp_queryEngine_saveCompareString | string | "true" | Save compare string |
| adp_cleanUpHistory | boolean | false | Clean up history |
| adp_queryEngine_numberOfDocuments | string | "adp_query_engine_documents_count" | Number of documents output |
| adp_queryEngine_siblingFields | string | "rm_attachmentroot" | Sibling fields |
| adp_queryEngine_engineQuery | string | "*" | Query string |
| adp_queryEngine_mainQueryType | null | null | Main query type |
| adp_queryEngine_waitForResult | boolean | false | Wait for result |
| adp_queryEngine_categoryDisplayName | string | "" | Category display name |
| adp_queryEngine_waitWhileCountIs | string | "{adp_oldDsDoccount}" | Wait while count matches |
| adp_taskTimeout | integer | 0 | Task timeout |
| adp_queryEngine_applicationType | string | "" | Application type |
| adp_queryEngine_exitOnValueChanged | boolean | true | Exit on value changed |

---

## Example Request

> Example below matches **exactly** the default configuration from API-SPEC.md

```json
{
  "taskType": "Query Engine",
  "taskConfiguration": {
    "adp_progressTaskTimeout": 0,
    "adp_queryEngine_fieldName": "virtual_filesize",
    "adp_queryEngine_enableSiblingExpansion": "false",
    "adp_queryEngine_engineName": "{adp_engineName}",
    "adp_taskActive": true,
    "adp_executionPersistent": true,
    "adp_queryEngine_engineUserPassword": "",
    "adp_abortWfOnFailure": true,
    "adp_loggingEnabled": true,
    "adp_queryEngine_engineTaxonomies": [],
    "adp_queryEngine_engineUserName": "{adp_user}",
    "adp_queryEngine_engineType": "true",
    "adp_queryEngine_saveVariable": "{engine_save}",
    "adp_queryEngine_categoryToDelete": "",
    "adp_queryEngine_activateCategoryDeletion": false,
    "adp_queryEngine_applicationIdentifier": "",
    "adp_queryEngine_taxonomyToDelete": "",
    "adp_queryEngine_successIfCountIs": "{adp_expectedDsDoccount}",
    "adp_queryEngine_category": "",
    "adp_queryEngine_activateTagging": false,
    "adp_queryEngine_globalSearchId": "",
    "adp_queryEngine_aggregatedValue": "adp_query_engine_aggregated_value",
    "adp_queryEngine_AdvancedRestrictions": [],
    "adp_queryEngine_taxonomy": "",
    "adp_queryEngine_globalSearchJson": "",
    "adp_queryEngine_saveCompareString": "true",
    "adp_cleanUpHistory": false,
    "adp_queryEngine_numberOfDocuments": "adp_query_engine_documents_count",
    "adp_queryEngine_siblingFields": "rm_attachmentroot",
    "adp_queryEngine_engineQuery": "*",
    "adp_queryEngine_mainQueryType": null,
    "adp_queryEngine_waitForResult": false,
    "adp_queryEngine_categoryDisplayName": "",
    "adp_queryEngine_waitWhileCountIs": "{adp_oldDsDoccount}",
    "adp_taskTimeout": 0,
    "adp_queryEngine_applicationType": "",
    "adp_queryEngine_exitOnValueChanged": true
  },
  "taskDescription": "Queries an engine",
  "taskDisplayName": "Query engine"
}
```

---

## Example Response

```json
{
  "executionId": "f9463001-dc1f-486a-a8a0-efaca8dd29cb",
  "taskType": "Query Engine",
  "loggingEnabled": "false",
  "progressMax": 1,
  "executionStatus": "success",
  "executionRootDir": "E:\\MindServer\\Projects\\adp.adp\\adpRootDir",
  "contextId": "2e5a47e4-d9c8-4547-aaba-45c0a3774d47",
  "executionPersistent": "false",
  "progressCurrent": 1,
  "progressPercentage": 1.0,
  "taskDisplayName": "",
  "executionMetaData": {
    "adp_query_engine_documents_count": "100",
    "adp_query_engine_aggregated_value": "500"
  }
}
```

---

## Response Fields

All responses include the common fields. Query Engine-specific `ExecutionMetaData` fields:

| Field | Type | Description |
|-------|------|-------------|
| adp_query_engine_documents_count | string | Number of documents matching the query |
| adp_query_engine_aggregated_value | string | Aggregated value result |
