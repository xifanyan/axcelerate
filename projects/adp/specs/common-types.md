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

| Field | Type | Description |
|-------|------|-------------|
| ExecutionID | string | Unique execution identifier (UUID) |
| TaskType | string | Task type (e.g., "List Entities", "Query Engine") |
| LoggingEnabled | boolean | Whether logging is enabled |
| ProgressMax | integer | Maximum progress value |
| ExecutionStatus | string | Status of execution ("success", "failed", "running") |
| ExecutionRootDir | string | Root directory for execution |
| ContextID | string | Context identifier (UUID) |
| ExecutionPersistent | boolean | Whether execution is persistent |
| ProgressCurrent | integer | Current progress value |
| ProgressPercentage | integer | Progress percentage (0-100) |
| TaskDisplayName | string | Display name of the task |
| ExecutionMetaData | object | Task-specific metadata (differs per task) |

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
  "ExecutionID": "uuid",
  "TaskType": "Task Name",
  "ExecutionStatus": "failed",
  "ExecutionMetaData": {
    "error": "Error message details"
  }
}
```
