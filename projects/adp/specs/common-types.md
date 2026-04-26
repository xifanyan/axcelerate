# Common Types

Shared input types used across multiple tasks.

For API contract types (envelopes, shared response types), see [api-contract.md](./api-contract.md).

---

## Task-Specific Shared Types

### EngineTaxonomyArg

Used by: Query Engine, Taxonomy Statistic, Create OCR Job

**Definition and CLI shorthand format:** See [api-contract.md](./api-contract.md#enginetaxonomyarg)

### OutputTaxonomiesArg

Used by: Taxonomy Statistic

**Definition and CLI format:** See [api-contract.md](./api-contract.md#outputtaxonomiesarg)

### FieldMappingArg

Used by: CSV Merge

```json
{
  "csvFieldName": "Column Header",
  "textType": "internal_field_name",
  "valueDelimiter": "|",
  "useDisplayName": "true"
}
```

| Field | Type | Description |
|-------|------|-------------|
| csvFieldName | string | CSV column header name |
| textType | string | Internal field name |
| valueDelimiter | string | Value delimiter |
| useDisplayName | string | Use display name ("true"/"false") |

#### CLI Format

JSON array format:
```bash
--fieldMappings '[{"csvFieldName":"Column A","textType":"field_a","valueDelimiter":"|","useDisplayName":"true"}]'
```