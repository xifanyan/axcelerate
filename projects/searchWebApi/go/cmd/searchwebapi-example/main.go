package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	searchwebapi "searchwebapi"
)

func main() {
	baseURL := flag.String("base-url", envOr("SEARCHWEBAPI_BASE_URL", ""), "Base URL")
	username := flag.String("username", envOr("SEARCHWEBAPI_USERNAME", ""), "Username")
	password := flag.String("password", envOr("SEARCHWEBAPI_PASSWORD", ""), "Password")
	bearerToken := flag.String("bearer-token", envOr("SEARCHWEBAPI_BEARER_TOKEN", ""), "Bearer token")
	project := flag.String("project", "", "Project ID")
	collection := flag.String("collection", "default", "Collection ID")
	query := flag.String("query", "*", "Search query")
	command := flag.String("command", "projects-list", "projects-list or records-search")
	flag.Parse()

	client, err := searchwebapi.NewClient(searchwebapi.Config{
		BaseURL:     *baseURL,
		Username:    *username,
		Password:    *password,
		BearerToken: *bearerToken,
	})
	if err != nil {
		fatal(err)
	}

	ctx := context.Background()
	switch *command {
	case "projects-list":
		result, err := client.ListProjects(ctx)
		fatalIf(err)
		printJSON(result)
	case "records-search":
		if *project == "" {
			fatal(fmt.Errorf("--project is required for records-search"))
		}
		result, err := client.SearchRecords(ctx, *project, *collection, searchwebapi.SearchRecordsOptions{Query: *query})
		fatalIf(err)
		printJSON(result)
	default:
		fatal(fmt.Errorf("unknown command %q", *command))
	}
}

func envOr(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func printJSON(v any) {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fatal(err)
	}
	_, _ = fmt.Fprintln(os.Stdout, string(data))
}

func fatalIf(err error) {
	if err != nil {
		fatal(err)
	}
}

func fatal(err error) {
	_, _ = fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
