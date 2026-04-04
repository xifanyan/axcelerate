# Start Application Task Specification

## Overview

| Property | Value |
|----------|-------|
| Task Type | `Start Application` |
| Description | Starts an application |
| Display Name | Start application |
| Subcommand | `start-application` |

---

## Semantic Inputs

These are the user-facing fields for the request-construction API.

| Field | Type | Default | Required | Description |
|-------|------|---------|----------|-------------|
| applicationIdentifier | string | "" | No | Application identifier |
| useHttps | boolean | false | No | Use HTTPS |

---

## Raw Default Configuration

> Configuration below shows **all fields with their exact default values** from [API-SPEC.md](../../API-SPEC.md). This is for reference only. Clients must not pre-populate all fields. See [request-construction.md](../request-construction.md).

```json
{
  "taskType": "Start Application",
  "taskConfiguration": {
    "adp_progressTaskTimeout": 0,
    "adp_loggingEnabled": true,
    "adp_taskActive": true,
    "adp_taskTimeout": 0,
    "adp_startApplication_useHttps": false,
    "adp_startApplication_applicationUrl": "adp_started_application_url",
    "adp_executionPersistent": true,
    "adp_abortWfOnFailure": true,
    "adp_cleanUpHistory": false,
    "adp_startApplication_applicationIdentifier": "{adp_create_application_application_identifier}"
  },
  "taskDescription": "Starts an application",
  "taskDisplayName": "Start application"
}
```

---

## Field Reference (with Defaults)

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| adp_progressTaskTimeout | integer | 0 | Progress task timeout |
| adp_loggingEnabled | boolean | true | Enable logging |
| adp_taskActive | boolean | true | Whether task is active |
| adp_taskTimeout | integer | 0 | Task timeout |
| adp_startApplication_useHttps | boolean | false | Use HTTPS |
| adp_startApplication_applicationUrl | string | "adp_started_application_url" | Application URL output variable |
| adp_executionPersistent | boolean | true | Persist execution |
| adp_abortWfOnFailure | boolean | true | Abort workflow on failure |
| adp_cleanUpHistory | boolean | false | Clean up history |
| adp_startApplication_applicationIdentifier | string | "{adp_create_application_application_identifier}" | Application identifier |

---

## Raw Example Request

> Example below matches **exactly** the default configuration from [API-SPEC.md](../../API-SPEC.md).

```json
{
  "taskType": "Start Application",
  "taskConfiguration": {
    "adp_progressTaskTimeout": 0,
    "adp_loggingEnabled": true,
    "adp_taskActive": true,
    "adp_taskTimeout": 0,
    "adp_startApplication_useHttps": false,
    "adp_startApplication_applicationUrl": "adp_started_application_url",
    "adp_executionPersistent": true,
    "adp_abortWfOnFailure": true,
    "adp_cleanUpHistory": false,
    "adp_startApplication_applicationIdentifier": "{adp_create_application_application_identifier}"
  },
  "taskDescription": "Starts an application",
  "taskDisplayName": "Start application"
}
```

---

## CLI Arguments

See [cli.md](../cli.md) for global flags and naming conventions.

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--applicationIdentifier` | string | "" | Application identifier |
| `--useHttps` | boolean | false | Use HTTPS |

### CLI Examples

```bash
# Start an application
adpgo start-application --applicationIdentifier "my-app-id"

# Start with HTTPS
adpgo start-application --applicationIdentifier "my-app-id" --useHttps
```

---

## Raw Example Response

```json
{
  "executionId": "bea9fb98-c8f1-4189-be46-14e6c1a77d8c",
  "taskType": "Start Application",
  "loggingEnabled": "true",
  "executionMetaData": {
    "adp_started_application_url": "https://vm-rhauswirth2.otxlab.net:8443/demo00001"
  },
  "progressMax": 1,
  "executionStatus": "success",
  "executionRootDir": "E:\\MindServer\\Projects\\adp.adp\\adpRootDir",
  "contextId": "9b22e627-e10b-412c-aa17-4f9bdfc48d5c",
  "executionPersistent": "true",
  "progressCurrent": 1,
  "progressPercentage": 1.0,
  "taskDisplayName": "Start Application"
}
```

---

## Decoded Result

### Result Type

```
StartApplicationResult {
    applicationUrl: string
}
```

### Decoding Rules

1. Map `executionMetaData.adp_started_application_url` directly to `applicationUrl`

---

## executionMetaData Contract

| Field | Type | Description |
|-------|------|-------------|
| adp_started_application_url | string | The URL of the started application |

---

## Failure Response

On `executionStatus: "failed"`:

```json
{
  "executionId": "bea9fb98-c8f1-4189-be46-14e6c1a77d8c",
  "taskType": "Start Application",
  "executionStatus": "failed",
  "errorMessage": "Error message details",
  "executionMetaData": null
}
```
