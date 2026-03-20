# CSV Merge Task Specification

## Overview

| Property | Value |
|----------|-------|
| Task Type | `CSV Merge` |
| Description | Merges content or updates natives/images by using a csv file. |
| Display Name | Csv merge task |

---

## Default Configuration

> Configuration below shows **all fields with their exact default values** from API-SPEC.md

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

| Argument | Type | Default | Description |
|----------|------|---------|-------------|
| applicationIdentifier | string | "" | Application identifier |
| csvFieldImageLocation | string | null | CSV field image location |
| csvFieldNativeLocation | string | null | CSV field native location |
| csvFile | string | null | CSV file path (Required) |
| csvIdFieldKey | string | null | CSV ID field key |
| csvMode | string | "append" | CSV mode (append/merge) |
| doNotChangeProtectedDocuments | string | "false" | Do not change protected documents (true/false) |
| engineName | string | null | Engine name |
| enginePassword | string | null | Engine password |
| engineUser | string | null | Engine user |
| fieldMappings | JSON array | [] | Field mappings as column definitions like `CSV Field Name|Text type|Value delimiter|Use display name` |
| fieldSeperator | string | ";" | Field separator |
| imageBasePath | string | null | Image base path |
| mergeType | string | "Merge content" | Merge type (`Merge content` or `Update natives/images`) |
| multiValueDelimiter | string | null | Multi-value delimiter |
| nativeBasePath | string | null | Native base path |
| textIndicator | string | "" | Text indicator |

---

## Example Request

> Example below matches **exactly** the default configuration from [API-SPEC.md](../../API-SPEC.md)

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

## Example Response

```json
{
  "executionId": "uuid",
  "taskType": "CSV Merge",
  "loggingEnabled": "false",
  "progressMax": 1,
  "executionStatus": "success",
  "executionRootDir": "path",
  "contextId": "uuid",
  "executionPersistent": "false",
  "progressCurrent": 1,
  "progressPercentage": 1.0,
  "taskDisplayName": "",
  "executionMetaData": {
    "adp_csvMerge_csvMergeConfiguration": "",
    "adp_csvMerge_csvMode": "append",
    "adp_csvMerge_fieldMappings": [],
    "adp_csvMerge_engineName": null,
    "adp_csvMerge_applicationIdentifier": "",
    "adp_csvMerge_mergeType": "Merge content"
  }
}
```

---

## Response Fields

> **Pending** - Actual `ExecutionMetaData` response fields not yet documented.
