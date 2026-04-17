# searchWebApi Skeleton Design

## Context

This document defines the initial repository setup for a new `searchWebApi` project under `projects/searchWebApi/`.

The project is spec-only for now. Its detailed API contract will come later from an OpenAPI document supplied by the user. No generated client or server code is in scope for this step.

## Goals

- Add `searchWebApi` as a first-class project in the repository.
- Create a minimal `projects/searchWebApi/specs/` skeleton that can become the home for the future OpenAPI-derived spec set.
- Keep the initial structure consistent with the repository's existing project conventions.
- Avoid inventing placeholder API contracts before the upstream OpenAPI file exists.

## Non-Goals

- Writing the actual `searchWebApi` endpoint, schema, or auth contract.
- Generating implementation code under `~/ai-generated`.
- Pre-creating an ADP-sized multi-file spec tree without source material.
- Defining language-specific bindings, CLIs, or runtime behavior.

## Recommended Approach

Create a minimal project skeleton now and defer detailed spec structure until the OpenAPI document is available.

Why:

- It matches the user's requested scope exactly.
- It avoids placeholder files that would be guesses rather than authoritative specs.
- It keeps the repository clean while still reserving the correct project location.
- It makes the next step straightforward once the OpenAPI file is uploaded.

Alternatives considered:

1. Pre-create an ADP-style multi-file spec tree.
   This would mirror the existing project more closely, but most files would be empty guesses until the OpenAPI document exists.
2. Add only a top-level project folder with no `specs/index.md`.
   This is smaller, but it does not establish an obvious source-of-truth entrypoint for the project.

## Project Layout

```text
projects/searchWebApi/
└── specs/
    └── index.md
```

## File Design

### `projects/searchWebApi/specs/index.md`

This file will:

- identify `searchWebApi` as the project name,
- state that it is the source of truth for the internal REST API specification,
- state that the concrete contract will be populated from a future OpenAPI file,
- and document that implementation code will be generated later under `~/ai-generated/searchWebApi/[language]/` rather than stored in this repository.

The file should stay intentionally short. It is a project entrypoint, not a guessed contract.

### `AGENTS.md`

Update the project reference table to add:

- project name: `searchWebApi`
- specs index path: `projects/searchWebApi/specs/index.md`

No other workflow changes are needed. The existing repository-wide guidance already matches a spec-first project that later generates code outside the repository.

## Error Handling

This setup step has very little runtime behavior. The main failure mode is structural drift.

To avoid that:

- the new project path should follow the existing `projects/[project]/specs/index.md` convention,
- the new `index.md` should avoid references to files that do not exist yet,
- and `AGENTS.md` should point to the correct spec entrypoint.

## Testing Expectations

Verification for this change is structural rather than executable:

- confirm that `projects/searchWebApi/specs/index.md` exists,
- confirm that `AGENTS.md` includes the new project table entry,
- and confirm that no extra placeholder contract files were added.

## Success Criteria

This design is successful when:

- `searchWebApi` exists as a new project under `projects/`,
- the project has a minimal `specs/index.md` entrypoint,
- `AGENTS.md` lists the project correctly,
- and the repository contains no guessed API details ahead of the future OpenAPI import.
