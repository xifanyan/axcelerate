# Matter Management Task Specification

## Overview

| Property | Value |
|----------|-------|
| Task Type | `Matter Management` |
| Description | Task to manage matters and saved searches |
| Display Name | Task to manage matters and saved searches |
| Subcommand | `matter-management` |

---

## Semantic Inputs

These are the user-facing fields for the request-construction API. These are the fields callers set directly. Raw field names are implementation details.

| Field | Type | Default | Required | Description |
|-------|------|---------|----------|-------------|
| matterId | string | null | No | Matter ID to manage |
| matterDisplayName | string | null | No | Display name for the matter |
| searchString | string | "*" | No | Search string to filter |
| searchFields | string[] | [] | No | Fields to search in |
| enableSearch | boolean | true | No | Enable search functionality |
| savedSearchId | string | null | No | Saved search ID |
| savedSearchDisplayName | string | null | No | Display name for saved search |
| matterSpecificApplication | string | null | No | Matter-specific application |
| matterSpecificUrlRegEx | string | null | No | URL regex for matter-specific matching |
| moveNativesToMatterStorage | boolean | false | No | Move native files to matter storage |
| copyNotCopiedNatives | boolean | false | No | Copy native files not yet copied |
| retryCopyNativesOnError | boolean | false | No | Retry copying natives on error |

---

## Raw Default Configuration

> Configuration below shows **all fields with their exact default values** from [API-SPEC.md](../../API-SPEC.md). This is for reference only. Clients must not pre-populate all fields — only send fields explicitly set by the caller. See [request-construction.md](../request-construction.md).

```json
{
  "taskType": "Matter Management",
  "taskConfiguration": {
    "adp_progressTaskTimeout": 0,
    "adp_mm_webservicePassword": null,
    "adp_mm_matterSpecificUrlRegEx": null,
    "adp_taskActive": true,
    "adp_executionPersistent": true,
    "adp_abortWfOnFailure": true,
    "adp_mm_moveNativesToMatterStorage": "false",
    "adp_mm_webserviceUrl": "http://{host_name}/{project_name}",
    "adp_mm_usedWebserviceUrl": "adp_url_used_for_matter_management",
    "adp_loggingEnabled": true,
    "adp_mm_savedSearchId": null,
    "adp_mm_webserviceRequestTimeoutSeconds": 900,
    "adp_mm_matterId": null,
    "adp_mm_searchString": "*",
    "adp_mm_retryCopyNativesOnError": "false",
    "adp_mm_matterDisplayName": null,
    "adp_mm_processedSavedSearchId": "adp_processed_saved_search_id",
    "adp_mm_webserviceUser": "{service_user}",
    "adp_mm_httpsAllowUntrustedHosts": "false",
    "adp_mm_httpsPassword": "",
    "adp_cleanUpHistory": false,
    "adp_mm_httpsKeystoreFile": null,
    "adp_mm_processedMatterId": "adp_processed_matter_id",
    "adp_mm_webserviceConnectTimeoutSeconds": 300,
    "adp_mm_httpsTrustCertificate": "",
    "adp_mm_enableSearch": "true",
    "adp_taskTimeout": 0,
    "adp_mm_copyNotCopiedNatives": "false",
    "adp_mm_searchFields": [],
    "adp_mm_matterSpecificApplication": null,
    "adp_mm_savedSearchDisplayName": null
  },
  "taskDescription": "",
  "taskDisplayName": "Task to manage matters and saved searches"
}
```

---

## Raw Example Request

> Example below matches **exactly** the default configuration from [API-SPEC.md](../../API-SPEC.md). This is the raw upstream shape.

```json
{
  "taskType": "Matter Management",
  "taskConfiguration": {
    "adp_progressTaskTimeout": 0,
    "adp_mm_webservicePassword": null,
    "adp_mm_matterSpecificUrlRegEx": null,
    "adp_taskActive": true,
    "adp_executionPersistent": true,
    "adp_abortWfOnFailure": true,
    "adp_mm_moveNativesToMatterStorage": "false",
    "adp_mm_webserviceUrl": "http://{host_name}/{project_name}",
    "adp_mm_usedWebserviceUrl": "adp_url_used_for_matter_management",
    "adp_loggingEnabled": true,
    "adp_mm_savedSearchId": null,
    "adp_mm_webserviceRequestTimeoutSeconds": 900,
    "adp_mm_matterId": null,
    "adp_mm_searchString": "*",
    "adp_mm_retryCopyNativesOnError": "false",
    "adp_mm_matterDisplayName": null,
    "adp_mm_processedSavedSearchId": "adp_processed_saved_search_id",
    "adp_mm_webserviceUser": "{service_user}",
    "adp_mm_httpsAllowUntrustedHosts": "false",
    "adp_mm_httpsPassword": "",
    "adp_cleanUpHistory": false,
    "adp_mm_httpsKeystoreFile": null,
    "adp_mm_processedMatterId": "adp_processed_matter_id",
    "adp_mm_webserviceConnectTimeoutSeconds": 300,
    "adp_mm_httpsTrustCertificate": "",
    "adp_mm_enableSearch": "true",
    "adp_taskTimeout": 0,
    "adp_mm_copyNotCopiedNatives": "false",
    "adp_mm_searchFields": [],
    "adp_mm_matterSpecificApplication": null,
    "adp_mm_savedSearchDisplayName": null
  },
  "taskDescription": "",
  "taskDisplayName": "Task to manage matters and saved searches"
}
```

---

## CLI Arguments

See [cli.md](../cli.md) for global flags and naming conventions.

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--matter-id` | string | null | Matter ID to manage |
| `--matter-display-name` | string | null | Display name for the matter |
| `--search-string` | string | "*" | Search string to filter |
| `--search-fields` | string | "" | Comma-separated fields to search in |
| `--enable-search` | boolean | true | Enable search functionality |
| `--saved-search-id` | string | null | Saved search ID |
| `--saved-search-display-name` | string | null | Display name for saved search |
| `--matter-specific-application` | string | null | Matter-specific application |
| `--matter-specific-url-regex` | string | null | URL regex for matter-specific matching |
| `--move-natives-to-matter-storage` | boolean | false | Move native files to matter storage |
| `--copy-not-copied-natives` | boolean | false | Copy native files not yet copied |
| `--retry-copy-natives-on-error` | boolean | false | Retry copying natives on error |

### CLI Examples

```bash
# Basic
adpgo matter-management

# With matter ID
adpgo matter-management --matter-id Demo_Review1

# With search options
adpgo matter-management --search-string "custom*" --search-fields "name,description"

# With display name
adpgo matter-management --matter-display-name "My Matter"
```

---

## Raw Example Response

```json
{
  "executionId": "uuid",
  "taskType": "Matter Management",
  "loggingEnabled": "true",
  "progressMax": 1,
  "executionStatus": "success",
  "executionRootDir": "path",
  "contextId": "uuid",
  "executionPersistent": "true",
  "progressCurrent": 1,
  "progressPercentage": 1.0,
  "taskDisplayName": "Task to manage matters and saved searches",
  "executionMetaData": {
    "adp_mm_savedSearchProcessingRequestId": "Matter Management",
    "adp_processed_saved_search_id": "Demo_Review1_saved_search",
    "adp_mm_savedSearchProcessingResponseMessage": "OK.",
    "adp_url_used_for_matter_management": "https://vm-rhauswirth2.otxlab.net:8443/demo00001/matterspecific",
    "adp_processed_matter_id": "Demo_Review1",
    "adp_mm_matterProcessingRequestId": "Matter Management",
    "adp_mm_matterProcessingResponseMessage": "OK."
  }
}
```

---

## Decoded Result

### Result Type

```
MatterManagementResult {
    processedMatterId: string
    matterProcessingRequestId: string
    matterProcessingResponseMessage: string
    processedSavedSearchId: string
    savedSearchProcessingRequestId: string
    savedSearchProcessingResponseMessage: string
    usedWebserviceUrl: string
}
```

### Decoding Rules

1. Map `executionMetaData.adp_processed_matter_id` to `processedMatterId`
2. Map `executionMetaData.adp_mm_matterProcessingRequestId` to `matterProcessingRequestId`
3. Map `executionMetaData.adp_mm_matterProcessingResponseMessage` to `matterProcessingResponseMessage`
4. Map `executionMetaData.adp_processed_saved_search_id` to `processedSavedSearchId`
5. Map `executionMetaData.adp_mm_savedSearchProcessingRequestId` to `savedSearchProcessingRequestId`
6. Map `executionMetaData.adp_mm_savedSearchProcessingResponseMessage` to `savedSearchProcessingResponseMessage`
7. Map `executionMetaData.adp_url_used_for_matter_management` to `usedWebserviceUrl`

---

## executionMetaData Contract

| Field | Type | Description |
|-------|------|-------------|
| adp_processed_matter_id | string | The processed matter ID |
| adp_mm_matterProcessingRequestId | string | Matter processing request identifier |
| adp_mm_matterProcessingResponseMessage | string | Matter processing response message |
| adp_processed_saved_search_id | string | The processed saved search ID |
| adp_mm_savedSearchProcessingRequestId | string | Saved search processing request identifier |
| adp_mm_savedSearchProcessingResponseMessage | string | Saved search processing response message |
| adp_url_used_for_matter_management | string | The URL used for matter management |

---

## Failure Response

On `executionStatus: "failed"`:

```json
{
  "executionId": "uuid",
  "taskType": "Matter Management",
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