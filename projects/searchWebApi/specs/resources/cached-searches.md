# Cached Searches Resource

## Operations

### List Cached Searches

- Raw operation: `GET /projects/{projectId}/collections/{collectionId}/cachedSearches`
- Raw operationId: `getCachedSearches`
- Result schema: array of `CachedSearchDescription`

### Drop Cached Searches

- Raw operation: `DELETE /projects/{projectId}/collections/{collectionId}/cachedSearches`
- Raw operationId: `dropCachedSearches`
- Result schema: array of `CachedSearchDescription`

## Shared Rules

- Cache inspection and deletion are session-scoped behaviors.
- `creationTraceIds` must stay a wire-level query parameter name.
