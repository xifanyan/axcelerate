# ADP Engine/Application Selector Validation Design

## Context

Several ADP tasks currently expose both `engineName` and `applicationIdentifier` as inputs.

In the current specs and Go implementation, these fields are inconsistent:

- some tasks treat `engineName` as required,
- some allow both fields to be omitted,
- some allow both fields to be provided together,
- and the Go CLI currently hard-codes `engineName` as parse-time required for some tasks even when `applicationIdentifier` is also present.

The user reported the intended rule: for tasks that define both selectors, one of them is required and the two are mutually exclusive.

## Scope

This rule applies to all four current tasks that define both `engineName` and `applicationIdentifier`:

- `Query Engine`
- `Taxonomy Statistic`
- `Create OCR Job`
- `CSV Merge`

## Goals

- Define a single consistent selector rule in the affected task specs.
- Align the Go library builders with that rule.
- Align the Go CLI with that rule.
- Ensure CLI and library users get clear validation errors before invalid requests are sent.

## Non-Goals

- Changing tasks that only expose `applicationIdentifier`.
- Changing upstream raw default request examples from `API-SPEC.md`.
- Adding new selector fields or new task-discovery behavior.

## Rule

For any task that defines both `engineName` and `applicationIdentifier`:

1. Exactly one of the two fields must be provided.
2. Providing neither field is invalid.
3. Providing both fields is invalid.
4. `engineName` and `applicationIdentifier` are alternative selectors for the same target.

This rule applies equally to the request-construction API and the CLI.

## Recommended Approach

Implement the rule in both the specs and the Go code, with validation at two layers:

- builder validation in the Go library,
- command validation in the Go CLI.

Why:

- builder validation keeps the Go library correct even when the CLI is not used,
- CLI validation gives command users immediate, task-specific errors,
- and keeping both layers aligned avoids drift.

Alternatives considered:

1. CLI-only validation.
   This would still allow invalid requests through the Go library.
2. Builder-only validation.
   This would work functionally, but CLI users would still see weaker or later errors.
3. Per-task ad hoc wording and validation.
   This duplicates the same rule four times in code with a higher drift risk.

## Spec Changes

Update these task specs:

- `projects/adp/specs/tasks/query-engine.md`
- `projects/adp/specs/tasks/taxonomy-statistic.md`
- `projects/adp/specs/tasks/create-ocr-job.md`
- `projects/adp/specs/tasks/csv-merge.md`

For each affected spec:

- keep both fields in the semantic input list,
- change the Required semantics so they no longer imply `engineName` alone is always required,
- add an explicit note that exactly one of `engineName` or `applicationIdentifier` must be provided,
- update CLI argument documentation to reflect the same rule,
- and update examples so they do not imply both are required or that `engineName` is the only valid selector.

Recommended wording:

> `engineName` and `applicationIdentifier` are mutually exclusive selectors. Exactly one must be provided.

## Go Library Changes

Update these builder implementations:

- `projects/adp/src/go/query_engine.go`
- `projects/adp/src/go/taxonomy_statistic.go`
- `projects/adp/src/go/create_ocr_job.go`
- `projects/adp/src/go/csv_merge.go`

Behavior:

- if neither selector is set, return a validation error,
- if both selectors are set, return a validation error,
- if exactly one selector is set, continue building the sparse request normally.

Use consistent error text across tasks so tests and users see the same rule.

Recommended error text:

- missing case: `exactly one of engineName or applicationIdentifier is required`
- both case: `engineName and applicationIdentifier are mutually exclusive`

The builders should continue omitting whichever selector was not provided.

## Go CLI Changes

Update:

- `projects/adp/src/go/cmd/adpgo/main.go`

Behavior:

- remove unconditional `Required: true` on `--engineName` for affected commands,
- validate `engineName` and `applicationIdentifier` together at command level,
- reject neither-provided and both-provided cases with clear user-facing errors,
- allow either `--engineName` or `--applicationIdentifier` to be the valid selector.

The CLI should enforce the same exact-one rule as the builders.

## Testing Changes

Update and extend tests in:

- `projects/adp/src/go/builders_test.go`
- `projects/adp/src/go/cmd/adpgo/main_test.go`

Required builder coverage for each affected task:

- neither selector provided -> error
- both selectors provided -> error
- engine-only -> valid
- application-only -> valid

Required CLI coverage for each affected command:

- neither selector provided -> parse/command error
- both selectors provided -> command error
- engine-only -> valid path
- application-only -> valid path

## Error Handling

Validation errors should happen before request execution.

For CLI users, the error must be user-facing and should not depend on ADP server responses.

For Go library users, the error must be returned from the builder execution path without producing a request.

## Success Criteria

This design is successful when:

- all four affected task specs describe the same selector rule,
- the Go builders enforce exactly-one selector semantics,
- the Go CLI enforces the same rule,
- and tests cover neither/both/engine-only/application-only cases for each affected task.
