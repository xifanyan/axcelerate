# searchWebApi Authentication And Session

## Authentication Modes

- HTTP basic authentication
- Bearer-token authentication

## Session Model

The API is stateful.

- A successful login creates a session.
- The session identifier is returned in response header `SWA-SESSION`.
- Subsequent requests must send `SWA-SESSION` to reuse the same session.
- If credentials or a bearer token are present, non-session endpoints may implicitly create a new session after timeout.

## Explicit Session Operations

- `POST /login` creates a new session explicitly.
- `DELETE /logout` closes the current session and associated resources.

## Session Type

- Clients may send `SWA-SESSION-TYPE`.
- Documented values include `MONITORING`, `USER`, and `DEFAULT`.
- Session type is a transport concern and should be configurable at the root client level.

## Tracing Headers

- Clients may send `SWA-MDC-TOKEN` to propagate trace identifiers.
- Clients may send `SWA-MDC-METHOD` to propagate logical method names.

## Client Responsibilities

- Preserve session state between requests when the caller chooses session reuse.
- Expose enough control for callers to clear, replace, or ignore session state.
- Treat session expiry as a first-class protocol condition, not as a hidden implementation detail.
