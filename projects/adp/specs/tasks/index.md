# Tasks Index

## Available Tasks

| # | Task | Subcommand | Description | executionMetaData |
|---|------|-----------|-------------|-------------------|
| 1 | [list-entities.md](./list-entities.md) | `list-entities` | Writes a list of entities to an output variable | outputFile, entities[] |
| 2 | [query-engine.md](./query-engine.md) | `query-engine` | Queries an engine | documentsCount, aggregatedValue |
| 3 | [start-application.md](./start-application.md) | `start-application` | Starts an application | applicationUrl |
| 4 | [taxonomy-statistic.md](./taxonomy-statistic.md) | `taxonomy-statistic` | Retrieves category counts for a taxonomy | outputFile, statistics |
| 5 | [csv-merge.md](./csv-merge.md) | `csv-merge` | Merges content/images via CSV file | empty array (no result fields) |
| 6 | [export-documents.md](./export-documents.md) | `export-documents` | Export documents in CSV format | exportFileName, exportPath, searchResultSize |
| 7 | [read-configuration.md](./read-configuration.md) | `read-configuration` | Read configurations into JSON or XML | outputFile, configuration |
| 8 | [create-ocr-job.md](./create-ocr-job.md) | `create-ocr-job` | Changes metaData by using regEx replacement | async-only, empty array |
| 9 | [cli.md](./cli.md) | `cli` | Runs a native task in its own process | cli_result, json_output, cli_error_path, cli_result_path |
| 10 | [create-data-source.md](./create-data-source.md) | `create-data-source` | Creates a new data source | adp_created_data_source_displayname, adp_hostname, adp_chosen_host_cpu_load, adp_created_data_source_name, adp_chosen_host_memory, adp_chosen_host_memory_ratio, adp_chosen_engine, adp_used_data_source_template |
| 11 | [matter-management.md](./matter-management.md) | `matter-management` | Task to manage matters and saved searches | processedMatterId, matterProcessingRequestId, matterProcessingResponseMessage, processedSavedSearchId, savedSearchProcessingRequestId, savedSearchProcessingResponseMessage, usedWebserviceUrl |
| 12 | [create-review-interface.md](./create-review-interface.md) | `create-review-interface` | Task to create a review interface | publishEngineId, publishApplicationId, applicationHost, applicationHostMemory, applicationHostMemoryRatio, engineHost, engineHostMemory, engineHostMemoryRatio, createReviewRequestId, createReviewResponseMessage, usedWebserviceUrl |
| 13 | [publish-to-review.md](./publish-to-review.md) | `publish-to-review` | Task to publish stuff | usedWebserviceUrl, publishApplicationId, publishApplicationUrl, publishEngineId, publishResponseMessage, publishRequestId |

---

## Adding New Tasks

1. Copy [TEMPLATE.md](./TEMPLATE.md)
2. Fill using [API-SPEC.md](../../API-SPEC.md)
3. Add entry to the [## Available Tasks](#available-tasks) table above — **this table must always reflect all current task specs**
4. Do NOT generate code — only update specs

> **Rule:** The "Available Tasks" table above is the authoritative list of all tasks. Every task spec must have an entry here. Keep it synchronized — an entry must be added, updated, or removed whenever a task spec changes.
