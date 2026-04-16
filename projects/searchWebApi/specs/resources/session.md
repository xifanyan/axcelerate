# Session Resource

## Operations

### Login

- Raw operation: `POST /login`
- Raw operationId: `login`
- Result schema: `LoginResult`

### Logout

- Raw operation: `DELETE /logout`
- Raw operationId: `logout`
- Result schema: `LogoutResult`

## Shared Rules

- Explicit login/logout complement, but do not replace, the implicit session behavior described in `../auth-and-session.md`.
- Session identifiers are transported via `SWA-SESSION` headers rather than JSON body fields.
