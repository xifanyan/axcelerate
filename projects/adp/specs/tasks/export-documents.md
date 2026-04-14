# Export Documents Task Specification

## Overview

| Property | Value |
|----------|-------|
| Task Type | `Export Documents` |
| Description | Export documents in CSV format |
| Display Name | Export documents task |
| Subcommand | `export-documents` |

---

## Semantic Inputs

These are the user-facing fields for the request-construction API.

| Field | Type | Default | Required | Description |
|-------|------|---------|----------|-------------|
| fieldSeparator | string | ";" | No | Field separator for CSV output |
| waitForExport | boolean | false | No | Wait for export to complete |
| query | string | "*" | No | Query to select documents |
| applicationIdentifier | string | "" | No | Application identifier |
| applicationType | string | "" | No | Application type |
| engineIdentifier | string | null | No | Engine identifier |
| engineUser | string | null | No | Engine username |
| enginePassword | string | null | No | Engine password |
| exportName | string | null | **Yes** | Export name (must be unique per run, cannot be reused) |
| exportFields | string | null | No | Fields to export as JSON map (e.g., `{"URI":"URI","rm_custodian":"Custodian"}` where key is internal field name, value is CSV header) |
| exportDirectory | string | null | No | Export directory |
| fileEnding | string | "csv" | No | File extension for export |
| imageField | boolean | false | No | Include image field in export |
| nativesField | boolean | false | No | Include natives field in export |
| textField | boolean | false | No | Include text field in export |

---

## Raw Default Configuration

> Configuration below shows **all fields with their exact default values** from [API-SPEC.md](../../API-SPEC.md). This is for reference only. Clients must not pre-populate all fields. See [request-construction.md](../request-construction.md).

```json
{
  "taskType": "Export Documents",
  "taskConfiguration": {
    "adp_progressTaskTimeout": 0,
    "adp_exportDocuments_field_separator": ";",
    "adp_exportDocuments_waitForExport": false,
    "adp_exportDocuments_image_field": null,
    "adp_exportDocuments_searchResultSize": "adp_exportDocuments_searchResultSize",
    "adp_taskActive": true,
    "adp_exportDocuments_File_Ending": "csv",
    "adp_exportDocuments_applicationType": "",
    "adp_executionPersistent": true,
    "adp_abortWfOnFailure": true,
    "adp_exportDocuments_query": "*",
    "adp_loggingEnabled": true,
    "adp_exportDocuments_exportName": null,
    "adp_exportDocuments_text_indicator": "\"",
    "adp_exportDocuments_natives_field": null,
    "adp_exportDocuments_multivalue_separator": "|",
    "adp_exportDocuments_line_break": "",
    "adp_exportDocuments_applicationIdentifier": "",
    "adp_exportDocuments_engineIdentifier": null,
    "adp_exportDocuments_exportFileName": "adp_exportDocuments_exportFileName",
    "adp_exportDocuments_engineUser": null,
    "adp_exportDocuments_image_volume": "Volume",
    "adp_exportDocuments_exportFields": null,
    "adp_exportDocuments_fullExportPath": "adp_exportDocuments_exportPath",
    "adp_exportDocuments_text_field": null,
    "adp_cleanUpHistory": false,
    "adp_exportDocuments_exportDirectory": null,
    "adp_taskTimeout": 0,
    "adp_exportDocuments_enginePassword": null,
    "adp_exportDocuments_adp_exportDocuments_mainQueryType": null
  },
  "taskDescription": "Export documents in CSV format.",
  "taskDisplayName": "Export documents task"
}
```

---

## Field Reference (with Defaults)

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| adp_exportDocuments_field_separator | string | ";" | Field separator for CSV output |
| adp_exportDocuments_waitForExport | boolean | false | Wait for export to complete |
| adp_exportDocuments_image_field | string | null | Image field to include in export |
| adp_exportDocuments_searchResultSize | string | "100000" | Maximum number of documents to export |
| adp_taskActive | boolean | true | Whether task is active |
| adp_exportDocuments_File_Ending | string | "csv" | File extension for export |
| adp_exportDocuments_applicationType | string | "" | Application type |
| adp_executionPersistent | boolean | true | Persist execution |
| adp_abortWfOnFailure | boolean | true | Abort workflow on failure |
| adp_exportDocuments_query | string | "*" | Query to select documents |
| adp_loggingEnabled | boolean | true | Enable logging |
| adp_exportDocuments_exportName | string | null | Export name (required, must be unique per run, cannot be reused) |
| adp_exportDocuments_text_indicator | string | "\"" | Text indicator for CSV |
| adp_exportDocuments_natives_field | string | null | Natives field to include |
| adp_exportDocuments_multivalue_separator | string | "|" | Multi-value separator |
| adp_exportDocuments_line_break | string | "" | Line break character |
| adp_exportDocuments_applicationIdentifier | string | "" | Application identifier |
| adp_exportDocuments_engineIdentifier | string | null | Engine identifier |
| adp_exportDocuments_exportFileName | string | "adp_exportDocuments_exportFileName" | Export file name variable |
| adp_exportDocuments_engineUser | string | null | Engine username |
| adp_exportDocuments_image_volume | string | "Volume" | Image volume |
| adp_exportDocuments_exportFields | string | null | Fields to export as JSON map (key = internal field name, value = CSV header) |
| adp_exportDocuments_fullExportPath | string | "adp_exportDocuments_exportPath" | Full export path variable |
| adp_exportDocuments_text_field | string | null | Text field to include |
| adp_cleanUpHistory | boolean | false | Clean up history |
| adp_exportDocuments_exportDirectory | string | null | Export directory |
| adp_taskTimeout | integer | 0 | Task timeout |
| adp_exportDocuments_enginePassword | string | null | Engine password |
| adp_exportDocuments_adp_exportDocuments_mainQueryType | string | null | Main query type |

---

## Raw Example Request

> Example below matches **exactly** the default configuration from [API-SPEC.md](../../API-SPEC.md).

```json
{
  "taskType": "Export Documents",
  "taskConfiguration": {
    "adp_progressTaskTimeout": 0,
    "adp_exportDocuments_field_separator": ";",
    "adp_exportDocuments_waitForExport": false,
    "adp_exportDocuments_image_field": null,
    "adp_exportDocuments_searchResultSize": "adp_exportDocuments_searchResultSize",
    "adp_taskActive": true,
    "adp_exportDocuments_File_Ending": "csv",
    "adp_exportDocuments_applicationType": "",
    "adp_executionPersistent": true,
    "adp_abortWfOnFailure": true,
    "adp_exportDocuments_query": "*",
    "adp_loggingEnabled": true,
    "adp_exportDocuments_exportName": null,
    "adp_exportDocuments_text_indicator": "\"",
    "adp_exportDocuments_natives_field": null,
    "adp_exportDocuments_multivalue_separator": "|",
    "adp_exportDocuments_line_break": "",
    "adp_exportDocuments_applicationIdentifier": "",
    "adp_exportDocuments_engineIdentifier": null,
    "adp_exportDocuments_exportFileName": "adp_exportDocuments_exportFileName",
    "adp_exportDocuments_engineUser": null,
    "adp_exportDocuments_image_volume": "Volume",
    "adp_exportDocuments_exportFields": null,
    "adp_exportDocuments_fullExportPath": "adp_exportDocuments_exportPath",
    "adp_exportDocuments_text_field": null,
    "adp_cleanUpHistory": false,
    "adp_exportDocuments_exportDirectory": null,
    "adp_taskTimeout": 0,
    "adp_exportDocuments_enginePassword": null,
    "adp_exportDocuments_adp_exportDocuments_mainQueryType": null
  },
  "taskDescription": "Export documents in CSV format.",
  "taskDisplayName": "Export documents task"
}
```

---

## CLI Arguments

See [cli.md](../cli.md) for global flags and naming conventions.

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--fieldSeparator` | string | ";" | Field separator for CSV output |
| `--waitForExport` | boolean | false | Wait for export to complete |
| `--query` | string | "*" | Query to select documents |
| `--applicationIdentifier` | string | "" | Application identifier |
| `--applicationType` | string | "" | Application type |
| `--engineIdentifier` | string | null | Engine identifier |
| `--engineUser` | string | null | Engine username |
| `--enginePassword` | string | null | Engine password |
| `--exportName` | string | null | **Required** — Export name (must be unique per run, cannot be reused) |
| `--exportFields` | string | null | Fields to export as JSON map (e.g., `{"URI":"URI","rm_custodian":"Custodian"}`) |
| `--exportDirectory` | string | null | Export directory |
| `--fileEnding` | string | "csv" | File extension for export |
| `--imageField` | boolean (switch) | false | Include image field in export |
| `--nativesField` | boolean (switch) | false | Include natives field in export |
| `--textField` | boolean (switch) | false | Include text field in export |

### CLI Examples

```bash
# Working example with application identifier, fields, and natives
# exportFields can be tricky, it needs to be a properly escaped JSON string. and powershell and command prompt might use different escaping rules. Below is an example for powershell, for command prompt you might need to adjust the escaping. 
.\bin\adpgo.exe export-documents --applicationIdentifier documentHold.demo00001 --exportFields '{"URI": "URI", "rm_custodian": "Custodian"}' --nativesField --exportName demo_exp_01
```

---

## Raw Example Response

```json
{
  "executionId": "uuid",
  "taskType": "Export Documents",
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
    "adp_exportDocuments_exportFileName": "export.csv",
    "adp_exportDocuments_exportPath": "/path/to/export/export.csv",
    "adp_exportDocuments_searchResultSize": "1000"
  }
}
```

---

## Decoded Result

### Result Type

```
ExportDocumentsResult {
    exportFileName: string
    exportPath: string
    searchResultSize: integer
}
```

### Decoding Rules

1. Map `executionMetaData.adp_exportDocuments_exportFileName` to `exportFileName`
2. Map `executionMetaData.adp_exportDocuments_exportPath` to `exportPath`
3. Parse `executionMetaData.adp_exportDocuments_searchResultSize` from string to integer

---

## executionMetaData Contract

| Field | Type | Description |
|-------|------|-------------|
| adp_exportDocuments_exportFileName | string | Name of the exported file |
| adp_exportDocuments_exportPath | string | Full path to the exported file |
| adp_exportDocuments_searchResultSize | string | Number of documents exported — coerce to integer |

---

## Failure Response

On `executionStatus: "failed"`:

```json
{
  "executionId": "uuid",
  "taskType": "Export Documents",
  "executionStatus": "failed",
  "errorMessage": "Error message details",
  "executionMetaData": null
}
```
