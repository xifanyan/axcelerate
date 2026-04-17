# Create Data Source Task Specification

## Overview

| Property | Value |
|----------|-------|
| Task Type | `Create Data Source` |
| Description | Creates a new data source |
| Display Name | Create data source |
| Subcommand | `create-data-source` |

---

## Semantic Inputs

These are the user-facing fields for the request-construction API.

| Field | Type | Default | Required | Description |
|-------|------|---------|----------|-------------|
| dataSourceName | string | "" | Yes | Data source name |
| dataSourceIdentifier | string | "{datasource_id}" | No | Data source identifier |
| dataSourceSystemTemplateDisplayName | string | "Server - file share" | No | System template display name |
| dataSourceTemplate | string | "" | No | Data source template |
| usedTemplate | string | "adp_used_data_source_template" | No | Output variable for used template |
| abortOnExistingDataSource | boolean | false | No | Abort if datasource already exists |
| hostIdentifier | string | null | No | Host identifier |
| hostCpuLoadThreshold | string | "50" | No | CPU load threshold percentage |
| maxNumberRunningCrawlers | string | "0" | No | Max number of running crawlers |
| hostMemoryLimit | string | "0" | No | Host memory limit |
| hostMemoryLimitRatio | string | "0" | No | Host memory limit ratio |
| retryMaxNumberRunningCrawlers | string | "30" | No | Retry max number of running crawlers |
| hostRolesBlackList | array | null | No | Host roles blacklist |
| engineIdentifier | string | null | Conditional | Engine identifier |
| engineBoxDocThreshold | string | "1000000" | No | Engine box document threshold |
| applicationIdentifier | string | null | Conditional | Application identifier |
| workspaceIdentifier | string | null | No | Workspace identifier |
| choosenHostNameParameter | string | "adp_hostname" | No | Output variable for hostname |
| choosenHostCpuLoad | string | "adp_chosen_host_cpu_load" | No | Output variable for CPU load |
| choosenHostMemory | string | "adp_chosen_host_memory" | No | Output variable for host memory |
| choosenHostMemoryRatio | string | "adp_chosen_host_memory_ratio" | No | Output variable for memory ratio |
| choosenEngineNameParameter | string | "adp_chosen_engine" | No | Output variable for engine name |
| createdDataSourceNameParameter | string | "adp_created_data_source_name" | No | Output variable for datasource name |
| createdDataSourceDisplaynameParameter | string | "adp_created_data_source_displayname" | No | Output variable for display name |

> `applicationIdentifier` and `engineIdentifier` are mutually exclusive selectors. Exactly one must be provided.

---

## Raw Default Configuration

> Configuration below shows **all fields with their exact default values** from [API-SPEC.md](../../API-SPEC.md). This is for reference only. Clients must not pre-populate all fields — only send fields explicitly set by the caller. See [request-construction.md](../request-construction.md).

```json
{
  "taskType": "Create Data Source",
  "taskConfiguration": {
    "adp_progressTaskTimeout": 0,
    "adp_createDataSource_abortOnExistingDataSource": false,
    "adp_createDataSource_applicationIdentifier": null,
    "adp_taskActive": true,
    "adp_createDataSource_choosenHostNameParameter": "adp_hostname",
    "adp_executionPersistent": true,
    "adp_createDataSource_choosenHostMemoryRatio": "adp_chosen_host_memory_ratio",
    "adp_abortWfOnFailure": true,
    "adp_createDataSource_choosenHostCpuLoad": "adp_chosen_host_cpu_load",
    "adp_loggingEnabled": true,
    "adp_createDataSource_dataSourceSystemTemplateDisplayName": "Server - file share",
    "adp_createDataSource_usedTemplate": "adp_used_data_source_template",
    "adp_createDataSource_hostCpuLoadThreshold": "50",
    "adp_createDataSource_createdDataSourceNameParameter": "adp_created_data_source_name",
    "adp_createDataSource_retryMaxNumberRunningCrawlers": "30",
    "adp_createDataSource_choosenHostMemory": "adp_chosen_host_memory",
    "adp_createDataSource_workspaceIdentifier": null,
    "adp_createDataSource_hostIdentifier": null,
    "adp_createDataSource_hostMemoryLimit": "0",
    "adp_createDataSource_maxNumberRunningCrawlers": "0",
    "adp_cleanUpHistory": false,
    "adp_createDataSource_engineIdentifier": null,
    "adp_createDataSource_engineBoxDocThreshold": "1000000",
    "adp_createDataSource_hostMemoryLimitRatio": "0",
    "adp_createDataSource_choosenEngineNameParameter": "adp_chosen_engine",
    "adp_createDataSource_hostRolesBlackList": null,
    "adp_createDataSource_dataSourceIdentifier": "{datasource_id}",
    "adp_taskTimeout": 0,
    "adp_createDataSource_createdDataSourceDisplaynameParameter": "adp_created_data_source_displayname",
    "adp_createDataSource_dataSourceTemplate": "",
    "adp_createDataSource_dataSourceName": "{datasource_name}"
  },
  "taskDescription": "Creates a new data source",
  "taskDisplayName": "Create data source"
}
```

---

## Raw Example Request

> Example below shows a request using `engineIdentifier` as the selector.

```json
{
  "taskType": "Create Data Source",
  "taskConfiguration": {
    "adp_createDataSource_dataSourceName": "file_demo_05",
    "adp_createDataSource_dataSourceSystemTemplateDisplayName": "Server - file share",
    "adp_createDataSource_engineIdentifier": "eng123"
  }
}
```

### Verified Example Request (from live API)

```json
{
  "taskType": "Create Data Source",
  "taskConfiguration": {
    "adp_createDataSource_dataSourceName": "file_demo_05",
    "adp_createDataSource_dataSourceSystemTemplateDisplayName": "Server - file share",
    "adp_createDataSource_choosenHostNameParameter": "adp_hostname",
    "adp_createDataSource_choosenHostCpuLoad": "adp_chosen_host_cpu_load",
    "adp_createDataSource_choosenHostMemory": "adp_chosen_host_memory",
    "adp_createDataSource_choosenHostMemoryRatio": "adp_chosen_host_memory_ratio",
    "adp_createDataSource_choosenEngineNameParameter": "adp_chosen_engine",
    "adp_createDataSource_createdDataSourceNameParameter": "adp_created_data_source_name",
    "adp_createDataSource_createdDataSourceDisplaynameParameter": "adp_created_data_source_displayname",
    "adp_createDataSource_usedTemplate": "adp_used_data_source_template"
  }
}
```

---

## CLI Arguments

See [cli.md](../cli.md) for global flags and naming conventions.

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--dataSourceName` | string | "" | Data source name |
| `--dataSourceIdentifier` | string | "{datasource_id}" | Data source identifier |
| `--dataSourceSystemTemplateDisplayName` | string | "Server - file share" | System template display name |
| `--dataSourceTemplate` | string | "" | Data source template |
| `--usedTemplate` | string | "adp_used_data_source_template" | Output variable for used template |
| `--abortOnExistingDataSource` | boolean | false | Abort if datasource already exists |
| `--hostIdentifier` | string | null | Host identifier |
| `--hostCpuLoadThreshold` | string | "50" | CPU load threshold percentage |
| `--maxNumberRunningCrawlers` | string | "0" | Max number of running crawlers |
| `--hostMemoryLimit` | string | "0" | Host memory limit |
| `--hostMemoryLimitRatio` | string | "0" | Host memory limit ratio |
| `--retryMaxNumberRunningCrawlers` | string | "30" | Retry max number of running crawlers |
| `--hostRolesBlackList` | array | null | Host roles blacklist |
| `--engineIdentifier` | string | null | Engine identifier (mutually exclusive with applicationIdentifier) |
| `--engineBoxDocThreshold` | string | "1000000" | Engine box document threshold |
| `--applicationIdentifier` | string | null | Application identifier (mutually exclusive with engineIdentifier) |
| `--workspaceIdentifier` | string | null | Workspace identifier |

> `--engineIdentifier` and `--applicationIdentifier` are mutually exclusive selectors. Exactly one must be provided.

### CLI Examples

```bash
# Using engine identifier
adpgo create-data-source --dataSourceName "myDS" --engineIdentifier "eng123"

# Using application identifier
adpgo create-data-source --dataSourceName "myDS" --applicationIdentifier "app123"

# With full options
adpgo create-data-source --dataSourceName "file_demo_05" --dataSourceSystemTemplateDisplayName "Server - file share"

# Using application identifier with display name
adpgo create-data-source --dataSourceName "file_demo_05" --applicationIdentifier "app123" --dataSourceSystemTemplateDisplayName "Server - file share"
```

---

## Raw Example Response

### Verified Response (from live API)

```json
{
  "executionId": "b0c0c8b7-a0b4-4c35-a0c4-3d411c98974d",
  "taskType": "Create Data Source",
  "loggingEnabled": "true",
  "progressMax": 0,
  "executionStatus": "success",
  "executionRootDir": "E:\\MindServer\\Projects\\adp.adp\\adpRootDir",
  "contextId": "860e2cdc-24ec-424b-a369-a4ce29b4f39d",
  "executionPersistent": "true",
  "progressCurrent": 0,
  "progressPercentage": 0,
  "taskDisplayName": "Create data source",
  "executionMetaData": {
    "adp_created_data_source_displayname": "file_demo_05",
    "adp_hostname": "vm-rhauswirth2.otxlab.net",
    "adp_chosen_host_cpu_load": "0.9790296",
    "adp_created_data_source_name": "dataSource.file_demo_05",
    "adp_chosen_host_memory": "57291833344",
    "adp_chosen_host_memory_ratio": "73",
    "adp_chosen_engine": "demo00001",
    "adp_used_data_source_template": "_Disney_Template_v1"
  }
}
```

---

## Decoded Result

### Result Type

```
CreateDataSourceResult {
    displayName: string
    hostname: string
    cpuLoad: string
    dataSourceName: string
    hostMemory: string
    hostMemoryRatio: string
    engineName: string
    usedTemplate: string
}
```

### Decoding Rules

1. Map `executionMetaData.adp_created_data_source_displayname` to `displayName`
2. Map `executionMetaData.adp_hostname` to `hostname`
3. Map `executionMetaData.adp_chosen_host_cpu_load` to `cpuLoad`
4. Map `executionMetaData.adp_created_data_source_name` to `dataSourceName`
5. Map `executionMetaData.adp_chosen_host_memory` to `hostMemory`
6. Map `executionMetaData.adp_chosen_host_memory_ratio` to `hostMemoryRatio`
7. Map `executionMetaData.adp_chosen_engine` to `engineName`
8. Map `executionMetaData.adp_used_data_source_template` to `usedTemplate`

---

## executionMetaData Contract

### Success Response (Verified)

| Field | Type | Description |
|-------|------|-------------|
| adp_created_data_source_displayname | string | Display name of created data source |
| adp_hostname | string | Host name |
| adp_chosen_host_cpu_load | string | CPU load (decimal string, e.g. "0.9790296") |
| adp_created_data_source_name | string | Name of created data source |
| adp_chosen_host_memory | string | Host memory in bytes (large integer as string) |
| adp_chosen_host_memory_ratio | string | Memory ratio percentage (e.g. "73") |
| adp_chosen_engine | string | Engine name |
| adp_used_data_source_template | string | Data source template used |

On failure:

| Field | Type | Description |
|-------|------|-------------|
| executionMetaData | null | Always null on failure |

---

## Failure Response

On `executionStatus: "failed"`:

```json
{
  "executionId": "uuid",
  "taskType": "Create Data Source",
  "executionStatus": "failed",
  "errorMessage": "Error message details",
  "executionMetaData": null
}
```