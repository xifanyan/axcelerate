# List Entities CLI Examples Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Add two approved `list-entities` CLI usage examples to the task-specific spec in `projects/adp/specs/tasks/list-entities.md`.

**Architecture:** Make a minimal documentation-only change in the existing `### CLI Examples` section of the `List Entities` task spec. Preserve the existing examples, append the two user-approved examples exactly as provided, and verify that no unrelated spec files change.

**Tech Stack:** Markdown, git diff/status

---

## File Map

- Modify: `projects/adp/specs/tasks/list-entities.md`
  - Append two task-specific CLI examples under the existing `### CLI Examples` section.

No other spec or code files should change.

---

### Task 1: Update The List Entities CLI Examples

**Files:**
- Modify: `projects/adp/specs/tasks/list-entities.md`

- [ ] **Step 1: Record the exact approved example block**

Use this exact content during the edit:

```bash
# get all datasources for ingestion application documentHold.demo00001
adpgo --debug=false list-entities --type dataSource --relatedEntity documentHold.demo00001

# get all running ingestion applications
adpgo list-entities --type docmentHold --status running
```

- [ ] **Step 2: Edit the existing `### CLI Examples` section in `projects/adp/specs/tasks/list-entities.md`**

Keep the existing examples and append the new examples so the section looks like this:

```md
### CLI Examples

```bash
# Basic
adpgo list-entities

# With type filter
adpgo list-entities --type singleMindServer

# With multiple options
adpgo list-entities --type singleMindServer --whiteList "id,displayName,processStatus"

# get all datasources for ingestion application documentHold.demo00001
adpgo --debug=false list-entities --type dataSource --relatedEntity documentHold.demo00001

# get all running ingestion applications
adpgo list-entities --type docmentHold --status running
```
```

Do not change the wording or values in the two newly added examples.

- [ ] **Step 3: Review the edited section for scope and formatting**

Check these exact conditions:

```text
- the new examples are only in projects/adp/specs/tasks/list-entities.md
- the examples appear under the existing ### CLI Examples heading
- the pre-existing examples remain unchanged
- the new examples match the approved text exactly
```

- [ ] **Step 4: Verify only the intended file changed**

Run:

```bash
git diff -- projects/adp/specs/tasks/list-entities.md
git status --short
```

Expected:

```text
- diff shows only the new example lines in list-entities.md
- no unrelated spec or code files are modified
```

- [ ] **Step 5: Commit the spec update**

```bash
git add projects/adp/specs/tasks/list-entities.md
git commit -m "docs: add List Entities CLI examples"
```

---

## Spec Coverage Check

- Add examples only to `projects/adp/specs/tasks/list-entities.md`: Task 1
- Preserve approved example text exactly: Task 1, Steps 1 and 2
- Avoid unrelated spec changes: Task 1, Steps 3 and 4

## Notes For The Implementer

- This is a docs-only update.
- Do not correct `docmentHold` to another value unless the user explicitly asks for it.
- Do not move the examples into `projects/adp/specs/cli.md`.
- Do not edit any Go code for this task.
