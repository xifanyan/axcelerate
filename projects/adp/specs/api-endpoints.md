# API Endpoints Specification

## Overview

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/executeAdpTask` | PUT | Execute task synchronously |
| `/executeAdpTaskAsync` | PUT | Execute task asynchronously |
| `/statusAndProgress` | PUT | Poll task status by execution ID |

---

## Common Response Format

> **IMPORTANT:** All field names in API responses use camelCase (e.g., `executionId`, `executionMetaData`), NOT PascalCase.

All endpoints return responses with these common fields:

| Field | Type | Description |
|-------|------|-------------|
| executionId | string | Unique execution identifier (UUID) |
| taskType | string | Task type (e.g., "List Entities") |
| loggingEnabled | string | Whether logging is enabled ("true"/"false") |
| progressMax | integer | Maximum progress value |
| executionStatus | string | Status: "success", "failed", "running" |
| executionRootDir | string | Root directory for execution |
| contextId | string | Context identifier (UUID) |
| executionPersistent | string | Whether execution is persistent ("true"/"false") |
| progressCurrent | integer | Current progress value |
| progressPercentage | float | Progress percentage (0-100) |
| taskDisplayName | string | Display name of the task |
| executionMetaData | object? | Task-specific metadata (null on failure) |
| errorMessage | string? | Error message (present on failure) |

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
    "adp_entities_output_file_name": "E:\\MindServer\\Projects\\adp.adp\\adpRootDir\\output.json",
    "adp_entities_json_output": "[...]"
  }
}
```

#### Example Failure Response

```json
{
  "executionId": "f9463001-dc1f-486a-a8a0-efaca8dd29cb",
  "taskType": "List Entities",
  "loggingEnabled": "true",
  "progressMax": 1,
  "executionStatus": "failed",
  "errorMessage": "Invalid entity type",
  "executionRootDir": "E:\\MindServer\\Projects\\adp.adp\\adpRootDir",
  "contextId": "2e5a47e4-d9c8-4547-aaba-45c0a3774d47",
  "executionPersistent": "true",
  "progressCurrent": 0,
  "progressPercentage": 0.0,
  "taskDisplayName": "List Entities",
  "executionMetaData": null
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
|-----------|------|----------|-------------|
| executionId | string | Yes | The execution ID from async task |

### Response

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
  "progressPercentage": 100.0,
  "taskDisplayName": "",
  "executionMetaData": {
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
