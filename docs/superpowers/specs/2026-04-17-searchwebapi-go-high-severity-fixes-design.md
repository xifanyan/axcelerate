# searchWebApi Go High Severity Fixes Design

## Context

The generated Go module at `C:\Users\pyan\ai-generated\searchWebApi\go` builds and its current tests pass, but review found four high-severity mismatches against `projects/searchWebApi/API-SPEC.md` and `projects/searchWebApi/specs/language-bindings/go.md`.

The goal of this change is to repair those mismatches with the smallest correct patch rather than refactoring the generated client broadly.

## Scope

This repair covers only these issues:

- preserve structured JSON result bodies on non-2xx HTTP responses,
- send `SWA-searchCacheControl` as a request header instead of a query parameter,
- correct multipart change request body shape for record change endpoints,
- and align the example CLI and README with the latest binding spec structure.

This repair does not attempt to address lower-severity type gaps or redesign the transport layer beyond what these fixes require.

## Design

### 1. Structured JSON Error Preservation

`doJSON` must stop treating every HTTP status `>= 400` as a pure transport error.

For JSON endpoints:

- if the response body is JSON and the caller provided an output struct, decode the body into that struct even on non-2xx responses,
- return an error that still exposes the HTTP status and raw body,
- and allow callers to inspect the decoded result together with the returned error.

This preserves the contract requirement that many operations return structured result bodies even when the backend reports failure.

Binary and NDJSON helpers remain error-returning on non-2xx responses because those high-severity findings only affected structured JSON result handling.

### 2. Search Cache Control Header

`SWA-searchCacheControl` must move from query construction to header construction.

Implementation rule:

- option structs may still hold a `SearchCacheControl` string,
- but their `values()` helpers must no longer place it in query parameters,
- and each affected client method must add it to request headers when present.

Affected operations include record search/fetch flows, binary search retrieval, in-document search, highlight lookup, measure cube access, and insert/remove workflows.

### 3. Multipart Change Request Shape

For multipart record change operations, the multipart form field named `request` must contain the JSON array of `ChangeRequest` values directly.

It must not be wrapped as `{"request": [...]}`.

The existing insert/remove multipart shape stays unchanged because that contract requires `request` to be an object of type `InsertRemoveRequest`.

To keep the patch small:

- retain the shared multipart writer,
- pass the raw `[]ChangeRequest` slice for record multipart change operations,
- and continue passing `InsertRemoveRequest` objects for insert/remove multipart operations.

### 4. CLI And README Alignment

The example CLI must match `projects/searchWebApi/specs/language-bindings/go.md`.

Required changes:

- rename subcommands to the current spec names such as `projects get`, `collections get`, `collections filter-list`, and `collections values`,
- add the missing resource groups and record subcommands required by the binding spec,
- keep the CLI small and direct, with each command calling the library without adding a second abstraction layer,
- and update README usage examples to the nested command form actually implemented by the CLI.

The CLI only needs enough flag coverage to satisfy the current binding spec and demonstrate usage of the library surface.

## Tests

Write failing tests first for:

- JSON non-2xx responses that still return structured bodies,
- `SWA-searchCacheControl` being emitted as a header and not a query parameter,
- multipart record change requests encoding the `request` form field as a JSON array,
- and CLI command structure matching the latest binding spec at least for the renamed and newly added top-level command groups.

After implementation:

- run `go test ./...`,
- run `go build ./...`,
- and verify the example CLI still builds through the module build.

## Success Criteria

This repair is successful when:

- callers can inspect decoded JSON result bodies even when the server returns a non-2xx status,
- `SWA-searchCacheControl` is sent only as a header on affected operations,
- multipart record change endpoints send the contract-correct `request` payload shape,
- and the example CLI plus README match the latest `go.md` binding spec structure.
