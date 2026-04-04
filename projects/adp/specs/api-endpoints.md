# API Endpoints Specification

> **Deprecated**: This file is retained for reference only. The authoritative API contract is now in [api-contract.md](./api-contract.md).

---

## Historical Content

The content below is superseded by [api-contract.md](./api-contract.md). Refer to that file for the current authoritative specification.

### Endpoints

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/executeAdpTask` | PUT | Execute task synchronously |
| `/executeAdpTaskAsync` | PUT | Execute task asynchronously |
| `/statusAndProgress` | PUT | Poll task status by execution ID |

### Common Response Format

> **IMPORTANT:** All field names in API responses use **camelCase** (e.g., `executionId`, `executionMetaData`), NOT PascalCase.

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

### ExecutionStatus Values

| Status | Description |
|--------|-------------|
| success | Task completed successfully |
| failed | Task failed to complete |
| running | Task is currently executing |
