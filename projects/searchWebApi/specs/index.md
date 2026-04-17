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
| [language-bindings/README.md](./language-bindings/README.md) | Rules for language-specific bindings |
| [language-bindings/go.md](./language-bindings/go.md) | Go CLI binding for the searchWebApi surface |

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
