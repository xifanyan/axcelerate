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
| engineName | string | "" | Conditional | Engine name |
| query | string | "*" | No | Query string |
| engineUserName | string | "" | No | Engine username |
| engineUserPassword | string | "" | No | Engine password |
| jobName | string | "" | No | Job name |
| jobDescription | string | "" | No | Job description |
| jobPriority | integer | 10 | No | Job priority |
| applicationIdentifier | string | "" | Conditional | Application identifier |
| applicationType | string | "" | No | Application type |
| wait | boolean | false | No | Wait for completion |
| engineType | string | "true" | No | Engine type |
| listOfJobProperties | string | "" | No | List of job properties |
| globalSearchJson | string | "" | No | Global search as JSON |
| globalSearchId | string | "" | No | Global search ID |
| restrictions | EngineTaxonomyArg[] | [] | No | Restrictions (see [common-types.md](../common-types.md)) |
| advancedRestrictions | EngineTaxonomyArg[] | [] | No | Advanced restrictions (see [common-types.md](../common-types.md)) |
| mainQueryType | string | null | No | Main query type |

> engineName and applicationIdentifier are mutually exclusive selectors. Exactly one must be provided.
>
> When applicationIdentifier is used, the client still treats it as the single effective selector, but for live ADP compatibility it explicitly serializes engineName as an empty string to clear the server-side default.
> This is a task-specific exception to the normal sparse-request rule described in [request-construction.md](../request-construction.md).
> Application-selected requests intentionally include `adp_createOcrJob_engineName: ""` alongside `adp_createOcrJob_applicationIdentifier` for live ADP compatibility.

---

## Raw Default Configuration

> Configuration below shows **all fields with their exact default values** from [API-SPEC.md](../../API-SPEC.md). This is for reference only. Clients must not pre-populate all fields. See [request-construction.md](../request-construction.md).
> These upstream defaults are shown as-is for reference. Real client-built requests must still provide exactly one of `engineName` or `applicationIdentifier`.

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
> This raw reference example mirrors upstream defaults. Real client-built requests must still provide exactly one of `engineName` or `applicationIdentifier`.

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

### Verified Example Request (from live API)

```json
{
  "taskType": "Create OCR Job",
  "taskConfiguration": {
    "adp_createOcrJob_engineName": "singleMindServer.demo00001",
    "adp_createOcrJob_jobName": "demo_ocr",
    "adp_createOcrJob_query": "*",
    "adp_createOcrJob_restrictions": [
      {
        "taxonomy": "rm_source",
        "negation": false,
        "query": "file_demo_04"
      },
      {
        "taxonomy": "meta_documentcharacteristics",
        "negation": false,
        "query": "Without+Text"
      },
      {
        "taxonomy": "rm_mimetype",
        "negation": false,
        "query": "image%2Ftiff OR application%2Fpdf"
      }
    ],
    "adp_createOcrJob_wait": "true"
  }
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
| `--restrictions` | string | "" | Restrictions (format: `Taxonomy=Query`, repeat for multiple) |
| `--advancedRestrictions` | string | "" | Advanced restrictions (format: `Taxonomy=Query`) |

> engineName and applicationIdentifier are mutually exclusive selectors. Exactly one must be provided.
>
> The same task-specific exception to the normal sparse-request rule applies to CLI-generated requests: when `--applicationIdentifier` is used, generated requests intentionally include `adp_createOcrJob_engineName: ""` alongside `adp_createOcrJob_applicationIdentifier` for live ADP compatibility.

### CLI Examples

```bash
# Basic
adpgo create-ocr-job --engineName "myEngine"

# With options
adpgo create-ocr-job --engineName "myEngine" --query "*"

# Using application identifier
adpgo create-ocr-job --applicationIdentifier "my-app-id" --jobName "demo_ocr" --query "*"

# Wait for completion
adpgo create-ocr-job --engineName "myEngine" --wait

# With restrictions
adpgo create-ocr-job --engineName "singleMindServer.demo00001" --jobName "demo_ocr" --restrictions "rm_source=file_demo_04" --restrictions "meta_documentcharacteristics=Without+Text"
```

---

## Raw Example Response

### Verified Async Initial Response (from live API)

```json
{
  "executionId": "fe0ead8d-5348-4c35-8caa-3d411c98974d",
  "taskType": "Create OCR Job",
  "loggingEnabled": "true",
  "progressMax": 0,
  "executionStatus": "",
  "executionRootDir": "E:\\MindServer\\Projects\\adp.adp\\adpRootDir",
  "contextId": "860e2cdc-24ec-424b-a369-a4ce29b4f39d",
  "executionPersistent": "true",
  "progressCurrent": 0,
  "progressPercentage": 0,
  "taskDisplayName": "",
  "executionMetaData": []
}
```

> **Note:** This task only supports async execution. The initial response returns `executionStatus: ""` (empty string) with empty `executionMetaData`. Poll using `executionId` to check completion.

---

## Decoded Result

### Result Type

On async completion, `executionMetaData` is an **empty array `[]`**. There is no meaningful result data to decode.

```
CreateOcrJobResult {
    // No meaningful result fields — executionMetaData is empty array on completion
}
```

### Decoding Rules

Poll using `GetTaskStatus` with the returned `executionId`. On completion, `executionMetaData` is an **empty array `[]`** — no fields to decode.

> **Pending verification:** Completion response `executionMetaData` shape is unconfirmed. The verified sample above reflects the **async initial response** only.

---

## executionMetaData Contract

### Async Initial Response (Verified)

| Field | Type | Description |
|-------|------|-------------|
| executionMetaData | `[]` | Empty array — no fields to decode |

### Async Completion Response (Pending Verification)

| Field | Type | Description |
|-------|------|-------------|
| executionMetaData | `[]` | Empty array — no fields to decode (unconfirmed) |

On failure:

| Field | Type | Description |
|-------|------|-------------|
| executionMetaData | null | Always null on failure |

---

## Failure Response

On `executionStatus: "failed"`:

```json
{
  "executionId": "fe0ead8d-5348-4c35-8caa-3d411c98974d",
  "taskType": "Create OCR Job",
  "executionStatus": "failed",
  "errorMessage": "Error message details",
  "executionMetaData": null
}
```
