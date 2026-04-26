# Go Language Binding

Maps the language-agnostic ADP contracts into idiomatic Go APIs.

---

## Module

```
github.com/xifanyan/adp/v2
```

---

## Generated Code Path

Generated code lives at: `$HOME/ai-generated/adp/go/`

---

## CLI Framework

Use [urfave/cli/v3](https://github.com/urfave/cli) for CLI implementation.

---

## Project Structure

```
go/
├── cmd/adpgo/
│   ├── main.go       # CLI entrypoint
│   └── main_test.go
├── client.go         # Client struct + NewClient
├── client_test.go
├── types.go          # Shared types + result types
├── list_entities.go  # List Entities task options + method
├── query_engine.go   # Query Engine task options + method
├── taxonomy_statistic.go  # Taxonomy Statistic task options + method
├── csv_merge.go      # CSV Merge task options + method
├── export_documents.go    # Export Documents task options + method
├── start_application.go  # Start Application task options + method
├── read_configuration.go # Read Configuration task options + method
├── create_ocr_job.go     # Create OCR Job task options + method
├── cli_task.go       # CLI task options + method
├── cli_helpers.go    # CLI helper functions
├── decode.go         # JSON response decoding
├── decode_test.go
├── options_test.go
├── parsing_test.go
├── go.mod
├── go.sum
├── AGENTS.md         # Agent guidelines
└── task_spec.json    # Task specification source
```

---

## Request Construction

Go clients use the functional options pattern for request construction.

### Functional Options

```go
// Use prefixed functional options to configure request
result, err := client.ListEntities(ctx,
    adp.WithListEntitiesType("singleMindServer"),
    adp.WithListEntitiesWhiteList("id,displayName"),
)
```

### Task-Specific Client Methods

Each task has its own client method returning the appropriate result type:

```go
// ListEntities -> *ListEntitiesResult
result, err := client.ListEntities(ctx, adp.WithListEntitiesType("singleMindServer"))

// QueryEngine -> *QueryEngineResult
result, err := client.QueryEngine(ctx, adp.WithQueryEngineEngineID("engine-123"))

// TaxonomyStatistic -> *TaxonomyStatisticResult
result, err := client.TaxonomyStatistic(ctx, adp.WithTaxonomyStatisticTaxonomyID("tax-456"))

// StartApplication -> *StartApplicationResult
result, err := client.StartApplication(ctx, adp.WithStartApplicationApplicationID("app-789"))
```

### Client

Go uses a concrete `Client` type with no interface. Users instantiate `Client` directly.

```go
type Client struct {
    httpClient *http.Client
    host       string
    user       string
    password   string
}

func NewClient(host, user, password string) *Client {
    return &Client{
        httpClient: &http.Client{},
        host:       host,
        user:       user,
        password:   password,
    }
}
```

### ListEntities Implementation

```go
func (c *Client) ListEntities(ctx context.Context, opts ...ListEntitiesOption) (*ListEntitiesResult, error) {
    cfg := &ListEntitiesConfig{}
    for _, opt := range opts {
        opt(cfg)
    }

    if err := cfg.Validate(); err != nil {
        return nil, err
    }

    taskReq := TaskRequest{
        TaskType:           "List Entities",
        TaskConfiguration:  cfg.ToTaskConfiguration(),
        TaskDescription:    cfg.TaskDescription,
        TaskDisplayName:    cfg.TaskDisplayName,
    }

    resp, err := c.doRequest(ctx, taskReq)
    if err != nil {
        return nil, err
    }

    return decodeListEntitiesResult(resp)
}

func (c *Client) doRequest(ctx context.Context, req TaskRequest) (*TaskResponse, error) {
    // HTTP POST to ADP endpoint
    // ...
}
```

### Common Configuration

All tasks share common fields. Each task config embeds `CommonConfig`. Defaults are set in `ToTaskConfiguration()`:

```go
type CommonConfig struct {
    LoggingEnabled       *bool    // default: false (nil means not set)
    TaskActive          *bool    // nil means use API default
    TaskTimeout         *time.Duration // nil means not set
    ExecutionPersistent *bool    // nil means not set
    AbortOnFailure      *bool    // nil means not set
    TaskDescription     string   // default: TaskType
    TaskDisplayName     string   // default: TaskType
}
```

### Functional Options Definition

```go
type ListEntitiesOption func(*ListEntitiesConfig)

type ListEntitiesConfig struct {
    CommonConfig  // embedded common fields
    Type          string
    ID            string
    RelatedEntity string
    WhiteList     string
    Workspace     string
    Status        string
}

// Common options apply to any config embedding CommonConfig
func WithLogging(enabled bool) func(*CommonConfig) {
    return func(c *CommonConfig) {
        c.LoggingEnabled = &enabled
    }
}

func WithTaskActive(active bool) func(*CommonConfig) {
    return func(c *CommonConfig) {
        c.TaskActive = &active
    }
}

func WithTaskTimeout(timeout time.Duration) func(*CommonConfig) {
    return func(c *CommonConfig) {
        c.TaskTimeout = &timeout
    }
}

func WithExecutionPersistent(persistent bool) func(*CommonConfig) {
    return func(c *CommonConfig) {
        c.ExecutionPersistent = &persistent
    }
}

func WithAbortOnFailure(abort bool) func(*CommonConfig) {
    return func(c *CommonConfig) {
        c.AbortOnFailure = &abort
    }
}

func WithTaskDescription(desc string) func(*CommonConfig) {
    return func(c *CommonConfig) {
        c.TaskDescription = desc
    }
}

func WithTaskDisplayName(name string) func(*CommonConfig) {
    return func(c *CommonConfig) {
        c.TaskDisplayName = name
    }
}

// Task-specific options
func WithListEntitiesType(t string) ListEntitiesOption {
    return func(c *ListEntitiesConfig) {
        c.Type = t
    }
}

func WithListEntitiesWhiteList(w string) ListEntitiesOption {
    return func(c *ListEntitiesConfig) {
        c.WhiteList = w
    }
}

func (c *ListEntitiesConfig) Validate() error { /* only validation, no defaults */ }

func (c *ListEntitiesConfig) ToTaskConfiguration() map[string]any {
    cfg := c.CommonConfig
    result := map[string]any{}

    // Only set if explicitly provided
    if cfg.LoggingEnabled != nil {
        result["adp_loggingEnabled"] = *cfg.LoggingEnabled
    }
    if cfg.TaskActive != nil {
        result["adp_taskActive"] = *cfg.TaskActive
    }
    if cfg.TaskTimeout != nil {
        result["adp_taskTimeout"] = *cfg.TaskTimeout
    }
    if cfg.ExecutionPersistent != nil {
        result["adp_executionPersistent"] = *cfg.ExecutionPersistent
    }
    if cfg.AbortOnFailure != nil {
        result["adp_abortWfOnFailure"] = *cfg.AbortOnFailure
    }

    // TaskDescription and TaskDisplayName default to TaskType
    if cfg.TaskDescription == "" {
        cfg.TaskDescription = "List Entities"
    }
    if cfg.TaskDisplayName == "" {
        cfg.TaskDisplayName = "List Entities"
    }

    // task-specific fields...

    return result
}
```

### Usage

```go
result, err := client.ListEntities(ctx,
    adp.WithLogging(true),                    // common option
    adp.WithTaskTimeout(30*time.Second),      // common option
    adp.WithListEntitiesType("singleMindServer"), // task-specific
)

---

## Error Handling

```go
type TaskExecutionError struct {
    ExecutionID     string
    TaskType        string
    ExecutionStatus string
    ErrorMessage    string
    ExecutionMetaData map[string]any
}

func (e *TaskExecutionError) Error() string {
    return fmt.Sprintf("task %s failed: %s (executionId=%s)", e.TaskType, e.ErrorMessage, e.ExecutionID)
}
```

### Public Surface

Go clients use a concrete `Client` type. Each task file contains:
- Option functions (`WithXXX`)
- Config struct (`XXXConfig`)
- Client method implementation on `*Client`

No interface required — users instantiate `Client` directly.

---

## JSON Serialization

- Use `omitempty` for optional fields in request config structs
- Unset fields must not be serialized (use `omitempty` or `nil` pointers)
- Do not pre-populate default values

---

## CLI Implementation

```go
cmd := &cli.Command{
    Name:  "list-entities",
    Usage: "List entities",
    Flags: []cli.Flag{
        &cli.StringFlag{Name: "type"},
        &cli.StringFlag{Name: "whiteList", Value: "id,displayName"},
    },
    Action: handleListEntities,
}
```
