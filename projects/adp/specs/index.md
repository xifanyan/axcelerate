# Specs Index

Single source of truth for all ADP specifications.

---

## Core Specs

| File | Description |
|------|-------------|
| [http-client.md](./http-client.md) | HTTP client configuration |
| [api-endpoints.md](./api-endpoints.md) | API endpoints |
| [common-types.md](./common-types.md) | Shared types |
| [languages.md](./languages.md) | Language-specific rules (including service layer) |

---

## Tasks

See [tasks/index.md](./tasks/index.md) for complete task listing.

| Task | Description |
|------|-------------|
| [list-entities.md](./tasks/list-entities.md) | List Entities |
| [query-engine.md](./tasks/query-engine.md) | Query Engine |
| [taxonomy-statistic.md](./tasks/taxonomy-statistic.md) | Taxonomy Statistic |

---

## Adding New Tasks

1. Copy [tasks/TEMPLATE.md](./tasks/TEMPLATE.md) to `tasks/new-task.md`
2. Fill in using [API-SPEC.md](../API-SPEC.md)
3. Add to [tasks/index.md](./tasks/index.md)

---

## ADP-Specific Rules

### Task Spec Rules
- Task specs must match [API-SPEC.md](../API-SPEC.md) exactly (field names, values, ordering)
- Example Request = Default Configuration (no custom values)
- **Task Configuration** - Only send fields that need to be changed from defaults; do not include all fields
- Use **Builder Pattern** for task configuration (e.g., `ListEntities().Type("x").WhiteList("y").Execute(client)`)

### Function Naming
- Default functions (e.g., `ListEntities`) are synchronous
- Add `Async` suffix for asynchronous variants (e.g., `ListEntitiesAsync`)

### Logging
- Logging must be **enabled by default** for all tasks
- Debug mode must trace request/response payloads (input and output)
- Enable via CLI: `--debug` or `-d` flag (e.g., `adpgo --debug ... listEntities`)

### CLI Interface
- Must support subcommands for each task
- Global flags: `--host`, `--port`, `--user`, `--password`, `--insecure`, `--debug`
- `--port` default: 8443
- Example: `adpgo --host example.com --user adp --password adp listEntities --type singleMindServer`
- CLI naming: `[project][lang]` (e.g., `adpgo` for Go, `adppy` for Python)

### CLI Output Parsing
- Each task returns different `ExecutionMetaData` fields
- CLI must parse output based on **Example Response** in each task's spec
- See [list-entities.md](./tasks/list-entities.md), [query-engine.md](./tasks/query-engine.md), [taxonomy-statistic.md](./tasks/taxonomy-statistic.md) for task-specific output formats

### CLI Output Rule
- On HTTP 200 and `ExecutionStatus: "success"` - output only the parsed task-specific data (e.g., JSON array for List Entities)
- On failure - output error details including `ExecutionID` and `ExecutionMetaData`

### Response Processing
- Use a **shared response handler** for common fields (`ExecutionID`, `TaskType`, `ExecutionStatus`, `ProgressMax`, `ProgressCurrent`, `ProgressPercentage`)
- Each task should only implement task-specific `ExecutionMetaData` parsing
- Do not duplicate common field handling in each command handler

### Response Format
`ExecutionID`, `TaskType`, `LoggingEnabled`, `ProgressMax`, `ExecutionStatus`, `ExecutionRootDir`, `ContextID`, `ExecutionPersistent`, `ProgressCurrent`, `ProgressPercentage`, `TaskDisplayName`, `ExecutionMetaData`

### Code Generation Conventions

When generating code for ADP client library:

1. **One task per file** - Each task (List Entities, Query Engine, Taxonomy Statistic) should be in its own file (e.g., `list_entities.go`, `query_engine.go`, `taxonomy_statistic.go`)
2. **Shared types** - Common types go in `types.go`
3. **Client** - HTTP client implementation in `client.go`
4. **Config structs** - Place in the same file as the task builder function
5. **CLI** - CLI implementation in `cmd/adpgo/main.go`

Example structure:
```
projects/adp/src/go/
├── cmd/adpgo/
│   └── main.go       # CLI entrypoint
├── client.go         # HTTP client
├── types.go          # Shared types
├── list_entities.go  # List Entities task
├── query_engine.go   # Query Engine task
└── taxonomy_statistic.go  # Taxonomy Statistic task
```

See also: [specs/tasks/TEMPLATE.md](./tasks/TEMPLATE.md), [specs/tasks/index.md](./tasks/index.md), and [languages.md](./languages.md)
