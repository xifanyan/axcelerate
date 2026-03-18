# ADP Spec Correction Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Update ADP spec files to accurately reflect the real API response format, avoiding field naming and type mismatches discovered during Rust implementation.

**Architecture:** Update specification documents in `projects/adp/specs/` to document actual API behavior based on real-world testing.

**Tech Stack:** Markdown documentation

---

## Background

During Rust client implementation, multiple discrepancies were found between the spec and actual API responses:

1. Field names use camelCase, not PascalCase
2. `LoggingEnabled` and `ExecutionPersistent` are strings, not booleans
3. `ProgressPercentage` is a float, not an integer
4. `executionMetaData` is optional (null on failure)
5. `errorMessage` field exists on failed responses (not documented)

---

## Files

- Modify: `projects/adp/specs/common-types.md`
- Modify: `projects/adp/specs/api-endpoints.md`

---

### Task 1: Update common-types.md Response Field Types

**Files:**
- Modify: `projects/adp/specs/common-types.md:66-82`

- [ ] **Step 1: Read current common-types.md response section**

```bash
read projects/adp/specs/common-types.md
```

- [ ] **Step 2: Replace Response Fields table with corrected types**

Replace lines 66-82 with:

```markdown
### Common Response Fields

| Field | Type | Description | Notes |
|-------|------|-------------|-------|
| executionId | string | Unique execution identifier (UUID) | camelCase |
| taskType | string | Task type (e.g., "List Entities") | |
| loggingEnabled | string | Whether logging is enabled | **String** ("true"/"false"), not boolean |
| progressMax | integer | Maximum progress value | |
| executionStatus | string | Status: "success", "failed", "running" | |
| executionRootDir | string | Root directory for execution | |
| contextId | string | Context identifier (UUID) | |
| executionPersistent | boolean | Whether execution is persistent | **String** ("true"/"false"), not boolean |
| progressCurrent | integer | Current progress value | |
| progressPercentage | float | Progress percentage (0-100) | **Float**, not integer |
| taskDisplayName | string | Display name of the task | |
| executionMetaData | object? | Task-specific metadata | **Optional** - null when status is "failed" |
| errorMessage | string? | Error message on failure | **Present when executionStatus is "failed"** |
```

- [ ] **Step 3: Update Error Response section**

Replace lines 143-156 with:

```markdown
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
```

- [ ] **Step 4: Commit changes**

```bash
git add projects/adp/specs/common-types.md
git commit -m "docs: fix response field types in common-types.md"
```

---

### Task 2: Update api-endpoints.md Response Field Names

**Files:**
- Modify: `projects/adp/specs/api-endpoints.md:14-31`

- [ ] **Step 1: Read current api-endpoints.md**

```bash
read projects/adp/specs/api-endpoints.md
```

- [ ] **Step 2: Add note about camelCase field names**

Add this note at the top of the "Common Response Format" section (before line 15):

```markdown
> **IMPORTANT:** All field names in API responses use camelCase (e.g., `executionId`, `executionMetaData`), NOT PascalCase.
```

- [ ] **Step 3: Fix Response table field names**

Replace lines 17-31:

```markdown
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
```

- [ ] **Step 4: Fix Example Responses to use camelCase**

Update all example JSON responses in the file (executeAdpTask, executeAdpTaskAsync, statusAndProgress sections) to use camelCase field names:
- `ExecutionID` → `executionId`
- `TaskType` → `taskType`
- `LoggingEnabled` → `loggingEnabled`
- `ProgressMax` → `progressMax`
- `ExecutionStatus` → `executionStatus`
- `ExecutionRootDir` → `executionRootDir`
- `ContextID` → `contextId`
- `ExecutionPersistent` → `executionPersistent`
- `ProgressCurrent` → `progressCurrent`
- `ProgressPercentage` → `progressPercentage`
- `TaskDisplayName` → `taskDisplayName`
- `ExecutionMetaData` → `executionMetaData`

- [ ] **Step 5: Add failure example with errorMessage**

In the executeAdpTask response section, add an example failure response showing `errorMessage`:

```json
{
  "executionId": "f9463001-dc1f-486a-a8a0-efaca8dd29cb",
  "taskType": "List Entities",
  "loggingEnabled": "true",
  "progressMax": 1,
  "executionStatus": "failed",
  "errorMessage": "Invalid entity type",
  "executionRootDir": "E:\\MindServer\\Projects\\adp.adp\\adpRootDir",
  "contextID": "2e5a47e4-d9c8-4547-aaba-45c0a3774d47",
  "executionPersistent": "true",
  "progressCurrent": 0,
  "progressPercentage": 0.0,
  "taskDisplayName": "List Entities",
  "executionMetaData": null
}
```

- [ ] **Step 6: Commit changes**

```bash
git add projects/adp/specs/api-endpoints.md
git commit -m "docs: fix response field names and types in api-endpoints.md"
```

---

### Task 3: Create API Verification Checklist

**Files:**
- Create: `projects/adp/specs/VERIFICATION.md`

- [ ] **Step 1: Create verification checklist**

```markdown
# API Response Verification Checklist

Before finalizing any spec, verify against actual API:

## Field Names
- [ ] All response fields use camelCase (e.g., `executionId`, not `ExecutionID`)
- [ ] Test with real API call using debug logging

## Field Types
- [ ] `loggingEnabled` - is it boolean or string?
- [ ] `executionPersistent` - is it boolean or string?
- [ ] `progressPercentage` - is it integer or float?
- [ ] `executionMetaData` - is it required or optional?

## Error Cases
- [ ] Test failure scenarios
- [ ] Document `errorMessage` field presence
- [ ] Document `executionMetaData` behavior on failure (null?)

## Testing
Always test with:
```bash
# Enable debug to see raw JSON
./client --debug command --args
```
```

- [ ] **Step 2: Commit**

```bash
git add projects/adp/specs/VERIFICATION.md
git commit -m "docs: add API verification checklist"
```

---

## Summary

After completing these tasks, the spec will accurately document:
1. Correct camelCase field names
2. Actual field types (string vs boolean vs float)
3. Optional nature of `executionMetaData`
4. Presence of `errorMessage` on failures
5. A checklist to avoid these issues in the future
