# Request Construction

Describes how clients should build task requests. This contract is language-agnostic. Clients must not eagerly populate all upstream default values — they must only serialize fields that were explicitly set.

---

## Core Principle: Sparse Requests

Clients should build requests that contain **only the fields explicitly set by the caller**.

Upstream default values from `API-SPEC.md` are documented for reference only. They are **not** pre-populated by the client.

**Why**: The upstream API applies its own defaults server-side. Sending all fields with their default values is unnecessary and makes payloads larger and less readable.

---

## Request Construction API

Clients must expose a progressive request-construction API that allows callers to set only the fields they need.

The concrete pattern is language-idiomatic. Examples:

**Go**: builder with method chaining
```go
result, err := client.ListEntitiesRequest().
    Type("singleMindServer").
    WhiteList("id,displayName").
    Execute(ctx)
```

**Go**: functional options
```go
result, err := client.ListEntities(ctx,
    adp.WithListEntitiesType("singleMindServer"),
    adp.WithListEntitiesWhiteList("id,displayName"),
)
```

**Rust**: builder
```rust
let result = client
    .list_entities()
    .task_type("singleMindServer")
    .white_list("id,displayName")
    .execute()
    .await?;
```

**Python**: builder or keyword arguments
```python
result = (
    client.list_entities()
    .type_("singleMindServer")
    .white_list("id,displayName")
    .execute()
)
```

---

## Request Envelope

The raw request envelope is defined in [api-contract.md](./api-contract.md).

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

When building a request, the client:
1. Sets `taskType` to the task's canonical name.
2. Sets `taskDisplayName` to the task's display name.
3. Sets `taskDescription` to the task's description.
4. In `taskConfiguration`, includes only the common fields the caller explicitly changed from their defaults, plus task-specific fields the caller set.

---

## Minimal Request Example

A minimal `List Entities` request that only sets the entity type:

```json
{
  "taskType": "List Entities",
  "taskConfiguration": {
    "adp_listEntities_type": "singleMindServer",
    "adp_loggingEnabled": true
  },
  "taskDescription": "Writes a list of entities ot an output variable",
  "taskDisplayName": "List entities"
}
```

Compare to the full default configuration in [tasks/list-entities.md](./tasks/list-entities.md). The full configuration has ~30+ fields. A client that only sends what the caller configured produces a much cleaner request.

---

## Common Fields Default Policy

The common fields have these default values (from `API-SPEC.md`):

| Field | Default |
|-------|---------|
| adp_taskActive | true |
| adp_taskTimeout | 0 |
| adp_executionPersistent | true |
| adp_abortWfOnFailure | true |
| adp_loggingEnabled | true |
| adp_cleanUpHistory | false |

These fields follow the same rule as all others: only include them in `taskConfiguration` if the caller explicitly set them to a non-default value.

**Exception**: `adp_loggingEnabled` may be included even when set to `true` if the client always explicitly sets it. However, callers should be able to omit it and rely on the server default.

---

## Task-Specific Fields

Task-specific fields are defined in each task spec in [tasks/](./tasks/). Each task spec documents:

1. **Semantic inputs** — the user-facing fields that make sense to set
2. **Raw field mapping** — how those map to the upstream `taskConfiguration` keys
3. **Defaults** — upstream defaults for reference

Clients should expose the semantic inputs as the primary API. The mapping to raw field names is an implementation detail.

---

## Validation

Clients should validate required fields before sending. Each task spec documents which fields are required.

---

## Async Requests

Async requests use the same request envelope. The only difference is which endpoint is called:

- Sync: `PUT /executeAdpTask`
- Async: `PUT /executeAdpTaskAsync`

The request body is identical.
