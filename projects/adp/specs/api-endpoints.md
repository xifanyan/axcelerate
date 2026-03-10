# API Endpoints Specification

## Overview

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/executeAdpTask` | PUT | Execute task synchronously |
| `/executeAdpTaskAsync` | PUT | Execute task asynchronously |
| `/statusAndProgress` | PUT | Poll task status by execution ID |

---

## Common Response Format

All endpoints return responses with these common fields:

| Field | Type | Description |
|-------|------|-------------|
| ExecutionID | string | Unique execution identifier (UUID) |
| TaskType | string | Task type (e.g., "List Entities") |
| LoggingEnabled | boolean | Whether logging is enabled |
| ProgressMax | integer | Maximum progress value |
| ExecutionStatus | string | Status: "success", "failed", "running" |
| ExecutionRootDir | string | Root directory for execution |
| ContextID | string | Context identifier (UUID) |
| ExecutionPersistent | boolean | Whether execution is persistent |
| ProgressCurrent | integer | Current progress value |
| ProgressPercentage | integer | Progress percentage (0-100) |
| TaskDisplayName | string | Display name of the task |
| ExecutionMetaData | object | Task-specific metadata (differs per task) |

---

## 1. Execute Task Synchronously

### Endpoint
```
PUT /executeAdpTask
```

### Description
Executes an ADP task and waits for completion. Returns the final result when the task finishes.

### Request Body

```json
{
  "taskType": "List Entities",
  "taskConfiguration": {
    "adp_taskActive": true,
    "adp_taskTimeout": 0,
    "adp_executionPersistent": true,
    "adp_abortWfOnFailure": true,
    "adp_loggingEnabled": true,
    "adp_cleanUpHistory": false
  },
  "taskDescription": "Writes a list of entities ot an output variable",
  "taskDisplayName": "List entities"
}
```

### Request Fields

| Field | Type | Required | Description |
|-----------|------|----------|-------------|
| taskType | string | Yes | Task type identifier (e.g., "List Entities") |
| taskConfiguration | object | Yes | Task-specific configuration |
| taskDescription | string | No | Task description |
| taskDisplayName | string | No | Task display name |

### Response

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
    "adp_entities_json_output": "[...]"
  }
}
```

---

## 2. Execute Task Asynchronously

### Endpoint
```
PUT /executeAdpTaskAsync
```

### Description
Executes an ADP task and returns immediately with an execution ID. Use `statusAndProgress` to poll for results.

### Request Body
Same as `executeAdpTask`.

### Response

```json
{
  "ExecutionID": "f9463001-dc1f-486a-a8a0-efaca8dd29cb",
  "TaskType": "List Entities",
  "LoggingEnabled": false,
  "ProgressMax": 1,
  "ExecutionStatus": "running",
  "ExecutionRootDir": "E:\\MindServer\\Projects\\adp.adp\\adpRootDir",
  "ContextID": "2e5a47e4-d9c8-4547-aaba-45c0a3774d47",
  "ExecutionPersistent": false,
  "ProgressCurrent": 0,
  "ProgressPercentage": 0,
  "TaskDisplayName": "",
  "ExecutionMetaData": {}
}
```

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
  "ExecutionID": "f9463001-dc1f-486a-a8a0-efaca8dd29cb"
}
```

| Field | Type | Required | Description |
|-----------|------|----------|-------------|
| ExecutionID | string | Yes | The execution ID from async task |

### Response

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
  "ProgressPercentage": 100,
  "TaskDisplayName": "",
  "ExecutionMetaData": {
    "adp_entities_output_file_name": "E:\\MindServer\\Projects\\adp.adp\\adpRootDir\\output.json",
    "adp_entities_json_output": "[...]"
  }
}
```

---

## ExecutionStatus Values

| Status | Description |
|--------|-------------|
| success | Task completed successfully |
| failed | Task failed to complete |
| running | Task is currently executing |
