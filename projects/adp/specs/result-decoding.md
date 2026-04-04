# Result Decoding

Describes how clients decode task execution results. This contract is language-agnostic. All sync and async calls follow the same decoding path once a terminal response is available.

---

## Overview

Every task execution returns a response envelope (defined in [api-contract.md](./api-contract.md)). The decoding process is:

1. Parse the common response envelope once for all tasks.
2. Check `executionStatus` for the terminal state.
3. On `success`: decode `executionMetaData` into the task-specific typed result.
4. On `failure`: return a failure object with `executionId`, `taskType`, `executionStatus`, `errorMessage`, and `executionMetaData` (which is `null` on failure).
5. On `running` (async poll): return the envelope as-is for continued polling.

---

## Common Response Envelope

All responses share these fields:

```
TaskResponse {
    executionId: string
    taskType: string
    loggingEnabled: string          # "true" or "false"
    progressMax: integer
    executionStatus: string         # "success" | "failed" | "running"
    executionRootDir: string
    contextId: string
    executionPersistent: string    # "true" or "false"
    progressCurrent: integer
    progressPercentage: float
    taskDisplayName: string
    executionMetaData: object | null
    errorMessage: string | absent
}
```

---

## Decoded Result Types

Each task has a typed result contract. These are the types that clients return to callers after decoding `executionMetaData`.

See individual task specs in [tasks/](./tasks/) for the per-task result contract.

### List Entities

```
ListEntitiesResult {
    outputFile: string
    entities: Entity[]
}
```

Where `Entity` contains fields from `adp_entities_json_output` (a JSON string that must be parsed first).

### Query Engine

```
QueryEngineResult {
    documentsCount: integer
    aggregatedValue: string
}
```

### Taxonomy Statistic

```
TaxonomyStatisticResult {
    outputFile: string
    statistics: StatisticsDocument   # parsed from adp_taxonomy_statistics_json_output
}
```

Where `StatisticsDocument` contains `date`, `searchParameter[]`, and `statistics.taxonomy[]`.

### Start Application

```
StartApplicationResult {
    applicationUrl: string
}
```

### Export Documents

```
ExportDocumentsResult {
    exportFileName: string
    exportPath: string
    searchResultSize: integer
}
```

### CSV Merge

> **Pending**: `executionMetaData` response fields not yet verified against actual API response.

---

## Decoding Flow

### Sync Call

```
1. POST request to /executeAdpTask
2. Receive response envelope
3. Parse common fields
4. If executionStatus == "success":
     a. Identify task type from taskType
     b. Select task-specific decoder
     c. Decode executionMetaData into typed result
     d. Return Result(executionId, taskType, result)
   Else if executionStatus == "failed":
     a. Return Failure(executionId, taskType, executionStatus, errorMessage, executionMetaData=null)
   Else (should not happen on sync):
     a. Treat as unexpected state
```

### Async Call

```
1. POST request to /executeAdpTaskAsync
2. Receive response envelope with executionStatus == "running"
3. Parse common fields, return executionId to caller
4. Poll GET /statusAndProgress with executionId until terminal state
5. On "success":
     a. Identify task type from taskType
     b. Select task-specific decoder
     c. Decode executionMetaData into typed result
     d. Return Result(executionId, taskType, result)
   Else on "failed":
     a. Return Failure(...)
```

### Polling

Clients may implement polling internally (hiding async nature) or expose it to callers. Either way, the decoding path from step 4 onward is identical.

---

## TaskDecoder Contract

Each task implements a decoder that transforms `executionMetaData: object` into the task's typed result.

Language-agnostic interface:

```
interface TaskDecoder[T] {
    decode(executionMetaData: Map<string, any>): T
}
```

Implementations must handle:
- Missing fields (return zero value or error depending on whether field is required)
- JSON string fields that must be parsed separately (e.g., `adp_entities_json_output`)
- Type coercion from string to integer/float where the API returns strings

---

## CLI Output

CLI output is based on **decoded task results**, not raw `executionMetaData`.

### On Success (HTTP 200 + executionStatus == "success")

Output only the parsed task-specific data:
- `List Entities`: the JSON array of entities
- `Query Engine`: `documentsCount` and `aggregatedValue`
- `Taxonomy Statistic`: the decoded statistics JSON
- `Start Application`: the application URL
- `Export Documents`: export file name, path, and count

### On Failure (HTTP 200 + executionStatus == "failed")

Output error details including `executionId` and `errorMessage`.

### Debug Mode

When `--debug` is enabled, additionally trace the raw request and response payloads.

---

## Error Handling

Failures return:
- `executionId`: for correlation
- `taskType`: for identifying which task failed
- `executionStatus`: always "failed"
- `errorMessage`: human-readable error from the server
- `executionMetaData`: `null` on failure

Clients should propagate all of these fields in the failure object.

---

## Type Coercion Rules

The ADP API returns many numeric and boolean values as strings. Decoders must coerce types:

| API returns | Decoder produces |
|-------------|------------------|
| string `"true"` / `"false"` | boolean `true` / `false` |
| string `"100"` | integer `100` |
| string `"1.0"` | float `1.0` |
| JSON string `"[...]"` | parsed array/object |
| `null` | language null / None / nil |

Always consult `VERIFICATION.md` and test with the live API to confirm actual types.
