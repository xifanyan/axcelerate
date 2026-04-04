# {Task Name} Task Specification

Use this template when creating new task specification files in `specs/tasks/`.

---

## Overview

| Property | Value |
|----------|-------|
| Task Type | `{Task Name}` |
| Description | {Task description from API-SPEC.md} |
| Display Name | {Display name from API-SPEC.md} |
| Subcommand | `{kebab-case-name}` |

---

## Semantic Inputs

These are the user-facing fields for the request-construction API. These are the fields callers set directly. Raw field names are implementation details.

| Field | Type | Default | Required | Description |
|-------|------|---------|----------|-------------|
| | | | | |

---

## Raw Default Configuration

> Configuration below shows **all fields with their exact default values** from [API-SPEC.md](../../API-SPEC.md). This is for reference only. Clients must not pre-populate all fields — only send fields explicitly set by the caller. See [request-construction.md](../request-construction.md).

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

## Raw Example Request

> Example below matches **exactly** the default configuration from [API-SPEC.md](../../API-SPEC.md). This is the raw upstream shape.

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

## CLI Arguments

See [cli.md](../cli.md) for global flags and naming conventions.

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| | | | |

### CLI Examples

```bash
# Basic
adpgo {kebab-case-name}

# With options
adpgo {kebab-case-name} --field value
```

---

## Raw Example Response

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

## Decoded Result

### Result Type

```
{Task}Result {
    // decoded fields
}
```

### Decoding Rules

Describe how to transform `executionMetaData` into the typed result:
1. Map field `X` to result field `Y`
2. Parse JSON string field `Z` into type `T`
3. Coerce string `"true"`/`"false"` to boolean
4. Coerce string `"100"` to integer

---

## executionMetaData Contract

| Field | Type | Description |
|-------|------|-------------|
| | | |

### JSON String Fields

Some `executionMetaData` fields contain JSON strings that must be parsed:

| Field | Parse As |
|-------|----------|
| | |

---

## Failure Response

On `executionStatus: "failed"`:

```json
{
  "executionId": "uuid",
  "taskType": "{Task Name}",
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
- **Decoded Result types must be language-agnostic** — use TypeScript-like notation:

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
