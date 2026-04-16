# searchWebApi API Contract

This file defines how the markdown client specs relate to `../API-SPEC.md`.

## Raw Contract Authority

`../API-SPEC.md` is authoritative for:

- path strings,
- HTTP methods,
- path/query/header parameters,
- request body schemas,
- response body schemas,
- content types,
- and component definitions under `schemas` and `parameters`.

## Path Inventory

The raw contract currently covers these path families:

- `/projects`
- `/projects/{projectId}`
- `/projects/{projectId}/collections`
- `/projects/{projectId}/collections/{collectionId}`
- `/projects/{projectId}/collections/{collectionId}/fields`
- `/projects/{projectId}/collections/{collectionId}/filters`
- `/projects/{projectId}/collections/{collectionId}/filters/{fieldId}`
- `/projects/{projectId}/collections/{collectionId}/filters/{fieldId}/values`
- `/projects/{projectId}/collections/{collectionId}/records`
- `/projects/{projectId}/collections/{collectionId}/searchToken`
- `/projects/{projectId}/collections/{collectionId}/searchToken/sortOrderSnapshot`
- `/projects/{projectId}/collections/{collectionId}/search/highlightExpression`
- `/projects/{projectId}/collections/{collectionId}/records/{recordId}`
- `/projects/{projectId}/collections/{collectionId}/records/{recordId}/content`
- `/projects/{projectId}/collections/{collectionId}/records/{recordId}/inDocumentSearch`
- `/projects/{projectId}/collections/{collectionId}/binary`
- `/projects/{projectId}/collections/{collectionId}/binary/{recordId}/content`
- `/projects/{projectId}/collections/{collectionId}/measures`
- `/projects/{projectId}/collections/{collectionId}/records/insertRemoveTransaction`
- `/projects/{projectId}/collections/{collectionId}/records/bulkInsertRemoveTransaction`
- `/projects/{projectId}/collections/{collectionId}/records/bulkInsertRemoveTransaction/{indexingBufferId}/end`
- `/projects/{projectId}/collections/{collectionId}/records/bulkInsertRemoveTransaction/{indexingBufferId}/end/{jobId}`
- `/projects/{projectId}/collections/{collectionId}/records/bulkInsertRemoveTransaction/{indexingBufferId}/buffer`
- `/projects/{projectId}/collections/{collectionId}/cachedSearches`
- `/projects/{projectId}/collections/{collectionId}/changes/queue`
- `/login`
- `/logout`

## Resource Mapping Rules

- Project and project-resource discovery belongs in `resources/projects.md`.
- Collection discovery, fields, filters, and folder values belong in `resources/collections.md`.
- Search, record fetch, highlighting, record changes, and in-document search belong in `resources/records.md`.
- Binary retrieval belongs in `resources/binary.md`.
- Measure cube aggregation belongs in `resources/measures.md`.
- Insert/remove and bulk transaction operations belong in `resources/insert-remove.md`.
- Cached search operations belong in `resources/cached-searches.md`.
- Change queue waiting belongs in `resources/change-queue.md`.
- Login and logout belong in `resources/session.md`.

## Normalization Rules

- Markdown specs must normalize duplicate raw `operationId` values into distinct client-facing operation names.
- Markdown specs must distinguish OpenAPI component schemas named like `QueryParameter` from parameter components with the same or similar names.
- Wire-level names remain unchanged even when client-facing names are clarified.
