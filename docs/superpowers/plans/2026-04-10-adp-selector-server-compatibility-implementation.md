# ADP Selector Server Compatibility Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Make application-selected `Query Engine`, `Taxonomy Statistic`, and `Create OCR Job` requests compatible with live ADP by explicitly clearing `engineName` while preserving exact-one selector validation.

**Architecture:** Keep the existing exact-one selector contract for callers, but change request serialization for the three affected tasks so application-selected requests include `engineName: ""` in addition to the chosen `applicationIdentifier`. Update the task specs first, then drive the Go changes with focused failing tests at the builder and CLI layers.

**Tech Stack:** Markdown specs, Go, `urfave/cli/v3`, `go test`

---

## File Map

- Modify: `projects/adp/specs/tasks/query-engine.md`
  - Document the live-ADP compatibility exception for application-selected requests.
- Modify: `projects/adp/specs/tasks/taxonomy-statistic.md`
  - Document the same compatibility exception.
- Modify: `projects/adp/specs/tasks/create-ocr-job.md`
  - Document the same compatibility exception.
- Modify: `projects/adp/src/go/query_engine.go`
  - Serialize `adp_queryEngine_engineName` as `""` when `applicationIdentifier` is the chosen selector.
- Modify: `projects/adp/src/go/taxonomy_statistic.go`
  - Serialize `adp_taxonomyStatistic_engineName` as `""` when `applicationIdentifier` is the chosen selector.
- Modify: `projects/adp/src/go/create_ocr_job.go`
  - Serialize `adp_createOcrJob_engineName` as `""` when `applicationIdentifier` is the chosen selector.
- Modify: `projects/adp/src/go/builders_test.go`
  - Add and update builder tests for explicit empty engine serialization on application-selected requests.
- Modify: `projects/adp/src/go/cmd/adpgo/main_test.go`
  - Add and update CLI tests to assert explicit empty engine serialization for the same tasks.

Do not modify `projects/adp/src/go/csv_merge.go` or `projects/adp/specs/tasks/csv-merge.md` for this compatibility change.

---

### Task 1: Update The Three Affected Task Specs

**Files:**
- Modify: `projects/adp/specs/tasks/query-engine.md`
- Modify: `projects/adp/specs/tasks/taxonomy-statistic.md`
- Modify: `projects/adp/specs/tasks/create-ocr-job.md`

- [ ] **Step 1: Write the approved compatibility note in your notes**

Use this exact text while editing:

```text
When applicationIdentifier is used, the client still treats it as the single effective selector, but for live ADP compatibility it explicitly serializes engineName as an empty string to clear the server-side default.
```

- [ ] **Step 2: Update `projects/adp/specs/tasks/query-engine.md`**

Make these concrete edits:

```md
- Keep the exact-one selector note already present.
- Add the compatibility note near the selector documentation or request-construction guidance.
- Clarify that application-selected requests intentionally include `adp_queryEngine_engineName: ""` for live ADP compatibility.
- Do not change the CSV Merge task spec in this task.
```

- [ ] **Step 3: Update `projects/adp/specs/tasks/taxonomy-statistic.md`**

Make these concrete edits:

```md
- Keep the exact-one selector note already present.
- Add the same compatibility note for live ADP.
- Clarify that application-selected requests intentionally include `adp_taxonomyStatistic_engineName: ""`.
```

- [ ] **Step 4: Update `projects/adp/specs/tasks/create-ocr-job.md`**

Make these concrete edits:

```md
- Keep the exact-one selector note already present.
- Add the same compatibility note for live ADP.
- Clarify that application-selected requests intentionally include `adp_createOcrJob_engineName: ""`.
```

- [ ] **Step 5: Review only the compatibility-spec diff**

Run:

```bash
git diff -- projects/adp/specs/tasks/query-engine.md projects/adp/specs/tasks/taxonomy-statistic.md projects/adp/specs/tasks/create-ocr-job.md
```

Expected:

```text
- all three task specs still describe exact-one selector semantics
- all three now document the explicit empty engineName compatibility exception
- csv-merge.md is unchanged
```

- [ ] **Step 6: Commit the spec changes**

```bash
git add projects/adp/specs/tasks/query-engine.md projects/adp/specs/tasks/taxonomy-statistic.md projects/adp/specs/tasks/create-ocr-job.md
git commit -m "docs: define ADP selector compatibility behavior"
```

---

### Task 2: Add Failing Builder Tests For Compatibility Serialization

**Files:**
- Modify: `projects/adp/src/go/builders_test.go`

- [ ] **Step 1: Update the Query Engine application-selector builder test**

Edit `TestQueryEngineAllowsApplicationIdentifierWithoutEngineName` so it now requires:

```go
if req.TaskConfiguration["adp_queryEngine_applicationIdentifier"] != "appA" {
	t.Fatalf("taskConfiguration = %#v", req.TaskConfiguration)
}
if got := req.TaskConfiguration["adp_queryEngine_engineName"]; got != "" {
	t.Fatalf("engineName = %#v", got)
}
```

Keep the neither/both validation assertions unchanged.

- [ ] **Step 2: Update the Taxonomy Statistic application-selector builder test**

Edit `TestTaxonomyStatisticAllowsApplicationIdentifierWithoutEngineName` so it now requires:

```go
if req.TaskConfiguration["adp_taxonomyStatistic_applicationIdentifier"] != "appA" {
	t.Fatalf("taskConfiguration = %#v", req.TaskConfiguration)
}
if got := req.TaskConfiguration["adp_taxonomyStatistic_engineName"]; got != "" {
	t.Fatalf("engineName = %#v", got)
}
```

- [ ] **Step 3: Update the Create OCR Job application-selector builder test**

Edit `TestCreateOcrJobAllowsApplicationIdentifierWithoutEngineName` so it now requires:

```go
if req.TaskConfiguration["adp_createOcrJob_applicationIdentifier"] != "appA" {
	t.Fatalf("taskConfiguration = %#v", req.TaskConfiguration)
}
if got := req.TaskConfiguration["adp_createOcrJob_engineName"]; got != "" {
	t.Fatalf("engineName = %#v", got)
}
```

- [ ] **Step 4: Add regression assertions for engine-selected requests still omitting application identifiers**

Verify these existing tests still assert omission of the application selector:

```go
TestQueryEngineAllowsEngineNameWhenApplicationIdentifierIsEmpty
TestTaxonomyStatisticAllowsEngineNameWhenApplicationIdentifierIsEmpty
TestCreateOcrJobAllowsEngineNameWhenApplicationIdentifierIsEmpty
```

If any do not assert omission yet, add:

```go
if _, ok := req.TaskConfiguration["<application-key>"]; ok {
	t.Fatalf("taskConfiguration should omit applicationIdentifier: %#v", req.TaskConfiguration)
}
```

- [ ] **Step 5: Run the focused builder tests and verify they fail**

Run:

```bash
go test ./... -run "Test(QueryEngineAllowsApplicationIdentifierWithoutEngineName|TaxonomyStatisticAllowsApplicationIdentifierWithoutEngineName|CreateOcrJobAllowsApplicationIdentifierWithoutEngineName|QueryEngineAllowsEngineNameWhenApplicationIdentifierIsEmpty|TaxonomyStatisticAllowsEngineNameWhenApplicationIdentifierIsEmpty|CreateOcrJobAllowsEngineNameWhenApplicationIdentifierIsEmpty)$" -count=1
```

Expected: FAIL because the builders still omit `engineName` on application-selected requests.

- [ ] **Step 6: Commit the failing builder tests**

```bash
git add projects/adp/src/go/builders_test.go
git commit -m "test: cover ADP selector compatibility builders"
```

---

### Task 3: Implement Builder-Level Compatibility Serialization

**Files:**
- Modify: `projects/adp/src/go/query_engine.go`
- Modify: `projects/adp/src/go/taxonomy_statistic.go`
- Modify: `projects/adp/src/go/create_ocr_job.go`

- [ ] **Step 1: Update `projects/adp/src/go/query_engine.go`**

In `buildRequest()`, keep the existing validation logic but change selector serialization to this shape:

```go
if hasEngine {
	cfg["adp_queryEngine_engineName"] = *b.engineName
}
if hasApplication {
	cfg["adp_queryEngine_engineName"] = ""
	cfg["adp_queryEngine_applicationIdentifier"] = *b.applicationIdentifier
}
```

This must leave engine-selected requests unchanged and make application-selected requests explicitly clear the upstream engine default.

- [ ] **Step 2: Update `projects/adp/src/go/taxonomy_statistic.go`**

In `buildRequest()`, keep the existing validation logic but change selector serialization to this shape:

```go
if hasEngine {
	cfg["adp_taxonomyStatistic_engineName"] = *b.engineName
}
if hasApplication {
	cfg["adp_taxonomyStatistic_engineName"] = ""
	cfg["adp_taxonomyStatistic_applicationIdentifier"] = *b.applicationIdentifier
}
```

- [ ] **Step 3: Update `projects/adp/src/go/create_ocr_job.go`**

In `buildRequest()`, keep the existing validation logic but change selector serialization to this shape:

```go
if hasEngine {
	cfg["adp_createOcrJob_engineName"] = *b.engineName
}
if hasApplication {
	cfg["adp_createOcrJob_engineName"] = ""
	cfg["adp_createOcrJob_applicationIdentifier"] = *b.applicationIdentifier
}
```

- [ ] **Step 4: Confirm `projects/adp/src/go/csv_merge.go` stays unchanged**

Review that no compatibility serialization is added to CSV Merge in this task.

- [ ] **Step 5: Run the focused builder tests and verify they pass**

Run the same command from Task 2 Step 5.

Expected: PASS.

- [ ] **Step 6: Run the existing selector-validation builder tests as regression coverage**

Run:

```bash
go test ./... -run "Test(QueryEngineRequiresExactlyOneSelector|QueryEngineRejectsBothSelectors|TaxonomyStatisticRequiresExactlyOneSelector|TaxonomyStatisticRejectsBothSelectors|CreateOcrJobRequiresExactlyOneSelector|CreateOcrJobRejectsBothSelectors|CSVMergeRequiresExactlyOneSelector|CSVMergeRejectsBothSelectors)$" -count=1
```

Expected: PASS.

- [ ] **Step 7: Commit the builder implementation**

```bash
git add projects/adp/src/go/query_engine.go projects/adp/src/go/taxonomy_statistic.go projects/adp/src/go/create_ocr_job.go projects/adp/src/go/builders_test.go
git commit -m "feat: add ADP selector compatibility serialization"
```

---

### Task 4: Add Failing CLI Tests For Compatibility Serialization

**Files:**
- Modify: `projects/adp/src/go/cmd/adpgo/main_test.go`

- [ ] **Step 1: Update the Query Engine application-selector CLI test**

Edit `TestQueryEngineCommandAllowsApplicationIdentifier` so the request assertion requires:

```go
if got := cfg["adp_queryEngine_applicationIdentifier"]; got != "appA" {
	t.Fatalf("adp_queryEngine_applicationIdentifier = %#v", got)
}
if got := cfg["adp_queryEngine_engineName"]; got != "" {
	t.Fatalf("adp_queryEngine_engineName = %#v", got)
}
```

- [ ] **Step 2: Update the Taxonomy Statistic application-selector CLI test**

Edit `TestTaxonomyStatisticCommandAllowsApplicationIdentifier` so the request assertion requires:

```go
if got := cfg["adp_taxonomyStatistic_applicationIdentifier"]; got != "appA" {
	t.Fatalf("adp_taxonomyStatistic_applicationIdentifier = %#v", got)
}
if got := cfg["adp_taxonomyStatistic_engineName"]; got != "" {
	t.Fatalf("adp_taxonomyStatistic_engineName = %#v", got)
}
```

- [ ] **Step 3: Update the Create OCR Job application-selector CLI test**

Edit `TestCreateOcrJobCommandAllowsApplicationIdentifierWithoutEngineName` so the request assertion requires:

```go
if got := cfg["adp_createOcrJob_applicationIdentifier"]; got != "appA" {
	t.Fatalf("adp_createOcrJob_applicationIdentifier = %#v", got)
}
if got := cfg["adp_createOcrJob_engineName"]; got != "" {
	t.Fatalf("adp_createOcrJob_engineName = %#v", got)
}
```

- [ ] **Step 4: Keep engine-selected CLI assertions unchanged**

Ensure these tests still assert that application identifiers are omitted when engine selection is used:

```go
TestQueryEngineCommandParsesTaxonomyFlags
TestTaxonomyStatisticCommandAllowsEngineName
TestCreateOcrJobCommandStartsWithoutWaitingByDefault
```

- [ ] **Step 5: Run the focused CLI compatibility tests and verify they fail**

Run:

```bash
go test ./cmd/adpgo -run "Test(QueryEngineCommandAllowsApplicationIdentifier|TaxonomyStatisticCommandAllowsApplicationIdentifier|CreateOcrJobCommandAllowsApplicationIdentifierWithoutEngineName|QueryEngineCommandParsesTaxonomyFlags|TaxonomyStatisticCommandAllowsEngineName|CreateOcrJobCommandStartsWithoutWaitingByDefault)$" -count=1
```

Expected: FAIL because the request bodies still omit `engineName` on application-selected paths.

- [ ] **Step 6: Commit the failing CLI tests**

```bash
git add projects/adp/src/go/cmd/adpgo/main_test.go
git commit -m "test: cover ADP selector compatibility CLI"
```

---

### Task 5: Make The CLI Compatibility Tests Pass

**Files:**
- Modify: `projects/adp/src/go/cmd/adpgo/main_test.go`
- Modify: `projects/adp/src/go/query_engine.go`
- Modify: `projects/adp/src/go/taxonomy_statistic.go`
- Modify: `projects/adp/src/go/create_ocr_job.go`

- [ ] **Step 1: Re-run the focused CLI compatibility tests after Task 3**

Run the same command from Task 4 Step 5.

Expected: PASS, because the CLI uses the updated builders and does not need new runtime flag logic.

- [ ] **Step 2: Run selector-validation CLI regressions**

Run:

```bash
go test ./cmd/adpgo -run "Test(QueryEngineCommandRequiresExactlyOneSelector|QueryEngineCommandRejectsBothSelectors|TaxonomyStatisticCommandRequiresExactlyOneSelector|TaxonomyStatisticCommandRejectsBothSelectors|CreateOcrJobCommandRequiresExactlyOneSelector|CreateOcrJobCommandRejectsBothSelectors|CSVMergeCommandRequiresExactlyOneSelector|CSVMergeCommandRejectsBothSelectors)$" -count=1
```

Expected: PASS.

- [ ] **Step 3: Run the shared command-output regression tests**

Run:

```bash
go test ./cmd/adpgo -run "Test(RunPrintsParserErrorsToStderr|RunPrintsTaxonomyStatisticParserErrorsToStderr|RunPrintsCreateOcrJobSelectorErrorsToStderr|RunPrintsCSVMergeSelectorErrorsToStderr|RunDoesNotDuplicateTaskExecutionErrors)$" -count=1
```

Expected: PASS.

- [ ] **Step 4: Commit the compatibility-ready CLI/test state**

```bash
git add projects/adp/src/go/cmd/adpgo/main_test.go projects/adp/src/go/query_engine.go projects/adp/src/go/taxonomy_statistic.go projects/adp/src/go/create_ocr_job.go projects/adp/src/go/builders_test.go
git commit -m "test: verify ADP selector compatibility paths"
```

If Task 4 Step 5 already passes immediately after Task 3 and no code changes beyond tests are required, this commit still records the validated compatibility state.

---

### Task 6: Final Verification

**Files:**
- Modify: none expected

- [ ] **Step 1: Run the full Go test suite**

Run:

```bash
go test ./... -count=1
```

Expected: PASS.

- [ ] **Step 2: Verify CLI help still renders normally**

Run:

```bash
go run ./cmd/adpgo --help
```

Expected: PASS and the command list is unchanged.

- [ ] **Step 3: Review the final feature diff and status**

Run:

```bash
git diff --stat HEAD~6 HEAD
git status --short
```

Expected:

```text
- only the planned spec, builder, and test files changed
- csv-merge implementation is unchanged for this feature
- no unexpected files are dirty beyond known unrelated files
```

- [ ] **Step 4: Record verification state**

If final verification changes files, commit them:

```bash
git add projects/adp/specs/tasks/query-engine.md projects/adp/specs/tasks/taxonomy-statistic.md projects/adp/specs/tasks/create-ocr-job.md projects/adp/src/go/query_engine.go projects/adp/src/go/taxonomy_statistic.go projects/adp/src/go/create_ocr_job.go projects/adp/src/go/builders_test.go projects/adp/src/go/cmd/adpgo/main_test.go
git commit -m "test: verify ADP selector compatibility"
```

If no files changed, do not create an empty commit. Report the verification commands in the final handoff instead.

---

## Spec Coverage Check

- Live ADP compatibility note for `Query Engine`, `Taxonomy Statistic`, and `Create OCR Job`: Task 1
- Explicit empty `engineName` serialization for application-selected builder requests: Tasks 2 and 3
- No compatibility broadening to `CSV Merge`: Tasks 1, 3, and 6
- CLI request-body coverage for the same compatibility behavior: Tasks 4 and 5
- Full regression and final verification: Tasks 3, 5, and 6

## Notes For The Implementer

- Keep exact-one validation behavior unchanged for callers.
- Do not add new CLI flags or config toggles for this compatibility fix.
- Do not modify `csv_merge.go` or `csv-merge.md` unless new live evidence requires it.
- Prefer the smallest possible serialization change in each builder.
