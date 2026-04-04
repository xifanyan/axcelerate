# Common Types Specification

> **Note**: This file contains shared types referenced across multiple specs. Core API types are now documented in [api-contract.md](./api-contract.md). Shared input types used by multiple tasks are documented below.

---

## Shared Input Types

These types are used across multiple tasks.

### EngineTaxonomyArg

Used by: Query Engine, Taxonomy Statistic, Create OCR Job

```json
{
  "taxonomy": "rm_source",
  "negation": false,
  "query": "email"
}
```

| Field | Type | Description |
|-------|------|-------------|
| taxonomy | string | Taxonomy name (e.g., "rm_source", "meta_documentcharacteristics") |
| negation | boolean | Negation flag. `false` = equals, `true` = not equals |
| query | string | Query string (URL-encoded) |

#### CLI Shorthand Format

For CLI, use shorthand format instead of JSON:

| Format | Description | Example |
|--------|-------------|---------|
| `Taxonomy=Query` | Equals (negation=false) | `rm_mimetype=pdf` |
| `Taxonomy!=Query` | Not equals (negation=true) | `rm_source!=email` |

Multiple taxonomies: repeat the flag.

**CLI Examples:**
```bash
--engineTaxonomies "rm_mimetype=pdf"
--engineTaxonomies "rm_source=email" --engineTaxonomies "rm_mimetype=pdf"
--engineTaxonomies "rm_source!=email"
```

### OutputTaxonomiesArg

Used by: Taxonomy Statistic

```json
{
  "taxonomy": "rm_source",
  "mode": "Category counts",
  "maximumNumberOfCategories": 10
}
```

| Field | Type | Description |
|-------|------|-------------|
| taxonomy | string | Taxonomy name |
| mode | string | "Aggregate counts" or "Category counts" |
| maximumNumberOfCategories | integer | Maximum number of categories to return |

#### CLI Format

Comma-separated taxonomy names or JSON array:
```bash
--outputTaxonomies "rm_source,meta_documentcharacteristics"
```
