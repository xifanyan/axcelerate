# searchWebApi Common Types

## Core Status Shape

Most structured JSON results include a `status` object with:

- `successful`
- `backendStatus`
- `httpStatus`
- `errorMessage`

Clients must preserve this structure and allow callers to inspect it even when the transport-level HTTP exchange succeeded.

## Shared Result Families

- Discovery results: projects, project resources, collections, collection resources, record resources, folder-field resources
- Search results: records, sponsored links, spelling suggestions
- Folder-value results
- Highlighting results
- Token results
- Change results
- Insert/remove results
- Bulk transaction results and job-status results
- Session results
- Measure cube results

## Field And Record Shapes

- `Record` is the reusable record-level response shape for both search results and record fetch results.
- `Field`, `FolderSet`, and `Folder` are shared nested shapes.
- `value` and `valueObject` must be documented as wire-level fields, not collapsed into one synthetic concept.

## Naming Rules

- Preserve schema identity from `API-SPEC.md`.
- Describe client-friendly aliases only as guidance, never as replacements for the raw schema names.
- When raw schema names and raw parameter component names overlap conceptually, the markdown spec must call out the difference explicitly.

## Error Interpretation

- A successful HTTP exchange does not imply a successful backend operation.
- Callers must be able to inspect both transport-level failure and `status`-level failure.
- Resource specs may define additional operation-specific error conditions such as token expiry or invalid queued-change waits.
