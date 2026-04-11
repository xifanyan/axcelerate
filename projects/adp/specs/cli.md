# CLI Interface

Describes the command-line interface contract. Language-agnostic.

---

## Global Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--host` | string | — | ADP server host (required) |
| `--port` | integer | 8443 | ADP server port |
| `--path` | string | `/adp/rest/api/task` | ADP task endpoint path |
| `--user` | string | — | Username (required) |
| `--password` | string | — | Password (required) |
| `--insecure` | boolean | false | Enables insecure mode; must support `--insecure=true|false` |
| `--debug` | boolean | false | Enables debug logging; must support `--debug=true|false` |
| `-d` | boolean | false | Shorthand for enabling debug logging |

Global flags may be provided by explicit CLI flags or `adp_config.json`.

## CLI Configuration File

All language CLIs must support `adp_config.json`.

`adp_config.json` is read from the current working directory.

Supported config keys:
- `host`
- `port`
- `path`
- `user`
- `password`
- `insecure`
- `debug`

Resolution precedence for all global flags:
1. Explicit command-line flag
2. `adp_config.json`
3. Built-in default

Boolean global flags must support both enable and explicit override forms: `--debug` and `-d` enable debug, `--debug=true|false` explicitly sets debug, `--insecure` enables insecure mode, and `--insecure=true|false` explicitly sets insecure mode.

If the same explicit global flag is provided multiple times, the last explicit value wins.

Built-in defaults:
- `port=8443`
- `path=/adp/rest/api/task`
- `insecure=false`
- `debug=false`

`host`, `user`, and `password` are required after resolution.

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
