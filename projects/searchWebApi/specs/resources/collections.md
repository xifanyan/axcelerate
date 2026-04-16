# Collections Resource

## Purpose

This resource area covers collection discovery, collection metadata, field metadata, filter metadata, and folder-value retrieval.

## Operations

### List Collections

- Raw operation: `GET /projects/{projectId}/collections`
- Raw operationId: `getCollections`
- Result schema: `CollectionsResult`

### Get Collection Resources

- Raw operation: `GET /projects/{projectId}/collections/{collectionId}`
- Raw operationId: `getCollectionResources`
- Result schema: `CollectionResourcesResult`

### Get Fields

- Raw operation: `GET /projects/{projectId}/collections/{collectionId}/fields`
- Raw operationId: `getFields`
- Result schema: `FieldsResult`

### Get Folder Fields

- Raw operation: `GET /projects/{projectId}/collections/{collectionId}/filters`
- Raw operationId: `getFolderFields`
- Result schema: `FieldsResult`

### Get Folder Field Resources

- Raw operation: `GET /projects/{projectId}/collections/{collectionId}/filters/{fieldId}`
- Raw operationId: `getFolderFieldResources`
- Result schema: `FolderFieldResourcesResult`

### Get Folder Values

- Raw operation: `GET /projects/{projectId}/collections/{collectionId}/filters/{fieldId}/values`
- Raw operationId: `getFolderValues`
- Result schema: `FolderValuesResult`

## Shared Rules

- `collectionId` and `fieldId` stay wire-exact.
- Folder-value requests reuse common search-style query parameters and add folder-specific paging and ordering rules.
- Folder collections and folder fields must be documented exactly as described by the raw field metadata.
