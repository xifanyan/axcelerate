# Tasks Index

## Available Tasks

| #   | Task                                        | Description                                  | Key Fields                                              |
| --- | ------------------------------------------- | -------------------------------------------- | ------------------------------------------------------- |
| 1   | [list-entities.md](./list-entities.md)      | Writes a list of entities to an output var   | workspace, type, status, whiteList, id, relatedEntity |
| 2   | [query-engine.md](./query-engine.md)        | Queries an engine                             | engineName, engineQuery, category, applicationIdentifier |
| 3   | [start-application.md](./start-application.md) | Starts an application                       | applicationIdentifier, useHttps                        |
| 4   | [taxonomy-statistic.md](./taxonomy-statistic.md) | Retrieves category counts for a taxonomy  | engineName, engineQuery, computeCounts                  |
| 5   | [csv-merge.md](./csv-merge.md)              | Merges content/images via CSV file            | csvFile, mergeType, engineName, csvIdFieldKey         |

---

## Adding New Tasks

1. Copy [TEMPLATE.md](./TEMPLATE.md)
2. Fill using [API-SPEC.md](../../API-SPEC.md)
3. Add entry to the [## Available Tasks](#available-tasks) table above — **this table must always reflect all current task specs**
4. Do NOT generate code - only update specs

> **Rule:** The "Available Tasks" table above is the authoritative list of all tasks. Every task spec must have an entry here. Keep it synchronized — an entry must be added, updated, or removed whenever a task spec changes.
