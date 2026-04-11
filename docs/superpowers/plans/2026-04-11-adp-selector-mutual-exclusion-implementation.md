# ADP Selector Mutual Exclusion Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Enforce the rule that `engineName` and `applicationIdentifier` are mutually exclusive alternative selectors, with exactly one required, across the affected ADP task specs and the Go implementation.

**Architecture:** Update the four affected task specs first so the contract is explicit, then add builder-level validation in the Go library and command-level validation in the Go CLI. Keep the logic consistent across tasks by using shared validation wording and focused tests for neither/both/engine-only/application-only cases.

**Tech Stack:** Markdown specs, Go, `urfave/cli/v3`, `go test`

---

## File Map

- Modify: `projects/adp/specs/tasks/query-engine.md`
  - Document the exact-one selector rule.
- Modify: `projects/adp/specs/tasks/taxonomy-statistic.md`
  - Document the exact-one selector rule.
- Modify: `projects/adp/specs/tasks/create-ocr-job.md`
  - Document the exact-one selector rule.
- Modify: `projects/adp/specs/tasks/csv-merge.md`
  - Document the exact-one selector rule.
- Modify: `projects/adp/src/go/query_engine.go`
  - Enforce selector validation in the builder.
- Modify: `projects/adp/src/go/taxonomy_statistic.go`
  - Enforce selector validation in the builder.
- Modify: `projects/adp/src/go/create_ocr_job.go`
  - Enforce selector validation in the builder.
- Modify: `projects/adp/src/go/csv_merge.go`
  - Enforce selector validation in the builder.
- Modify: `projects/adp/src/go/cmd/adpgo/main.go`
  - Enforce selector validation in the CLI and remove conflicting parse-time `engineName` required flags.
- Modify: `projects/adp/src/go/builders_test.go`
  - Add builder validation tests for neither/both/engine-only/application-only cases.
- Modify: `projects/adp/src/go/cmd/adpgo/main_test.go`
  - Add CLI validation and valid-path tests for affected commands.

---

### Task 1: Update The Affected Task Specs

**Files:**
- Modify: `projects/adp/specs/tasks/query-engine.md`
- Modify: `projects/adp/specs/tasks/taxonomy-statistic.md`
- Modify: `projects/adp/specs/tasks/create-ocr-job.md`
- Modify: `projects/adp/specs/tasks/csv-merge.md`

- [ ] **Step 1: Write the exact selector rule text in your notes**

Use this exact wording while editing:

```text
engineName and applicationIdentifier are mutually exclusive selectors. Exactly one must be provided.
```

- [ ] **Step 2: Update `query-engine.md`**

Make these concrete edits:

```md
- In the Semantic Inputs table, stop implying `engineName` alone is unconditionally required.
- Add a note immediately below the table:

> `engineName` and `applicationIdentifier` are mutually exclusive selectors. Exactly one must be provided.

- In the CLI Arguments section, add the same note.
- Add one example using `--applicationIdentifier` as the valid selector.
- Do not add any example that passes both selectors.
```

- [ ] **Step 3: Update `taxonomy-statistic.md`**

Make these concrete edits:

```md
- In the Semantic Inputs table, stop implying `engineName` alone is unconditionally required.
- Add the exact-one selector note below Semantic Inputs.
- Add the same note in CLI Arguments.
- Add one valid example using `--applicationIdentifier`.
- Do not add any example that passes both selectors.
```

- [ ] **Step 4: Update `create-ocr-job.md`**

Make these concrete edits:

```md
- Add the exact-one selector note below Semantic Inputs.
- Add the same note in CLI Arguments.
- Keep existing engine-based examples.
- Add one valid application-based example.
- Do not add any example that passes both selectors.
```

- [ ] **Step 5: Update `csv-merge.md`**

Make these concrete edits:

```md
- Add the exact-one selector note below Semantic Inputs.
- Add the same note in CLI Arguments.
- Add one valid example using `--applicationIdentifier`.
- Do not add any example that passes both selectors.
```

- [ ] **Step 6: Review the spec-only diff**

Run:

```bash
git diff -- projects/adp/specs/tasks/query-engine.md projects/adp/specs/tasks/taxonomy-statistic.md projects/adp/specs/tasks/create-ocr-job.md projects/adp/specs/tasks/csv-merge.md
```

Expected:

```text
- all four task specs contain the same selector rule
- none of the edited examples use both selectors
- each task still documents both fields
```

- [ ] **Step 7: Commit the spec changes**

```bash
git add projects/adp/specs/tasks/query-engine.md projects/adp/specs/tasks/taxonomy-statistic.md projects/adp/specs/tasks/create-ocr-job.md projects/adp/specs/tasks/csv-merge.md
git commit -m "docs: define ADP task selector rules"
```

---

### Task 2: Add Failing Builder Tests For Selector Validation

**Files:**
- Modify: `projects/adp/src/go/builders_test.go`

- [ ] **Step 1: Add failing Query Engine builder tests**

Add tests with these exact names:

```go
func TestQueryEngineRequiresExactlyOneSelector(t *testing.T) {}
func TestQueryEngineRejectsBothSelectors(t *testing.T) {}
func TestQueryEngineAllowsApplicationIdentifierWithoutEngineName(t *testing.T) {}
```

Required assertions:

- neither selector -> error `exactly one of engineName or applicationIdentifier is required`
- both selectors -> error `engineName and applicationIdentifier are mutually exclusive`
- application-only -> request includes `adp_queryEngine_applicationIdentifier` and omits `adp_queryEngine_engineName`

- [ ] **Step 2: Add failing Taxonomy Statistic builder tests**

Add tests with these exact names:

```go
func TestTaxonomyStatisticRequiresExactlyOneSelector(t *testing.T) {}
func TestTaxonomyStatisticRejectsBothSelectors(t *testing.T) {}
func TestTaxonomyStatisticAllowsApplicationIdentifierWithoutEngineName(t *testing.T) {}
```

- [ ] **Step 3: Add failing Create OCR Job builder tests**

Add tests with these exact names:

```go
func TestCreateOcrJobRequiresExactlyOneSelector(t *testing.T) {}
func TestCreateOcrJobRejectsBothSelectors(t *testing.T) {}
func TestCreateOcrJobAllowsApplicationIdentifierWithoutEngineName(t *testing.T) {}
```

- [ ] **Step 4: Add failing CSV Merge builder tests**

Add tests with these exact names:

```go
func TestCSVMergeRequiresExactlyOneSelector(t *testing.T) {}
func TestCSVMergeRejectsBothSelectors(t *testing.T) {}
func TestCSVMergeAllowsApplicationIdentifierWithoutEngineName(t *testing.T) {}
```

- [ ] **Step 5: Run the focused builder tests and verify they fail**

Run:

```bash
go test ./... -run "Test(QueryEngineRequiresExactlyOneSelector|QueryEngineRejectsBothSelectors|QueryEngineAllowsApplicationIdentifierWithoutEngineName|TaxonomyStatisticRequiresExactlyOneSelector|TaxonomyStatisticRejectsBothSelectors|TaxonomyStatisticAllowsApplicationIdentifierWithoutEngineName|CreateOcrJobRequiresExactlyOneSelector|CreateOcrJobRejectsBothSelectors|CreateOcrJobAllowsApplicationIdentifierWithoutEngineName|CSVMergeRequiresExactlyOneSelector|CSVMergeRejectsBothSelectors|CSVMergeAllowsApplicationIdentifierWithoutEngineName)$" -count=1
```

Expected: FAIL for missing validation support.

- [ ] **Step 6: Commit the failing builder tests**

```bash
git add projects/adp/src/go/builders_test.go
git commit -m "test: cover ADP selector builder validation"
```

---

### Task 3: Implement Builder-Level Exact-One Selector Validation

**Files:**
- Modify: `projects/adp/src/go/query_engine.go`
- Modify: `projects/adp/src/go/taxonomy_statistic.go`
- Modify: `projects/adp/src/go/create_ocr_job.go`
- Modify: `projects/adp/src/go/csv_merge.go`

- [ ] **Step 1: Implement Query Engine builder validation**

Update the request-building validation so it follows this shape:

```go
hasEngine := b.engineName != nil && *b.engineName != ""
hasApplication := b.applicationIdentifier != nil && *b.applicationIdentifier != ""

if !hasEngine && !hasApplication {
	return rawTaskRequest{}, errors.New("exactly one of engineName or applicationIdentifier is required")
}
if hasEngine && hasApplication {
	return rawTaskRequest{}, errors.New("engineName and applicationIdentifier are mutually exclusive")
}
```

Include only the selected field in the resulting sparse configuration.

- [ ] **Step 2: Implement the same validation in `taxonomy_statistic.go`**

Use the same error text and omit the unselected field from the sparse request.

- [ ] **Step 3: Implement the same validation in `create_ocr_job.go`**

Use the same error text and omit the unselected field from the sparse request.

- [ ] **Step 4: Implement the same validation in `csv_merge.go`**

Use the same error text and omit the unselected field from the sparse request.

- [ ] **Step 5: Run the focused builder tests and verify they pass**

Run the same command from Task 2 Step 5.

Expected: PASS.

- [ ] **Step 6: Run existing builder regressions**

Run:

```bash
go test ./... -run "Test(QueryEngineBuildsSparseRequestAndDecodesResult|TaxonomyStatisticBuildsRequestAndDecodesResult|CreateOcrJobBuildsSparseAsyncRequest|CSVMergeBuildsSparseRequest)$" -count=1
```

Expected: PASS.

- [ ] **Step 7: Commit the builder implementation**

```bash
git add projects/adp/src/go/query_engine.go projects/adp/src/go/taxonomy_statistic.go projects/adp/src/go/create_ocr_job.go projects/adp/src/go/csv_merge.go projects/adp/src/go/builders_test.go
git commit -m "feat: enforce ADP selector validation in builders"
```

---

### Task 4: Add Failing CLI Tests For Selector Validation

**Files:**
- Modify: `projects/adp/src/go/cmd/adpgo/main_test.go`

- [ ] **Step 1: Add failing CLI tests for Query Engine**

Add tests with these exact names:

```go
func TestQueryEngineCommandRequiresExactlyOneSelector(t *testing.T) {}
func TestQueryEngineCommandRejectsBothSelectors(t *testing.T) {}
func TestQueryEngineCommandAllowsApplicationIdentifier(t *testing.T) {}
```

Required assertions:

- neither selector -> error text contains `exactly one of engineName or applicationIdentifier is required`
- both selectors -> error text contains `engineName and applicationIdentifier are mutually exclusive`
- application-only -> request succeeds and includes only the application selector

- [ ] **Step 2: Add failing CLI tests for Taxonomy Statistic**

Add tests with these exact names:

```go
func TestTaxonomyStatisticCommandRequiresExactlyOneSelector(t *testing.T) {}
func TestTaxonomyStatisticCommandRejectsBothSelectors(t *testing.T) {}
func TestTaxonomyStatisticCommandAllowsApplicationIdentifier(t *testing.T) {}
```

- [ ] **Step 3: Add failing CLI tests for Create OCR Job**

Add tests with these exact names:

```go
func TestCreateOcrJobCommandRequiresExactlyOneSelector(t *testing.T) {}
func TestCreateOcrJobCommandRejectsBothSelectors(t *testing.T) {}
func TestCreateOcrJobCommandAllowsApplicationIdentifierWithoutEngineName(t *testing.T) {}
```

- [ ] **Step 4: Add failing CLI tests for CSV Merge**

Add tests with these exact names:

```go
func TestCSVMergeCommandRequiresExactlyOneSelector(t *testing.T) {}
func TestCSVMergeCommandRejectsBothSelectors(t *testing.T) {}
func TestCSVMergeCommandAllowsApplicationIdentifierWithoutEngineName(t *testing.T) {}
```

- [ ] **Step 5: Run the focused CLI selector tests and verify they fail**

Run:

```bash
go test ./cmd/adpgo -run "Test(QueryEngineCommandRequiresExactlyOneSelector|QueryEngineCommandRejectsBothSelectors|QueryEngineCommandAllowsApplicationIdentifier|TaxonomyStatisticCommandRequiresExactlyOneSelector|TaxonomyStatisticCommandRejectsBothSelectors|TaxonomyStatisticCommandAllowsApplicationIdentifier|CreateOcrJobCommandRequiresExactlyOneSelector|CreateOcrJobCommandRejectsBothSelectors|CreateOcrJobCommandAllowsApplicationIdentifierWithoutEngineName|CSVMergeCommandRequiresExactlyOneSelector|CSVMergeCommandRejectsBothSelectors|CSVMergeCommandAllowsApplicationIdentifierWithoutEngineName)$" -count=1
```

Expected: FAIL before CLI implementation is updated.

- [ ] **Step 6: Commit the failing CLI tests**

```bash
git add projects/adp/src/go/cmd/adpgo/main_test.go
git commit -m "test: cover ADP selector CLI validation"
```

---

### Task 5: Implement CLI-Level Exact-One Selector Validation

**Files:**
- Modify: `projects/adp/src/go/cmd/adpgo/main.go`

- [ ] **Step 1: Add a shared selector validation helper**

Add a small helper in `main.go` with this shape:

```go
func validateExclusiveSelectors(cmd *cli.Command) error {
	hasEngine := strings.TrimSpace(cmd.String("engineName")) != ""
	hasApplication := strings.TrimSpace(cmd.String("applicationIdentifier")) != ""

	if !hasEngine && !hasApplication {
		return errors.New("exactly one of engineName or applicationIdentifier is required")
	}
	if hasEngine && hasApplication {
		return errors.New("engineName and applicationIdentifier are mutually exclusive")
	}
	return nil
}
```

- [ ] **Step 2: Remove conflicting `Required: true` usage for affected `engineName` CLI flags**

In these commands, make `engineName` a normal string flag so the exact-one helper can decide validity:

```go
query-engine
taxonomy-statistic
create-ocr-job
csv-merge
```

Do not change unrelated required flags such as `csvFile` or `batchScriptPath`.

- [ ] **Step 3: Call the shared selector validator in each affected command action**

Before building the request, add:

```go
if err := validateExclusiveSelectors(cmd); err != nil {
	return err
}
```

Do this in:

- `query-engine`
- `taxonomy-statistic`
- `create-ocr-job`
- `csv-merge`

- [ ] **Step 4: Keep successful selector paths sparse**

Ensure command actions still only apply whichever selector flag the user actually set.

- [ ] **Step 5: Run the focused CLI selector tests and verify they pass**

Run the same command from Task 4 Step 5.

Expected: PASS.

- [ ] **Step 6: Run existing CLI regressions**

Run:

```bash
go test ./cmd/adpgo -run "Test(QueryEngineCommandParsesTaxonomyFlags|TestRunPrintsParserErrorsToStderr|TestRunPrintsTaxonomyStatisticParserErrorsToStderr|TestCreateOcrJobCommandStartsWithoutWaitingByDefault|TestCreateOcrJobCommandWaitsWhenRequested|TestCSVMergeCommandParsesFieldMappingsJSON)$" -count=1
```

Expected: PASS.

- [ ] **Step 7: Commit the CLI implementation**

```bash
git add projects/adp/src/go/cmd/adpgo/main.go projects/adp/src/go/cmd/adpgo/main_test.go
git commit -m "feat: enforce ADP selector validation in CLI"
```

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

- [ ] **Step 2: Review the full feature diff**

Run:

```bash
git diff --stat HEAD~5 HEAD
git status --short
```

Expected:

```text
- only the planned spec, builder, CLI, and test files changed
- no unexpected files are dirty
```

- [ ] **Step 3: Verify CLI help still works**

Run:

```bash
go run ./cmd/adpgo --help
```

Expected: PASS and commands still render normally.

- [ ] **Step 4: Record verification state**

If final verification changes files, commit them:

```bash
git add projects/adp/specs/tasks/query-engine.md projects/adp/specs/tasks/taxonomy-statistic.md projects/adp/specs/tasks/create-ocr-job.md projects/adp/specs/tasks/csv-merge.md projects/adp/src/go/query_engine.go projects/adp/src/go/taxonomy_statistic.go projects/adp/src/go/create_ocr_job.go projects/adp/src/go/csv_merge.go projects/adp/src/go/builders_test.go projects/adp/src/go/cmd/adpgo/main.go projects/adp/src/go/cmd/adpgo/main_test.go
git commit -m "test: verify ADP selector validation"
```

If no files changed, do not create an empty commit. Report the verification commands in the final handoff instead.

---

## Spec Coverage Check

- exact-one selector rule for all four affected tasks: Task 1
- builder enforcement for all four tasks: Tasks 2 and 3
- CLI enforcement for all four commands: Tasks 4 and 5
- neither/both/engine-only/application-only coverage: Tasks 2, 3, 4, and 5
- final end-to-end verification: Task 6

## Notes For The Implementer

- Use the same validation wording across builders and CLI.
- Do not leave `engineName` as unconditionally required in the CLI for affected tasks.
- Keep successful requests sparse: include only the selector the user chose.
- Do not touch tasks outside the four scoped by the approved design.
