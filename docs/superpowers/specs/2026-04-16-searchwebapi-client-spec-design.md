# searchWebApi Client Spec Design

## Context

This document defines how to turn `projects/searchWebApi/API-SPEC.md` into a language-agnostic markdown spec set for a generated REST API client.

`API-SPEC.md` is the authoritative upstream contract. It currently contains the full Search Web API surface in OpenAPI form, including 33 operations across project discovery, collection discovery, search, record access, binary access, change operations, insert/remove transactions, cached searches, queue control, and session lifecycle.

The repository now needs a client-oriented spec layer similar in role to `adp`, but adapted to a stateful HTTP REST API rather than task execution endpoints.

## Goals

- Define a language-agnostic markdown spec set for the full `searchWebApi` client surface.
- Keep `API-SPEC.md` as the source of truth for raw endpoint, parameter, header, and schema details.
- Describe a resource-centric client shape rather than a flat list of HTTP operations.
- Cover transport, authentication, session state, request conventions, response models, and operation grouping in a way that multiple language generators can follow.
- Preserve traceability from markdown specs back to the raw OpenAPI source.

## Non-Goals

- Replacing the OpenAPI file as the raw source of truth.
- Generating implementation code in this design step.
- Committing to one language's class, interface, or package conventions.
- Inventing behavior not present in `API-SPEC.md`.
- Flattening every operation into one giant spec file.

## Recommended Approach

Create a layered markdown spec set with a root client contract, shared transport/session rules, shared types, and resource-focused operation specs.

Why:

- It matches the desired client abstraction: a root client with nested resource helpers.
- It keeps HTTP details explicit without forcing endpoint-by-endpoint duplication into the main index.
- It gives generators a stable conceptual model across languages while leaving naming and packaging idiomatic.
- It keeps `API-SPEC.md` authoritative for low-level facts and uses markdown specs for client structure and derived rules.

Alternatives considered:

1. Direct tag-by-tag OpenAPI translation.
   This is easier to derive mechanically, but it produces a flatter and less ergonomic client contract.
2. Thin transport-only contract.
   This preserves maximum freedom, but it leaves too many client-shape decisions implicit for consistent generated libraries.

## Spec Layout

The spec set should live under `projects/searchWebApi/specs/` and use a structure like this:

```text
projects/searchWebApi/specs/
├── index.md
├── api-contract.md
├── transport.md
├── auth-and-session.md
├── request-conventions.md
├── common-types.md
├── resources/
│   ├── projects.md
│   ├── collections.md
│   ├── records.md
│   ├── binary.md
│   ├── measures.md
│   ├── cached-searches.md
│   ├── change-queue.md
│   ├── insert-remove.md
│   └── session.md
└── language-bindings/
    └── README.md
```

This is a client-spec layout, not a generated code layout.

## Source-Of-Truth Rules

### Raw Contract Authority

`API-SPEC.md` remains authoritative for:

- path strings,
- HTTP methods,
- query/path/header parameters,
- request body shapes,
- response body shapes,
- content types,
- and schema component definitions.

### Markdown Spec Authority

The markdown spec set becomes authoritative for:

- resource-centric client organization,
- operation grouping,
- language-agnostic naming guidance,
- session and transport behavior as exposed by the client,
- shared request/response interpretation rules,
- and generator-facing expectations for client ergonomics.

If a markdown spec conflicts with the raw OpenAPI facts, the OpenAPI facts must win and the markdown spec must be corrected.

## Client Architecture

### Root Client

The spec should define one top-level client responsible for:

- base URL configuration,
- authentication configuration,
- session header management,
- request execution,
- shared error interpretation,
- and access to nested resource areas.

The root client should expose resource-centric navigation conceptually similar to:

- projects discovery,
- project-scoped collection access,
- collection-scoped search and change operations,
- record-scoped fetch and change operations,
- binary retrieval,
- and session lifecycle operations.

The markdown specs should describe these concepts without prescribing exact class or method syntax for any one language.

### Resource Boundaries

Resource specs should group operations by client-facing workflow rather than by raw tag alone.

- `projects.md` covers `/projects` and `/projects/{projectId}` discovery.
- `collections.md` covers collection discovery, field discovery, filter discovery, folder values, and collection resource metadata.
- `records.md` covers record search, highlighting, record fetch, record-level changes, and in-document search.
- `binary.md` covers binary retrieval by search and by record id.
- `measures.md` covers measure cube access.
- `insert-remove.md` covers insert/remove transactions, bulk transaction lifecycle, flush job status, and buffering.
- `cached-searches.md` covers listing and dropping cached searches.
- `change-queue.md` covers queue-drain waiting behavior.
- `session.md` covers login, logout, and shared session semantics.

This grouping keeps related workflows together even when the raw OpenAPI tags are broad.

## Shared Transport And Session Design

### Transport

The transport spec should define:

- default base path `/searchWebApi`,
- support for JSON requests and responses,
- support for `application/x-www-form-urlencoded` as an alternative to GET for eligible read-oriented endpoints,
- support for `multipart/form-data` where the OpenAPI contract allows binary uploads,
- support for `application/x-ndjson` streaming responses for record search,
- and explicit header behavior for authentication and state propagation.

### Authentication And Session

The auth/session specs should make the stateful nature of the API explicit.

They should define:

- login via HTTP basic auth or bearer token,
- implicit session creation when authenticated requests arrive after timeout,
- extraction and reuse of `SWA-SESSION`,
- optional `SWA-SESSION-TYPE`,
- optional MDC tracing headers such as `SWA-MDC-TOKEN` and `SWA-MDC-METHOD`,
- explicit logout behavior,
- and client expectations for preserving and updating session state between calls.

Because this API is stateful, session handling must be a first-class client concern rather than an incidental header note.

## Request And Response Rules

### Requests

The request-conventions spec should capture cross-cutting rules that appear repeatedly in the OpenAPI source:

- shared path parameters like `projectId`, `collectionId`, `recordId`, and folder field identifiers,
- reusable search query parameters like `query`, `language`, `joinRestriction`, `order`, paging, field selection, and cache control,
- change operations with array-based change request payloads,
- binary and multipart request handling,
- and token-based workflows such as search result token creation, renewal, and deletion.

### Responses

The common-types spec should define the reusable response families and interpretation rules for:

- discovery results,
- search results,
- records,
- folder values,
- change results,
- token responses,
- job-status results,
- highlighting responses,
- and default-response error handling where the API may still return a JSON body on failure.

The spec should stay language-agnostic by describing shapes in neutral notation and by separating raw schema identity from generated-type naming guidance.

## Naming And Language-Agnostic Rules

The markdown specs should define naming rules suitable for generated clients across languages:

- use resource-centric names rather than raw Swagger-style grouping names,
- preserve raw parameter and field names where they are API-facing,
- allow generated type names to be language-idiomatic while tracing back to the OpenAPI schema names,
- and separate client-surface names from HTTP wire-format names.

Like `adp`, the spec set should be explicit about what is wire-level truth versus what is generator-facing design guidance.

## Error Handling

The spec set should explicitly describe:

- operations with only a `default` response in OpenAPI,
- schema-backed error payloads when returned inside otherwise normal JSON responses,
- transport-level failures,
- session-expiry behavior,
- token-expiry behavior such as `410` for touched search tokens,
- and content-negotiation edge cases like NDJSON streaming versus normal JSON responses.

The client design should require callers to be able to inspect both structured API result bodies and transport/protocol failures.

## Testing Expectations

When this design is implemented, verification should cover:

- mapping every OpenAPI path to exactly one resource spec section,
- shared parameter definitions reused consistently across resource specs,
- stateful session handling behavior,
- GET-versus-form-encoded POST compatibility where documented,
- JSON versus multipart versus NDJSON handling,
- and traceability from markdown specs back to `API-SPEC.md`.

## Success Criteria

This design is successful when:

- `searchWebApi` has a complete language-agnostic client spec set under `projects/searchWebApi/specs/`,
- every operation in `API-SPEC.md` is represented in the markdown spec structure,
- the client contract is resource-centric rather than a flat endpoint dump,
- `API-SPEC.md` remains the authoritative raw contract,
- and future generators can implement clients in multiple languages without inventing the client shape from scratch.
