# searchWebApi Skeleton Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Add `searchWebApi` as a new spec-only project with a minimal `projects/searchWebApi/specs/index.md` entrypoint and register it in `AGENTS.md`.

**Architecture:** Keep the change intentionally small. Create one new spec entrypoint file under `projects/searchWebApi/specs/` and update the repository's project reference table in `AGENTS.md`. Do not create any additional spec files or generated code until the user provides the OpenAPI document.

**Tech Stack:** Markdown, repository metadata in `AGENTS.md`, `git`, `rg`

---

## File Map

- Create: `projects/searchWebApi/specs/index.md`
  - Minimal source-of-truth entrypoint for the new internal REST API project.
- Modify: `AGENTS.md`
  - Add `searchWebApi` to the project reference table.

---

### Task 1: Create The Minimal searchWebApi Spec Entrypoint

**Files:**
- Create: `projects/searchWebApi/specs/index.md`

- [ ] **Step 1: Confirm the project path is not already in use**

Run:

```bash
rg --files projects/searchWebApi
```

Expected: no output because the project does not exist yet.

- [ ] **Step 2: Create `projects/searchWebApi/specs/index.md` with the exact starter content**

Add this file:

```md
# searchWebApi Project Specifications

Single source of truth for the `searchWebApi` internal REST API specification.

This project is spec-only in this repository.

The detailed API contract will be added from an OpenAPI document supplied later.

Implementation code is generated outside this repository under `~/ai-generated/searchWebApi/[language]/`.
```

- [ ] **Step 3: Verify the new file contains the intended scope markers**

Run:

```bash
rg -n "searchWebApi|spec-only|OpenAPI|ai-generated" projects/searchWebApi/specs/index.md
```

Expected: matches for the project name, the spec-only note, the OpenAPI note, and the external generation path.

- [ ] **Step 4: Commit the new spec entrypoint**

```bash
git add projects/searchWebApi/specs/index.md
git commit -m "docs: add searchWebApi spec entrypoint"
```

---

### Task 2: Register searchWebApi In Repository Metadata

**Files:**
- Modify: `AGENTS.md`

- [ ] **Step 1: Add the new project row to the project reference table**

Update the table in `AGENTS.md` to this exact shape:

```md
| Project | Specs Index |
|---------|-------------|
| adp | [projects/adp/specs/index.md](./projects/adp/specs/index.md) |
| searchWebApi | [projects/searchWebApi/specs/index.md](./projects/searchWebApi/specs/index.md) |
```

- [ ] **Step 2: Verify `AGENTS.md` points to the new spec entrypoint**

Run:

```bash
rg -n "searchWebApi" AGENTS.md
```

Expected: one matching project-reference row containing `projects/searchWebApi/specs/index.md`.

- [ ] **Step 3: Review the exact diff for only the intended metadata change**

Run:

```bash
git diff -- AGENTS.md
```

Expected: a single added table row for `searchWebApi` and no unrelated workflow edits.

- [ ] **Step 4: Commit the metadata update**

```bash
git add AGENTS.md
git commit -m "docs: register searchWebApi project"
```

---

### Task 3: Verify The Skeleton Stayed Minimal

**Files:**
- Verify: `projects/searchWebApi/specs/index.md`
- Verify: `AGENTS.md`

- [ ] **Step 1: Confirm the new project contains only the single spec entrypoint file**

Run:

```bash
rg --files projects/searchWebApi
```

Expected:

```text
projects/searchWebApi/specs/index.md
```

- [ ] **Step 2: Check for whitespace or patch-format issues in the changed files**

Run:

```bash
git diff --check -- AGENTS.md projects/searchWebApi/specs/index.md
```

Expected: no output.

- [ ] **Step 3: Review the final combined diff before handing off**

Run:

```bash
git diff -- AGENTS.md projects/searchWebApi/specs/index.md
```

Expected: only two changes are present:

```text
1. a new `projects/searchWebApi/specs/index.md` file
2. a new `searchWebApi` row in the `AGENTS.md` project table
```

- [ ] **Step 4: Record completion in the final status note**

Use this completion note:

```text
Created the spec-only `searchWebApi` project skeleton and registered it in `AGENTS.md`. No additional spec files or generated code were added.
```
