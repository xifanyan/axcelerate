# ADP Selector Server Compatibility Design

## Context

The current ADP selector design enforces an exact-one rule for tasks that expose both `engineName` and `applicationIdentifier`.

That rule is correct at the client API level, but live ADP behavior shows an additional compatibility constraint for some tasks.

Observed evidence:

- `query-engine --applicationIdentifier <value>` produced a sparse request that only sent `adp_queryEngine_applicationIdentifier`.
- The live ADP server still rejected the request with `Both engine and application are defined.`
- In `projects/adp/API-SPEC.md`, the upstream defaults for these tasks include a non-empty engine placeholder:
  - `Query Engine`: `adp_queryEngine_engineName = "{adp_engineName}"`
  - `Taxonomy Statistic`: `adp_taxonomyStatistic_engineName = "{adp_engine_name}"`
  - `Create OCR Job`: `adp_createOcrJob_engineName = "{adp_engineName}"`
- `CSV Merge` differs because its upstream default is `adp_csvMerge_engineName = null`.

The most likely explanation is that the ADP server materializes omitted upstream defaults before validating selectors. For affected tasks, omitting `engineName` while sending `applicationIdentifier` is not sufficient. The client must explicitly clear `engineName` with an empty string.

## Scope

Apply this compatibility fix to:

- `Query Engine`
- `Taxonomy Statistic`
- `Create OCR Job`

Do not apply this change to `CSV Merge` unless live-server evidence shows the same need.

## Goals

- Preserve the exact-one selector rule at the client API level.
- Make application-selected requests work against the live ADP server.
- Keep the workaround minimal and task-scoped.
- Document the compatibility exception in the affected task specs.

## Non-Goals

- Changing the selector semantics for callers.
- Relaxing validation to allow both selectors.
- Applying the workaround to unrelated tasks.
- Changing `CSV Merge` without evidence.

## Rule

For `Query Engine`, `Taxonomy Statistic`, and `Create OCR Job`:

1. The client still treats `engineName` and `applicationIdentifier` as mutually exclusive selectors.
2. Exactly one effective selector must be provided by the caller.
3. When `engineName` is the effective selector, serialize only the engine selector.
4. When `applicationIdentifier` is the effective selector, serialize:
   - `applicationIdentifier` with the provided value
   - `engineName` as `""` to explicitly clear the upstream server default

This is a server-compatibility exception to the normal sparse-request rule.

## Alternatives Considered

1. Keep pure sparse requests.

   This matches the language-level request-construction rule, but it does not work against the live server for affected tasks.

2. Add a compatibility flag or environment switch.

   This would support both behaviors, but it adds avoidable complexity without evidence that any real ADP environment requires the sparse form.

3. Apply the explicit-clear workaround to all four selector tasks.

   This is unnecessary for `CSV Merge` based on the current upstream defaults and would broaden the exception without evidence.

## Recommended Approach

Implement a task-scoped compatibility override in the three affected builders:

- `projects/adp/src/go/query_engine.go`
- `projects/adp/src/go/taxonomy_statistic.go`
- `projects/adp/src/go/create_ocr_job.go`

Behavior:

- keep exact-one validation unchanged,
- keep empty string inputs treated as unset for validation,
- when application selection is used, explicitly serialize the corresponding `..._engineName` field as `""`,
- and continue serializing the chosen `applicationIdentifier` field.

Do not change `projects/adp/src/go/csv_merge.go` for this compatibility issue.

## Spec Changes

Update these task specs:

- `projects/adp/specs/tasks/query-engine.md`
- `projects/adp/specs/tasks/taxonomy-statistic.md`
- `projects/adp/specs/tasks/create-ocr-job.md`

For each spec:

- preserve the exact-one selector rule,
- add a note that live ADP requires explicit clearing of `engineName` when application selection is used,
- clarify that this is a compatibility exception to sparse serialization,
- and ensure examples remain valid under the exact-one rule.

Do not add this compatibility note to `csv-merge.md` unless implementation scope changes.

## Go Library Changes

Builder behavior for the three affected tasks:

- neither selector -> error: `exactly one of engineName or applicationIdentifier is required`
- both non-empty selectors -> error: `engineName and applicationIdentifier are mutually exclusive`
- engine-only -> serialize only the engine selector
- application-only -> serialize application selector and `engineName: ""`

This keeps caller semantics stable while adapting request serialization to live-server behavior.

## Go CLI Changes

The CLI should continue using the same exact-one validation already in place.

No new user-facing flags are needed.

The main CLI change is test coverage that confirms application-selected commands now send explicit empty engine fields for:

- `query-engine`
- `taxonomy-statistic`
- `create-ocr-job`

`csv-merge` CLI behavior remains unchanged.

## Testing Changes

Update tests in:

- `projects/adp/src/go/builders_test.go`
- `projects/adp/src/go/cmd/adpgo/main_test.go`

Required builder coverage for each of the three affected tasks:

- application-only requests include `applicationIdentifier`
- application-only requests include the corresponding `engineName` key with value `""`
- engine-only requests do not include `applicationIdentifier`
- exact-one validation still rejects neither and both

Required CLI coverage for each of the three affected commands:

- application-only command path sends `applicationIdentifier`
- application-only command path sends explicit empty `engineName`
- engine-only command path still omits `applicationIdentifier`

`CSV Merge` tests should remain unchanged for this compatibility issue.

## Success Criteria

This design is successful when:

- application-selected `query-engine`, `taxonomy-statistic`, and `create-ocr-job` requests work against live ADP environments that materialize upstream engine defaults,
- exact-one selector validation remains intact for callers,
- `CSV Merge` remains unchanged,
- and specs and tests clearly document the compatibility exception.
