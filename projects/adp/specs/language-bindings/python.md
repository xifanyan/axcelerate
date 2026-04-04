# Python Language Binding

Maps the language-agnostic ADP contracts into idiomatic Python APIs.

---

## Project Structure

```
src/python/
├── adppy/
│   ├── __init__.py
│   ├── client.py      # HTTP client
│   ├── types.py       # Shared types
│   ├── error.py       # Error types
│   ├── tasks/
│   │   ├── __init__.py
│   │   ├── list_entities.py
│   │   ├── query_engine.py
│   │   ├── taxonomy_statistic.py
│   │   ├── start_application.py
│   │   ├── export_documents.py
│   │   └── csv_merge.py
├── pyproject.toml
└── adppy/__main__.py  # CLI entrypoint
```

---

## Request Construction

Python clients use a **builder pattern** with keyword arguments.

```python
# Via builder
result = (
    client.list_entities()
    .type_("singleMindServer")
    .white_list("id,displayName")
    .execute()
)

# Async variant
handle = await client.list_entities_async(
    type_="singleMindServer",
)
```

---

## Serialization

Use `dataclasses` with `field(default=None)` and `omit_none=True` to ensure unset fields are not serialized:

```python
from dataclasses import dataclass, field
from typing import Optional

@dataclass
class ListEntitiesConfig:
    adp_list_entities_type: Optional[str] = field(default=None, metadata={"json_key": "adp_listEntities_type"})
    adp_list_entities_white_list: Optional[str] = field(default=None, metadata={"json_key": "adp_listEntities_whiteList"})
    adp_logging_enabled: Optional[bool] = field(default=None, metadata={"json_key": "adp_loggingEnabled"})
```

---

## Response Handling

### Common Response

```python
from dataclasses import dataclass
from typing import Optional, Any, Dict

@dataclass
class TaskResponse:
    execution_id: str
    task_type: str
    logging_enabled: str
    progress_max: int
    execution_status: str
    execution_root_dir: str
    context_id: str
    execution_persistent: str
    progress_current: int
    progress_percentage: float
    task_display_name: str
    execution_meta_data: Optional[Dict[str, Any]]
    error_message: Optional[str] = None
```

### Task-Specific Results

```python
from dataclasses import dataclass
from typing import List, Optional, Dict

# List Entities
@dataclass
class Entity:
    id: str
    display_name: str
    process_status: str
    host_id: str
    host_name: str
    source_for_create_from_existing: bool

@dataclass
class ListEntitiesResult:
    output_file: str
    entities: List[Entity]

# Query Engine
@dataclass
class QueryEngineResult:
    documents_count: int
    aggregated_value: str

# Taxonomy Statistic
@dataclass
class SearchParameter:
    key: str
    value: str

@dataclass
class Category:
    id: str
    display_name: str
    count: int
    properties: Optional[Dict[str, List[str]]] = None

@dataclass
class TaxonomyEntry:
    id: str
    category: List[Category]

@dataclass
class Statistics:
    taxonomy: List[TaxonomyEntry]

@dataclass
class TaxonomyStatisticsDocument:
    date: str
    search_parameter: List[SearchParameter]
    statistics: Statistics

@dataclass
class TaxonomyStatisticResult:
    output_file: str
    statistics: TaxonomyStatisticsDocument

# Start Application
@dataclass
class StartApplicationResult:
    application_url: str

# Export Documents
@dataclass
class ExportDocumentsResult:
    export_file_name: str
    export_path: str
    search_result_size: int

# CSV Merge
# Pending verification
@dataclass
class CSVMergeResult:
    # Pending verification
    pass
```

---

## Error Handling

```python
class TaskExecutionError(Exception):
    def __init__(
        self,
        execution_id: str,
        task_type: str,
        execution_status: str,
        error_message: str,
        execution_meta_data: Optional[Dict[str, Any]] = None,
    ):
        self.execution_id = execution_id
        self.task_type = task_type
        self.execution_status = execution_status
        self.error_message = error_message
        self.execution_meta_data = execution_meta_data
        super().__init__(
            f"task {task_type} failed: {error_message} (executionId={execution_id})"
        )
```

---

## CLI Implementation

Use [click](https://click.palletsprojects.com/) for CLI argument parsing.

```python
import click

@click.group()
@click.option("--host", required=True)
@click.option("--port", default=8443)
@click.option("--user", required=True)
@click.option("--password", required=True)
@click.option("--insecure", is_flag=True)
@click.option("--debug", is_flag=True)
def cli(host, port, user, password, insecure, debug):
    pass

@cli.command("list-entities")
@click.option("--type")
@click.option("--white-list", default="id,displayName")
def list_entities(type_, white_list):
    pass
```
