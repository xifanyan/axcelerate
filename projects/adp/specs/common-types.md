# Common Types Specification

## TaskRequest

Base structure for all ADP task requests.

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

### Fields

| Field | Type | Required | Default | Description |
|-------|------|----------|---------|-------------|
| taskType | string | Yes | - | Task type identifier |
| taskConfiguration | object | Yes | - | Task-specific configuration |
| taskDescription | string | No | - | Task description |
| taskDisplayName | string | No | - | Task display name |

### Task Configuration Common Fields

| Field | Type | Default | Description |
|-------|------|---------|-------------|
| adp_taskActive | boolean | true | Whether task is active |
| adp_taskTimeout | integer | 0 | Timeout in seconds |
| adp_executionPersistent | boolean | true | Persist execution |
| adp_abortWfOnFailure | boolean | true | Abort workflow on failure |
| adp_loggingEnabled | boolean | true | Enable logging |
| adp_cleanUpHistory | boolean | false | Clean up history |

---

## TaskResponse

Base structure for all ADP task responses. All responses contain these common fields:

```json
{
  "ExecutionID": "uuid",
  "TaskType": "Task Name",
  "LoggingEnabled": false,
  "ProgressMax": 1,
  "ExecutionStatus": "success",
  "ExecutionRootDir": "path/to/root",
  "ContextID": "uuid",
  "ExecutionPersistent": false,
  "ProgressCurrent": 1,
  "ProgressPercentage": 1,
  "TaskDisplayName": "",
  "ExecutionMetaData": {}
}
```

### Common Response Fields

| Field | Type | Description | Notes |
|-------|------|-------------|-------|
| executionId | string | Unique execution identifier (UUID) | camelCase |
| taskType | string | Task type (e.g., "List Entities", "Query Engine") | |
| loggingEnabled | string | Whether logging is enabled | **String** ("true"/"false"), not boolean |
| progressMax | integer | Maximum progress value | |
| executionStatus | string | Status of execution ("success", "failed", "running") | |
| executionRootDir | string | Root directory for execution | |
| contextId | string | Context identifier (UUID) | |
| executionPersistent | string | Whether execution is persistent | **String** ("true"/"false"), not boolean |
| progressCurrent | integer | Current progress value | |
| progressPercentage | float | Progress percentage (0-100) | **Float**, not integer |
| taskDisplayName | string | Display name of the task | |
| executionMetaData | object? | Task-specific metadata (differs per task) | **Optional** - null when status is "failed" |
| errorMessage | string? | Error message on failure | **Present when executionStatus is "failed"** |

---

## ExecutionMetaData by Task

> The `ExecutionMetaData` object contains task-specific fields. Each task has different fields.

### List Entities

```json
{
  "adp_entities_output_file_name": "path/to/output.json",
  "adp_entities_json_output": "[{\"id\":\"...\",\"displayName\":\"...\",...}]"
}
```

| Field | Type | Description |
|-------|------|-------------|
| adp_entities_output_file_name | string | Output file path |
| adp_entities_json_output | string | JSON string containing array of entities |

### Query Engine

```json
{
  "adp_query_engine_documents_count": "10",
  "adp_query_engine_aggregated_value": "value"
}
```

| Field | Type | Description |
|-------|------|-------------|
| adp_query_engine_documents_count | string | Number of documents matching query |
| adp_query_engine_aggregated_value | string | Aggregated value result |

### Taxonomy Statistic

```json
{
  "adp_taxonomy_statistics_json_output": "[{\"category\":\"...\",\"count\":10}]",
  "adp_taxonomy_statistics_json_file_path": "path/to/file.json"
}
```

| Field | Type | Description |
|-------|------|-------------|
| adp_taxonomy_statistics_json_output | string | JSON string containing taxonomy statistics |
| adp_taxonomy_statistics_json_file_path | string | Output file path |

---

## Shared Input Types

These types are used across multiple tasks.

### EngineTaxonomyArg

Used by: Query Engine, Taxonomy Statistic

```go
type EngineTaxonomyArg struct {
    Taxonomy string // Taxonomy name (e.g., "rm_source", "meta_documentcharacteristics")
    Negation bool   // Negation flag
    Query    string // Query string (URL-encoded)
}
```

### OutputTaxonomiesArg

Used by: Taxonomy Statistic

```go
type OutputTaxonomiesArg struct {
    Taxonomy                  string // Taxonomy name
    Mode                      string // "Aggregate counts" or "Category counts"
    MaximumNumberOfCategories int    // Maximum number of categories to return
}
```

---

## ExecutionStatus Values

| Value | Description |
|-------|-------------|
| success | Task completed successfully |
| failed | Task failed to complete |
| running | Task is currently executing |

---

## Error Response

When execution fails, the response may include:

```json
{
  "executionId": "uuid",
  "taskType": "Task Name",
  "executionStatus": "failed",
  "errorMessage": "Error message details",
  "executionMetaData": null
}
```

| Field | Type | Description |
|-------|------|-------------|
| executionId | string | Unique execution identifier (UUID) |
| taskType | string | Task type |
| executionStatus | string | Will be "failed" |
| errorMessage | string | Error details |
| executionMetaData | null | Is null on failure |
