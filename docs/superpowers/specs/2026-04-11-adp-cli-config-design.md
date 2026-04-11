# ADP CLI Global Config Design

## Context

This document defines a language-agnostic update to the ADP CLI contract in `projects/adp/specs/cli.md` and the matching global summary in `projects/adp/specs/index.md`.

The current CLI spec already defines shared global flags, but it does not yet specify:

- the global `--path` flag,
- a universal `adp_config.json` config file,
- or the precedence rule between config values and command-line flags.

The goal of this update is to make all ADP CLIs behave consistently regardless of implementation language.

## Goals

- Add `--path` to the universal ADP CLI global flags.
- Define `adp_config.json` as the shared CLI config file for all language implementations.
- Allow the config file to provide defaults for all global flags.
- Define a single precedence rule where explicit command-line flags override config-file values.
- Keep the contract language-agnostic so Go, Python, and Rust CLIs can all implement the same behavior.

## Non-Goals

- Defining task-specific config-file entries.
- Designing per-language config search paths beyond the shared file name contract.
- Changing task subcommands or task-specific flags.
- Implementing the behavior in code in this design step.

## Recommended Approach

Update the core CLI contract in one place and keep task specs indirect.

Why:

- `projects/adp/specs/cli.md` is already the authoritative CLI contract.
- Task specs already defer to `cli.md` for global flags and naming conventions.
- This keeps the new config-file behavior centralized instead of repeating the same rules across every task spec.

Alternatives considered:

1. Update only `cli.md`.
   This is smaller, but leaves `projects/adp/specs/index.md` inconsistent with the CLI contract summary.
2. Repeat the config-file rules in every task spec.
   This is more explicit, but it duplicates the same cross-cutting rule and increases drift risk.

## Contract Update

### Global Flags

All language CLIs must support these global flags:

- `--host` string: ADP server host
- `--port` integer: ADP server port, default `8443`
- `--path` string: ADP task API path, default `"/adp/rest/api/task"`
- `--user` string: ADP username
- `--password` string: ADP password
- `--insecure` boolean: skip TLS certificate verification
- `--debug`, `-d` boolean: enable debug logging

`--host`, `--user`, and `--password` remain required overall, but they may be satisfied either by explicit command-line flags or by `adp_config.json`.

### Config File

All language CLIs must support a config file named exactly `adp_config.json`.

This file may define defaults for any or all global CLI settings:

```json
{
  "host": "example.com",
  "port": 8443,
  "path": "/adp/rest/api/task",
  "user": "adp",
  "password": "secret",
  "insecure": false,
  "debug": false
}
```

The config keys map directly to the global flag names without the `--` prefix.

### Precedence Rules

The CLI must resolve global configuration in this order:

1. Explicit command-line flag
2. `adp_config.json`
3. Built-in default, when one exists

Built-in defaults remain:

- `port = 8443`
- `path = "/adp/rest/api/task"`
- `insecure = false`
- `debug = false`

There is no built-in default for:

- `host`
- `user`
- `password`

If one of those required values is missing after applying the precedence rules, the CLI must fail with a user-facing error.

### Debug Behavior

Debug behavior does not change semantically:

- when debug resolves to `true`, the CLI traces request and response payloads,
- when debug resolves to `false`, it does not.

The only new rule is that `debug` now participates in the same precedence rules as every other global flag.

## Spec File Changes

### `projects/adp/specs/cli.md`

Update this file to:

- add `--path` to the global flag table,
- document `adp_config.json`,
- define the shared config keys,
- define the precedence order,
- clarify that required global values may come from either CLI flags or config,
- and keep examples aligned with the new global contract.

### `projects/adp/specs/index.md`

Update the global CLI summary to:

- include `--path` in the shared global flag list,
- note `adp_config.json` as the shared CLI config file,
- and state that command-line flags override config-file values.

Task specs should remain unchanged unless a task spec currently duplicates outdated global-flag text.

## Error Handling

If `adp_config.json` is present but invalid JSON, the CLI should fail with a clear user-facing configuration error.

If `adp_config.json` is valid JSON but contains values of the wrong type for the global fields, the CLI should fail with a clear user-facing configuration error.

If both config file and CLI flags are present, the CLI must use the command-line value for each explicitly provided global flag without treating that as an error.

## Testing Expectations

When this design is implemented in language-specific CLIs, verification should cover:

- config-only resolution for required fields,
- CLI override of each global field from config,
- built-in defaults for `port`, `path`, `insecure`, and `debug`,
- invalid `adp_config.json` handling,
- missing required values after resolution,
- and `--debug` / `-d` override behavior.

## Success Criteria

This design is successful when:

- the ADP CLI spec clearly defines the full shared global flag set,
- all language implementations can follow one config-file contract,
- command-line override behavior is explicit for every global flag,
- and the global CLI rules in `cli.md` and `index.md` no longer conflict.
