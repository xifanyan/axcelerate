# Publish To Review Task Specification

## Overview

| Property | Value |
|----------|-------|
| Task Type | `Publish To Review` |
| Description |  |
| Display Name | Task to publish stuff |
| Subcommand | `publish-to-review` |

---

## Semantic Inputs

These are the user-facing fields for the request-construction API. These are the fields callers set directly. Raw field names are implementation details.

| Field | Type | Default | Required | Description |
|-------|------|---------|----------|-------------|
| matterId | string | "default" | No | Matter ID |
| matterSpecificUrlRegEx | string | null | No | URL regex for matter-specific matching |
| searchDetails | any[] | [] | No | Search detail filters |
| searchString | string | null | No | Search string |
| waitForMatterCompletion | boolean | true | No | Wait for matter completion before publishing |
| secondsBetweenNextTryToWaitForMatterCompletion | integer | 100 | No | Polling interval in seconds while waiting for matter completion |
| ecaEngine | string | "{project_engine}" | No | ECA engine identifier |
| ecaApplication | string | "{project_name}" | No | ECA application identifier |
| ecaMasterHost | string | "" | No | ECA master host |
| ecaMasterPort | string | "" | No | ECA master port |
| mode | string | "all" | No | Publish mode |
| deduplicate | boolean | false | No | Deduplicate documents before publishing |
| enforceDeduplication | boolean | false | No | Enforce deduplication |
| startLearner | boolean | false | No | Start learner after publishing |
| ecaPublish | boolean | false | No | Enable ECA publish mode |
| matterSpecificApplication | string | null | No | Matter-specific application |

---

## Raw Default Configuration

> Configuration below shows **all fields with their exact default values** from [API-SPEC.md](../../API-SPEC.md). This is for reference only. Clients must not pre-populate all fields — only send fields explicitly set by the caller. See [request-construction.md](../request-construction.md).

```json
{
  "taskType": "Publish To Review",
  "taskConfiguration": {
    "adp_progressTaskTimeout": 0,
    "adp_ptr_matterSpecificUrlRegEx": null,
    "adp_taskActive": true,
    "adp_ptr_matterId": "default",
    "adp_executionPersistent": true,
    "adp_ptr_webserviceRequestTimeoutSeconds": 900,
    "adp_abortWfOnFailure": true,
    "adp_ptr_searchDetails": [],
    "adp_loggingEnabled": true,
    "adp_ptr_publishEngineId": "adp_publish_engine_id",
    "adp_ptr_startLearner": "false",
    "adp_ptr_searchString": null,
    "adp_ptr_waitForMatterCompletion": true,
    "adp_ptr_usedWebserviceUrl": "adp_url_used_for_publish_to_review",
    "adp_ptr_secondsBetweenNextTryToWaitForMatterCompletion": "100",
    "adp_ptr_ecaEngine": "{project_engine}",
    "adp_ptr_ecaMasterPort": "",
    "adp_ptr_httpsKeystoreFile": null,
    "adp_ptr_mode": "all",
    "adp_ptr_webserviceConnectTimeoutSeconds": 300,
    "adp_ptr_publishApplicationId": "adp_publish_application_id",
    "adp_ptr_httpsPassword": "",
    "adp_ptr_webservicePassword": "",
    "adp_ptr_webserviceUrl": "http://{host_name}/{project_name}",
    "adp_ptr_publishApplicationUrl": "adp_publish_application_url",
    "adp_ptr_enforceDeduplication": false,
    "adp_cleanUpHistory": false,
    "adp_ptr_matterSpecificApplication": null,
    "adp_ptr_webserviceUser": "{service_user}",
    "adp_taskTimeout": 0,
    "adp_ptr_deduplicate": false,
    "adp_ptr_ecaApplication": "{project_name}",
    "adp_ptr_httpsAllowUntrustedHosts": "false",
    "adp_ptr_ecaPublish": false,
    "adp_ptr_ecaMasterHost": "",
    "adp_ptr_httpsTrustCertificate": ""
  },
  "taskDescription": "",
  "taskDisplayName": "Task to publish stuff"
}
```

---

## Raw Example Request

> Example below matches **exactly** the default configuration from [API-SPEC.md](../../API-SPEC.md). This is the raw upstream shape.

```json
{
  "taskType": "Publish To Review",
  "taskConfiguration": {
    "adp_progressTaskTimeout": 0,
    "adp_ptr_matterSpecificUrlRegEx": null,
    "adp_taskActive": true,
    "adp_ptr_matterId": "default",
    "adp_executionPersistent": true,
    "adp_ptr_webserviceRequestTimeoutSeconds": 900,
    "adp_abortWfOnFailure": true,
    "adp_ptr_searchDetails": [],
    "adp_loggingEnabled": true,
    "adp_ptr_publishEngineId": "adp_publish_engine_id",
    "adp_ptr_startLearner": "false",
    "adp_ptr_searchString": null,
    "adp_ptr_waitForMatterCompletion": true,
    "adp_ptr_usedWebserviceUrl": "adp_url_used_for_publish_to_review",
    "adp_ptr_secondsBetweenNextTryToWaitForMatterCompletion": "100",
    "adp_ptr_ecaEngine": "{project_engine}",
    "adp_ptr_ecaMasterPort": "",
    "adp_ptr_httpsKeystoreFile": null,
    "adp_ptr_mode": "all",
    "adp_ptr_webserviceConnectTimeoutSeconds": 300,
    "adp_ptr_publishApplicationId": "adp_publish_application_id",
    "adp_ptr_httpsPassword": "",
    "adp_ptr_webservicePassword": "",
    "adp_ptr_webserviceUrl": "http://{host_name}/{project_name}",
    "adp_ptr_publishApplicationUrl": "adp_publish_application_url",
    "adp_ptr_enforceDeduplication": false,
    "adp_cleanUpHistory": false,
    "adp_ptr_matterSpecificApplication": null,
    "adp_ptr_webserviceUser": "{service_user}",
    "adp_taskTimeout": 0,
    "adp_ptr_deduplicate": false,
    "adp_ptr_ecaApplication": "{project_name}",
    "adp_ptr_httpsAllowUntrustedHosts": "false",
    "adp_ptr_ecaPublish": false,
    "adp_ptr_ecaMasterHost": "",
    "adp_ptr_httpsTrustCertificate": ""
  },
  "taskDescription": "",
  "taskDisplayName": "Task to publish stuff"
}
```

---

## CLI Arguments

See [cli.md](../cli.md) for global flags and naming conventions.

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--matter-id` | string | "default" | Matter ID |
| `--matter-specific-url-regex` | string | null | URL regex for matter-specific matching |
| `--search-details` | string | "[]" | JSON array of search detail filters |
| `--search-string` | string | null | Search string |
| `--wait-for-matter-completion` | boolean | true | Wait for matter completion before publishing |
| `--seconds-between-next-try-to-wait-for-matter-completion` | integer | 100 | Polling interval in seconds while waiting for matter completion |
| `--eca-engine` | string | "{project_engine}" | ECA engine identifier |
| `--eca-application` | string | "{project_name}" | ECA application identifier |
| `--eca-master-host` | string | "" | ECA master host |
| `--eca-master-port` | string | "" | ECA master port |
| `--mode` | string | "all" | Publish mode |
| `--deduplicate` | boolean | false | Deduplicate documents before publishing |
| `--enforce-deduplication` | boolean | false | Enforce deduplication |
| `--start-learner` | boolean | false | Start learner after publishing |
| `--eca-publish` | boolean | false | Enable ECA publish mode |
| `--matter-specific-application` | string | null | Matter-specific application |

### CLI Examples

```bash
# Basic
adpgo publish-to-review

# With matter ID
adpgo publish-to-review --matter-id Demo_Review1

# With search string and deduplication
adpgo publish-to-review --search-string "tag:review" --deduplicate --enforce-deduplication
```

---

## Raw Example Response

> The sample below is the observed client-facing response shape provided by the user. `ExecutionMetaData` is a byte array that must be converted to a UTF-8 string and then parsed as JSON.

```text
{ExecutionID:786deb13-a965-4f51-b4b9-fe5a0296e7d9 TaskType:Publish To Review LoggingEnabled:true ProgressMax:11 ExecutionStatus:success ExecutionRootDir:E:\MindServer\Projects\adp.adp\adpRootDir ContextID:71afda79-9408-46a1-883c-5b3553da6e26 ExecutionPersistent:true ProgressCurrent:11 ProgressPercentage:1 TaskDisplayName: ExecutionMetaData:[123 34 97 100 112 95 117 114 108 95 117 115 101 100 95 102 111 114 95 112 117 98 108 105 115 104 95 116 111 95 114 101 118 105 101 119 34 58 34 104 116 116 112 115 58 47 47 118 109 45 114 104 97 117 115 119 105 114 116 104 50 46 111 116 120 108 97 98 46 110 101 116 58 56 52 52 51 47 100 101 109 111 48 48 48 48 49 47 109 97 116 116 101 114 115 112 101 99 105 102 105 99 34 44 34 97 100 112 95 112 117 98 108 105 115 104 95 97 112 112 108 105 99 97 116 105 111 110 95 105 100 34 58 34 97 120 99 101 108 101 114 97 116 101 46 100 101 109 111 48 48 48 48 49 95 68 101 109 111 95 82 101 118 105 101 119 49 34 44 34 97 100 112 95 112 116 114 95 80 117 98 108 105 115 104 82 101 115 112 111 110 115 101 77 101 115 115 97 103 101 34 58 34 79 75 46 34 44 34 97 100 112 95 112 117 98 108 105 115 104 95 97 112 112 108 105 99 97 116 105 111 110 95 117 114 108 34 58 34 104 116 116 112 115 58 47 47 118 109 45 114 104 97 117 115 119 105 114 116 104 50 46 111 116 120 108 97 98 46 110 101 116 58 56 52 52 51 47 100 101 109 111 48 48 48 48 49 95 68 101 109 111 95 82 101 118 105 101 119 49 34 44 34 97 100 112 95 112 117 98 108 105 115 104 95 101 110 103 105 110 101 95 105 100 34 58 34 115 105 110 103 108 101 77 105 110 100 83 101 114 118 101 114 46 100 101 109 111 48 48 48 48 49 95 68 101 109 111 95 82 101 118 105 101 119 49 34 44 34 97 100 112 95 112 116 114 95 80 117 98 108 105 115 104 82 101 113 117 101 115 116 73 100 34 58 34 80 117 98 108 105 115 104 32 84 111 32 82 101 118 105 101 119 34 125]}
```

---

## Decoded Result

### Result Type

```
PublishToReviewResult {
    usedWebserviceUrl: string
    publishApplicationId: string
    publishApplicationUrl: string
    publishEngineId: string
    publishResponseMessage: string
    publishRequestId: string
}
```

### Decoding Rules

1. Read `ExecutionMetaData` from the response.
2. Convert the byte array to a UTF-8 string.
3. Parse the decoded UTF-8 string as JSON into an object.
4. Map `executionMetaData.adp_url_used_for_publish_to_review` to `usedWebserviceUrl`.
5. Map `executionMetaData.adp_publish_application_id` to `publishApplicationId`.
6. Map `executionMetaData.adp_publish_application_url` to `publishApplicationUrl`.
7. Map `executionMetaData.adp_publish_engine_id` to `publishEngineId`.
8. Map `executionMetaData.adp_ptr_PublishResponseMessage` to `publishResponseMessage`.
9. Map `executionMetaData.adp_ptr_PublishRequestId` to `publishRequestId`.

---

## executionMetaData Contract

| Field | Type | Description |
|-------|------|-------------|
| adp_url_used_for_publish_to_review | string | The URL used for publish to review |
| adp_publish_application_id | string | Published application identifier |
| adp_publish_application_url | string | Published application URL |
| adp_publish_engine_id | string | Published engine identifier |
| adp_ptr_PublishResponseMessage | string | Publish response message |
| adp_ptr_PublishRequestId | string | Publish request identifier |

### JSON String Fields

This task does not expose nested JSON string fields after byte-array decoding.

---

## Failure Response

On `executionStatus: "failed"`:

```json
{
  "executionId": "uuid",
  "taskType": "Publish To Review",
  "executionStatus": "failed",
  "errorMessage": "Error details",
  "executionMetaData": null
}
```

---

## Adding a New Task

1. Copy this template
2. Fill using [API-SPEC.md](../../API-SPEC.md)
3. Add entry to [index.md](./index.md) tasks table — **the tasks table must always reflect all current task specs**
4. Do NOT generate code — only update specs

---

## Rules

- Raw Default Configuration must match API-SPEC.md exactly (field names, values, ordering)
- Example Request must match Default Configuration exactly (no custom values)
- Preserved exact field names, values, and ordering from source
- Use camelCase for all response field names
- **Decoded Result types must be language-agnostic** — use TypeScript-like notation

### Type Notation Standard

| This notation | NOT this | Reason |
|---------------|----------|--------|
| `Type[]` | `[]Type` | TS/Go/Rust/Java/C# all understand `Type[]` |
| `Record<K, V>` | `map[K]V` | TypeScript/Rust standard; clearer than Go-specific |
| `string` | `String` | TypeScript/Rust/Python convention |
| `integer` | `int` | Language-neutral numeric type |
| `boolean` | `bool` | More widely recognized |
| `any` | `interface{}` | TypeScript/Rust `any`; Go-specific `interface{}` not allowed |
| `absent` | `null` | "absent" describes optionality; "null" is a value |
| `Cell[][]` | `[][]Cell` | Arrays of arrays in TS notation |
```
