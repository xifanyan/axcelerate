# searchWebApi Request Conventions

## Shared Path Parameters

- `projectId`
- `collectionId`
- `fieldId` for folder-field routes
- `recordId`
- `indexingBufferId`
- `jobId`

These names must match the raw API exactly.

## Shared Search Parameters

Common search-oriented parameters include:

- `query`
- `language`
- `joinRestriction`
- `order`
- `page`
- `limit`
- `offset`
- `fields`
- `folderFields`
- `folderProperties`
- `body`
- `highlight`
- `sponsoredLinks`
- `spellingSuggestions`
- `SearchCacheControl`

Resource specs must define which of these apply to each operation without renaming their wire format.

## Request Body Conventions

- Record change operations accept arrays of `ChangeRequest`.
- Insert/remove operations accept `InsertRemoveRequest`.
- Measure operations accept arrays of `DimensionRequest`.
- Search token lifecycle operations use `SearchResultToken` payloads where defined.
- Bulk transaction operations use `StartTransactionRequest`, `FinishTransactionRequest`, and insert/remove payloads as defined in `API-SPEC.md`.

## Multipart Rules

- Multipart operations must document the relationship between the structured request object and the uploaded binary list.
- Where the raw contract uses binary indices inside field values, markdown specs must state that explicitly.

## Naming Normalization

- Client-facing operation names must be unique even when the raw `operationId` values are duplicated.
- Parameter component names and schema component names with overlapping wording must be described in separate sections to avoid client-generator confusion.
