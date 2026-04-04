# Create OCR Job Task Specification

## Overview

| Property | Value |
|----------|-------|
| Task Type | `Create OCR Job` |
| Description | Changes metaData by using regEx replacement |
| Display Name | Create OCR Job |
| Subcommand | `create-ocr-job` |

---

## Semantic Inputs

These are the user-facing fields for the request-construction API.

| Field | Type | Default | Required | Description |
|-------|------|---------|----------|-------------|
| engineName | string | "" | No | Engine name |
| query | string | "*" | No | Query string |
| engineUserName | string | "" | No | Engine username |
| engineUserPassword | string | "" | No | Engine password |
| jobName | string | "" | No | Job name |
| jobDescription | string | "" | No | Job description |
| jobPriority | integer | 10 | No | Job priority |
| applicationIdentifier | string | "" | No | Application identifier |
| applicationType | string | "" | No | Application type |
| wait | boolean | false | No | Wait for completion |
| engineType | string | "true" | No | Engine type |
| listOfJobProperties | string | "" | No | List of job properties |
| globalSearchJson | string | "" | No | Global search as JSON |
| globalSearchId | string | "" | No | Global search ID |
| restrictions | array | [] | No | Restrictions |
| advancedRestrictions | array | [] | No | Advanced restrictions |
| mainQueryType | string | null | No | Main query type |

---

## Raw Default Configuration

> Configuration below shows **all fields with their exact default values** from [API-SPEC.md](../../API-SPEC.md). This is for reference only. Clients must not pre-populate all fields. See [request-construction.md](../request-construction.md).

```json
{
  "taskType": "Create OCR Job",
  "taskConfiguration": {
    "adp_progressTaskTimeout": 0,
    "adp_createOcrJob_engineUserPassword": "",
    "adp_createOcrJob_query": "*",
    "adp_taskActive": true,
    "adp_createOcrJob_listOfJobProperties": "",
    "adp_executionPersistent": true,
    "adp_createOcrJob_engineType": "true",
    "adp_abortWfOnFailure": true,
    "adp_loggingEnabled": true,
    "adp_createOcrJob_AdvancedRestrictions": [],
    "adp_createOcrJob_globalSearchJson": "",
    "adp_createOcrJob_wait": "false",
    "adp_createOcrJob_engineName": "",
    "adp_createOcrJob_jobDescription": "",
    "adp_createOcrJob_applicationIdentifier": "",
    "adp_createOcrJob_jobPriority": "10",
    "adp_createOcrJob_jobName": "",
    "adp_createOcrJob_restrictions": [],
    "adp_cleanUpHistory": false,
    "adp_createOcrJob_engineUserName": "",
    "adp_createOcrJob_mainQueryType": null,
    "adp_createOcrJob_applicationType": "",
    "adp_createOcrJob_globalSearchId": "",
    "adp_taskTimeout": 0,
    "adp_createOcrJob_jsonOutputVariable": "adp_createOcrJob_json_output"
  },
  "taskDescription": "Changes metaData by using regEx replacement.",
  "taskDisplayName": "Create OCR Job"
}
```

---

## Field Reference (with Defaults)

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| adp_progressTaskTimeout | integer | 0 | Progress task timeout |
| adp_createOcrJob_engineUserPassword | string | "" | Engine user password |
| adp_createOcrJob_query | string | "*" | Query string |
| adp_taskActive | boolean | true | Whether task is active |
| adp_createOcrJob_listOfJobProperties | string | "" | List of job properties |
| adp_executionPersistent | boolean | true | Persist execution |
| adp_createOcrJob_engineType | string | "true" | Engine type |
| adp_abortWfOnFailure | boolean | true | Abort workflow on failure |
| adp_loggingEnabled | boolean | true | Enable logging |
| adp_createOcrJob_AdvancedRestrictions | array | [] | Advanced restrictions |
| adp_createOcrJob_globalSearchJson | string | "" | Global search as JSON |
| adp_createOcrJob_wait | string | "false" | Wait for completion — coerce to boolean |
| adp_createOcrJob_engineName | string | "" | Engine name |
| adp_createOcrJob_jobDescription | string | "" | Job description |
| adp_createOcrJob_applicationIdentifier | string | "" | Application identifier |
| adp_createOcrJob_jobPriority | string | "10" | Job priority — coerce to integer |
| adp_createOcrJob_jobName | string | "" | Job name |
| adp_createOcrJob_restrictions | array | [] | Restrictions |
| adp_cleanUpHistory | boolean | false | Clean up history |
| adp_createOcrJob_engineUserName | string | "" | Engine username |
| adp_createOcrJob_mainQueryType | string | null | Main query type |
| adp_createOcrJob_applicationType | string | "" | Application type |
| adp_createOcrJob_globalSearchId | string | "" | Global search ID |
| adp_taskTimeout | integer | 0 | Task timeout |
| adp_createOcrJob_jsonOutputVariable | string | "adp_createOcrJob_json_output" | Output variable name |

---

## Raw Example Request

> Example below matches **exactly** the default configuration from [API-SPEC.md](../../API-SPEC.md).

```json
{
  "taskType": "Create OCR Job",
  "taskConfiguration": {
    "adp_progressTaskTimeout": 0,
    "adp_createOcrJob_engineUserPassword": "",
    "adp_createOcrJob_query": "*",
    "adp_taskActive": true,
    "adp_createOcrJob_listOfJobProperties": "",
    "adp_executionPersistent": true,
    "adp_createOcrJob_engineType": "true",
    "adp_abortWfOnFailure": true,
    "adp_loggingEnabled": true,
    "adp_createOcrJob_AdvancedRestrictions": [],
    "adp_createOcrJob_globalSearchJson": "",
    "adp_createOcrJob_wait": "false",
    "adp_createOcrJob_engineName": "",
    "adp_createOcrJob_jobDescription": "",
    "adp_createOcrJob_applicationIdentifier": "",
    "adp_createOcrJob_jobPriority": "10",
    "adp_createOcrJob_jobName": "",
    "adp_createOcrJob_restrictions": [],
    "adp_cleanUpHistory": false,
    "adp_createOcrJob_engineUserName": "",
    "adp_createOcrJob_mainQueryType": null,
    "adp_createOcrJob_applicationType": "",
    "adp_createOcrJob_globalSearchId": "",
    "adp_taskTimeout": 0,
    "adp_createOcrJob_jsonOutputVariable": "adp_createOcrJob_json_output"
  },
  "taskDescription": "Changes metaData by using regEx replacement.",
  "taskDisplayName": "Create OCR Job"
}
```

---

## CLI Arguments

See [cli.md](../cli.md) for global flags and naming conventions.

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--engineName` | string | "" | Engine name |
| `--query` | string | "*" | Query string |
| `--engineUserName` | string | "" | Engine username |
| `--engineUserPassword` | string | "" | Engine password |
| `--jobName` | string | "" | Job name |
| `--jobDescription` | string | "" | Job description |
| `--jobPriority` | integer | 10 | Job priority |
| `--applicationIdentifier` | string | "" | Application identifier |
| `--applicationType` | string | "" | Application type |
| `--wait` | boolean | false | Wait for completion |
| `--engineType` | string | "true" | Engine type |
| `--listOfJobProperties` | string | "" | List of job properties |
| `--globalSearchJson` | string | "" | Global search as JSON |
| `--globalSearchId` | string | "" | Global search ID |

### CLI Examples

```bash
# Basic
adpgo create-ocr-job

# With options
adpgo create-ocr-job --engineName "myEngine" --query "*"

# Wait for completion
adpgo create-ocr-job --engineName "myEngine" --wait
```

---

## Raw Example Response

> **Pending**: `executionMetaData` response fields not yet verified against actual API response.

```json
{
  "executionId": "uuid",
  "taskType": "Create OCR Job",
  "loggingEnabled": "false",
  "progressMax": 1,
  "executionStatus": "success",
  "executionRootDir": "path",
  "contextId": "uuid",
  "executionPersistent": "false",
  "progressCurrent": 1,
  "progressPercentage": 1.0,
  "taskDisplayName": "",
  "executionMetaData": {
    "adp_createOcrJob_json_output": "{...}"
  }
}
```

---

## Decoded Result

### Result Type

> **Pending**: `executionMetaData` response fields not yet verified against actual API response.

```
CreateOcrJobResult {
    output: any  # parsed from adp_createOcrJob_json_output — pending verification
}
```

### Decoding Rules

> **Pending**: Awaiting verification against actual API response.

---

## executionMetaData Contract

> **Pending**: `executionMetaData` response fields not yet verified against actual API response.

| Field | Type | Description |
|-------|------|-------------|
| adp_createOcrJob_json_output | string | JSON output — pending verification |

### JSON String Fields

| Field | Parse As |
|-------|----------|
| adp_createOcrJob_json_output | `any` — pending verification |

---

## Failure Response

On `executionStatus: "failed"`:

```json
{
  "executionId": "uuid",
  "taskType": "Create OCR Job",
  "executionStatus": "failed",
  "errorMessage": "Error message details",
  "executionMetaData": null
}
```
