# searchwebapi

Generated Go client for the `searchWebApi` contract.

## Build

```bash
go test ./...
go build ./...
```

## Example CLI

```bash
go run ./cmd/searchwebapi-example --base-url https://example.test --command projects-list
go run ./cmd/searchwebapi-example --base-url https://example.test --command records-search --project my-project --collection default --query hello
```
