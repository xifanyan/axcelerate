# Task Spec Template

Use this template when creating new task specification files in `specs/tasks/`.

## Template

```markdown
# {Task Name} Task Specification

## Overview

| Property | Value |
|----------|-------|
| Task Type | `{Task Name}` |
| Description | {Task description from API-SPEC.md} |
| Display Name | {Display name from API-SPEC.md} |

---

## Default Configuration

> Configuration below shows **all fields with their exact default values** from API-SPEC.md

```json
{
  "taskType": "{Task Name}",
  "taskConfiguration": {
    // COPY ALL FIELDS FROM API-SPEC.md HERE
  },
  "taskDescription": "{description}",
  "taskDisplayName": "{display name}"
}
```

---

## Field Reference (with Defaults)

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| // | // | // | // |

---

## Example Request

> Example below matches **exactly** the default configuration from [API-SPEC.md](../../API-SPEC.md)

```json
{
  "taskType": "{Task Name}",
  "taskConfiguration": {
    // SAME AS DEFAULT CONFIGURATION
  },
  "taskDescription": "{description}",
  "taskDisplayName": "{display name}"
}
```

---

## Example Response

```json
{
  "executionId": "uuid",
  "taskType": "{Task Name}",
  "loggingEnabled": "false",
  "progressMax": 1,
  "executionStatus": "success",
  "executionRootDir": "path",
  "contextId": "uuid",
  "executionPersistent": "false",
  "progressCurrent": 1,
  "progressPercentage": 1.0,
  "taskDisplayName": "",
  "executionMetaData": {
    // Task-specific fields here
  }
}
```

---

## Response Fields

All responses include the common fields. {Task Name}-specific `ExecutionMetaData` fields:

| Field | Type | Description |
|-------|------|-------------|
| | | |
```

---

## How to Fill the Template

1. Find the task in `API-SPEC.md`
2. Copy the exact JSON from `taskConfiguration`
3. Fill in the overview table
4. Create the Field Reference table
5. Copy default config to Example Request (no changes)
6. Determine ExecutionMetaData fields from actual API response
7. Document ExecutionMetaData fields in Response Fields table

---

## Rules

- Default Configuration must match API-SPEC.md exactly
- Example Request must match Default Configuration exactly (no custom values)
- Preserve exact field names, values, and ordering from source
- Use actual response format: executionId, taskType, executionStatus, executionMetaData, etc. (camelCase)
