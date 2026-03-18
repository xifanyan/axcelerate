# Language-Specific Rules

## CLI Framework

| Language | CLI Library | Rationale |
|----------|-------------|-----------|
| Go | [urfave/cli](https://github.com/urfave/cli) | Most popular Go CLI library, well-maintained |
| Python | [click](https://click.palletsprojects.com/) | Most popular Python CLI library, decorator-based |

## CLI Argument Naming

- **All CLI arguments must use camelCase** (e.g., `--relatedEntity`, `--whiteList`)
- This applies to both Go and Python implementations

## Go Rules

### CLI Library: urfave/cli

- Use `urfave/cli/v3` for CLI implementation
- Define commands using `&cli.Command` struct
- Global flags via `app.Flags`
- Subcommands via `app.Commands`
- Use `action` function for command logic

### Project Structure

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

## Python Rules

### CLI Library: click

- Use `@click.group()`, `@click.command()`, `@click.option()` decorators
- Define commands as functions with decorators
- Global options via `@click.group()` with `@click.option()`
- Use `click.echo()` for output

### Project Structure

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
