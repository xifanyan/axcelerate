# List Entities Task Specification

## Overview

| Property | Value |
|----------|-------|
| Task Type | `List Entities` |
| Description | Writes a list of entities ot an output variable |
| Display Name | List entities |

---

## Default Configuration

> Configuration below shows **all fields with their exact default values** from [API-SPEC.md](../../API-SPEC.md)

```json
{
  "taskType": "List Entities",
  "taskConfiguration": {
    "adp_progressTaskTimeout": 0,
    "adp_listEntities_file": "output.json",
    "adp_listEntities_numberOfEntities": "-1",
    "adp_listEntities_axcRequestTimeoutSeconds": 900,
    "adp_taskActive": true,
    "adp_listEntities_userHasAccess": "",
    "adp_listEntities_whiteList": "id,displayName",
    "adp_executionPersistent": true,
    "adp_abortWfOnFailure": true,
    "adp_listEntities_relatedEntity": "",
    "adp_listEntities_workspace": "",
    "adp_loggingEnabled": true,
    "adp_listEntities_status": "",
    "adp_listEntities_axcServiceCoreAddress": "",
    "adp_listEntities_relatedEntityType": "",
    "adp_listEntities_type": "",
    "adp_listEntities_httpsKeystoreFile": null,
    "adp_listEntities_httpsPassword": "",
    "adp_listEntities_axcConnectTimeoutSeconds": 300,
    "adp_listEntities_axcServicePassword": "",
    "adp_listEntities_startingEntity": "1",
    "adp_listEntities_outputJson": "adp_entities_json_output",
    "adp_cleanUpHistory": false,
    "adp_listEntities_descriptionSettingFilterValueDateFormat": "yyyy-MM-dd",
    "adp_listEntities_descriptionFilters": [],
    "adp_listEntities_axcServiceUser": "",
    "adp_listEntities_axcFields": "",
    "adp_taskTimeout": 0,
    "adp_listEntities_httpsTrustCertificate": "",
    "adp_listEntities_host": "",
    "adp_listEntities_outputFilename": "adp_entities_output_file_name",
    "adp_listEntities_id": "",
    "adp_listEntities_httpsAllowUntrustedHosts": "true"
  },
  "taskDescription": "Writes a list of entities ot an output variable",
  "taskDisplayName": "List entities"
}
```

---

## Field Reference (with Defaults)

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| adp_progressTaskTimeout | integer | 0 | Progress task timeout |
| adp_listEntities_file | string | "output.json" | Output file name |
| adp_listEntities_numberOfEntities | string | "-1" | Number of entities to retrieve |
| adp_listEntities_axcRequestTimeoutSeconds | integer | 900 | Request timeout in seconds |
| adp_taskActive | boolean | true | Whether task is active |
| adp_listEntities_userHasAccess | string | "" | User access filter |
| adp_listEntities_whiteList | string | "id,displayName" | Fields to include |
| adp_executionPersistent | boolean | true | Persist execution |
| adp_abortWfOnFailure | boolean | true | Abort workflow on failure |
| adp_listEntities_relatedEntity | string | "" | Related entity ID |
| adp_listEntities_workspace | string | "" | Workspace name |
| adp_loggingEnabled | boolean | true | Enable logging |
| adp_listEntities_status | string | "" | Entity status filter |
| adp_listEntities_axcServiceCoreAddress | string | "" | Service core address |
| adp_listEntities_relatedEntityType | string | "" | Related entity type |
| adp_listEntities_type | string | "" | Entity type to filter |
| adp_listEntities_httpsKeystoreFile | null | null | HTTPS keystore file |
| adp_listEntities_httpsPassword | string | "" | HTTPS password |
| adp_listEntities_axcConnectTimeoutSeconds | integer | 300 | Connection timeout |
| adp_listEntities_axcServicePassword | string | "" | Service password |
| adp_listEntities_startingEntity | string | "1" | Starting entity index |
| adp_listEntities_outputJson | string | "adp_entities_json_output" | Output variable name |
| adp_cleanUpHistory | boolean | false | Clean up history |
| adp_listEntities_descriptionSettingFilterValueDateFormat | string | "yyyy-MM-dd" | Date format |
| adp_listEntities_descriptionFilters | array | [] | Description filters |
| adp_listEntities_axcServiceUser | string | "" | Service username |
| adp_listEntities_axcFields | string | "" | Fields to retrieve |
| adp_taskTimeout | integer | 0 | Task timeout |
| adp_listEntities_httpsTrustCertificate | string | "" | HTTPS trust certificate |
| adp_listEntities_host | string | "" | Target host URL |
| adp_listEntities_outputFilename | string | "adp_entities_output_file_name" | Output filename parameter |
| adp_listEntities_id | string | "" | Specific entity ID |
| adp_listEntities_httpsAllowUntrustedHosts | string | "true" | Allow untrusted hosts |

---

## Example Request

> Example below matches **exactly** the default configuration from API-SPEC.md

```json
{
  "taskType": "List Entities",
  "taskConfiguration": {
    "adp_progressTaskTimeout": 0,
    "adp_listEntities_file": "output.json",
    "adp_listEntities_numberOfEntities": "-1",
    "adp_listEntities_axcRequestTimeoutSeconds": 900,
    "adp_taskActive": true,
    "adp_listEntities_userHasAccess": "",
    "adp_listEntities_whiteList": "id,displayName",
    "adp_executionPersistent": true,
    "adp_abortWfOnFailure": true,
    "adp_listEntities_relatedEntity": "",
    "adp_listEntities_workspace": "",
    "adp_loggingEnabled": true,
    "adp_listEntities_status": "",
    "adp_listEntities_axcServiceCoreAddress": "",
    "adp_listEntities_relatedEntityType": "",
    "adp_listEntities_type": "",
    "adp_listEntities_httpsKeystoreFile": null,
    "adp_listEntities_httpsPassword": "",
    "adp_listEntities_axcConnectTimeoutSeconds": 300,
    "adp_listEntities_axcServicePassword": "",
    "adp_listEntities_startingEntity": "1",
    "adp_listEntities_outputJson": "adp_entities_json_output",
    "adp_cleanUpHistory": false,
    "adp_listEntities_descriptionSettingFilterValueDateFormat": "yyyy-MM-dd",
    "adp_listEntities_descriptionFilters": [],
    "adp_listEntities_axcServiceUser": "",
    "adp_listEntities_axcFields": "",
    "adp_taskTimeout": 0,
    "adp_listEntities_httpsTrustCertificate": "",
    "adp_listEntities_host": "",
    "adp_listEntities_outputFilename": "adp_entities_output_file_name",
    "adp_listEntities_id": "",
    "adp_listEntities_httpsAllowUntrustedHosts": "true"
  },
  "taskDescription": "Writes a list of entities ot an output variable",
  "taskDisplayName": "List entities"
}
```

---

## Example Response

```json
{
  "ExecutionID": "f9463001-dc1f-486a-a8a0-efaca8dd29cb",
  "TaskType": "List Entities",
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
    "adp_entities_output_file_name": "E:\\MindServer\\Projects\\adp.adp\\adpRootDir\\output.json",
    "adp_entities_json_output": "[{\"id\":\"axcelerate.demo00001_Demo_Review4\",\"displayName\":\"demo00001_Demo_Review4\",\"processStatus\":\"Running\",\"hostId\":\"vm-rhauswirth2.otxlab.net\",\"hostName\":\"vm-rhauswirth2.otxlab.net\",\"sourceForCreateFromExisting\":false},{\"id\":\"axcelerate.demo_01_review\",\"displayName\":\"demo_01_review\",\"processStatus\":\"Killed\",\"hostId\":\"vm-rhauswirth2.otxlab.net\",\"hostName\":\"vm-rhauswirth2.otxlab.net\",\"sourceForCreateFromExisting\":false}]"
  }
}
```

---

## Response Fields

| Field | Type | Description |
|-------|------|-------------|
| ExecutionID | string | Unique execution identifier |
| TaskType | string | Task type ("List Entities") |
| LoggingEnabled | boolean | Whether logging is enabled |
| ProgressMax | integer | Maximum progress value |
| ExecutionStatus | string | Status of execution ("success", "failed", etc.) |
| ExecutionRootDir | string | Root directory for execution |
| ContextID | string | Context identifier |
| ExecutionPersistent | boolean | Whether execution is persistent |
| ProgressCurrent | integer | Current progress value |
| ProgressPercentage | integer | Progress percentage |
| TaskDisplayName | string | Display name of the task |
| ExecutionMetaData | object | Task-specific metadata |
| ExecutionMetaData.adp_entities_output_file_name | string | Output file path |
| ExecutionMetaData.adp_entities_json_output | string | JSON string of entities array |

---

## Entity Fields

> The `adp_entities_json_output` field in `ExecutionMetaData` is a JSON string that must be parsed. Each entity contains:

| Field | Type | Description |
|-------|------|-------------|
| id | string | Entity identifier |
| displayName | string | Entity display name |
| processStatus | string | Process status (e.g., "Running", "Killed", "Not running") |
| hostId | string | Host ID |
| hostName | string | Host name |
| sourceForCreateFromExisting | boolean | Whether source is for creating from existing |

---

## Parsing Example

```json
// ExecutionMetaData.adp_entities_json_output contains:
"[{\"id\":\"axcelerate.demo00001_Demo_Review4\",\"displayName\":\"demo00001_Demo_Review4\",\"processStatus\":\"Running\",...}]"

// After parsing as JSON array:
[
  {
    "id": "axcelerate.demo00001_Demo_Review4",
    "displayName": "demo00001_Demo_Review4",
    "processStatus": "Running",
    "hostId": "vm-rhauswirth2.otxlab.net",
    "hostName": "vm-rhauswirth2.otxlab.net",
    "sourceForCreateFromExisting": false
  }
]
```
