# List Entities CLI Examples Design

## Context

The `List Entities` task spec already contains a `### CLI Examples` section in `projects/adp/specs/tasks/list-entities.md`.

The user wants to add two concrete usage examples for common ingestion-related lookups:

- getting all data sources for ingestion application `documentHold.demo00001`,
- getting all running ingestion applications.

This is a task-specific documentation update, not a change to the global CLI contract.

## Goals

- Add the requested examples to `projects/adp/specs/tasks/list-entities.md`.
- Keep the examples aligned with the existing documented flag names.
- Preserve the examples as practical usage snippets rather than turning them into normative contract text.

## Non-Goals

- Changing `projects/adp/specs/cli.md`.
- Changing the Go CLI implementation.
- Renaming any flags or entity types.
- Correcting domain values unless explicitly requested.

## Recommended Approach

Update only the `### CLI Examples` section of `projects/adp/specs/tasks/list-entities.md`.

Why:

- The examples are specific to the `list-entities` task.
- The file already has an examples section in the correct place.
- Keeping them local avoids duplicating task-specific usage in the shared CLI spec.

Alternatives considered:

1. Add the examples to `projects/adp/specs/cli.md`.
   This makes the global CLI spec noisier and mixes task-specific usage into a cross-task contract.
2. Add the examples to both `cli.md` and `list-entities.md`.
   This increases drift risk without adding much value.

## Proposed Change

Append these examples under the existing `### CLI Examples` section in `projects/adp/specs/tasks/list-entities.md`:

```bash
# get all datasources for ingestion application documentHold.demo00001
adpgo --debug=false list-entities --type dataSource --relatedEntity documentHold.demo00001

# get all running ingestion applications
adpgo list-entities --type docmentHold --status running
```

## Notes

- The examples should be added exactly as provided by the user.
- Existing examples in the file should remain in place.
- No other sections need to change.

## Success Criteria

This update is successful when:

- `projects/adp/specs/tasks/list-entities.md` includes the two requested examples,
- the examples appear in the task-specific CLI examples section,
- and no unrelated spec files are modified.
