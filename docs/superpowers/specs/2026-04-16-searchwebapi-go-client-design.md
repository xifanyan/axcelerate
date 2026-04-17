# searchWebApi Go Client Design

## Context

This document defines the generated Go implementation for the `searchWebApi` specification in `projects/searchWebApi/`.

The repository already contains the language-agnostic client spec set and the raw OpenAPI contract in `projects/searchWebApi/API-SPEC.md`. The Go output will be generated outside the repository at `C:\Users\pyan\ai-generated\searchWebApi\go` as required by `AGENTS.md`.

The generated output should include:

- a reusable Go library for the full `searchWebApi` surface,
- a small example CLI in the same Go module,
- resource-centric client APIs that follow the repo spec layout,
- and transport/session behavior that preserves the wire protocol exactly.

## Goals

- Generate a working Go module at `C:\Users\pyan\ai-generated\searchWebApi\go`.
- Expose the API as a typed Go library with a root client and resource-focused methods.
- Preserve wire-level names for paths, query parameters, headers, and JSON fields.
- Support the full documented API surface: projects, collections, records, binary, measures, cached searches, change queue, insert/remove, and session.
- Support stateful session reuse via `SWA-SESSION`.
- Include a small example CLI that demonstrates basic configuration, login-compatible usage, project listing, and record search.

## Non-Goals

- Building a full-featured production CLI for every endpoint.
- Generating code inside this repository.
- Hiding protocol details behind a large abstraction layer.
- Re-shaping server JSON into heavily customized Go-only models.
- Adding speculative retries, caching, or background session refresh logic not required by the spec.

## Recommended Approach

Generate a small, explicit Go SDK with a single root `Client`, resource-focused files, shared request helpers, and mostly direct schema mappings from the API spec.

Why this approach:

- It matches the resource-centric design already established in `projects/searchWebApi/specs/`.
- It keeps the code size controlled while still giving callers an idiomatic Go surface.
- It minimizes spec drift because request construction and response decoding stay close to the wire format.
- It keeps session handling visible and testable rather than implicit magic.

Alternatives considered:

1. Thin raw HTTP wrapper.
   This would be quick to generate, but it would leave too much request construction, typing, and session handling to callers.
2. Rich builder-heavy SDK.
   This could be ergonomic for some workflows, but it would add a lot of extra code and many invented abstractions not required by the spec.

## Output Structure

The generated Go module should use this layout:

```text
C:\Users\pyan\ai-generated\searchWebApi\go
├── go.mod
├── README.md
├── client.go
├── transport.go
├── auth.go
├── types.go
├── projects.go
├── collections.go
├── records.go
├── binary.go
├── measures.go
├── cached_searches.go
├── change_queue.go
├── insert_remove.go
├── session.go
└── cmd/
    └── searchwebapi-example/
        └── main.go
```

File responsibilities:

- `client.go`: root client config, constructor, shared state, and resource entry points.
- `transport.go`: request creation, response handling, URL/query encoding, JSON decoding, and common header application.
- `auth.go`: basic auth, bearer auth, session handling, and optional transport-level headers.
- `types.go`: shared schema types and shared option structs.
- Resource files: endpoint-specific request/response methods grouped by the spec resource layout.
- `cmd/searchwebapi-example/main.go`: small demonstration CLI using the generated library.

## Client Architecture

### Root Client

The root `Client` will own:

- the normalized base URL,
- the underlying `http.Client`,
- authentication settings,
- optional session and tracing headers,
- and all common request execution behavior.

The client will be constructed from a config struct rather than many positional arguments. That config should support:

- base URL,
- username/password for basic auth,
- bearer token,
- optional fixed session id,
- optional session type,
- optional MDC token and MDC method,
- and a custom `http.Client`.

The client surface should be flat enough to stay simple. Methods can live directly on `Client` and be grouped by file, instead of inventing nested service structs unless a clear need appears during implementation.

### Resource Methods

Methods should be grouped by resource area in code and naming, for example:

- `ListProjects`
- `GetProjectResources`
- `ListCollections`
- `SearchRecords`
- `SearchRecordsStream`
- `FetchRecordContent`
- `GetBinary`
- `GetBinaryByRecordID`
- `Login`
- `Logout`

Where the raw API has duplicate or awkward names, the Go surface should normalize them while keeping traceability to the spec. For example:

- `FetchRecordContent` for raw operation `fetch`
- `ChangeRecordContent` for the `PUT .../content` variant of raw operation `change`
- `ChangeRecordInDocumentContext` for the `PUT .../inDocumentSearch` variant of raw operation `change`

## Transport Design

### Request Construction

Every operation should accept `context.Context` as the first parameter.

Request helpers should:

- build paths safely from path parameters,
- encode query parameters only when present,
- apply headers consistently,
- set `Accept` and `Content-Type` explicitly when needed,
- and decode JSON responses only when the endpoint is documented to return JSON.

The transport layer should not hide HTTP semantics completely. Binary and NDJSON operations should expose streaming results explicitly.

### Response Handling

Standard JSON endpoints should decode directly into typed result structs.

Binary endpoints should return a small response wrapper containing:

- response headers,
- the content type,
- an optional filename if derivable from response metadata,
- and an `io.ReadCloser` body.

NDJSON search should use a dedicated stream type with:

- one decoded initial metadata record,
- subsequent decoded `Record` values,
- and a close method for the underlying response body.

This is preferred over returning raw lines because callers still get typed records while the streaming behavior remains explicit.

## Authentication And Session Design

The Go client must support:

- HTTP basic auth,
- bearer auth,
- explicit `Login`,
- explicit `Logout`,
- and session reuse via `SWA-SESSION`.

Session behavior:

- If a response contains `SWA-SESSION`, the client should store that value for subsequent requests.
- If the caller provides a fixed session id, the client should send it.
- The client should expose accessors to read, set, and clear the stored session id.
- The client should not silently persist sessions to disk.

Concurrency rule:

- Session state should be guarded so the client can be used concurrently without data races.

## Type Design

The Go types should stay close to the raw schema names and field names.

Rules:

- Exported Go field names should be idiomatic, but JSON tags must preserve the wire-level field names.
- Shared types like `StatusObject`, `Record`, `Field`, `FolderSet`, `Folder`, `SearchResult`, `ChangeRequest`, `InsertRemoveRequest`, and `MeasureCube` should be defined once in `types.go`.
- `valueObject` should be represented in a way that preserves arbitrary JSON shape, most likely as `any`.
- Optional fields should generally use `omitempty` where appropriate for request payloads, but response decoding should tolerate missing values without adding special compatibility logic.

For request parameters, small option structs are preferred over long method signatures for complex endpoints. Simple endpoints can keep direct arguments for path identifiers plus a single options struct for optional query/header parameters.

## Endpoint Coverage

The generated library should cover all operations described by the current spec set:

- projects and project resources,
- collections, fields, filters, folder field resources, and folder values,
- records search, token lifecycle, sort order snapshot, highlight expressions, record resources, record content fetch, record changes, and in-document search,
- binary retrieval by search and record id,
- measure cube retrieval,
- insert/remove transaction workflows,
- cached search listing and deletion,
- change queue waiting,
- login and logout.

Multipart upload support is required for the endpoints that accept `multipart/form-data`, specifically record change and insert/remove workflows that reference uploaded binaries by index.

## Example CLI

The example CLI should be intentionally small and readable.

Scope:

- accept base URL and auth settings from flags and environment variables,
- build a `Client`,
- expose a `projects list` command,
- expose a `records search` command with `project`, `collection`, and `query` inputs,
- print JSON-like output suitable for manual inspection,
- and demonstrate session reuse naturally through the shared client.

The CLI is a usage example, not a second abstraction layer. It should call the library directly and avoid adding its own business logic.

## Error Handling

The generated code should preserve both classes of failure:

- transport/protocol failures such as request creation, network errors, invalid URLs, or non-decodable bodies,
- and successful HTTP exchanges whose decoded result contains `status.successful = false` or related backend error information.

The library should not automatically convert every backend status failure into a Go error if doing so would hide the structured result. A small helper error type for transport failures is fine, but decoded result bodies should remain inspectable by the caller.

For special cases:

- token touch `410` responses should remain distinguishable,
- NDJSON streaming errors should surface during iteration,
- and multipart request-building errors should fail before the request is sent.

## Verification Expectations

Implementation is complete when:

- the Go module builds successfully,
- the example CLI builds successfully,
- every documented endpoint has a corresponding library method,
- session header propagation is covered by tests,
- JSON, binary, NDJSON, and multipart flows are each covered by tests,
- and the README explains basic setup and CLI usage.

## Success Criteria

This design is successful when the generated Go output:

- lives at `C:\Users\pyan\ai-generated\searchWebApi\go`,
- provides a usable typed library for the full API surface,
- includes a small working example CLI,
- preserves the raw API protocol and schema semantics,
- and remains small enough that a caller can understand the transport and session model without reading the entire implementation.
