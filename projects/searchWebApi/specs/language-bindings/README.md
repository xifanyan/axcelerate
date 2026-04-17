# searchWebApi Language Bindings

This directory contains language-specific bindings.

The core markdown specs in the parent directory remain language-agnostic.

Language bindings may define:

- idiomatic client naming,
- package or module structure,
- type mappings,
- streaming abstractions,
- and session-state handling conventions for a specific language.

Language bindings must not change wire-level paths, parameters, headers, request bodies, or response shapes defined by `../../API-SPEC.md` and the core specs.
