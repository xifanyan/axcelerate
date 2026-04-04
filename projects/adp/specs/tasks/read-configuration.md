# Read Configuration Task Specification

## Overview

| Property | Value |
|----------|-------|
| Task Type | `Read Configuration` |
| Description | A Task to read configurations into JSON or XML |
| Display Name | Read Configuration |
| Subcommand | `read-configuration` |

---

## Semantic Inputs

These are the user-facing fields for the request-construction API.

| Field | Type | Default | Required | Description |
|-------|------|---------|----------|-------------|
| entityIdToRead | string | "" | No | Entity ID to read |
| configsToRead | array | [] | No | Configs to read |
| fileFormat | string | "JSON" | No | Output format: "JSON" or "XML" |

---

## Raw Default Configuration

> Configuration below shows **all fields with their exact default values** from [API-SPEC.md](../../API-SPEC.md). This is for reference only. Clients must not pre-populate all fields. See [request-construction.md](../request-construction.md).

```json
{
  "taskType": "Read Configuration",
  "taskConfiguration": {
    "adp_readConfiguration_outputJson": "adp_entities_json_output",
    "adp_progressTaskTimeout": 0,
    "adp_readConfiguration_configsToRead": [],
    "adp_taskActive": true,
    "adp_readConfiguration_outputFilename": "adp_entities_output_file_name",
    "adp_executionPersistent": true,
    "adp_abortWfOnFailure": true,
    "adp_cleanUpHistory": false,
    "adp_readConfiguration_entityIdToRead": "",
    "adp_loggingEnabled": true,
    "adp_readConfiguration_file": "output.json",
    "adp_taskTimeout": 0,
    "adp_readConfiguration_fileFormat": "JSON"
  },
  "taskDescription": "A Task to read configurations into JSON or XML.",
  "taskDisplayName": "Read Configuration"
}
```

---

## Field Reference (with Defaults)

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| adp_readConfiguration_outputJson | string | "adp_entities_json_output" | Output JSON variable name |
| adp_progressTaskTimeout | integer | 0 | Progress task timeout |
| adp_readConfiguration_configsToRead | array | [] | Configs to read |
| adp_taskActive | boolean | true | Whether task is active |
| adp_readConfiguration_outputFilename | string | "adp_entities_output_file_name" | Output filename variable |
| adp_executionPersistent | boolean | true | Persist execution |
| adp_abortWfOnFailure | boolean | true | Abort workflow on failure |
| adp_cleanUpHistory | boolean | false | Clean up history |
| adp_readConfiguration_entityIdToRead | string | "" | Entity ID to read |
| adp_loggingEnabled | boolean | true | Enable logging |
| adp_readConfiguration_file | string | "output.json" | Output file name |
| adp_taskTimeout | integer | 0 | Task timeout |
| adp_readConfiguration_fileFormat | string | "JSON" | Output format: "JSON" or "XML" |

---

## Raw Example Request

> Example below matches **exactly** the default configuration from [API-SPEC.md](../../API-SPEC.md).

```json
{
  "taskType": "Read Configuration",
  "taskConfiguration": {
    "adp_readConfiguration_outputJson": "adp_entities_json_output",
    "adp_progressTaskTimeout": 0,
    "adp_readConfiguration_configsToRead": [],
    "adp_taskActive": true,
    "adp_readConfiguration_outputFilename": "adp_entities_output_file_name",
    "adp_executionPersistent": true,
    "adp_abortWfOnFailure": true,
    "adp_cleanUpHistory": false,
    "adp_readConfiguration_entityIdToRead": "",
    "adp_loggingEnabled": true,
    "adp_readConfiguration_file": "output.json",
    "adp_taskTimeout": 0,
    "adp_readConfiguration_fileFormat": "JSON"
  },
  "taskDescription": "A Task to read configurations into JSON or XML.",
  "taskDisplayName": "Read Configuration"
}
```

---

## CLI Arguments

See [cli.md](../cli.md) for global flags and naming conventions.

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--entityIdToRead` | string | "" | Entity ID to read |
| `--configsToRead` | string | "" | Configs to read (comma-separated or JSON array) |
| `--fileFormat` | string | "JSON" | Output format: JSON or XML |

### CLI Examples

```bash
# Basic
adpgo read-configuration

# With entity ID
adpgo read-configuration --entityIdToRead "my-entity-id"

# With configs
adpgo read-configuration --configsToRead "config1,config2"

# As XML
adpgo read-configuration --fileFormat "XML"
```

---

## Raw Example Response

> **Pending**: `executionMetaData` response fields not yet verified against actual API response.

```json
{
  "executionId": "uuid",
  "taskType": "Read Configuration",
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
    "adp_entities_output_file_name": "path/to/output.json",
    "adp_entities_json_output": "{...}"
  }
}
```

---

## Decoded Result

### Result Type

> **Pending**: `executionMetaData` response fields not yet verified against actual API response.

```
ReadConfigurationResult {
    # fields to be confirmed
}
```

### Decoding Rules

> **Pending**: Awaiting verification against actual API response.

---

## executionMetaData Contract

> **Pending**: `executionMetaData` response fields not yet verified against actual API response.

| Field | Type | Description |
|-------|------|-------------|
| adp_entities_output_file_name | string | Output file path — pending verification |
| adp_entities_json_output | string | JSON output — pending verification |

---

## Failure Response

On `executionStatus: "failed"`:

```json
{
  "executionId": "uuid",
  "taskType": "Read Configuration",
  "executionStatus": "failed",
  "errorMessage": "Error message details",
  "executionMetaData": null
}
```
