# CLI Task Specification

## Overview

| Property | Value |
|----------|-------|
| Task Type | `CLI` |
| Description | Runs a native task in its own process |
| Display Name | Command Line Task |
| Subcommand | `cli` |

---

## Semantic Inputs

These are the user-facing fields for the request-construction API.

| Field | Type | Default | Required | Description |
|-------|------|---------|----------|-------------|
| batchScriptPath | string | null | Yes | Path to the script/executable |
| batchScriptParameters | CLIBatchParameter[] | [] | No | Parameters for the script |
| workingDirectory | string | null | No | Working directory for the script |
| batchScriptJsonLogOutput | string | "" | No | JSON output variable name |
| batchScriptRedirectLogging | boolean | false | No | Redirect logging |
| batchScriptPositiveExecutionCodes | string | "0" | No | Comma-separated success codes |
| batchScriptFilterPasswords | boolean | true | No | Filter passwords from logs |
| batchScriptLoggingDirectory | string | null | No | Logging directory |
| batchScriptResultCode | string | "cli_result" | No | Result code field name |
| batchScriptResultLogPath | string | "cli_result_path" | No | Result log path field name |
| batchScriptErrorLogPath | string | "cli_error_path" | No | Error log path field name |

### CLIBatchParameter

| Field | Type | Description |
|-------|------|-------------|
| Parameter | string | Parameter value |

---

## Raw Default Configuration

> Configuration below shows **all fields with their exact default values** from [API-SPEC.md](../../API-SPEC.md). This is for reference only. Clients must not pre-populate all fields. See [request-construction.md](../request-construction.md).

```json
{
  "taskType": "CLI",
  "taskConfiguration": {
    "adp_batchScriptRedirectLogging": "false",
    "adp_progressTaskTimeout": 0,
    "adp_envVariables": null,
    "adp_taskActive": true,
    "adp_pathEntries": null,
    "adp_batchScriptPositiveExecutionCodes": "0",
    "adp_executionPersistent": true,
    "adp_abortWfOnFailure": true,
    "adp_batchScriptResultCode": "cli_result",
    "adp_cleanUpHistory": false,
    "adp_batchScriptResultLogPath": "cli_result_path",
    "adp_loggingEnabled": true,
    "adp_batchScriptFilterPasswords": true,
    "adp_batchScriptLoggingDirectory": null,
    "adp_batchScriptErrorLogPath": "cli_error_path",
    "adp_taskTimeout": 0,
    "adp_batchScriptPath": null,
    "adp_workingDirectory": null,
    "adp_batchScriptParameters": [],
    "adp_batchScriptJsonLogOutput": ""
  },
  "taskDescription": "Runs a native task in its own process",
  "taskDisplayName": "Command Line Task"
}
```

---

## Field Reference (with Defaults)

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| adp_batchScriptRedirectLogging | string | "false" | Redirect logging — coerce to boolean |
| adp_progressTaskTimeout | integer | 0 | Progress task timeout |
| adp_envVariables | string | null | Environment variables |
| adp_taskActive | boolean | true | Whether task is active |
| adp_pathEntries | string | null | Path entries |
| adp_batchScriptPositiveExecutionCodes | string | "0" | Comma-separated success codes |
| adp_executionPersistent | boolean | true | Persist execution |
| adp_abortWfOnFailure | boolean | true | Abort workflow on failure |
| adp_batchScriptResultCode | string | "cli_result" | Result code field name |
| adp_cleanUpHistory | boolean | false | Clean up history |
| adp_batchScriptResultLogPath | string | "cli_result_path" | Result log path field name |
| adp_loggingEnabled | boolean | true | Enable logging |
| adp_batchScriptFilterPasswords | boolean | true | Filter passwords from logs |
| adp_batchScriptLoggingDirectory | string | null | Logging directory |
| adp_batchScriptErrorLogPath | string | "cli_error_path" | Error log path field name |
| adp_taskTimeout | integer | 0 | Task timeout |
| adp_batchScriptPath | string | null | Path to script/executable |
| adp_workingDirectory | string | null | Working directory |
| adp_batchScriptParameters | array | [] | Script parameters |
| adp_batchScriptJsonLogOutput | string | "" | JSON output variable name |

---

## Raw Example Request

> Example below matches **exactly** the default configuration from [API-SPEC.md](../../API-SPEC.md).

```json
{
  "taskType": "CLI",
  "taskConfiguration": {
    "adp_batchScriptRedirectLogging": "false",
    "adp_progressTaskTimeout": 0,
    "adp_envVariables": null,
    "adp_taskActive": true,
    "adp_pathEntries": null,
    "adp_batchScriptPositiveExecutionCodes": "0",
    "adp_executionPersistent": true,
    "adp_abortWfOnFailure": true,
    "adp_batchScriptResultCode": "cli_result",
    "adp_cleanUpHistory": false,
    "adp_batchScriptResultLogPath": "cli_result_path",
    "adp_loggingEnabled": true,
    "adp_batchScriptFilterPasswords": true,
    "adp_batchScriptLoggingDirectory": null,
    "adp_batchScriptErrorLogPath": "cli_error_path",
    "adp_taskTimeout": 0,
    "adp_batchScriptPath": null,
    "adp_workingDirectory": null,
    "adp_batchScriptParameters": [],
    "adp_batchScriptJsonLogOutput": ""
  },
  "taskDescription": "Runs a native task in its own process",
  "taskDisplayName": "Command Line Task"
}
```

---

## CLI Arguments

See [cli.md](../cli.md) for global flags and naming conventions.

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--batchScriptPath` | string | null | Path to script/executable (Required) |
| `--batchScriptParameters` | string | "" | Parameters as JSON array |
| `--workingDirectory` | string | null | Working directory |
| `--batchScriptJsonLogOutput` | string | "" | JSON output variable name |
| `--batchScriptRedirectLogging` | boolean | false | Redirect logging |
| `--batchScriptPositiveExecutionCodes` | string | "0" | Comma-separated success codes |
| `--batchScriptFilterPasswords` | boolean | true | Filter passwords from logs |
| `--batchScriptLoggingDirectory` | string | null | Logging directory |

### CLI Examples

```bash
# Basic PowerShell script
adpgo cli --batchScriptPath "C:\Windows\SysWOW64\WindowsPowerShell\v1.0\powershell.exe" --batchScriptParameters '[{"Parameter":"-File"},{"Parameter":"c:\\demo\\test.ps1"}]'

# With working directory
adpgo cli --batchScriptPath "c:\demo\script.ps1" --workingDirectory "c:\demo"

# With JSON output
adpgo cli --batchScriptPath "c:\demo\script.ps1" --batchScriptJsonLogOutput "json_output"
```

---

## Raw Example Response

> **Pending**: Full `executionMetaData` response not yet verified. Partial structure provided.

```json
{
  "executionId": "uuid",
  "taskType": "CLI",
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
    "cli_result": 0,
    "json_output": "{...}",
    "cli_error_path": "path/to/error.log",
    "cli_result_path": "path/to/result.log"
  }
}
```

### Commonly Used executionMetaData Fields

| Field | Type | Description |
|-------|------|-------------|
| cli_result | integer | Exit code (0 = success) |
| json_output | string | JSON string with stdout/stderr output |
| cli_error_path | string | Path to error log file |
| cli_result_path | string | Path to result log file |

### json_output Structure

When `batchScriptJsonLogOutput` is set, `json_output` contains:

```json
{
  "stdout": "command output",
  "errout": "error output"
}
```

---

## Decoded Result

### Result Type

```
CLIResult {
    result: integer
    jsonOutput: Record<string, any> | absent  # parsed from json_output if set
    errorPath: string
    resultPath: string
}
```

### Decoding Rules

1. Map `executionMetaData.cli_result` to `result` (integer)
2. If `executionMetaData.json_output` is present, parse as JSON into `jsonOutput`
3. Map `executionMetaData.cli_error_path` to `errorPath`
4. Map `executionMetaData.cli_result_path` to `resultPath`

---

## executionMetaData Contract

| Field | Type | Description |
|-------|------|-------------|
| cli_result | integer | Exit code (0 = success) |
| json_output | string | JSON string with stdout/stderr — parse if present |
| cli_error_path | string | Path to error log file |
| cli_result_path | string | Path to result log file |

### JSON String Fields

| Field | Parse As |
|-------|----------|
| json_output | `Record<string, string>` with `stdout` and `errout` fields |

---

## Failure Response

On `executionStatus: "failed"`:

```json
{
  "executionId": "uuid",
  "taskType": "CLI",
  "executionStatus": "failed",
  "errorMessage": "Error message details",
  "executionMetaData": null
}
```
