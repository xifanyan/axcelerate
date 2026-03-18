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
├── service.go        # Service layer (optional fluent API)
├── types.go          # Shared types
├── list_entities.go # Task: List Entities
├── query_engine.go  # Task: Query Engine
└── taxonomy_statistic.go  # Task: Taxonomy Statistic
```

### Service Layer (Optional)

```go
svc := adp.NewService(client)
result, err := svc.ListEntities().Type("server").Execute(ctx)
```

#### Rules

1. **Service wraps client** - Service holds a reference to the HTTP client
2. **One method per task** - Each task has a service method returning its builder
3. **Builder retains fluent setters** - Use existing builder pattern
4. **Execute method** - Builder must implement `Execute(context)` that calls client
5. **Optional Async** - Builders may also have `ExecuteAsync(context)`

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
│   ├── service.py     # Service layer (optional fluent API)
│   ├── types.py       # Shared types
│   ├── list_entities.py  # Task: List Entities
│   ├── query_engine.py  # Task: Query Engine
│   └── taxonomy_statistic.py  # Task: Taxonomy Statistic
├── pyproject.toml
└── adppy/__main__.py  # CLI entrypoint
```

### Service Layer (Optional)

```python
svc = adp.Service(client)
result = svc.list_entities().type_("server").execute()
```
