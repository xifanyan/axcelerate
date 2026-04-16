# Records Resource

## Purpose

This resource area covers search, search-result token workflows, search highlighting, record resource discovery, record content fetch, record changes, and in-document search.

## Operations

### Search Records

- Raw operation: `GET /projects/{projectId}/collections/{collectionId}/records`
- Raw operationId: `getRecords`
- Result schema: `SearchResult`

### Change All Records In Search Result

- Raw operation: `PUT /projects/{projectId}/collections/{collectionId}/records`
- Raw operationId: `changeAllInSearchResult`
- Result schema: `ChangeResult`

### Create Search Token

- Raw operation: `GET /projects/{projectId}/collections/{collectionId}/searchToken`
- Raw operationId: `getSearchResultToken`
- Result schema: `SearchResultTokenResponse`

### Delete Search Token

- Raw operation: `DELETE /projects/{projectId}/collections/{collectionId}/searchToken`
- Raw operationId: `deleteSearchResultToken`
- Request schema: `SearchResultToken`

### Touch Search Token

- Raw operation: `PUT /projects/{projectId}/collections/{collectionId}/searchToken`
- Raw operationId: `touchSearchResultToken`
- Result schema: `SearchResultTokenResponse`

### Create Sort Order Snapshot

- Raw operation: `POST /projects/{projectId}/collections/{collectionId}/searchToken/sortOrderSnapshot`
- Raw operationId: `createSortOrderSnapshot`
- Request schema: `SearchResultToken`

### Get Search Highlight Expressions

- Raw operation: `GET /projects/{projectId}/collections/{collectionId}/search/highlightExpression`
- Raw operationId: `getHighlightingForSearchResult`
- Result schema: `SearchResultHighlightingResult`

### Get Record Resources

- Raw operation: `GET /projects/{projectId}/collections/{collectionId}/records/{recordId}`
- Raw operationId: `getRecordResources`
- Result schema: `RecordResourcesResult`

### Fetch Record Content

- Raw operation: `GET /projects/{projectId}/collections/{collectionId}/records/{recordId}/content`
- Raw operationId: `fetch`
- Result schema: `Record`
- Client-facing normalized name: `fetchRecordContent`

### Change One Record

- Raw operation: `PUT /projects/{projectId}/collections/{collectionId}/records/{recordId}/content`
- Raw operationId: `change`
- Result schema: `ChangeResult`
- Client-facing normalized name: `changeRecordContent`

### Search In Document Text

- Raw operation: `GET /projects/{projectId}/collections/{collectionId}/records/{recordId}/inDocumentSearch`
- Raw operationId: `searchInDocumentText`
- Result schema: `HighlightedWordResult`

### Change One Record Through In-Document Route

- Raw operation: `PUT /projects/{projectId}/collections/{collectionId}/records/{recordId}/inDocumentSearch`
- Raw operationId: `change`
- Result schema: `ChangeResult`
- Client-facing normalized name: `changeRecordInDocumentContext`

## Shared Rules

- NDJSON streaming applies to the records search endpoint only where requested by the caller.
- Search token expiry and `410` semantics must be preserved in client behavior.
- Duplicate raw `operationId` values must not become duplicate client-surface names.
