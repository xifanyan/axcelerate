# List Entities Task Specification

## Overview

| Property | Value |
|----------|-------|
| Task Type | `List Entities` |
| Description | Writes a list of entities to an output variable |
| Display Name | List entities |
| Subcommand | `list-entities` |

---

## Semantic Inputs

These are the user-facing fields for the request-construction API.

| Field | Type | Default | Required | Description |
|-------|------|---------|----------|-------------|
| type | string | "" | No | Entity type to filter |
| id | string | "" | No | Specific entity ID |
| relatedEntity | string | "" | No | Related entity ID |
| whiteList | string | "id,displayName" | No | Fields to include in output |
| workspace | string | "" | No | Workspace name |
| status | string | "" | No | Entity status filter |

---

## Raw Default Configuration

> Configuration below shows **all fields with their exact default values** from [API-SPEC.md](../../API-SPEC.md). This is for reference only. Clients must not pre-populate all fields. See [request-construction.md](../request-construction.md).

```json
{
  "taskType": "List Entities",
  "taskConfiguration": {
    "adp_progressTaskTimeout": 0,
    "adp_listEntities_file": "output.json",
    "adp_listEntities_numberOfEntities": "-1",
    "adp_listEntities_axcRequestTimeoutSeconds": 900,
    "adp_taskActive": true,
    "adp_listEntities_userHasAccess": "",
    "adp_listEntities_whiteList": "id,displayName",
    "adp_executionPersistent": true,
    "adp_abortWfOnFailure": true,
    "adp_listEntities_relatedEntity": "",
    "adp_listEntities_workspace": "",
    "adp_loggingEnabled": true,
    "adp_listEntities_status": "",
    "adp_listEntities_axcServiceCoreAddress": "",
    "adp_listEntities_relatedEntityType": "",
    "adp_listEntities_type": "",
    "adp_listEntities_httpsKeystoreFile": null,
    "adp_listEntities_httpsPassword": "",
    "adp_listEntities_axcConnectTimeoutSeconds": 300,
    "adp_listEntities_axcServicePassword": "",
    "adp_listEntities_startingEntity": "1",
    "adp_listEntities_outputJson": "adp_entities_json_output",
    "adp_cleanUpHistory": false,
    "adp_listEntities_descriptionSettingFilterValueDateFormat": "yyyy-MM-dd",
    "adp_listEntities_descriptionFilters": [],
    "adp_listEntities_axcServiceUser": "",
    "adp_listEntities_axcFields": "",
    "adp_taskTimeout": 0,
    "adp_listEntities_httpsTrustCertificate": "",
    "adp_listEntities_host": "",
    "adp_listEntities_outputFilename": "adp_entities_output_file_name",
    "adp_listEntities_id": "",
    "adp_listEntities_httpsAllowUntrustedHosts": "true"
  },
  "taskDescription": "Writes a list of entities ot an output variable",
  "taskDisplayName": "List entities"
}
```

---

## Field Reference (with Defaults)

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| adp_progressTaskTimeout | integer | 0 | Progress task timeout |
| adp_listEntities_file | string | "output.json" | Output file name |
| adp_listEntities_numberOfEntities | string | "-1" | Number of entities to retrieve |
| adp_listEntities_axcRequestTimeoutSeconds | integer | 900 | Request timeout in seconds |
| adp_taskActive | boolean | true | Whether task is active |
| adp_listEntities_userHasAccess | string | "" | User access filter |
| adp_listEntities_whiteList | string | "id,displayName" | Fields to include |
| adp_executionPersistent | boolean | true | Persist execution |
| adp_abortWfOnFailure | boolean | true | Abort workflow on failure |
| adp_listEntities_relatedEntity | string | "" | Related entity ID |
| adp_listEntities_workspace | string | "" | Workspace name |
| adp_loggingEnabled | boolean | true | Enable logging |
| adp_listEntities_status | string | "" | Entity status filter |
| adp_listEntities_axcServiceCoreAddress | string | "" | Service core address |
| adp_listEntities_relatedEntityType | string | "" | Related entity type |
| adp_listEntities_type | string | "" | Entity type to filter |
| adp_listEntities_httpsKeystoreFile | null | null | HTTPS keystore file |
| adp_listEntities_httpsPassword | string | "" | HTTPS password |
| adp_listEntities_axcConnectTimeoutSeconds | integer | 300 | Connection timeout |
| adp_listEntities_axcServicePassword | string | "" | Service password |
| adp_listEntities_startingEntity | string | "1" | Starting entity index |
| adp_listEntities_outputJson | string | "adp_entities_json_output" | Output variable name |
| adp_cleanUpHistory | boolean | false | Clean up history |
| adp_listEntities_descriptionSettingFilterValueDateFormat | string | "yyyy-MM-dd" | Date format |
| adp_listEntities_descriptionFilters | array | [] | Description filters |
| adp_listEntities_axcServiceUser | string | "" | Service username |
| adp_listEntities_axcFields | string | "" | Fields to retrieve |
| adp_taskTimeout | integer | 0 | Task timeout |
| adp_listEntities_httpsTrustCertificate | string | "" | HTTPS trust certificate |
| adp_listEntities_host | string | "" | Target host URL |
| adp_listEntities_outputFilename | string | "adp_entities_output_file_name" | Output filename parameter |
| adp_listEntities_id | string | "" | Specific entity ID |
| adp_listEntities_httpsAllowUntrustedHosts | string | "true" | Allow untrusted hosts |

---

## Raw Example Request

> Example below matches **exactly** the default configuration from [API-SPEC.md](../../API-SPEC.md).

```json
{
  "taskType": "List Entities",
  "taskConfiguration": {
    "adp_progressTaskTimeout": 0,
    "adp_listEntities_file": "output.json",
    "adp_listEntities_numberOfEntities": "-1",
    "adp_listEntities_axcRequestTimeoutSeconds": 900,
    "adp_taskActive": true,
    "adp_listEntities_userHasAccess": "",
    "adp_listEntities_whiteList": "id,displayName",
    "adp_executionPersistent": true,
    "adp_abortWfOnFailure": true,
    "adp_listEntities_relatedEntity": "",
    "adp_listEntities_workspace": "",
    "adp_loggingEnabled": true,
    "adp_listEntities_status": "",
    "adp_listEntities_axcServiceCoreAddress": "",
    "adp_listEntities_relatedEntityType": "",
    "adp_listEntities_type": "",
    "adp_listEntities_httpsKeystoreFile": null,
    "adp_listEntities_httpsPassword": "",
    "adp_listEntities_axcConnectTimeoutSeconds": 300,
    "adp_listEntities_axcServicePassword": "",
    "adp_listEntities_startingEntity": "1",
    "adp_listEntities_outputJson": "adp_entities_json_output",
    "adp_cleanUpHistory": false,
    "adp_listEntities_descriptionSettingFilterValueDateFormat": "yyyy-MM-dd",
    "adp_listEntities_descriptionFilters": [],
    "adp_listEntities_axcServiceUser": "",
    "adp_listEntities_axcFields": "",
    "adp_taskTimeout": 0,
    "adp_listEntities_httpsTrustCertificate": "",
    "adp_listEntities_host": "",
    "adp_listEntities_outputFilename": "adp_entities_output_file_name",
    "adp_listEntities_id": "",
    "adp_listEntities_httpsAllowUntrustedHosts": "true"
  },
  "taskDescription": "Writes a list of entities ot an output variable",
  "taskDisplayName": "List entities"
}
```

---

## CLI Arguments

See [cli.md](../cli.md) for global flags and naming conventions.

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--type` | string | "" | Entity type to filter |
| `--id` | string | "" | Specific entity ID |
| `--relatedEntity` | string | "" | Related entity ID |
| `--whiteList` | string | "id,displayName" | Fields to include |
| `--workspace` | string | "" | Workspace name |
| `--status` | string | "" | Entity status filter |

### CLI Examples

```bash
# Basic
adpgo list-entities

# With type filter
adpgo list-entities --type singleMindServer

# With multiple options
adpgo list-entities --type singleMindServer --whiteList "id,displayName,processStatus"

# get all datasources for ingestion application documentHold.demo00001
adpgo --debug=false list-entities --type dataSource --relatedEntity documentHold.demo00001

# get all running ingestion applications
adpgo list-entities --type docmentHold --status running
```

---

## Raw Example Response

```json
{
  "executionId": "f9463001-dc1f-486a-a8a0-efaca8dd29cb",
  "taskType": "List Entities",
  "loggingEnabled": "false",
  "progressMax": 1,
  "executionStatus": "success",
  "executionRootDir": "E:\\MindServer\\Projects\\adp.adp\\adpRootDir",
  "contextId": "2e5a47e4-d9c8-4547-aaba-45c0a3774d47",
  "executionPersistent": "false",
  "progressCurrent": 1,
  "progressPercentage": 1.0,
  "taskDisplayName": "",
  "executionMetaData": {
    "adp_entities_output_file_name": "E:\\MindServer\\Projects\\adp.adp\\adpRootDir\\output.json",
    "adp_entities_json_output": "[{\"id\":\"axcelerate.demo00001_Demo_Review4\",\"displayName\":\"demo00001_Demo_Review4\",\"processStatus\":\"Running\",\"hostId\":\"vm-rhauswirth2.otxlab.net\",\"hostName\":\"vm-rhauswirth2.otxlab.net\",\"sourceForCreateFromExisting\":false}]"
  }
}
```

---

## Decoded Result

### Result Type

```
ListEntitiesResult {
    outputFile: string
    entities: Entity[]
}
```

### Entity Type

```
Entity {
    id: string
    displayName: string
    processStatus: string
    hostId: string
    hostName: string
    sourceForCreateFromExisting: boolean
}
```

### Decoding Rules

1. Map `executionMetaData.adp_entities_output_file_name` to `outputFile`
2. Parse `executionMetaData.adp_entities_json_output` as a JSON string into `entities[]`
3. Each entity field type is inferred from the JSON value (string, boolean)

---

## executionMetaData Contract

| Field | Type | Description |
|-------|------|-------------|
| adp_entities_output_file_name | string | Output file path |
| adp_entities_json_output | string | JSON string containing array of entities — must be parsed |

### JSON String Fields

| Field | Parse As |
|-------|----------|
| adp_entities_json_output | `Entity[]` (JSON array) |

---

## Failure Response

On `executionStatus: "failed"`:

```json
{
  "executionId": "f9463001-dc1f-486a-a8a0-efaca8dd29cb",
  "taskType": "List Entities",
  "executionStatus": "failed",
  "errorMessage": "Error message details",
  "executionMetaData": null
}
```
