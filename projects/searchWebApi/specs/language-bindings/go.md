# Go CLI Bindings Specification

## Overview

The Go CLI (`searchwebapi-example`) provides a command-line interface to the searchWebApi using a nested command structure.

## Command Structure

```
searchwebapi-example [global flags] <command> <subcommand> [args] [command flags]

Commands:
  projects         List projects and get project resources
  collections     List collections, fields, filters, and folder values
  records         Search, fetch, and change records
  binary          Get binary content by search or record id
  measures        Get measure cube data
  cached-searches  List and drop cached searches
  change-queue    Wait for pending changes to complete
  insert-remove   Insert/remove transactions and bulk operations
  session         Session operations (login/logout)
```

### Projects

```
searchwebapi-example projects list
searchwebapi-example projects get <projectId>
```

| Subcommand | Operation | HTTP Method | Path |
|------------|-----------|-------------|------|
| `list` | ListProjects | GET | /projects |
| `get <projectId>` | GetProjectResources | GET | /projects/{projectId} |

### Collections

```
searchwebapi-example collections list <projectId>
searchwebapi-example collections get <projectId> <collectionId>
searchwebapi-example collections fields <projectId> <collectionId>
searchwebapi-example collections filter-list <projectId> <collectionId>
searchwebapi-example collections filter-get <projectId> <collectionId> <fieldId>
searchwebapi-example collections values <projectId> <collectionId> <fieldId> [options]
```

| Subcommand | Operation | HTTP Method | Path |
|-----------|-----------|-------------|------|
| `list <projectId>` | ListCollections | GET | /projects/{projectId}/collections |
| `get <projectId> <collectionId>` | GetCollectionResources | GET | /projects/{projectId}/collections/{collectionId} |
| `fields <projectId> <collectionId>` | GetFields | GET | /projects/{projectId}/collections/{collectionId}/fields |
| `filter-list <projectId> <collectionId>` | GetFolderFields | GET | /projects/{projectId}/collections/{collectionId}/filters |
| `filter-get <projectId> <collectionId> <fieldId>` | GetFolderFieldResources | GET | /projects/{projectId}/collections/{collectionId}/filters/{fieldId} |
| `values <projectId> <collectionId> <fieldId>` | GetFolderValues | GET | /projects/{projectId}/collections/{collectionId}/filters/{fieldId}/values |

### Records

```
searchwebapi-example records search <projectId> [-c collection] [-q query] [--language code] [--order expr] [--page N] [--limit N]
searchwebapi-example records search-stream <projectId> [-c collection] [-q query] [--language code] [--order expr] [--page N] [--limit N]
searchwebapi-example records get <projectId> <recordId> [-c collection]
searchwebapi-example records content <projectId> <recordId> [-c collection] [--body] [--page N] [--fields csv] [--summarize]
searchwebapi-example records change <projectId> <recordId> [-c collection] --body-file <path> [--block-until-complete]
searchwebapi-example records change-all <projectId> [-c collection] --body-file <path> [-q query] [--language code] [--block-until-complete]
searchwebapi-example records in-doc-search <projectId> <recordId> [-c collection] [--highlight-search-term-query q] [--highlight-user-terms csv]
searchwebapi-example records token create <projectId> [-c collection] [-q query] [--language code] [--join-restriction expr] [--order expr]
searchwebapi-example records token touch <projectId> [-c collection] --body-file <path>
searchwebapi-example records token delete <projectId> [-c collection] --body-file <path>
searchwebapi-example records token snapshot <projectId> [-c collection] --body-file <path> [--top-n N]
searchwebapi-example records highlight <projectId> [-c collection] [-q query] [--language code] [--join-restriction expr] [--search-cache-control value]
```

| Subcommand | Operation | HTTP Method | Path |
|-----------|-----------|-------------|------|
| `search <projectId>` | SearchRecords | GET | /projects/{projectId}/collections/{collectionId}/records |
| `search-stream <projectId>` | SearchRecordsStream | GET | /projects/{projectId}/collections/{collectionId}/records (NDJSON) |
| `get <projectId> <recordId>` | GetRecordResources | GET | /projects/{projectId}/collections/{collectionId}/records/{recordId} |
| `content <projectId> <recordId>` | FetchRecordContent | GET | /projects/{projectId}/collections/{collectionId}/records/{recordId}/content |
| `change <projectId> <recordId>` | ChangeRecordContent | PUT | /projects/{projectId}/collections/{collectionId}/records/{recordId}/content |
| `change-all <projectId>` | ChangeAllInSearchResult | PUT | /projects/{projectId}/collections/{collectionId}/records |
| `in-doc-search <projectId> <recordId>` | SearchInDocumentText | GET | /projects/{projectId}/collections/{collectionId}/records/{recordId}/inDocumentSearch |
| `token create <projectId>` | CreateSearchToken | GET | /projects/{projectId}/collections/{collectionId}/searchToken |
| `token touch <projectId>` | TouchSearchToken | PUT | /projects/{projectId}/collections/{collectionId}/searchToken |
| `token delete <projectId>` | DeleteSearchToken | DELETE | /projects/{projectId}/collections/{collectionId}/searchToken |
| `token snapshot <projectId>` | CreateSortOrderSnapshot | POST | /projects/{projectId}/collections/{collectionId}/searchToken/sortOrderSnapshot |
| `highlight <projectId>` | GetSearchHighlightExpressions | GET | /projects/{projectId}/collections/{collectionId}/search/highlightExpression |

### Binary

```
searchwebapi-example binary get <projectId> [-c collection] --field <fieldId> [-q query] [--language code] [--join-restriction expr] [--order expr] [--selected-index N] [--search-cache-control value]
searchwebapi-example binary get-by-id <projectId> <recordId> [-c collection] --field <fieldId>
```

| Subcommand | Operation | HTTP Method | Path |
|-----------|-----------|-------------|------|
| `get <projectId>` | GetBinaryBySearch | GET | /projects/{projectId}/collections/{collectionId}/binary |
| `get-by-id <projectId> <recordId>` | GetBinaryByRecordId | GET | /projects/{projectId}/collections/{collectionId}/binary/{recordId}/content |

### Measures

```
searchwebapi-example measures get <projectId> [-c collection] --body-file <path> [-q query] [--language code] [--join-restriction expr] [--measure-type json] [--search-cache-control value]
```

| Subcommand | Operation | HTTP Method | Path |
|-----------|-----------|-------------|------|
| `get <projectId>` | GetMeasureCube | POST | /projects/{projectId}/collections/{collectionId}/measures |

### Cached Searches

```
searchwebapi-example cached-searches list <projectId> [-c collection]
searchwebapi-example cached-searches drop <projectId> [-c collection] [--creation-trace-ids csv]
```

| Subcommand | Operation | HTTP Method | Path |
|-----------|-----------|-------------|------|
| `list <projectId>` | ListCachedSearches | GET | /projects/{projectId}/collections/{collectionId}/cachedSearches |
| `drop <projectId>` | DropCachedSearches | DELETE | /projects/{projectId}/collections/{collectionId}/cachedSearches |

### Change Queue

```
searchwebapi-example change-queue wait <projectId> [-c collection] [--timeout-millis N] [--high-priority]
```

| Subcommand | Operation | HTTP Method | Path |
|-----------|-----------|-------------|------|
| `wait <projectId>` | WaitForPendingChanges | GET | /projects/{projectId}/collections/{collectionId}/changes/queue |

### Insert Remove

```
searchwebapi-example insert-remove do <projectId> [-c collection] --body-file <path> [--search-cache-control value]
searchwebapi-example insert-remove bulk start <projectId> [-c collection] --body-file <path>
searchwebapi-example insert-remove bulk add <projectId> <bufferId> [-c collection] --body-file <path> [--search-cache-control value]
searchwebapi-example insert-remove bulk commit <projectId> <bufferId> [-c collection] --body-file <path>
searchwebapi-example insert-remove bulk status <projectId> <bufferId> <jobId> [-c collection]
```

| Subcommand | Operation | HTTP Method | Path |
|-----------|-----------|-------------|------|
| `do <projectId>` | InsertRemoveTransaction | POST | /projects/{projectId}/collections/{collectionId}/records/insertRemoveTransaction |
| `bulk start <projectId>` | StartBulkInsertRemoveTransaction | POST | /projects/{projectId}/collections/{collectionId}/records/bulkInsertRemoveTransaction |
| `bulk add <projectId> <bufferId>` | AddToBulkInsertRemoveTransaction | POST | /projects/{projectId}/collections/{collectionId}/records/bulkInsertRemoveTransaction/{indexingBufferId}/buffer |
| `bulk commit <projectId> <bufferId>` | CommitBulkInsertRemoveTransaction | POST | /projects/{projectId}/collections/{collectionId}/records/bulkInsertRemoveTransaction/{indexingBufferId}/end |
| `bulk status <projectId> <bufferId> <jobId>` | GetFlushJobStatus | GET | /projects/{projectId}/collections/{collectionId}/records/bulkInsertRemoveTransaction/{indexingBufferId}/end/{jobId} |

### Session

```
searchwebapi-example session login
searchwebapi-example session logout
```

| Subcommand | Operation | HTTP Method | Path |
|-----------|-----------|-------------|------|
| `login` | Login | POST | /login |
| `logout` | Logout | DELETE | /logout |

## Global Flags

| Flag | Alias | Env Var | Description |
|------|-------|---------|-------------|
| `--base-url` | `-b` | SEARCHWEBAPI_BASE_URL | API base URL |
| `--username` | `-u` | SEARCHWEBAPI_USERNAME | Username for auth |
| `--password` | `-p` | SEARCHWEBAPI_PASSWORD | Password for auth |
| `--bearer-token` | `-t` | SEARCHWEBAPI_BEARER_TOKEN | Bearer token for auth |
| `--session-id` | | SEARCHWEBAPI_SESSION_ID | Existing `SWA-SESSION` value |
| `--session-type` | | SEARCHWEBAPI_SESSION_TYPE | Optional `SWA-SESSION-TYPE` value |
| `--mdc-token` | | SEARCHWEBAPI_MDC_TOKEN | Optional `SWA-MDC-TOKEN` value |
| `--mdc-method` | | SEARCHWEBAPI_MDC_METHOD | Optional `SWA-MDC-METHOD` value |
| `--insecure` | `-k` | - | Skip TLS verification (for self-signed certs) |

## Command Flags

### records search / records search-stream

| Flag | Alias | Default | Description |
|------|-------|---------|-------------|
| `--collection` | `-c` | default | Collection ID |
| `--query` | `-q` | * | Search query |
| `--language` | | | Search language |
| `--join-restriction` | | | Join restriction |
| `--order` | | | Sort order |
| `--fields` | | | Comma-separated field list |
| `--folder-fields` | | | Comma-separated folder field list |
| `--folder-fields-with-properties` | | | Comma-separated folder field list with properties |
| `--body` | | false | Include body content |
| `--highlight` | | true | Enable highlighting |
| `--page` | | 1 | Page number |
| `--limit` | | 20 | Result limit |
| `--sponsored-links` | | false | Include sponsored links |
| `--spelling-suggestions` | | false | Include spelling suggestions |
| `--search-cache-control` | | | Value for `SWA-searchCacheControl` on `records search` |

### records content

| Flag | Alias | Default | Description |
|------|-------|---------|-------------|
| `--collection` | `-c` | default | Collection ID |
| `--body` | | false | Include body content |
| `--page` | | 1 | Page number |
| `--fields` | | | Comma-separated field list |
| `--folder-fields` | | | Comma-separated folder field list |
| `--folder-fields-with-properties` | | | Comma-separated folder field list with properties |
| `--highlight-search-term-query` | | | Highlight search term query |
| `--highlight-search-term-language` | | | Highlight search term language |
| `--highlight-search-term-join-restriction` | | | Highlight join restriction |
| `--fields-highlighted` | | | Comma-separated highlighted field list |
| `--highlight-hit-navigation` | | | Highlight hit navigation mode |
| `--highlight-user-terms` | | | Comma-separated user terms to highlight |
| `--highlight-folder-field-list` | | | Folder fields used for highlighting |
| `--summarize` | | false | Request summarization |
| `--search-cache-control` | | | Value for `SWA-searchCacheControl` |

### records change / records change-all

| Flag | Alias | Default | Description |
|------|-------|---------|-------------|
| `--collection` | `-c` | default | Collection ID |
| `--body-file` | | | JSON file containing `ChangeRequest[]` |
| `--block-until-complete` | | false | Wait for indexing changes to complete on `records change` |
| `--query` | `-q` | * | Search query for `change-all` |
| `--language` | | | Search language for `change-all` |
| `--body-file` | | | JSON file containing `ChangeRequest[]` for `change-all` |
| `--block-until-complete` | | false | Wait for indexing changes to complete on `change-all` |
| `--search-cache-control` | | | Value for `SWA-searchCacheControl` on `change-all` |

### records token create

| Flag | Alias | Default | Description |
|------|-------|---------|-------------|
| `--collection` | `-c` | default | Collection ID |
| `--query` | `-q` | * | Search query |
| `--language` | | | Search language |
| `--join-restriction` | | | Join restriction |
| `--order` | | | Sort order |

### collections values

| Flag | Alias | Default | Description |
|------|-------|---------|-------------|
| `--query` | `-q` | * | Search query |
| `--language` | | | Search language |
| `--join-restriction` | | | Join restriction |
| `--prefix` | | | Filter by name prefix |
| `--restrict-folders-by-query` | | | Restrict folders by query |
| `--return-empty-folders` | | false | Include empty folders |
| `--limit` | | 20 | Max results |
| `--offset` | | 0 | Result offset |
| `--order` | | | Sort order (`count`, `relevance`, `name`, `name:asc`, `name:desc`) |
| `--search-cache-control` | | | Value for `SWA-searchCacheControl` |

### change-queue wait

| Flag | Alias | Default | Description |
|------|-------|---------|-------------|
| `--collection` | `-c` | default | Collection ID |
| `--timeout-millis` | | 60000 | Timeout in milliseconds |
| `--high-priority` | | false | Only wait for high priority changes |

### records in-doc-search

| Flag | Alias | Default | Description |
|------|-------|---------|-------------|
| `--collection` | `-c` | default | Collection ID |
| `--highlight-search-term-query` | | | Highlight search term query |
| `--highlight-search-term-language` | | | Highlight search term language |
| `--highlight-search-term-join-restriction` | | | Highlight join restriction |
| `--highlight-user-terms` | | | Comma-separated user terms to highlight |
| `--highlight-folder-field-list` | | | Folder fields used for highlighting |
| `--highlight-folder-fields-aggregation` | | | Folder field aggregation for highlighting |
| `--content-field-names` | | | Content field names |
| `--page-tag` | | | Page tag |
| `--omit-hits-per-page` | | false | Omit hits per page |
| `--request-hit-locations-page-relative` | | false | Request page-relative hit locations |
| `--request-hit-locations-document-relative` | | false | Request document-relative hit locations |
| `--search-cache-control` | | enabled | Value for `SWA-searchCacheControl` |

### records highlight

| Flag | Alias | Default | Description |
|------|-------|---------|-------------|
| `--collection` | `-c` | default | Collection ID |
| `--query` | `-q` | * | Search query |
| `--language` | | | Search language |
| `--join-restriction` | | | Join restriction |
| `--search-cache-control` | | enabled | Value for `SWA-searchCacheControl` |

### binary get / binary get-by-id

| Flag | Alias | Default | Description |
|------|-------|---------|-------------|
| `--collection` | `-c` | default | Collection ID |
| `--field` | | | Binary field name |
| `--query` | `-q` | * | Search query for `binary get` |
| `--language` | | | Search language for `binary get` |
| `--join-restriction` | | | Join restriction for `binary get` |
| `--order` | | | Sort order for `binary get` |
| `--selected-index` | | 1 | Search hit index for `binary get` |
| `--search-cache-control` | | enabled | Value for `SWA-searchCacheControl` on `binary get` |

### measures get

| Flag | Alias | Default | Description |
|------|-------|---------|-------------|
| `--collection` | `-c` | default | Collection ID |
| `--body-file` | | | JSON file containing `DimensionRequest[]` |
| `--query` | `-q` | * | Search query |
| `--language` | | | Search language |
| `--join-restriction` | | | Join restriction |
| `--measure-type` | | | JSON value for `measureType` |
| `--search-cache-control` | | enabled | Value for `SWA-searchCacheControl` |

### cached-searches drop

| Flag | Alias | Default | Description |
|------|-------|---------|-------------|
| `--collection` | `-c` | default | Collection ID |
| `--creation-trace-ids` | | | Comma-separated creation trace ids |

### insert-remove bulk add / commit / status

| Flag | Alias | Default | Description |
|------|-------|---------|-------------|
| `--collection` | `-c` | default | Collection ID |
| `--search-cache-control` | | enabled | Value for `SWA-searchCacheControl` on `bulk add` |

### body-bearing commands

| Command | Input |
|---------|-------|
| `records change` | `--body-file <path>` containing JSON `ChangeRequest[]` |
| `records change-all` | `--body-file <path>` containing JSON `ChangeRequest[]` |
| `records token touch` | `--body-file <path>` containing JSON `SearchResultToken` |
| `records token delete` | `--body-file <path>` containing JSON `SearchResultToken` |
| `records token snapshot` | `--body-file <path>` containing JSON `SearchResultToken` |
| `measures get` | `--body-file <path>` containing JSON `DimensionRequest[]` |
| `insert-remove do` | `--body-file <path>` containing JSON `InsertRemoveRequest` |
| `insert-remove bulk start` | `--body-file <path>` containing JSON `StartTransactionRequest` |
| `insert-remove bulk add` | `--body-file <path>` containing JSON `InsertRemoveRequest` |
| `insert-remove bulk commit` | `--body-file <path>` containing JSON `FinishTransactionRequest` |

## Usage Examples

```bash
# List all projects
searchwebapi-example projects list

# Get project resources
searchwebapi-example projects get myproject

# List collections in a project
searchwebapi-example collections list myproject

# Search records (JSON)
searchwebapi-example records search myproject -q "status:active"

# Search records (NDJSON streaming)
searchwebapi-example records search-stream myproject -q "status:active"

# Get record content
searchwebapi-example records content myproject record123 --body --page 1 --fields title,author

# Change a record
searchwebapi-example records change myproject record123 -c mycollection

# Get binary by search result
searchwebapi-example binary get myproject -c mycollection --field native -q "id:record123" --selected-index 1

# Get binary by record ID
searchwebapi-example binary get-by-id myproject record123 -c mycollection --field native

# Get measure cube
searchwebapi-example measures get myproject -c mycollection --body-file dims.json

# List cached searches
searchwebapi-example cached-searches list myproject -c mycollection

# Drop cached searches
searchwebapi-example cached-searches drop myproject -c mycollection

# Wait for pending changes
searchwebapi-example change-queue wait myproject -c mycollection --timeout-millis 60000

# Session login
searchwebapi-example session login

# Session logout
searchwebapi-example session logout
```

## Implementation Notes

- Uses `github.com/urfave/cli/v3` for CLI framework
- Global flags are inherited by all subcommands
- JSON output is printed to stdout
- `search-stream` prints one compact JSON object per line: first the search metadata object, then one line per record
- Errors are printed to stderr with exit code 1
- TLS verification can be skipped with `--insecure` for testing against servers with self-signed certificates
