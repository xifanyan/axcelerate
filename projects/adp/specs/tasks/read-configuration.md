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
| configsToRead | ConfigArg[] | [] | No | Configs to read |
| fileFormat | string | "JSON" | No | Output format: "JSON" or "XML" |

### ConfigArg

| Field | Type | Description |
|-------|------|-------------|
| ConfigurationID | string | Configuration identifier |
| DynamicComponentNames | string | Dynamic component names |
| FieldList | string | Comma-separated field list |
| NameValueList | string | Comma-separated name value list |
| ApplicationType | string | Application type |
| EntityType | string | Entity type |

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

### Verified Example Request (from live API)

```json
{
  "taskType": "Read Configuration",
  "taskDescription": "",
  "taskDisplayName": "",
  "taskConfiguration": {
    "adp_readConfiguration_configsToRead": [
      {
        "Configuration ID": "dataSource.file_demo_01",
        "Dynamic Component Names": "x",
        "Field list": "name,value,cells",
        "Name value list": "crawlLocationClassifierRules,uriPerlPatterns",
        "Application type": "",
        "Entity type": ""
      }
    ],
    "adp_loggingEnabled": true
  }
}
```

---

## CLI Arguments

See [cli.md](../cli.md) for global flags and naming conventions.

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--entityIdToRead` | string | "" | Entity ID to read |
| `--configsToRead` | string | "" | Configs to read (JSON array) |
| `--fileFormat` | string | "JSON" | Output format: JSON or XML |

### CLI Examples

```bash
# Basic
adpgo read-configuration

# With entity ID
adpgo read-configuration --entityIdToRead "my-entity-id"

# With configs
adpgo read-configuration --configsToRead '[{"Configuration ID":"dataSource.file_demo_01","Field list":"name,value,cells"}]'

# As XML
adpgo read-configuration --fileFormat "XML"
```

---

## Raw Example Response

```json
{
  "executionId": "uuid",
  "taskType": "Read Configuration",
  "loggingEnabled": "false",
  "progressMax": 1,
  "executionStatus": "success",
  "executionRootDir": "E:\\MindServer\\Projects\\adp.adp\\adpRootDir",
  "contextId": "uuid",
  "executionPersistent": "false",
  "progressCurrent": 1,
  "progressPercentage": 1.0,
  "taskDisplayName": "",
  "executionMetaData": {
    "adp_entities_output_file_name": "E:\\MindServer\\Projects\\adp.adp\\adpRootDir\\output.json",
    "adp_entities_json_output": "{...configuration JSON...}"
  }
}
```

### Verified executionMetaData (from live API)

```json
{
  "adp_entities_output_file_name": "E:\\MindServer\\Projects\\adp.adp\\adpRootDir\\output.json",
  "adp_entities_json_output": "{\"dataSource.file_demo_01\" : { \"DynamicComponents\" : { }, \"Global\" : { \"Static\" : { \"Parameters\" : [ {\"cells\":[[{\"value\":\".*\",\"name\":\"Pattern\"},{\"value\":\"source1\",\"name\":\"Label\"},{\"value\":\"rm_source\",\"name\":\"Text Type\"}],[{\"value\":\".*\",\"name\":\"Pattern\"},{\"value\":\"mmt5\",\"name\":\"Label\"},{\"value\":\"rm_custodian\",\"name\":\"Text Type\"}],[{\"value\":\".*\",\"name\":\"Pattern\"},{\"value\":\"batch1\",\"name\":\"Label\"},{\"value\":\"rm_batch\",\"name\":\"Text Type\"}]],\"name\":\"crawlLocationClassifierRules\",\"value\":null}, {\"cells\":[],\"name\":\"uriPerlPatterns\",\"value\":null} ] } } }"
}
```

---

## Decoded Result

### Result Type

```
ReadConfigurationResult {
    outputFile: string
    configuration: Record<string, ConfigurationInfo>
}

ConfigurationInfo {
    DynamicComponents: Record<string, any>
    Global: {
        Static: {
            Parameters: ConfigurationParameter[]
        }
    }
}

ConfigurationParameter {
    Cells: Cell[][]     // 2D array
    Name: string
    Value: any
}

Cell {
    Value: any
    Name: string
}
```

> **Note:** Types use TypeScript-like notation for language-agnostic representation. `Record<K, V>` = map/dictionary, `any` = language-appropriate any/void*/dynamic type.

### Decoding Rules

1. Map `executionMetaData.adp_entities_output_file_name` to `outputFile`
2. Parse `executionMetaData.adp_entities_json_output` as a JSON string into `ReadConfigurationResult` (map of configuration key → ConfigurationInfo)

---

## executionMetaData Contract

| Field | Type | Description |
|-------|------|-------------|
| adp_entities_output_file_name | string | Output file path |
| adp_entities_json_output | string | JSON string containing configuration data — must be parsed |

### JSON String Fields

| Field | Parse As |
|-------|----------|
| adp_entities_json_output | `ReadConfigurationResult` (map[string]ConfigurationInfo) |

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
