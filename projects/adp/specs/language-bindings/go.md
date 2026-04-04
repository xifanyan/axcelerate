# Go Language Binding

Maps the language-agnostic ADP contracts into idiomatic Go APIs.

---

## Module

```
github.com/xifanyan/axcelerate/adp
```

---

## CLI Framework

Use [urfave/cli/v3](https://github.com/urfave/cli) for CLI implementation.

---

## Project Structure

```
src/go/
├── cmd/adpgo/
│   └── main.go       # CLI entrypoint
├── client.go         # HTTP client
├── types.go          # Shared types
├── list_entities.go  # List Entities task
├── query_engine.go   # Query Engine task
├── taxonomy_statistic.go  # Taxonomy Statistic task
├── csv_merge.go      # CSV Merge task
├── export_documents.go    # Export Documents task
└── start_application.go  # Start Application task
```

---

## Request Construction

Go clients currently use a direct builder pattern.

### Direct Builder

```go
// Via builder
result, err := adp.NewListEntitiesBuilder(client).
    Type("singleMindServer").
    WhiteList("id,displayName").
    Execute(ctx)
```

### Async Variant

```go
result, err := adp.NewListEntitiesBuilder(client).
    Type("singleMindServer").
    ExecuteAsync(ctx)
```

---

## Response Handling

### Common Response

```go
type TaskResponse struct {
    ExecutionID         string         `json:"executionId"`
    TaskType            string         `json:"taskType"`
    LoggingEnabled      string         `json:"loggingEnabled"`
    ProgressMax         int            `json:"progressMax"`
    ExecutionStatus     string         `json:"executionStatus"`
    ExecutionRootDir    string         `json:"executionRootDir"`
    ContextID           string         `json:"contextId"`
    ExecutionPersistent string         `json:"executionPersistent"`
    ProgressCurrent     int            `json:"progressCurrent"`
    ProgressPercentage  float64        `json:"progressPercentage"`
    TaskDisplayName     string         `json:"taskDisplayName"`
    ExecutionMetaData   map[string]any `json:"executionMetaData"`
    ErrorMessage        string         `json:"errorMessage,omitempty"`
}
```

### Task-Specific Results

```go
// List Entities
type ListEntitiesResult struct {
    OutputFile string
    Entities   []Entity
}

type Entity struct {
    ID                     string `json:"id"`
    DisplayName            string `json:"displayName"`
    ProcessStatus          string `json:"processStatus"`
    HostID                 string `json:"hostId"`
    HostName               string `json:"hostName"`
    SourceForCreateFromExisting bool `json:"sourceForCreateFromExisting"`
}

// Query Engine
type QueryEngineResult struct {
    DocumentsCount  int
    AggregatedValue string
}

// Taxonomy Statistic
type TaxonomyStatisticResult struct {
    OutputFile string
    Statistics TaxonomyStatisticsDocument
}

type TaxonomyStatisticsDocument struct {
    Date           string
    SearchParameter []SearchParameter
    Statistics     Statistics
}

type SearchParameter struct {
    Key   string `json:"key"`
    Value string `json:"value"`
}

type Statistics struct {
    Taxonomy []TaxonomyEntry
}

type TaxonomyEntry struct {
    ID       string
    Category []Category
}

type Category struct {
    ID          string
    DisplayName string
    Count       int
    Properties  map[string][]string // present only when listCategoryProperties is enabled
}

// Start Application
type StartApplicationResult struct {
    ApplicationURL string
}

// Export Documents
type ExportDocumentsResult struct {
    ExportFileName   string
    ExportPath       string
    SearchResultSize int
}

// CSV Merge
// Pending verification
type CSVMergeResult struct {
    // Pending verification
}
```

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

The current Go layout uses task builders directly and does not require a separate `service.go` wrapper.

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
