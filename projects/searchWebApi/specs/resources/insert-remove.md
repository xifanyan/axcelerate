# Insert Remove Resource

## Purpose

This resource area covers direct insert/remove requests and the multi-step bulk insert/remove transaction workflow.

## Operations

### Insert Remove Transaction

- Raw operation: `POST /projects/{projectId}/collections/{collectionId}/records/insertRemoveTransaction`
- Raw operationId: `insertRemoveTransaction`
- Request schema: `InsertRemoveRequest`
- Result schema: `InsertRemoveResult`

### Start Bulk Insert Remove Transaction

- Raw operation: `POST /projects/{projectId}/collections/{collectionId}/records/bulkInsertRemoveTransaction`
- Raw operationId: `startInsertRemoveTransaction`
- Request schema: `StartTransactionRequest`
- Result schema: `StartTransactionResult`

### Commit Bulk Insert Remove Transaction

- Raw operation: `POST /projects/{projectId}/collections/{collectionId}/records/bulkInsertRemoveTransaction/{indexingBufferId}/end`
- Raw operationId: `commitInsertRemoveTransaction`
- Request schema: `FinishTransactionRequest`
- Result schema: `FinishTransactionResponse`

### Get Flush Job Status

- Raw operation: `GET /projects/{projectId}/collections/{collectionId}/records/bulkInsertRemoveTransaction/{indexingBufferId}/end/{jobId}`
- Raw operationId: `getFlushJobStatus`
- Result schema: `JobStatusResponse`

### Add To Bulk Insert Remove Transaction

- Raw operation: `POST /projects/{projectId}/collections/{collectionId}/records/bulkInsertRemoveTransaction/{indexingBufferId}/buffer`
- Raw operationId: `addToInsertRemoveTransaction`
- Request schema: `InsertRemoveRequest`
- Result schema: `InsertRemoveResult`

## Shared Rules

- Multipart and JSON insertion modes must both be documented.
- Binary index references inside `FieldData` remain wire-level behavior and must not be hidden.
- Bulk workflows must preserve the explicit sequence: start buffer, add payloads, end transaction, poll job status.
