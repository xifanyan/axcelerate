# CSV Merge Task Specification

## Overview

| Property | Value |
|----------|-------|
| Task Type | `CSV Merge` |
| Description | Merges content or updates natives/images by using a CSV file |
| Display Name | Csv merge task |
| Subcommand | `csv-merge` |

---

## Semantic Inputs

These are the user-facing fields for the request-construction API.

| Field | Type | Default | Required | Description |
|-------|------|---------|----------|-------------|
| csvFile | string | null | Yes | CSV file path |
| csvIdFieldKey | string | null | No | CSV ID field key |
| mergeType | string | "Merge content" | No | Merge type (`Merge content` or `Update natives/images`) |
| csvMode | string | "append" | No | CSV mode (append/merge) |
| engineName | string | null | Conditional | Engine name |
| engineUser | string | null | No | Engine user |
| enginePassword | string | null | No | Engine password |
| engineIdFieldKey | string | null | No | Engine ID field key |
| applicationIdentifier | string | "" | Conditional | Application identifier |
| fieldMappings | array | [] | No | Field mappings as column definitions |
| fieldSeparator | string | ";" | No | Field separator |
| imageBasePath | string | null | No | Image base path |
| nativeBasePath | string | null | No | Native base path |
| csvFieldImageLocation | string | null | No | CSV field image location |
| csvFieldNativeLocation | string | null | No | CSV field native location |
| multiValueDelimiter | string | null | No | Multi-value delimiter |
| textIndicator | string | "" | No | Text indicator |
| doNotChangeProtectedDocuments | boolean | false | No | Do not change protected documents |

> engineName and applicationIdentifier are mutually exclusive selectors. Exactly one must be provided.

---

## Raw Default Configuration

> Configuration below shows **all fields with their exact default values** from [API-SPEC.md](../../API-SPEC.md). This is for reference only. Clients must not pre-populate all fields. See [request-construction.md](../request-construction.md).
> These upstream defaults are shown as-is for reference. Real client-built requests must still provide exactly one of `engineName` or `applicationIdentifier`.

```json
{
  "taskType": "CSV Merge",
  "taskConfiguration": {
    "adp_csvMerge_noUniqueMatch": "false",
    "adp_csvMerge_noFlushAfterMerge": "false",
    "adp_csvMerge_forceChange": "false",
    "adp_csvMerge_engineName": null,
    "adp_csvMerge_applicationIdentifier": "",
    "adp_csvMerge_mergeType": "Merge content",
    "adp_csvMerge_displayNameMappingErrorFile": null,
    "adp_csvMerge_lockDocumentChanges": "false",
    "adp_csvMerge_csvFieldNativeLocation": null,
    "adp_executionPersistent": true,
    "adp_csvMerge_doNotChangeProtectedDocuments": "false",
    "adp_csvMerge_imageBasePath": null,
    "adp_loggingEnabled": true,
    "adp_csvMerge_nullValue": "",
    "adp_csvMerge_deduplicateWhenAppending": "true",
    "adp_csvMerge_lineSeparatorForFulltext": "",
    "adp_csvMerge_allowUpdateMultipleDocs": "false",
    "adp_csvMerge_csvIdPostfix": null,
    "adp_csvMerge_origFile": null,
    "adp_csvMerge_maxCategoryLength": "128",
    "adp_csvMerge_fieldSeperator": ";",
    "adp_csvMerge_applicationType": "",
    "adp_csvMerge_enginePassword": null,
    "adp_taskTimeout": 0,
    "adp_csvMerge_charset": "UTF-8",
    "adp_progressTaskTimeout": 0,
    "adp_csvMerge_textIndicator": "",
    "adp_csvMerge_noFlushBeforeMerge": "false",
    "adp_taskActive": true,
    "adp_csvMerge_engineIdFieldKey": null,
    "adp_abortWfOnFailure": true,
    "adp_csvMerge_customLineSeparator": "U+000DU+000A",
    "adp_csvMerge_multiValueDelimiter": null,
    "adp_csvMerge_csvFile": null,
    "adp_csvMerge_csvIdFieldKey": null,
    "adp_csvMerge_textFileRefIndicator": "",
    "adp_csvMerge_engineUser": null,
    "adp_csvMerge_textFileCharset": "",
    "adp_csvMerge_doNotAddTimestampToOutputFiles": "false",
    "adp_csvMerge_nativeBasePath": null,
    "adp_csvMerge_csvMode": "append",
    "adp_csvMerge_fieldMappings": [],
    "adp_cleanUpHistory": false,
    "adp_csvMerge_csvMergeConfiguration": "",
    "adp_csvMerge_errorLogFile": null,
    "adp_csvMerge_dryRun": "false",
    "adp_csvMerge_csvFieldImageLocation": null,
    "adp_csvMerge_matchesLogFile": null,
    "adp_csvMerge_csvIdPrefix": null,
    "adp_csvMerge_expandToFamily": "false"
  },
  "taskDescription": "Merges metaData or natives/images by using a csv file.",
  "taskDisplayName": "Csv merge task"
}
```

---

## Field Reference (with Defaults)

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| adp_csvMerge_noUniqueMatch | string | "false" | No unique match |
| adp_csvMerge_noFlushAfterMerge | string | "false" | No flush after merge |
| adp_csvMerge_forceChange | string | "false" | Force change |
| adp_csvMerge_engineName | string | null | Engine name |
| adp_csvMerge_applicationIdentifier | string | "" | Application identifier |
| adp_csvMerge_mergeType | string | "Merge content" | Merge type |
| adp_csvMerge_displayNameMappingErrorFile | string | null | Display name mapping error file |
| adp_csvMerge_lockDocumentChanges | string | "false" | Lock document changes |
| adp_csvMerge_csvFieldNativeLocation | string | null | CSV field native location |
| adp_executionPersistent | boolean | true | Persist execution |
| adp_csvMerge_doNotChangeProtectedDocuments | string | "false" | Do not change protected documents |
| adp_csvMerge_imageBasePath | string | null | Image base path |
| adp_loggingEnabled | boolean | true | Enable logging |
| adp_csvMerge_nullValue | string | "" | Null value |
| adp_csvMerge_deduplicateWhenAppending | string | "true" | Deduplicate when appending |
| adp_csvMerge_lineSeparatorForFulltext | string | "" | Line separator for fulltext |
| adp_csvMerge_allowUpdateMultipleDocs | string | "false" | Allow update multiple docs |
| adp_csvMerge_csvIdPostfix | string | null | CSV ID postfix |
| adp_csvMerge_origFile | string | null | Original file |
| adp_csvMerge_maxCategoryLength | string | "128" | Max category length |
| adp_csvMerge_fieldSeperator | string | ";" | Field separator |
| adp_csvMerge_applicationType | string | "" | Application type |
| adp_csvMerge_enginePassword | string | null | Engine password |
| adp_taskTimeout | integer | 0 | Task timeout |
| adp_csvMerge_charset | string | "UTF-8" | Charset |
| adp_progressTaskTimeout | integer | 0 | Progress task timeout |
| adp_csvMerge_textIndicator | string | "" | Text indicator |
| adp_csvMerge_noFlushBeforeMerge | string | "false" | No flush before merge |
| adp_taskActive | boolean | true | Whether task is active |
| adp_csvMerge_engineIdFieldKey | string | null | Engine ID field key |
| adp_abortWfOnFailure | boolean | true | Abort workflow on failure |
| adp_csvMerge_customLineSeparator | string | "U+000DU+000A" | Custom line separator |
| adp_csvMerge_multiValueDelimiter | string | null | Multi-value delimiter |
| adp_csvMerge_csvFile | string | null | CSV file path |
| adp_csvMerge_csvIdFieldKey | string | null | CSV ID field key |
| adp_csvMerge_textFileRefIndicator | string | "" | Text file reference indicator |
| adp_csvMerge_engineUser | string | null | Engine user |
| adp_csvMerge_textFileCharset | string | "" | Text file charset |
| adp_csvMerge_doNotAddTimestampToOutputFiles | string | "false" | Do not add timestamp to output files |
| adp_csvMerge_nativeBasePath | string | null | Native base path |
| adp_csvMerge_csvMode | string | "append" | CSV mode |
| adp_csvMerge_fieldMappings | array | [] | Field mappings |
| adp_cleanUpHistory | boolean | false | Clean up history |
| adp_csvMerge_csvMergeConfiguration | string | "" | CSV merge configuration |
| adp_csvMerge_errorLogFile | string | null | Error log file |
| adp_csvMerge_dryRun | string | "false" | Dry run |
| adp_csvMerge_csvFieldImageLocation | string | null | CSV field image location |
| adp_csvMerge_matchesLogFile | string | null | Matches log file |
| adp_csvMerge_csvIdPrefix | string | null | CSV ID prefix |
| adp_csvMerge_expandToFamily | string | "false" | Expand to family |

---

## Raw Example Request

> Example below matches **exactly** the default configuration from [API-SPEC.md](../../API-SPEC.md).
> This raw reference example mirrors upstream defaults. Real client-built requests must still provide exactly one of `engineName` or `applicationIdentifier`.

```json
{
  "taskType": "CSV Merge",
  "taskConfiguration": {
    "adp_csvMerge_noUniqueMatch": "false",
    "adp_csvMerge_noFlushAfterMerge": "false",
    "adp_csvMerge_forceChange": "false",
    "adp_csvMerge_engineName": null,
    "adp_csvMerge_applicationIdentifier": "",
    "adp_csvMerge_mergeType": "Merge content",
    "adp_csvMerge_displayNameMappingErrorFile": null,
    "adp_csvMerge_lockDocumentChanges": "false",
    "adp_csvMerge_csvFieldNativeLocation": null,
    "adp_executionPersistent": true,
    "adp_csvMerge_doNotChangeProtectedDocuments": "false",
    "adp_csvMerge_imageBasePath": null,
    "adp_loggingEnabled": true,
    "adp_csvMerge_nullValue": "",
    "adp_csvMerge_deduplicateWhenAppending": "true",
    "adp_csvMerge_lineSeparatorForFulltext": "",
    "adp_csvMerge_allowUpdateMultipleDocs": "false",
    "adp_csvMerge_csvIdPostfix": null,
    "adp_csvMerge_origFile": null,
    "adp_csvMerge_maxCategoryLength": "128",
    "adp_csvMerge_fieldSeperator": ";",
    "adp_csvMerge_applicationType": "",
    "adp_csvMerge_enginePassword": null,
    "adp_taskTimeout": 0,
    "adp_csvMerge_charset": "UTF-8",
    "adp_progressTaskTimeout": 0,
    "adp_csvMerge_textIndicator": "",
    "adp_csvMerge_noFlushBeforeMerge": "false",
    "adp_taskActive": true,
    "adp_csvMerge_engineIdFieldKey": null,
    "adp_abortWfOnFailure": true,
    "adp_csvMerge_customLineSeparator": "U+000DU+000A",
    "adp_csvMerge_multiValueDelimiter": null,
    "adp_csvMerge_csvFile": null,
    "adp_csvMerge_csvIdFieldKey": null,
    "adp_csvMerge_textFileRefIndicator": "",
    "adp_csvMerge_engineUser": null,
    "adp_csvMerge_textFileCharset": "",
    "adp_csvMerge_doNotAddTimestampToOutputFiles": "false",
    "adp_csvMerge_nativeBasePath": null,
    "adp_csvMerge_csvMode": "append",
    "adp_csvMerge_fieldMappings": [],
    "adp_cleanUpHistory": false,
    "adp_csvMerge_csvMergeConfiguration": "",
    "adp_csvMerge_errorLogFile": null,
    "adp_csvMerge_dryRun": "false",
    "adp_csvMerge_csvFieldImageLocation": null,
    "adp_csvMerge_matchesLogFile": null,
    "adp_csvMerge_csvIdPrefix": null,
    "adp_csvMerge_expandToFamily": "false"
  },
  "taskDescription": "Merges metaData or natives/images by using a csv file.",
  "taskDisplayName": "Csv merge task"
}
```

---

## CLI Arguments

See [cli.md](../cli.md) for global flags and naming conventions.

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--applicationIdentifier` | string | "" | Application identifier |
| `--csvFieldImageLocation` | string | null | CSV field image location |
| `--csvFieldNativeLocation` | string | null | CSV field native location |
| `--csvFile` | string | null | CSV file path (Required) |
| `--csvIdFieldKey` | string | null | CSV ID field key |
| `--csvMode` | string | "append" | CSV mode (append/merge) |
| `--doNotChangeProtectedDocuments` | boolean | false | Do not change protected documents |
| `--engineIdFieldKey` | string | null | Engine ID field key |
| `--engineName` | string | null | Engine name |
| `--enginePassword` | string | null | Engine password |
| `--engineUser` | string | null | Engine user |
| `--fieldMappings` | JSON array | [] | Field mappings |
| `--fieldSeparator` | string | ";" | Field separator |
| `--imageBasePath` | string | null | Image base path |
| `--mergeType` | string | "Merge content" | Merge type |
| `--multiValueDelimiter` | string | null | Multi-value delimiter |
| `--nativeBasePath` | string | null | Native base path |
| `--textIndicator` | string | "" | Text indicator |

> engineName and applicationIdentifier are mutually exclusive selectors. Exactly one must be provided.

### CLI Examples

```bash
# Basic merge
adpgo csv-merge --csvFile "/path/to/data.csv" --engineName "myEngine"

# Merge with options
adpgo csv-merge --csvFile "/path/to/data.csv" --engineName "myEngine" --csvIdFieldKey "id" --mergeType "Merge content"

# Merge using application identifier
adpgo csv-merge --csvFile "/path/to/data.csv" --applicationIdentifier "my-app-id" --csvIdFieldKey "id"
```

---

## Raw Example Response

### Verified Response (from live API)

```json
{
  "executionId": "ed45a147-18b7-45ec-a0e3-ceb7ff658130",
  "taskType": "CSV Merge",
  "loggingEnabled": "true",
  "progressMax": 2,
  "executionStatus": "success",
  "executionRootDir": "E:\\MindServer\\Projects\\adp.adp\\adpRootDir",
  "contextId": "4437498a-e67c-4873-95e3-c8e5f625bcd4",
  "executionPersistent": "true",
  "progressCurrent": 2,
  "progressPercentage": 1.0,
  "taskDisplayName": "",
  "executionMetaData": []
}
```

---

## Decoded Result

### Result Type

```
CSVMergeResult {
    // No meaningful result fields — executionMetaData is empty array on success
}
```

### Decoding Rules

On `executionStatus: "success"`, `executionMetaData` is an **empty array `[]`**. There is no meaningful result data to decode.

---

## executionMetaData Contract

On success, `executionMetaData` is an **empty array `[]`** — no fields to decode.

On failure:

| Field | Type | Description |
|-------|------|-------------|
| executionMetaData | `[]` \| `""` \| `null` | Empty array, empty string, or null on failure |

---

## Failure Response

On `executionStatus: "failed"`:

```json
{
  "executionId": "ed45a147-18b7-45ec-a0e3-ceb7ff658130",
  "taskType": "CSV Merge",
  "executionStatus": "failed",
  "errorMessage": "Error message details",
  "executionMetaData": null
}
```
