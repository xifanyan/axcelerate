# ADP Project Specifications

Single source of truth for all ADP specifications. This document describes the overall structure and global rules. Language-agnostic by design.

---

## Spec Structure

| File | Description |
|------|-------------|
| [index.md](./index.md) | This file — global rules and spec map |
| [api-contract.md](./api-contract.md) | Raw API transport, endpoints, request/response envelopes |
| [request-construction.md](./request-construction.md) | How clients build sparse requests |
| [result-decoding.md](./result-decoding.md) | How clients decode task outputs (sync and async) |
| [common-types.md](./common-types.md) | Shared input types referenced across multiple specs |
| [cli.md](./cli.md) | CLI interface, subcommands, flags, output rules |
| [tasks/index.md](./tasks/index.md) | Available tasks and add-task rules |
| [http-client.md](./http-client.md) | HTTP client transport details |
| [VERIFICATION.md](./VERIFICATION.md) | API response verification checklist |

### Language Bindings

| File | Description |
|------|-------------|
| [language-bindings/go.md](./language-bindings/go.md) | Go-specific API conventions |
| [language-bindings/rust.md](./language-bindings/rust.md) | Rust-specific API conventions |
| [language-bindings/python.md](./language-bindings/python.md) | Python-specific API conventions |
| [languages.md](./languages.md) | **Deprecated** — superseded by language-bindings/ |

Language bindings are non-authoritative — they map the contracts defined in core specs into idiomatic APIs for each language.

---

## Global Rules

### Source of Truth

`API-SPEC.md` is the authoritative source for upstream task default values and raw field names. All spec documentation must match `API-SPEC.md` exactly.

Markdown specs in `specs/` are the authoritative design documents. Do not maintain parallel output schemas or generator scripts under `scripts/` for the same contracts.

### Task Spec Rules

- Task specs must match [API-SPEC.md](../API-SPEC.md) exactly (field names, values, ordering)
- Example Request = Default Configuration (no custom values) — used only for documentation of the raw upstream shape
- **Task Configuration** — Only send fields that need to be changed from defaults; do not include all fields
- Use **progressive request-construction API** for task configuration (e.g., builder or options-based, language-idiomatic)
- **Result types must use language-agnostic notation** — TypeScript-like notation only:

| Use | Not |
|-----|-----|
| `Type[]` | `[]Type` |
| `Record<K, V>` | `map[K]V` |
| `string`, `integer`, `boolean` | `String`, `int`, `bool` |
| `any` | `interface{}` |
| `absent` | `null` |

### Field Casing

All API request and response field names use **camelCase** exclusively (e.g., `executionId`, `executionMetaData`). This applies to:
- Raw API request bodies
- Raw API responses
- Spec documentation of API fields

PascalCase or other casing must not appear in API-facing documentation.

### Logging

Logging must be **enabled by default** for all tasks. Debug mode must trace request/response payloads. Enable via CLI `--debug` or `-d` flag.

### Function Naming

- Default functions (e.g., `ListEntities`) are synchronous
- Add `Async` suffix for asynchronous variants (e.g., `ListEntitiesAsync`)

### CLI Interface

- Must support subcommands for each task
- Global flags: `--host`, `--port`, `--path`, `--user`, `--password`, `--insecure`, `--debug`, `-d` (`--debug` shorthand to enable)
- All language CLIs must support `adp_config.json` for global flag values
- `adp_config.json` is read from the current working directory
- Precedence for all global flags: explicit CLI flag > `adp_config.json` > built-in default
- `--port` default: 8443
- `--path` default: `/adp/rest/api/task`
- `--insecure` default: false
- `--debug` default: false
- Boolean overrides must support explicit CLI forms for config override behavior: `--insecure=true|false`, `--debug=true|false`
- `host`, `user`, and `password` are required after resolution
- Example: `adpgo --host example.com --user adp --password adp list-entities --type singleMindServer`
- CLI naming: `[project][lang]` (e.g., `adpgo` for Go, `adppy` for Python, `adprs` for Rust)

### Code Generation Conventions

When generating code for ADP client library:

1. **One task per file** — Each task in its own file (e.g., `list_entities.go`, `query_engine.go`)
2. **Shared types** — Common types go in `types.go` or the equivalent
3. **Client** — HTTP client implementation in `client.go`
4. **Config structs** — Place in the same file as the task builder function
5. **CLI** — CLI implementation in `cmd/adpgo/main.go`

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

See also:
- [api-contract.md](./api-contract.md)
- [request-construction.md](./request-construction.md)
- [result-decoding.md](./result-decoding.md)
- [cli.md](./cli.md)
- [tasks/index.md](./tasks/index.md)
