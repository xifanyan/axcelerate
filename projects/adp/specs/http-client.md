# HTTP Client Specification

## Base Configuration

| Property | Value | Required |
|----------|-------|----------|
| Base URL | Built from `--host`, `--port`, `--path` | Yes |
| HTTP Method | PUT | Yes |
| Content-Type | `application/json` | Yes |

## CLI Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--host` | string | required | ADP host |
| `--port` | int | 8443 | ADP port |
| `--path` | string | /adp/rest/api/task | API context path |
| `--user` | string | required | Username (sets Auth-Username header) |
| `--password` | string | required | Password (sets Auth-Password header) |
| `--insecure` | bool | false | Skip SSL certificate verification |

## Request Format

All requests must:
1. Use HTTP PUT method
2. Set `Content-Type: application/json` header
3. Include authentication headers (`--user`, `--password` flags)
4. Send JSON body

## Example Request

```
PUT /executeAdpTask HTTP/1.1
Host: <host>:<port>
Content-Type: application/json
Auth-Username: <user>
Auth-Password: <password>
```

## Client Options (Library)

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| baseUrl | string | required | ADP API base URL (e.g., https://host:8443/path) |
| username | string | required | Authentication username |
| password | string | required | Authentication password |
| insecure | bool | false | Skip SSL verification |
| timeout | duration | 30s | Request timeout |
