# Create Review Interface Task Specification

## Overview

| Property | Value |
|----------|-------|
| Task Type | `Create Review Interface` |
| Description | Task to create a review interface |
| Display Name | Task to create a review interface |
| Subcommand | `create-review-interface` |

---

## Semantic Inputs

These are the user-facing fields for the request-construction API. These are the fields callers set directly. Raw field names are implementation details.

| Field | Type | Default | Required | Description |
|-------|------|---------|----------|-------------|
| matterId | string | "default" | No | Matter ID |
| engineHostMemoryLimit | integer | 0 | No | Engine host memory limit |
| engineHostMemoryLimitRatio | integer | 0 | No | Engine host memory limit ratio |
| appHostMemoryLimit | integer | 0 | No | Application host memory limit |
| appHostMemoryLimitRatio | integer | 0 | No | Application host memory limit ratio |
| expectedNumberOfDocuments | integer | 1000000 | No | Expected number of documents |
| predictiveCodingEnabled | boolean | false | No | Enable predictive coding |
| reduceMemoryFootPrint | boolean | false | No | Reduce memory footprint |
| matterSpecificApplication | string | null | No | Matter-specific application |
| matterSpecificEngineHostId | string | "localhost" | No | Matter-specific engine host ID |
| matterSpecificApplicationHostId | string | "" | No | Matter-specific application host ID |
| matterSpecificWorkspace | string | null | No | Matter-specific workspace |
| matterSpecificTemplate | string | null | No | Matter-specific template |
| matterSpecificSmallProject | boolean | false | No | Matter is a small project |
| matterSpecificUrlRegEx | string | null | No | URL regex for matter-specific matching |
| appHostDetection | boolean | true | No | Enable application host detection |

---

## Raw Default Configuration

> Configuration below shows **all fields with their exact default values** from [API-SPEC.md](../../API-SPEC.md). This is for reference only. Clients must not pre-populate all fields — only send fields explicitly set by the caller. See [request-construction.md](../request-construction.md).

```json
{
  "taskType": "Create Review Interface",
  "taskConfiguration": {
    "adp_cri_appHostMemoryLimitRatio": "0",
    "adp_cri_matterSpecificWorkspace": null,
    "adp_progressTaskTimeout": 0,
    "adp_cri_matterId": "default",
    "adp_cri_matterSpecificEngineHostId": "localhost",
    "adp_cri_httpsPassword": "",
    "adp_cri_matterSpecificTemplate": null,
    "adp_taskActive": true,
    "adp_cri_reduceMemoryFootPrint": false,
    "adp_executionPersistent": true,
    "adp_abortWfOnFailure": true,
    "adp_cri_engineHostMemoryLimit": "0",
    "adp_loggingEnabled": true,
    "adp_cri_chosenEngineHostMemoryRatio": "adp_create_review_engine_host_memory_ratio",
    "adp_cri_webserviceRequestTimeoutSeconds": 900,
    "adp_cri_webserviceUrl": "http://{host_name}/{project_name}",
    "adp_cri_publishEngineId": "adp_created_publish_engine_id",
    "adp_cri_appHostMemoryLimit": "0",
    "adp_cri_publishApplicationId": "adp_created_publish_application_id",
    "adp_cri_chosenApplicationHostMemory": "adp_create_review_application_host_memory",
    "adp_cri_chosenEngineHostMemory": "adp_create_review_engine_host_memory",
    "adp_cri_chosenEngineHostNameParameter": "adp_create_review_engine_host",
    "adp_cri_httpsKeystoreFile": null,
    "adp_cri_matterSpecificApplicationHostId": "",
    "adp_cri_matterSpecificApplication": null,
    "adp_cri_usedWebserviceUrl": "adp_url_used_for_create_review_interface",
    "adp_cri_appHostDetection": "true",
    "adp_cri_matterSpecificSmallProject": false,
    "adp_cri_engineHostMemoryLimitRatio": "0",
    "adp_cri_expectedNumberOfDocuments": 1000000,
    "adp_cri_predictiveCodingEnabled": false,
    "adp_cri_chosenApplicationHostNameParameter": "adp_create_review_application_host",
    "adp_cleanUpHistory": false,
    "adp_cri_httpsTrustCertificate": "",
    "adp_cri_httpsAllowUntrustedHosts": "false",
    "adp_cri_matterSpecificUrlRegEx": null,
    "adp_cri_chosenApplicationHostMemoryRatio": "adp_create_review_application_host_memory_ratio",
    "adp_taskTimeout": 0,
    "adp_cri_webserviceUser": "{service_user}",
    "adp_cri_webserviceConnectTimeoutSeconds": 300,
    "adp_cri_webservicePassword": null
  },
  "taskDescription": "",
  "taskDisplayName": "Task to create a review interface"
}
```

---

## Raw Example Request

> Example below matches **exactly** the default configuration from [API-SPEC.md](../../API-SPEC.md). This is the raw upstream shape.

```json
{
  "taskType": "Create Review Interface",
  "taskConfiguration": {
    "adp_cri_appHostMemoryLimitRatio": "0",
    "adp_cri_matterSpecificWorkspace": null,
    "adp_progressTaskTimeout": 0,
    "adp_cri_matterId": "default",
    "adp_cri_matterSpecificEngineHostId": "localhost",
    "adp_cri_httpsPassword": "",
    "adp_cri_matterSpecificTemplate": null,
    "adp_taskActive": true,
    "adp_cri_reduceMemoryFootPrint": false,
    "adp_executionPersistent": true,
    "adp_abortWfOnFailure": true,
    "adp_cri_engineHostMemoryLimit": "0",
    "adp_loggingEnabled": true,
    "adp_cri_chosenEngineHostMemoryRatio": "adp_create_review_engine_host_memory_ratio",
    "adp_cri_webserviceRequestTimeoutSeconds": 900,
    "adp_cri_webserviceUrl": "http://{host_name}/{project_name}",
    "adp_cri_publishEngineId": "adp_created_publish_engine_id",
    "adp_cri_appHostMemoryLimit": "0",
    "adp_cri_publishApplicationId": "adp_created_publish_application_id",
    "adp_cri_chosenApplicationHostMemory": "adp_create_review_application_host_memory",
    "adp_cri_chosenEngineHostMemory": "adp_create_review_engine_host_memory",
    "adp_cri_chosenEngineHostNameParameter": "adp_create_review_engine_host",
    "adp_cri_httpsKeystoreFile": null,
    "adp_cri_matterSpecificApplicationHostId": "",
    "adp_cri_matterSpecificApplication": null,
    "adp_cri_usedWebserviceUrl": "adp_url_used_for_create_review_interface",
    "adp_cri_appHostDetection": "true",
    "adp_cri_matterSpecificSmallProject": false,
    "adp_cri_engineHostMemoryLimitRatio": "0",
    "adp_cri_expectedNumberOfDocuments": 1000000,
    "adp_cri_predictiveCodingEnabled": false,
    "adp_cri_chosenApplicationHostNameParameter": "adp_create_review_application_host",
    "adp_cleanUpHistory": false,
    "adp_cri_httpsTrustCertificate": "",
    "adp_cri_httpsAllowUntrustedHosts": "false",
    "adp_cri_matterSpecificUrlRegEx": null,
    "adp_cri_chosenApplicationHostMemoryRatio": "adp_create_review_application_host_memory_ratio",
    "adp_taskTimeout": 0,
    "adp_cri_webserviceUser": "{service_user}",
    "adp_cri_webserviceConnectTimeoutSeconds": 300,
    "adp_cri_webservicePassword": null
  },
  "taskDescription": "",
  "taskDisplayName": "Task to create a review interface"
}
```

---

## CLI Arguments

See [cli.md](../cli.md) for global flags and naming conventions.

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--matter-id` | string | "default" | Matter ID |
| `--engine-host-memory-limit` | integer | 0 | Engine host memory limit |
| `--engine-host-memory-limit-ratio` | integer | 0 | Engine host memory limit ratio |
| `--app-host-memory-limit` | integer | 0 | Application host memory limit |
| `--app-host-memory-limit-ratio` | integer | 0 | Application host memory limit ratio |
| `--expected-number-of-documents` | integer | 1000000 | Expected number of documents |
| `--predictive-coding-enabled` | boolean | false | Enable predictive coding |
| `--reduce-memory-footprint` | boolean | false | Reduce memory footprint |
| `--matter-specific-application` | string | null | Matter-specific application |
| `--matter-specific-engine-host-id` | string | "localhost" | Matter-specific engine host ID |
| `--matter-specific-application-host-id` | string | "" | Matter-specific application host ID |
| `--matter-specific-workspace` | string | null | Matter-specific workspace |
| `--matter-specific-template` | string | null | Matter-specific template |
| `--matter-specific-small-project` | boolean | false | Matter is a small project |
| `--matter-specific-url-regex` | string | null | URL regex for matter-specific matching |
| `--app-host-detection` | boolean | true | Enable application host detection |

### CLI Examples

```bash
# Basic
adpgo create-review-interface

# With matter ID
adpgo create-review-interface --matter-id Demo_Review1

# With custom settings
adpgo create-review-interface --expected-number-of-documents 500000 --predictive-coding-enabled
```

---

## Raw Example Response

```json
{
  "executionId": "uuid",
  "taskType": "Create Review Interface",
  "loggingEnabled": "true",
  "progressMax": 1,
  "executionStatus": "success",
  "executionRootDir": "path",
  "contextId": "uuid",
  "executionPersistent": "true",
  "progressCurrent": 1,
  "progressPercentage": 1.0,
  "taskDisplayName": "Task to create a review interface",
  "executionMetaData": {
    "adp_create_review_application_host_memory_ratio": "76",
    "adp_create_review_engine_host_memory_ratio": "76",
    "adp_create_review_application_host": "vm-rhauswirth2.otxlab.net",
    "adp_create_review_application_host_memory": "59604877312",
    "adp_cri_CreateReviewResponseMessage": "OK.",
    "adp_created_publish_engine_id": "singleMindServer.demo00001_Demo_Review1",
    "adp_create_review_engine_host": "vm-rhauswirth2.otxlab.net",
    "adp_cri_CreateReviewRequestId": "Create Review Interface",
    "adp_url_used_for_create_review_interface": "https://vm-rhauswirth2.otxlab.net:8443/demo00001/matterspecific",
    "adp_create_review_engine_host_memory": "59604877312"
  }
}
```

---

## Decoded Result

### Result Type

```
CreateReviewInterfaceResult {
    publishEngineId: string
    publishApplicationId: string
    applicationHost: string
    applicationHostMemory: string
    applicationHostMemoryRatio: string
    engineHost: string
    engineHostMemory: string
    engineHostMemoryRatio: string
    createReviewRequestId: string
    createReviewResponseMessage: string
    usedWebserviceUrl: string
}
```

### Decoding Rules

1. Map `executionMetaData.adp_created_publish_engine_id` to `publishEngineId`
2. Map `executionMetaData.adp_created_publish_application_id` to `publishApplicationId`
3. Map `executionMetaData.adp_create_review_application_host` to `applicationHost`
4. Map `executionMetaData.adp_create_review_application_host_memory` to `applicationHostMemory`
5. Map `executionMetaData.adp_create_review_application_host_memory_ratio` to `applicationHostMemoryRatio`
6. Map `executionMetaData.adp_create_review_engine_host` to `engineHost`
7. Map `executionMetaData.adp_create_review_engine_host_memory` to `engineHostMemory`
8. Map `executionMetaData.adp_create_review_engine_host_memory_ratio` to `engineHostMemoryRatio`
9. Map `executionMetaData.adp_cri_CreateReviewRequestId` to `createReviewRequestId`
10. Map `executionMetaData.adp_cri_CreateReviewResponseMessage` to `createReviewResponseMessage`
11. Map `executionMetaData.adp_url_used_for_create_review_interface` to `usedWebserviceUrl`

---

## executionMetaData Contract

| Field | Type | Description |
|-------|------|-------------|
| adp_created_publish_engine_id | string | Created publish engine ID |
| adp_created_publish_application_id | string | Created publish application ID |
| adp_create_review_application_host | string | Application host name |
| adp_create_review_application_host_memory | string | Application host memory in bytes |
| adp_create_review_application_host_memory_ratio | string | Application host memory ratio percentage |
| adp_create_review_engine_host | string | Engine host name |
| adp_create_review_engine_host_memory | string | Engine host memory in bytes |
| adp_create_review_engine_host_memory_ratio | string | Engine host memory ratio percentage |
| adp_cri_CreateReviewRequestId | string | Create review request identifier |
| adp_cri_CreateReviewResponseMessage | string | Create review response message |
| adp_url_used_for_create_review_interface | string | The URL used for create review interface |

### JSON String Fields

| Field | Parse As |
|-------|----------|
| adp_created_publish_engine_id | `string` |
| adp_created_publish_application_id | `string` |

---

## Failure Response

On `executionStatus: "failed"`:

```json
{
  "executionId": "uuid",
  "taskType": "Create Review Interface",
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