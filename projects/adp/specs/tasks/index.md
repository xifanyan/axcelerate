# Tasks Index

Quick nav: [Discovery](#discovery) · [Processing](#processing) · [Management](#management) · [CLI](#cli)

---

## Discovery

| # | Task | Subcommand | Description |
|---|------|-----------|-------------|
| 1 | [list-entities.md](./list-entities.md) | `list-entities` | Writes a list of entities to an output variable |
| 2 | [query-engine.md](./query-engine.md) | `query-engine` | Queries an engine |

---

## Processing

| # | Task | Subcommand | Description |
|---|------|-----------|-------------|
| 3 | [start-application.md](./start-application.md) | `start-application` | Starts an application |
| 4 | [taxonomy-statistic.md](./taxonomy-statistic.md) | `taxonomy-statistic` | Retrieves category counts for a taxonomy |
| 5 | [csv-merge.md](./csv-merge.md) | `csv-merge` | Merges content/images via CSV file |
| 6 | [export-documents.md](./export-documents.md) | `export-documents` | Export documents in CSV format |
| 7 | [read-configuration.md](./read-configuration.md) | `read-configuration` | Read configurations into JSON or XML |
| 8 | [create-ocr-job.md](./create-ocr-job.md) | `create-ocr-job` | Changes metaData by using regEx replacement |

---

## Management

| # | Task | Subcommand | Description |
|---|------|-----------|-------------|
| 9 | [create-data-source.md](./create-data-source.md) | `create-data-source` | Creates a new data source |
| 10 | [matter-management.md](./matter-management.md) | `matter-management` | Task to manage matters and saved searches |
| 11 | [create-review-interface.md](./create-review-interface.md) | `create-review-interface` | Task to create a review interface |
| 12 | [publish-to-review.md](./publish-to-review.md) | `publish-to-review` | Task to publish stuff |

---

## CLI

| # | Task | Subcommand | Description |
|---|------|-----------|-------------|
| 13 | [cli.md](./cli.md) | `cli` | Runs a native task in its own process |

---

## Adding New Tasks

1. Copy [TEMPLATE.md](./TEMPLATE.md)
2. Fill using [API-SPEC.md](../../API-SPEC.md)
3. Add entry to the appropriate section above — **group by category**
4. Do NOT generate code — only update specs

> **Rule:** Each section is the authoritative list of its tasks. Keep sections synchronized — an entry must be added, updated, or removed whenever a task spec changes.
