# searchWebApi Transport

## Base URL

- Default base path: `/searchWebApi`
- Client configuration must allow callers to supply scheme, host, port, and optional path prefix around this base path.

## Supported Content Types

- `application/json` for standard request and response bodies
- `application/x-www-form-urlencoded` as an alternative encoding for operations where `API-SPEC.md` explicitly allows it
- `multipart/form-data` for operations that upload binary or streamed content
- `application/octet-stream` for binary download responses
- `application/x-ndjson` for streaming search results on the records search endpoint

## HTTP Method Rules

- Use the exact HTTP method defined by `API-SPEC.md`.
- If `API-SPEC.md` documents form-encoded POST as an alternative transport for an operation, the client may expose that option without changing the semantic operation name.

## Headers

- Clients must support standard authentication headers.
- Clients must support `SWA-SESSION` for stateful session reuse.
- Clients must support optional `SWA-SESSION-TYPE`.
- Clients must support optional MDC tracing headers `SWA-MDC-TOKEN` and `SWA-MDC-METHOD`.
- Clients must allow callers to request NDJSON via the `Accept` header where supported.

## Streaming Rules

- Binary endpoints return `application/octet-stream` and should be represented as streaming/binary responses in generated clients.
- The records search endpoint may return `application/x-ndjson`; in that mode, the first record is search metadata and later lines are records.
- NDJSON mode bypasses normal search cache control behavior where the raw contract says so.
