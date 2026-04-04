# API Contract

Describes the raw HTTP transport, endpoints, and envelope shapes. All field names use **camelCase**. This contract is language-agnostic.

---

## Endpoints

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/executeAdpTask` | PUT | Execute task synchronously — waits for completion and returns final result |
| `/executeAdpTaskAsync` | PUT | Execute task asynchronously — returns immediately with an execution ID |
| `/statusAndProgress` | PUT | Poll task status by execution ID |

---

## Raw Request Envelope

All task requests use the same envelope structure:

```json
{
  "taskType": "Task Name",
  "taskConfiguration": {
    "adp_taskActive": true,
    "adp_taskTimeout": 0,
    "adp_executionPersistent": true,
    "adp_abortWfOnFailure": true,
    "adp_loggingEnabled": true,
    "adp_cleanUpHistory": false
  },
  "taskDescription": "Task description",
  "taskDisplayName": "Task Display Name"
}
```

### Request Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| taskType | string | Yes | Task type identifier (e.g., "List Entities") |
| taskConfiguration | object | Yes | Task-specific configuration. Only include fields being changed from defaults. |
| taskDescription | string | No | Task description |
| taskDisplayName | string | No | Task display name |

### Task Configuration Common Fields

These fields appear in `taskConfiguration` for every task:

| Field | Type | Default | Description |
|-------|------|---------|-------------|
| adp_taskActive | boolean | true | Whether task is active |
| adp_taskTimeout | integer | 0 | Timeout in seconds |
| adp_executionPersistent | boolean | true | Persist execution |
| adp_abortWfOnFailure | boolean | true | Abort workflow on failure |
| adp_loggingEnabled | boolean | true | Enable logging |
| adp_cleanUpHistory | boolean | false | Clean up history |

> **Note:** The `taskConfiguration` object in client requests should only include fields explicitly set by the caller. Unset fields are omitted. See [request-construction.md](./request-construction.md) for details.

---

## Raw Response Envelope

All API responses return the same envelope structure. Field names are **camelCase only**.

### Success Response

```json
{
  "executionId": "f9463001-dc1f-486a-a8a0-efaca8dd29cb",
  "taskType": "List Entities",
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
    "adp_entities_output_file_name": "path/to/output.json",
    "adp_entities_json_output": "[...]"
  }
}
```

### Failure Response

```json
{
  "executionId": "f9463001-dc1f-486a-a8a0-efaca8dd29cb",
  "taskType": "List Entities",
  "loggingEnabled": "true",
  "progressMax": 1,
  "executionStatus": "failed",
  "executionRootDir": "E:\\MindServer\\Projects\\adp.adp\\adpRootDir",
  "contextId": "2e5a47e4-d9c8-4547-aaba-45c0a3774d47",
  "executionPersistent": "true",
  "progressCurrent": 0,
  "progressPercentage": 0.0,
  "taskDisplayName": "List Entities",
  "errorMessage": "Error message details",
  "executionMetaData": null
}
```

### Async Initial Response

```json
{
  "executionId": "f9463001-dc1f-486a-a8a0-efaca8dd29cb",
  "taskType": "List Entities",
  "loggingEnabled": "false",
  "progressMax": 1,
  "executionStatus": "running",
  "executionRootDir": "E:\\MindServer\\Projects\\adp.adp\\adpRootDir",
  "contextId": "2e5a47e4-d9c8-4547-aaba-45c0a3774d47",
  "executionPersistent": "false",
  "progressCurrent": 0,
  "progressPercentage": 0.0,
  "taskDisplayName": "",
  "executionMetaData": {}
}
```

### Common Response Fields

| Field | Type | Description |
|-------|------|-------------|
| executionId | string | Unique execution identifier (UUID) |
| taskType | string | Task type (e.g., "List Entities", "Query Engine") |
| loggingEnabled | string | Whether logging is enabled — **string** ("true"/"false"), not boolean |
| progressMax | integer | Maximum progress value |
| executionStatus | string | Status of execution: "success", "failed", "running" |
| executionRootDir | string | Root directory for execution |
| contextId | string | Context identifier (UUID) |
| executionPersistent | string | Whether execution is persistent — **string** ("true"/"false"), not boolean |
| progressCurrent | integer | Current progress value |
| progressPercentage | float | Progress percentage (0-100) — **float**, not integer |
| taskDisplayName | string | Display name of the task |
| executionMetaData | object? | Task-specific metadata. **null on failure.** Contains different fields per task. |
| errorMessage | string? | Error message on failure. **Present when executionStatus is "failed".** |

---

## 3. Poll Task Status

### Endpoint

```
PUT /statusAndProgress
```

### Description

Polls the status of a previously submitted asynchronous task.

### Request Body

```json
{
  "executionId": "f9463001-dc1f-486a-a8a0-efaca8dd29cb"
}
```

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| executionId | string | Yes | The execution ID from async task |

### Response

Same as the sync response envelope. Returns `executionStatus: "success"` when complete, `executionStatus: "running"` if still in progress, or `executionStatus: "failed"` on failure.

---

## ExecutionStatus Values

| Value | Description |
|-------|-------------|
| success | Task completed successfully |
| failed | Task failed to complete |
| running | Task is currently executing (async) |

---

## ExecutionMetaData by Task

The `executionMetaData` object contains task-specific fields. Each task returns different fields. See individual task specs in [tasks/](./tasks/) for the per-task `executionMetaData` contract.

Common patterns:
- **JSON string fields**: Some `executionMetaData` fields contain JSON strings that must be parsed (e.g., `adp_entities_json_output`, `adp_taxonomy_statistics_json_output`). These must be decoded separately.
- **File path fields**: Some tasks return output file paths in `executionMetaData`.
- **Value fields**: Some tasks return counts or aggregated values as strings.

See [result-decoding.md](./result-decoding.md) for how to decode these into typed results.

---

## Shared Input Types

These types are used across multiple tasks.

### EngineTaxonomyArg

Used by: Query Engine, Taxonomy Statistic

```json
{
  "taxonomy": "rm_source",
  "negation": false,
  "query": "email"
}
```

| Field | Type | Description |
|-------|------|-------------|
| taxonomy | string | Taxonomy name (e.g., "rm_source", "meta_documentcharacteristics") |
| negation | boolean | Negation flag. `false` = equals, `true` = not equals |
| query | string | Query string (URL-encoded) |

#### CLI Shorthand Format

For CLI, use shorthand format instead of JSON:

| Format | Description | Example |
|--------|-------------|---------|
| `Taxonomy=Query` | Equals (negation=false) | `rm_mimetype=pdf` |
| `Taxonomy!=Query` | Not equals (negation=true) | `rm_source!=email` |

Multiple taxonomies: repeat the flag.

### OutputTaxonomiesArg

Used by: Taxonomy Statistic

```json
{
  "taxonomy": "rm_source",
  "mode": "Category counts",
  "maximumNumberOfCategories": 10
}
```

| Field | Type | Description |
|-------|------|-------------|
| taxonomy | string | Taxonomy name |
| mode | string | "Aggregate counts" or "Category counts" |
| maximumNumberOfCategories | integer | Maximum number of categories to return |

---

## Field Type Notes

> **Important**: Many fields that look like booleans or numbers in the API are actually strings.
>
> - `loggingEnabled` — string ("true"/"false")
> - `executionPersistent` — string ("true"/"false")
> - `progressPercentage` — float
> - `executionMetaData` — null on failure, object on success

Always test with debug logging enabled to verify actual payload shapes.
