# ADP Go Client Design

## Context

This document defines the initial Go implementation for the `adp` project under `projects/adp/src/go/`.

The implementation will:

- cover all currently specified ADP tasks,
- provide a full Go client library and CLI,
- include minimal but meaningful automated tests,
- follow the written ADP specs strictly, including less-certain tasks,
- expose both async primitives and convenience wait methods.

## Goals

- Generate a complete, compilable Go codebase for the ADP client.
- Keep the implementation aligned with the existing ADP specs and Go binding.
- Preserve sparse request construction so only caller-set fields are serialized.
- Decode task-specific `executionMetaData` into typed Go results.
- Provide a usable `adpgo` CLI that mirrors the task specs.

## Non-Goals

- Auto-generating code from specs.
- Adding extra frameworks beyond `urfave/cli/v3`.
- Refactoring the specs themselves.
- Hiding all upstream quirks behind heavy abstraction.

## Recommended Approach

Use thin task files over a shared transport and decoder core.

Why:

- It matches the repo rule of one task per file.
- It keeps each task directly traceable to its spec.
- It minimizes duplication in transport, polling, logging, and error handling.
- It is smaller and easier to maintain than a registry-heavy generic framework.

Alternatives considered:

1. Per-task standalone implementations.
   This is explicit but duplicates transport, async polling, and decoding behavior.
2. Heavy generic registry framework.
   This reduces some wiring but adds abstraction that is not justified for the current scope.

## Project Layout

```text
projects/adp/src/go/
├── go.mod
├── client.go
├── types.go
├── decode.go
├── list_entities.go
├── query_engine.go
├── taxonomy_statistic.go
├── start_application.go
├── csv_merge.go
├── export_documents.go
├── read_configuration.go
├── create_ocr_job.go
├── cli_task.go
├── cli_helpers.go
├── client_test.go
├── decode_test.go
├── parsing_test.go
└── cmd/adpgo/
    └── main.go
```

Notes:

- `client.go` owns HTTP transport, authentication headers, debug tracing, request execution, async polling, and shared error handling.
- `types.go` owns common envelopes, shared input types, shared result types, and error types.
- `decode.go` owns small reusable decode and coercion helpers used by task decoders.
- Each task file owns its builder type, semantic setters, validation, raw field mapping, and typed result decoding.
- `cli_helpers.go` owns repeated CLI parsing and output formatting helpers.

## Module And Dependencies

- Module path: `github.com/xifanyan/axcelerate/adp`
- Go version: set `go.mod` to the version supported by the locally installed Go toolchain discovered during implementation.
- External dependency: `github.com/urfave/cli/v3`
- Standard library for HTTP, JSON, TLS configuration, polling, and tests.

## Core Client Design

### Client Configuration

The client will expose a constructor using explicit configuration:

```go
type ClientConfig struct {
    BaseURL  string
    Username string
    Password string
    Insecure bool
    Timeout  time.Duration
    Debug    bool
    DebugOut io.Writer
}
```

`Client` will retain:

- base URL,
- auth headers,
- configured `http.Client`,
- timeout,
- debug flag and writer.

The CLI will build `BaseURL` from `--host`, `--port`, and `--path`, with `/adp/rest/api/task` as the default path.

### HTTP Behavior

All requests will:

- use HTTP `PUT`,
- send JSON,
- set `Content-Type: application/json`,
- set `Auth-Username` and `Auth-Password` headers,
- target one of:
  - `/executeAdpTask`
  - `/executeAdpTaskAsync`
  - `/statusAndProgress`

### Common Envelopes

The Go types will mirror the shared contract exactly, including string-encoded boolean fields in responses:

```go
type TaskResponse struct {
    ExecutionID         string      `json:"executionId"`
    TaskType            string      `json:"taskType"`
    LoggingEnabled      string      `json:"loggingEnabled"`
    ProgressMax         int         `json:"progressMax"`
    ExecutionStatus     string      `json:"executionStatus"`
    ExecutionRootDir    string      `json:"executionRootDir"`
    ContextID           string      `json:"contextId"`
    ExecutionPersistent string      `json:"executionPersistent"`
    ProgressCurrent     int         `json:"progressCurrent"`
    ProgressPercentage  float64     `json:"progressPercentage"`
    TaskDisplayName     string      `json:"taskDisplayName"`
    ExecutionMetaData   any         `json:"executionMetaData"`
    ErrorMessage        string      `json:"errorMessage,omitempty"`
}
```

`ExecutionMetaData` stays flexible in the common envelope because some tasks return objects, some return JSON strings inside objects, and some return empty arrays.

### Error Handling

Failed task responses will become:

```go
type TaskExecutionError struct {
    ExecutionID     string
    TaskType        string
    ExecutionStatus string
    ErrorMessage    string
    ExecutionMetaData any
}
```

This preserves the correlation fields required by the specs.

## Request Construction

### Sparse Requests

Each task uses a builder that records only fields explicitly set by the caller. The library will not pre-populate the full upstream default configuration.

Each builder will:

- store semantic values in typed fields,
- track whether a setter was called,
- translate semantic inputs into raw `taskConfiguration` keys only during execution,
- include common fields only when explicitly changed.

### Common Builder Fields

Every task builder will support shared setters for:

- `TaskActive(bool)`
- `TaskTimeout(int)`
- `ExecutionPersistent(bool)`
- `AbortWfOnFailure(bool)`
- `LoggingEnabled(bool)`
- `CleanUpHistory(bool)`

These setters apply sparse serialization rules just like task-specific fields.

### Execution Methods

Each builder will expose concrete, task-specific execution methods.

Sync-capable tasks will expose:

- `Execute(ctx context.Context) (<TaskResult>, error)` returning the task's typed result.

Async support will expose:

- `Start(ctx context.Context) (*TaskResponse, error)` for the initial async request,
- `Poll(ctx context.Context, executionID string) (*TaskResponse, error)` on `Client`,
- `Wait(ctx context.Context, executionID string, interval time.Duration) (*TaskResponse, error)` on `Client`,
- task-specific convenience methods that wait and then decode the task's typed result.

`Create OCR Job` is async-only. Its builder will expose async start and wait flows, but not a synchronous `Execute` path.

## Task API Design

### List Entities

Public builder:

- `Type(string)`
- `ID(string)`
- `RelatedEntity(string)`
- `WhiteList(string)`
- `Workspace(string)`
- `Status(string)`

Validation:

- no required semantic fields.

Decoded result:

```go
type ListEntitiesResult struct {
    OutputFile string   `json:"outputFile"`
    Entities   []Entity `json:"entities"`
}
```

### Query Engine

Public builder:

- `EngineName(string)`
- `EngineQuery(string)`
- `EngineUserName(string)`
- `EngineUserPassword(string)`
- `EngineTaxonomies([]EngineTaxonomyArg)` and append-style helpers
- `ApplicationIdentifier(string)`

Validation:

- `engineName` is required.

Decoded result:

```go
type QueryEngineResult struct {
    DocumentsCount  int    `json:"documentsCount"`
    AggregatedValue string `json:"aggregatedValue"`
}
```

### Taxonomy Statistic

Public builder:

- `EngineName(string)`
- `EngineQuery(string)`
- `ComputeCounts(bool)`
- `ListCategoryProperties(bool)`
- `EngineTaxonomies([]EngineTaxonomyArg)` and append-style helpers
- `OutputTaxonomies([]OutputTaxonomiesArg)` and append-style helpers
- `ApplicationIdentifier(string)`

Validation:

- `engineName` is required.

Decoded result includes the nested statistics document with optional category properties.

### Start Application

Public builder:

- `ApplicationIdentifier(string)`
- `UseHTTPS(bool)`

Validation:

- no required semantic fields in the written spec.

Decoded result:

```go
type StartApplicationResult struct {
    ApplicationURL string `json:"applicationUrl"`
}
```

### CSV Merge

Public builder maps the semantic fields from the task spec, including:

- `CSVFile(string)`
- `CSVIDFieldKey(string)`
- `MergeType(string)`
- `CSVMode(string)`
- `EngineName(string)`
- `EngineUser(string)`
- `EnginePassword(string)`
- `EngineIDFieldKey(string)`
- `ApplicationIdentifier(string)`
- `FieldMappings([]map[string]any)`
- `FieldSeparator(string)`
- `ImageBasePath(string)`
- `NativeBasePath(string)`
- `CSVFieldImageLocation(string)`
- `CSVFieldNativeLocation(string)`
- `MultiValueDelimiter(string)`
- `TextIndicator(string)`
- `DoNotChangeProtectedDocuments(bool)`

Validation:

- `csvFile` is required.

Decoded result is an empty struct because success metadata is an empty array.

### Export Documents

Public builder:

- `FieldSeparator(string)`
- `WaitForExport(bool)`
- `Query(string)`
- `ApplicationIdentifier(string)`
- `ApplicationType(string)`
- `EngineIdentifier(string)`
- `EngineUser(string)`
- `EnginePassword(string)`
- `ExportName(string)`
- `ExportFields(string)`
- `ExportDirectory(string)`
- `FileEnding(string)`

Validation:

- no required semantic fields in the spec.

Decoded result:

```go
type ExportDocumentsResult struct {
    ExportFileName   string `json:"exportFileName"`
    ExportPath       string `json:"exportPath"`
    SearchResultSize int    `json:"searchResultSize"`
}
```

### Read Configuration

Public builder:

- `EntityIDToRead(string)`
- `ConfigsToRead([]ConfigArg)`
- `FileFormat(string)`

Validation:

- no required semantic fields in the spec.

Important mapping rule:

- public `ConfigArg` uses camelCase fields,
- raw request payload converts those to the upstream object keys with spaces such as `Configuration ID` and `Field list`.

Decoded result returns the output file and parsed configuration payload.

### Create OCR Job

Public builder:

- `EngineName(string)`
- `Query(string)`
- `EngineUserName(string)`
- `EngineUserPassword(string)`
- `JobName(string)`
- `JobDescription(string)`
- `JobPriority(int)`
- `ApplicationIdentifier(string)`
- `ApplicationType(string)`
- `Wait(bool)`
- `EngineType(string)`
- `ListOfJobProperties(string)`
- `GlobalSearchJSON(string)`
- `GlobalSearchID(string)`
- `Restrictions([]EngineTaxonomyArg)` and append-style helpers
- `AdvancedRestrictions([]EngineTaxonomyArg)` and append-style helpers
- `MainQueryType(string)`

Validation:

- follow the spec literally and do not invent extra required fields.

Execution model:

- async-only,
- initial response may contain empty `executionStatus`,
- completion polling uses the standard status endpoint,
- typed result is an empty struct.

### CLI Task

Public builder:

- `BatchScriptPath(string)`
- `BatchScriptParameters([]CLIBatchParameter)`
- `WorkingDirectory(string)`
- `BatchScriptJSONLogOutput(string)`
- `BatchScriptRedirectLogging(bool)`
- `BatchScriptPositiveExecutionCodes(string)`
- `BatchScriptFilterPasswords(bool)`
- `BatchScriptLoggingDirectory(string)`
- `BatchScriptResultCode(string)`
- `BatchScriptResultLogPath(string)`
- `BatchScriptErrorLogPath(string)`

Validation:

- `batchScriptPath` is required.

Decoded result:

```go
type CLIResult struct {
    Result     int                    `json:"result"`
    JSONOutput map[string]any         `json:"jsonOutput,omitempty"`
    ErrorPath  string                 `json:"errorPath"`
    ResultPath string                 `json:"resultPath"`
}
```

## Shared Types

### EngineTaxonomyArg

```go
type EngineTaxonomyArg struct {
    Taxonomy string `json:"taxonomy"`
    Negation bool   `json:"negation"`
    Query    string `json:"query"`
}
```

### OutputTaxonomiesArg

```go
type OutputTaxonomiesArg struct {
    Taxonomy                  string `json:"taxonomy"`
    Mode                      string `json:"mode"`
    MaximumNumberOfCategories int    `json:"maximumNumberOfCategories"`
}
```

### Read Configuration Semantic Input

```go
type ConfigArg struct {
    ConfigurationID     string
    DynamicComponentNames string
    FieldList           string
    NameValueList       string
    ApplicationType     string
    EntityType          string
}
```

### CLI Parameter Type

```go
type CLIBatchParameter struct {
    Parameter string `json:"parameter"`
}
```

## Decoding Rules

Shared decode helpers will support:

- object extraction from `executionMetaData`,
- empty-array acceptance for tasks that intentionally return `[]`,
- JSON-string decoding into typed structs and maps,
- integer coercion from strings,
- boolean coercion helpers where needed for parsed inner payloads,
- null-safe handling for failure responses.

Task decoders will be strict enough to catch malformed payloads but not invent undocumented defaults.

## CLI Design

### Global Flags

`adpgo` will support:

- `--host`
- `--port` with default `8443`
- `--path` with default `/adp/rest/api/task`
- `--user`
- `--password`
- `--insecure`
- `--debug` and `-d`

### Subcommands

The CLI will register all current task commands:

- `list-entities`
- `query-engine`
- `taxonomy-statistic`
- `start-application`
- `csv-merge`
- `export-documents`
- `read-configuration`
- `create-ocr-job`
- `cli`

### Flag Mapping

- Use camelCase flags exactly as the specs require.
- Parse repeatable taxonomy filters from `Taxonomy=Query` and `Taxonomy!=Query`.
- Parse `outputTaxonomies` from either a comma-separated list or JSON array.
- Parse `configsToRead`, `fieldMappings`, and `batchScriptParameters` from JSON arrays.

### CLI Output

Success output uses decoded results only:

- `list-entities`: JSON array of entities
- `query-engine`: JSON object with `documentsCount` and `aggregatedValue`
- `taxonomy-statistic`: decoded statistics JSON
- `start-application`: plain URL string
- `export-documents`: JSON object
- `read-configuration`: decoded configuration JSON
- `csv-merge`: empty JSON object `{}`
- `create-ocr-job`: empty JSON object `{}` after successful completion, while low-level async data remains available in the library
- `cli`: decoded JSON object

Failure output prints:

- `Error: <errorMessage>`
- `ExecutionID: <executionId>`
- `TaskType: <taskType>`

Debug mode will additionally trace full request and response JSON through the shared client path.

## Testing Strategy

This implementation will include minimal but meaningful tests.

### Unit Tests

- sparse request serialization for representative tasks,
- decode helpers for object, JSON string, empty array, and numeric string cases,
- taxonomy shorthand parsing,
- JSON-array CLI input parsing,
- key task validators for required fields.

### HTTP Tests

Use `httptest` to verify:

- sync success flow,
- sync failure flow,
- async start and polling flow,
- debug logging includes request and response bodies,
- auth headers and PUT method are sent correctly.

### Coverage Focus

Prioritize tests for the most failure-prone behavior:

- sparse request bodies,
- async polling,
- JSON-string metadata decoding,
- mixed metadata shapes across tasks.

## Implementation Risks And Responses

### Mixed `executionMetaData` Shapes

Risk:

- some tasks return objects,
- some return empty arrays,
- some embed JSON strings inside objects.

Response:

- keep the shared envelope flexible,
- centralize metadata shape checks in decode helpers,
- keep each task decoder explicit.

### Spec Fidelity Versus Go Idioms

Risk:

- upstream raw keys are inconsistent in places.

Response:

- keep public Go types idiomatic,
- keep raw-key translation local to request and decode boundaries.

### Async Initial Response Variance

Risk:

- `Create OCR Job` may return an empty execution status initially.

Response:

- treat a present `executionId` as sufficient to begin polling for this task,
- avoid imposing extra assumptions beyond the written spec.

## Implementation Sequence

1. Create module, shared client, types, and decode helpers.
2. Implement sync task builders and decoders.
3. Implement async and polling support, including `Create OCR Job`.
4. Implement CLI helpers and all subcommands.
5. Add focused tests for serialization, decoding, parsing, and async transport.

## Acceptance Criteria

The Go implementation is considered complete when:

- all nine current tasks are implemented,
- the code compiles as a Go module,
- `adpgo` exposes all specified global flags and subcommands,
- builders send sparse requests only,
- task results decode into typed Go values,
- async primitives and convenience wait methods are both available,
- minimal tests cover core transport, parsing, and decode behavior.
