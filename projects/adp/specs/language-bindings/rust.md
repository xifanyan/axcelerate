# Rust Language Binding

Maps the language-agnostic ADP contracts into idiomatic Rust APIs.

---

## Project Structure

```
src/rust/
├── Cargo.toml
├── src/
│   ├── main.rs          # CLI entrypoint
│   ├── client.rs        # HTTP client
│   ├── types.rs         # Shared types
│   ├── error.rs         # Error types
│   ├── tasks/
│   │   ├── mod.rs
│   │   ├── list_entities.rs
│   │   ├── query_engine.rs
│   │   ├── taxonomy_statistic.rs
│   │   ├── start_application.rs
│   │   ├── export_documents.rs
│   │   └── csv_merge.rs
```

---

## Generated Code Path

Generated code lives at: `$HOME/ai-generated/adp/rust/`

---

## Request Construction

Rust clients use a **builder pattern**.

```rust
// Via builder
let result = client
    .list_entities()
    .task_type("singleMindServer")
    .white_list("id,displayName")
    .execute()
    .await?;

// Async variant
let handle = client
    .list_entities_async()
    .task_type("singleMindServer")
    .execute()
    .await?;
```

---

## Serialization

Use `serde` with `skip_serializing_if = "Option::is_none"` to ensure unset fields are not serialized:

```rust
#[derive(Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
struct ListEntitiesConfig {
    #[serde(skip_serializing_if = "Option::is_none")]
    adp_list_entities_type: Option<String>,
    #[serde(skip_serializing_if = "Option::is_none")]
    adp_list_entities_white_list: Option<String>,
    #[serde(skip_serializing_if = "Option::is_none")]
    adp_logging_enabled: Option<bool>,
    // ... only include fields that are set
}
```

---

## Response Handling

### Common Response

```rust
#[derive(Debug, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct TaskResponse {
    pub execution_id: String,
    pub task_type: String,
    pub logging_enabled: String,
    pub progress_max: i32,
    pub execution_status: String,
    pub execution_root_dir: String,
    pub context_id: String,
    pub execution_persistent: String,
    pub progress_current: i32,
    pub progress_percentage: f64,
    pub task_display_name: String,
    pub execution_meta_data: Option<Map<String, Value>>,
    #[serde(rename = "errorMessage")]
    pub error_message: Option<String>,
}
```

### Task-Specific Results

```rust
// List Entities
#[derive(Debug)]
pub struct ListEntitiesResult {
    pub output_file: String,
    pub entities: Vec<Entity>,
}

#[derive(Debug, Deserialize)]
pub struct Entity {
    pub id: String,
    #[serde(rename = "displayName")]
    pub display_name: String,
    #[serde(rename = "processStatus")]
    pub process_status: String,
    #[serde(rename = "hostId")]
    pub host_id: String,
    #[serde(rename = "hostName")]
    pub host_name: String,
    #[serde(rename = "sourceForCreateFromExisting")]
    pub source_for_create_from_existing: bool,
}

// Query Engine
#[derive(Debug)]
pub struct QueryEngineResult {
    pub documents_count: i32,
    pub aggregated_value: String,
}

// Taxonomy Statistic
#[derive(Debug)]
pub struct TaxonomyStatisticResult {
    pub output_file: String,
    pub statistics: TaxonomyStatisticsDocument,
}

#[derive(Debug, Deserialize)]
pub struct TaxonomyStatisticsDocument {
    pub date: String,
    pub search_parameter: Vec<SearchParameter>,
    pub statistics: Statistics,
}

#[derive(Debug, Deserialize)]
pub struct SearchParameter {
    pub key: String,
    pub value: String,
}

#[derive(Debug, Deserialize)]
pub struct Statistics {
    pub taxonomy: Vec<TaxonomyEntry>,
}

#[derive(Debug, Deserialize)]
pub struct TaxonomyEntry {
    pub id: String,
    pub category: Vec<Category>,
}

#[derive(Debug, Deserialize)]
pub struct Category {
    pub id: String,
    pub display_name: String,
    pub count: i32,
    #[serde(default)]
    pub properties: Option<Map<String, Vec<String>>>,
}

// Start Application
#[derive(Debug)]
pub struct StartApplicationResult {
    pub application_url: String,
}

// Export Documents
#[derive(Debug)]
pub struct ExportDocumentsResult {
    pub export_file_name: String,
    pub export_path: String,
    pub search_result_size: i32,
}

// CSV Merge
// Pending verification
#[derive(Debug)]
pub struct CSVMergeResult {
    // Pending verification
}
```

---

## Error Handling

```rust
#[derive(Debug)]
pub struct TaskExecutionError {
    pub execution_id: String,
    pub task_type: String,
    pub execution_status: String,
    pub error_message: String,
    pub execution_meta_data: Option<Map<String, Value>>,
}

impl std::fmt::Display for TaskExecutionError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(
            f,
            "task {} failed: {} (executionId={})",
            self.task_type, self.error_message, self.execution_id
        )
    }
}

impl std::error::Error for TaskExecutionError {}
```

---

## CLI Implementation

Use [clap](https://clap.rs/) for CLI argument parsing.

```rust
use clap::{Parser, Subcommand};

#[derive(Parser)]
#[command(name = "adprs")]
struct Cli {
    #[command(flatten)]
    global: GlobalArgs,
    #[command(subcommand)]
    command: Commands,
}

#[derive(Subcommand)]
enum Commands {
    ListEntities(ListEntitiesArgs),
    QueryEngine(QueryEngineArgs),
    // ...
}
```
