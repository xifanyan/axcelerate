# Language-Specific Rules

> **Deprecated**: This file is retained for reference only. Language-specific conventions are now documented in [language-bindings/go.md](./language-bindings/go.md), [language-bindings/rust.md](./language-bindings/rust.md), and [language-bindings/python.md](./language-bindings/python.md).
>
> The authoritative language-agnostic contracts are defined in:
> - [api-contract.md](./api-contract.md)
> - [request-construction.md](./request-construction.md)
> - [result-decoding.md](./result-decoding.md)
> - [cli.md](./cli.md)

---

## Historical Content

The content below was previously in this file and is now superseded.

### CLI Framework

| Language | CLI Library | Rationale |
|----------|-------------|-----------|
| Go | [urfave/cli](https://github.com/urfave/cli) | Most popular Go CLI library, well-maintained |
| Python | [click](https://click.palletsprojects.com/) | Most popular Python CLI library, decorator-based |

### CLI Argument Naming

- **All CLI arguments must use camelCase** (e.g., `--relatedEntity`, `--whiteList`)
- This applies to both Go and Python implementations

### Project Structure (Go - Historical)

```
src/go/
├── cmd/adpgo/
│   └── main.go       # CLI entrypoint (app definition)
├── client.go         # HTTP client
├── types.go          # Shared types
├── list_entities.go # Task: List Entities
├── query_engine.go  # Task: Query Engine
└── taxonomy_statistic.go  # Task: Taxonomy Statistic
```

### Project Structure (Python - Historical)

```
src/python/
├── adppy/
│   ├── __init__.py
│   ├── client.py      # HTTP client
│   ├── types.py       # Shared types
│   ├── list_entities.py  # Task: List Entities
│   ├── query_engine.py  # Task: Query Engine
│   └── taxonomy_statistic.py  # Task: Taxonomy Statistic
├── pyproject.toml
└── adppy/__main__.py  # CLI entrypoint
```
