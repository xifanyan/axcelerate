# searchwebapi

Generated Go client for the `searchWebApi` contract.

## Build

```bash
go test ./...
go build ./...
```

## Library Usage

Create a client:

```go
client, err := searchwebapi.NewClient(searchwebapi.Config{
	BaseURL:  "https://example.test",
	Username: "user",
	Password: "pass",
})
if err != nil {
	log.Fatal(err)
}
```

List projects:

```go
projects, err := client.ListProjects(context.Background())
if err != nil {
	log.Fatal(err)
}

for _, p := range projects.Results {
	log.Println("project:", p.ID)
}
```

Search records:

```go
result, err := client.SearchRecords(
	context.Background(),
	"my-project",
	"default",
	searchwebapi.SearchRecordsOptions{Query: "hello world"},
)
if err != nil {
	log.Fatal(err)
}

for _, record := range result.Results {
	log.Println("record:", record.ID)
}
```

Stream NDJSON results:

```go
stream, err := client.SearchRecordsStream(
	context.Background(),
	"my-project",
	"default",
	searchwebapi.SearchRecordsOptions{Query: "hello world"},
)
if err != nil {
	log.Fatal(err)
}
defer stream.Close()

for {
	record, err := stream.Next()
	if err != nil {
		break
	}
	log.Println("streamed record:", record.ID)
}
```

## Example CLI

```bash
go run ./cmd/searchwebapi-example --base-url https://example.test --command projects-list
go run ./cmd/searchwebapi-example --base-url https://example.test --command records-search --project my-project --collection default --query hello
```
