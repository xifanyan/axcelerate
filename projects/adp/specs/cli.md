# CLI Interface

Describes the command-line interface contract. Language-agnostic.

---

## Global Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--host` | string | — | ADP server host (required) |
| `--port` | integer | 8443 | ADP server port |
| `--user` | string | — | Username (required) |
| `--password` | string | — | Password (required) |
| `--insecure` | boolean | false | Skip TLS certificate verification |
| `--debug` | boolean | false | Enable debug logging — traces request/response payloads |

---

## CLI Naming Convention

| Language | Binary Name |
|----------|--------------|
| Go | `adpgo` |
| Python | `adppy` |
| Rust | `adprs` |

---

## Subcommands

Each task has a subcommand. Subcommand names use `kebab-case` (e.g., `list-entities`, not `listEntities`).

| Task | Subcommand |
|------|-----------|
| List Entities | `list-entities` |
| Query Engine | `query-engine` |
| Taxonomy Statistic | `taxonomy-statistic` |
| Start Application | `start-application` |
| Export Documents | `export-documents` |
| CSV Merge | `csv-merge` |

---

## CLI Argument Naming

- All CLI arguments must use **camelCase** (e.g., `--relatedEntity`, `--whiteList`)
- This applies to all language implementations

---

## Output Rules

### On Success (HTTP 200 + executionStatus == "success")

Output only the parsed task-specific data:
- `list-entities`: JSON array of entities
- `query-engine`: documents count and aggregated value
- `taxonomy-statistic`: decoded statistics JSON
- `start-application`: application URL
- `export-documents`: export file name, path, and count

The raw `executionMetaData` is not output on success.

### On Failure (HTTP 200 + executionStatus == "failed")

Output error details:
```
Error: <errorMessage>
ExecutionID: <executionId>
TaskType: <taskType>
```

### Debug Mode (`--debug`)

When debug is enabled, additionally output:
- Full request payload (JSON)
- Full response payload (JSON)

Debug is off by default. Enable with `--debug` or `-d`.

---

## CLI Examples

```bash
# List entities
adpgo --host example.com --user adp --password adp list-entities --type singleMindServer

# Query engine with taxonomy filter
adpgo --host example.com --user adp --password adp query-engine --engineName "myEngine" --engineTaxonomies "rm_mimetype=pdf"

# Taxonomy statistic
adpgo --host example.com --user adp --password adp taxonomy-statistic --engineName "myEngine" --outputTaxonomies "rm_source,meta_documentcharacteristics"

# Start application
adpgo --host example.com --user adp --password adp start-application --applicationIdentifier "my-app-id" --useHttps

# With debug
adpgo --debug --host example.com --user adp --password adp list-entities --type singleMindServer
```

---

## CLI Flag Mappings

See individual task specs in [tasks/](./tasks/) for the complete list of CLI flags per task.

---

## Shared Input Type CLI Formats

### EngineTaxonomyArg

Shorthand format: `Taxonomy=Query` or `Taxonomy!=Query`

| Format | Description | Example |
|--------|-------------|---------|
| `Taxonomy=Query` | Equals (negation=false) | `rm_mimetype=pdf` |
| `Taxonomy!=Query` | Not equals (negation=true) | `rm_source!=email` |

Repeatable flag. Example:
```bash
--engineTaxonomies "rm_source=email" --engineTaxonomies "rm_mimetype=pdf"
```

### OutputTaxonomiesArg

Comma-separated taxonomy names or JSON array:
```bash
--outputTaxonomies "rm_source,meta_documentcharacteristics"
```
