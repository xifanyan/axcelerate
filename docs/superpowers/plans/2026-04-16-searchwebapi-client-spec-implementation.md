# searchWebApi Client Spec Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Build a complete language-agnostic markdown spec set for the `searchWebApi` REST client under `projects/searchWebApi/specs/`, using `API-SPEC.md` as the authoritative raw contract.

**Architecture:** Keep `API-SPEC.md` untouched and derive a layered client-spec tree from it. Update `index.md` into the project entrypoint, add shared transport/auth/request/type specs, and add resource-oriented spec files for each major workflow area. Where the OpenAPI source has awkward raw details such as duplicate operation IDs or overlapping parameter/schema names, normalize them in the markdown specs through explicit mapping rules instead of copying them blindly.

**Tech Stack:** Markdown, OpenAPI source in `projects/searchWebApi/API-SPEC.md`, `git`, `rg`

---

## File Map

- Modify: `projects/searchWebApi/specs/index.md`
  - Replace the temporary placeholder with the authoritative spec entrypoint and file map.
- Create: `projects/searchWebApi/specs/api-contract.md`
  - Define source-of-truth rules, path inventory, content types, and operation-to-resource mapping guidance.
- Create: `projects/searchWebApi/specs/transport.md`
  - Define base path, HTTP method rules, request execution, headers, and streaming behavior.
- Create: `projects/searchWebApi/specs/auth-and-session.md`
  - Define authentication, session lifecycle, session headers, and tracing headers.
- Create: `projects/searchWebApi/specs/request-conventions.md`
  - Define shared parameter semantics, common query patterns, request-body conventions, and normalization rules.
- Create: `projects/searchWebApi/specs/common-types.md`
  - Define reusable result families, status handling, shared schema conventions, and naming rules.
- Create: `projects/searchWebApi/specs/resources/projects.md`
  - Cover project discovery and project-level resources.
- Create: `projects/searchWebApi/specs/resources/collections.md`
  - Cover collection discovery, fields, filters, and folder values.
- Create: `projects/searchWebApi/specs/resources/records.md`
  - Cover record search, highlighting, record resources, record content fetch, record changes, and in-document search.
- Create: `projects/searchWebApi/specs/resources/binary.md`
  - Cover binary retrieval by search and by record id.
- Create: `projects/searchWebApi/specs/resources/measures.md`
  - Cover measure cube aggregation.
- Create: `projects/searchWebApi/specs/resources/cached-searches.md`
  - Cover cached search listing and dropping.
- Create: `projects/searchWebApi/specs/resources/change-queue.md`
  - Cover wait-for-pending-changes behavior.
- Create: `projects/searchWebApi/specs/resources/insert-remove.md`
  - Cover insert/remove operations and bulk transaction lifecycle.
- Create: `projects/searchWebApi/specs/resources/session.md`
  - Cover login/logout operations and their client exposure.
- Create: `projects/searchWebApi/specs/language-bindings/README.md`
  - State the purpose and boundaries of future language-binding docs.

---

### Task 1: Replace The Placeholder Index With The Real Spec Entry Point

**Files:**
- Modify: `projects/searchWebApi/specs/index.md`

- [ ] **Step 1: Confirm the current placeholder content before replacing it**

Run:

```bash
git diff -- projects/searchWebApi/specs/index.md
```

Expected: no output if the file has not been edited since the skeleton was added.

- [ ] **Step 2: Replace `projects/searchWebApi/specs/index.md` with this content**

```md
# searchWebApi Project Specifications

Single source of truth for the language-agnostic `searchWebApi` REST client specification.

---

## Source Of Truth

`API-SPEC.md` is the authoritative raw contract for:

- paths,
- HTTP methods,
- request parameters,
- headers,
- request bodies,
- response bodies,
- content types,
- and schema component definitions.

The markdown specs in `specs/` are authoritative for:

- resource-centric client organization,
- language-agnostic naming guidance,
- transport and session behavior as exposed by generated clients,
- shared request and response interpretation rules,
- and generator-facing client ergonomics.

If a markdown spec conflicts with `API-SPEC.md`, `API-SPEC.md` wins.

---

## Spec Structure

| File | Description |
|------|-------------|
| [index.md](./index.md) | This file - global rules and spec map |
| [api-contract.md](./api-contract.md) | Raw contract boundaries, path inventory, and operation mapping rules |
| [transport.md](./transport.md) | HTTP transport, headers, content negotiation, and execution rules |
| [auth-and-session.md](./auth-and-session.md) | Authentication and stateful session behavior |
| [request-conventions.md](./request-conventions.md) | Shared parameter, request-body, and request-shape rules |
| [common-types.md](./common-types.md) | Shared result families, status handling, and naming guidance |
| [resources/projects.md](./resources/projects.md) | Project discovery and project resource access |
| [resources/collections.md](./resources/collections.md) | Collection discovery, fields, filters, and folder values |
| [resources/records.md](./resources/records.md) | Record search, fetch, change, highlighting, and in-document search |
| [resources/binary.md](./resources/binary.md) | Binary retrieval operations |
| [resources/measures.md](./resources/measures.md) | Measure cube operations |
| [resources/cached-searches.md](./resources/cached-searches.md) | Cached search inspection and deletion |
| [resources/change-queue.md](./resources/change-queue.md) | Waiting for queued changes |
| [resources/insert-remove.md](./resources/insert-remove.md) | Insert/remove and bulk transaction workflows |
| [resources/session.md](./resources/session.md) | Login and logout operations |
| [language-bindings/README.md](./language-bindings/README.md) | Rules for future language-specific bindings |

---

## Global Rules

### Client Shape

- The client contract is resource-centric, not tag-centric.
- A root client owns transport, authentication, session state, and access to nested resources.
- Resource specs define conceptual client areas without prescribing one language's exact syntax.

### Wire Names Versus Client Names

- Wire-level path, query, header, and schema names must match `API-SPEC.md` exactly.
- Generated client names may be language-idiomatic, but each spec must keep traceability back to the raw OpenAPI names.
- Duplicate or ambiguous raw names from `API-SPEC.md` must be normalized in markdown guidance instead of copied directly into the client surface.

### Implementation Placement

This project is spec-only in this repository.

Generated client implementations live outside this repository under `~/ai-generated/searchWebApi/[language]/`.
```

- [ ] **Step 3: Verify the new index includes the complete file map**

Run:

```bash
rg -n "api-contract|transport|auth-and-session|request-conventions|common-types|resources/records|language-bindings" projects/searchWebApi/specs/index.md
```

Expected: matches for all shared spec files, at least one resource spec entry, and the language-bindings entry.

- [ ] **Step 4: Commit the index rewrite**

```bash
git add projects/searchWebApi/specs/index.md
git commit -m "specs: define searchWebApi spec index"
```

---

### Task 2: Add The Shared Contract And Transport Specs

**Files:**
- Create: `projects/searchWebApi/specs/api-contract.md`
- Create: `projects/searchWebApi/specs/transport.md`
- Create: `projects/searchWebApi/specs/auth-and-session.md`
- Create: `projects/searchWebApi/specs/request-conventions.md`
- Create: `projects/searchWebApi/specs/common-types.md`

- [ ] **Step 1: Create `projects/searchWebApi/specs/api-contract.md`**

Write this file:

```md
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
```

- [ ] **Step 2: Create `projects/searchWebApi/specs/transport.md`**

Write this file:

```md
# searchWebApi Transport

## Base URL

- Default base path: `/searchWebApi`
- Client configuration must allow callers to supply scheme, host, port, and optional path prefix around this base path.

## Supported Content Types

- `application/json` for standard request and response bodies
- `application/x-www-form-urlencoded` as an alternative encoding for eligible read-oriented operations documented in `API-SPEC.md`
- `multipart/form-data` for operations that upload binary or streamed content
- `application/octet-stream` for binary download responses
- `application/x-ndjson` for streaming search results on the records search endpoint

## HTTP Method Rules

- Use the exact HTTP method defined by `API-SPEC.md`.
- If `API-SPEC.md` documents form-encoded POST as an alternative to GET for an eligible operation, the client may expose that as a transport option without changing the semantic operation name.

## Headers

- Clients must support standard authentication headers.
- Clients must support `SWA-SESSION` for stateful session reuse.
- Clients must support optional `SWA-SESSION-TYPE`.
- Clients must support optional MDC tracing headers `SWA-MDC-TOKEN` and `SWA-MDC-METHOD`.
- Clients must allow callers to request NDJSON via the `Accept` header where supported.

## Streaming Rules

- Binary endpoints return `application/octet-stream` and should be represented as streaming/binary responses in generated clients.
- The records search endpoint may return `application/x-ndjson`; in that mode, the first record is search metadata and later lines are records.
- NDJSON mode bypasses normal search cache control behavior where the raw contract says so.
```

- [ ] **Step 3: Create `projects/searchWebApi/specs/auth-and-session.md`**

Write this file:

```md
# searchWebApi Authentication And Session

## Authentication Modes

- HTTP basic authentication
- Bearer-token authentication

## Session Model

The API is stateful.

- A successful login creates a session.
- The session identifier is returned in response header `SWA-SESSION`.
- Subsequent requests must send `SWA-SESSION` to reuse the same session.
- If credentials or a bearer token are present, non-session endpoints may implicitly create a new session after timeout.

## Explicit Session Operations

- `POST /login` creates a new session explicitly.
- `DELETE /logout` closes the current session and associated resources.

## Session Type

- Clients may send `SWA-SESSION-TYPE`.
- Documented values include `MONITORING`, `USER`, and `DEFAULT`.
- Session type is a transport concern and should be configurable at the root client level.

## Tracing Headers

- Clients may send `SWA-MDC-TOKEN` to propagate trace identifiers.
- Clients may send `SWA-MDC-METHOD` to propagate logical method names.

## Client Responsibilities

- Preserve session state between requests when the caller chooses session reuse.
- Expose enough control for callers to clear, replace, or ignore session state.
- Treat session expiry as a first-class protocol condition, not as a hidden implementation detail.
```

- [ ] **Step 4: Create `projects/searchWebApi/specs/request-conventions.md`**

Write this file:

```md
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
```

- [ ] **Step 5: Create `projects/searchWebApi/specs/common-types.md`**

Write this file:

```md
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
```

- [ ] **Step 6: Verify the shared spec files reference the known transport and naming edge cases**

Run:

```bash
rg -n "SWA-SESSION|application/x-ndjson|duplicate raw `operationId`|valueObject|multipart" projects/searchWebApi/specs
```

Expected: matches across the new shared spec files for session state, NDJSON, operation-id normalization, `valueObject`, and multipart handling.

- [ ] **Step 7: Commit the shared specs**

```bash
git add projects/searchWebApi/specs/api-contract.md projects/searchWebApi/specs/transport.md projects/searchWebApi/specs/auth-and-session.md projects/searchWebApi/specs/request-conventions.md projects/searchWebApi/specs/common-types.md
git commit -m "specs: add shared searchWebApi client rules"
```

---

### Task 3: Add Project And Collection Resource Specs

**Files:**
- Create: `projects/searchWebApi/specs/resources/projects.md`
- Create: `projects/searchWebApi/specs/resources/collections.md`

- [ ] **Step 1: Create `projects/searchWebApi/specs/resources/projects.md`**

Write this file:

```md
# Projects Resource

## Purpose

This resource area covers project discovery and project-level resource discovery.

## Client Shape

- Root client exposes project discovery.
- Project-scoped access starts with a `projectId` returned by the discovery endpoints.

## Operations

### List Projects

- Raw operation: `GET /projects`
- Raw operationId: `getProjects`
- Result schema: `ProjectsResult`
- Purpose: discover all available projects in the installation.

### Get Project Resources

- Raw operation: `GET /projects/{projectId}`
- Raw operationId: `getProjectResources`
- Result schema: `ProjectResourcesResult`
- Purpose: discover project-level resources for a selected project.

## Shared Rules

- `projectId` is the wire-level identifier and must not be renamed in request construction.
- Project discovery is read-oriented and follows the normal session and authentication rules.
```

- [ ] **Step 2: Create `projects/searchWebApi/specs/resources/collections.md`**

Write this file:

```md
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
```

- [ ] **Step 3: Verify project and collection resource coverage**

Run:

```bash
rg -n "getProjects|getProjectResources|getCollections|getCollectionResources|getFields|getFolderFields|getFolderFieldResources|getFolderValues" projects/searchWebApi/specs/resources
```

Expected: one match for each operationId in either `projects.md` or `collections.md`.

- [ ] **Step 4: Commit the project and collection specs**

```bash
git add projects/searchWebApi/specs/resources/projects.md projects/searchWebApi/specs/resources/collections.md
git commit -m "specs: define searchWebApi project and collection resources"
```

---

### Task 4: Add Records And Binary Resource Specs

**Files:**
- Create: `projects/searchWebApi/specs/resources/records.md`
- Create: `projects/searchWebApi/specs/resources/binary.md`

- [ ] **Step 1: Create `projects/searchWebApi/specs/resources/records.md`**

Write this file:

```md
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
```

- [ ] **Step 2: Create `projects/searchWebApi/specs/resources/binary.md`**

Write this file:

```md
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
```

- [ ] **Step 3: Verify records and binary coverage including normalized duplicate names**

Run:

```bash
rg -n "getRecords|changeAllInSearchResult|getSearchResultToken|deleteSearchResultToken|touchSearchResultToken|createSortOrderSnapshot|getHighlightingForSearchResult|getRecordResources|fetchRecordContent|changeRecordContent|searchInDocumentText|changeRecordInDocumentContext|getBinary|getBinaryByRecordId" projects/searchWebApi/specs/resources
```

Expected: matches for all listed raw or normalized operation names.

- [ ] **Step 4: Commit the records and binary specs**

```bash
git add projects/searchWebApi/specs/resources/records.md projects/searchWebApi/specs/resources/binary.md
git commit -m "specs: define searchWebApi record and binary resources"
```

---

### Task 5: Add Measures, Cached Searches, Change Queue, And Session Specs

**Files:**
- Create: `projects/searchWebApi/specs/resources/measures.md`
- Create: `projects/searchWebApi/specs/resources/cached-searches.md`
- Create: `projects/searchWebApi/specs/resources/change-queue.md`
- Create: `projects/searchWebApi/specs/resources/session.md`

- [ ] **Step 1: Create `projects/searchWebApi/specs/resources/measures.md`**

Write this file:

```md
# Measures Resource

## Operation

### Get Measure Cube

- Raw operation: `POST /projects/{projectId}/collections/{collectionId}/measures`
- Raw operationId: `getMeasureCube`
- Request schema: array of `DimensionRequest`
- Result schema: `MeasureCube`

## Shared Rules

- Measure requests may contain zero, one, or two dimension requests.
- Measure behavior depends on `MeasureTypeParameter` and must be documented without collapsing count and aggregate cases into one simplified description.
```

- [ ] **Step 2: Create `projects/searchWebApi/specs/resources/cached-searches.md`**

Write this file:

```md
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
```

- [ ] **Step 3: Create `projects/searchWebApi/specs/resources/change-queue.md`**

Write this file:

```md
# Change Queue Resource

## Operation

### Wait For Pending Changes

- Raw operation: `GET /projects/{projectId}/collections/{collectionId}/changes/queue`
- Raw operationId: `waitForAllCurrentlyScheduledChangeRequests`
- Result schema: `WaitForPendingChangesResult`

## Shared Rules

- This operation waits for changes that are queued or in progress at the time of submission.
- `timeoutMillis` and `onlyHighPriorityChanges` stay wire-exact.
- The result must preserve the backend's distinction between request success and queue-drain completion.
```

- [ ] **Step 4: Create `projects/searchWebApi/specs/resources/session.md`**

Write this file:

```md
# Session Resource

## Operations

### Login

- Raw operation: `POST /login`
- Raw operationId: `login`
- Result schema: `LoginResult`

### Logout

- Raw operation: `DELETE /logout`
- Raw operationId: `logout`
- Result schema: `LogoutResult`

## Shared Rules

- Explicit login/logout complement, but do not replace, the implicit session behavior described in `../auth-and-session.md`.
- Session identifiers are transported via `SWA-SESSION` headers rather than JSON body fields.
```

- [ ] **Step 5: Verify coverage of the smaller resource groups**

Run:

```bash
rg -n "getMeasureCube|getCachedSearches|dropCachedSearches|waitForAllCurrentlyScheduledChangeRequests|login|logout" projects/searchWebApi/specs/resources
```

Expected: one match for each operation name across the new resource files.

- [ ] **Step 6: Commit the measures, cache, queue, and session specs**

```bash
git add projects/searchWebApi/specs/resources/measures.md projects/searchWebApi/specs/resources/cached-searches.md projects/searchWebApi/specs/resources/change-queue.md projects/searchWebApi/specs/resources/session.md
git commit -m "specs: add remaining searchWebApi resource groups"
```

---

### Task 6: Add Insert/Remove And Language-Binding Boundary Specs

**Files:**
- Create: `projects/searchWebApi/specs/resources/insert-remove.md`
- Create: `projects/searchWebApi/specs/language-bindings/README.md`

- [ ] **Step 1: Create `projects/searchWebApi/specs/resources/insert-remove.md`**

Write this file:

```md
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
```

- [ ] **Step 2: Create `projects/searchWebApi/specs/language-bindings/README.md`**

Write this file:

```md
# searchWebApi Language Bindings

This directory is reserved for future language-specific bindings.

The core markdown specs in the parent directory remain language-agnostic.

Language bindings may define:

- idiomatic client naming,
- package or module structure,
- type mappings,
- streaming abstractions,
- and session-state handling conventions for a specific language.

Language bindings must not change wire-level paths, parameters, headers, request bodies, or response shapes defined by `../API-SPEC.md` and the core specs.
```

- [ ] **Step 3: Verify the bulk workflow and language-binding boundary are documented**

Run:

```bash
rg -n "start buffer|add payloads|end transaction|poll job status|language-agnostic|wire-level" projects/searchWebApi/specs
```

Expected: matches in `resources/insert-remove.md` and `language-bindings/README.md`.

- [ ] **Step 4: Commit the insert/remove and language-binding specs**

```bash
git add projects/searchWebApi/specs/resources/insert-remove.md projects/searchWebApi/specs/language-bindings/README.md
git commit -m "specs: add searchWebApi bulk transaction rules"
```

---

### Task 7: Final Coverage And Consistency Review

**Files:**
- Verify: `projects/searchWebApi/specs/index.md`
- Verify: `projects/searchWebApi/specs/api-contract.md`
- Verify: `projects/searchWebApi/specs/transport.md`
- Verify: `projects/searchWebApi/specs/auth-and-session.md`
- Verify: `projects/searchWebApi/specs/request-conventions.md`
- Verify: `projects/searchWebApi/specs/common-types.md`
- Verify: `projects/searchWebApi/specs/resources/*.md`
- Verify: `projects/searchWebApi/specs/language-bindings/README.md`

- [ ] **Step 1: Confirm every raw operationId is represented in the markdown spec set**

Run:

```bash
rg -n "getProjects|getProjectResources|getCollections|getCollectionResources|getFields|getFolderFields|getFolderFieldResources|getFolderValues|getRecords|changeAllInSearchResult|getSearchResultToken|deleteSearchResultToken|touchSearchResultToken|createSortOrderSnapshot|getHighlightingForSearchResult|getRecordResources|fetch|change|searchInDocumentText|getBinary|getBinaryByRecordId|getMeasureCube|insertRemoveTransaction|startInsertRemoveTransaction|commitInsertRemoveTransaction|getFlushJobStatus|addToInsertRemoveTransaction|getCachedSearches|dropCachedSearches|waitForAllCurrentlyScheduledChangeRequests|login|logout" projects/searchWebApi/specs
```

Expected: every listed raw operationId appears somewhere in the spec set, and the duplicate raw `change` operations are accompanied by normalized client-facing names in `resources/records.md`.

- [ ] **Step 2: Confirm the normalized duplicate-operation guidance exists**

Run:

```bash
rg -n "changeRecordContent|changeRecordInDocumentContext|duplicate raw `operationId`" projects/searchWebApi/specs
```

Expected: matches in `request-conventions.md`, `api-contract.md`, and `resources/records.md`.

- [ ] **Step 3: Check the spec tree for formatting issues**

Run:

```bash
git diff --check -- projects/searchWebApi/specs
```

Expected: no output.

- [ ] **Step 4: Review the final diff for only the intended `searchWebApi` spec additions**

Run:

```bash
git diff -- projects/searchWebApi/specs
```

Expected: only the planned spec files under `projects/searchWebApi/specs/` are changed.

- [ ] **Step 5: Commit the final review pass**

```bash
git add projects/searchWebApi/specs
git commit -m "specs: complete searchWebApi client specification"
```
