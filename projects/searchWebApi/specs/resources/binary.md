# Binary Resource

## Purpose

This resource area covers binary retrieval by search result selection and by explicit record id.

## Operations

### Get Binary By Search

- Raw operation: `GET /projects/{projectId}/collections/{collectionId}/binary`
- Raw operationId: `getBinary`
- Response content type: `application/octet-stream`

### Get Binary By Record Id

- Raw operation: `GET /projects/{projectId}/collections/{collectionId}/binary/{recordId}/content`
- Raw operationId: `getBinaryByRecordId`
- Response content type: `application/octet-stream`

## Shared Rules

- Binary responses are streaming/binary responses, not JSON documents.
- The `field` query parameter stays wire-exact and identifies the binary field.
- Search-based binary retrieval must preserve ordering and selection semantics such as `selectedIndex` and stream selection where defined by `API-SPEC.md`.
