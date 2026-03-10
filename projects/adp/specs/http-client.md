# HTTP Client Specification

## Base Configuration

| Property | Value | Required |
|----------|-------|----------|
| Base URL | Configurable via `--url` | Yes |
| HTTP Method | PUT | Yes |
| Content-Type | `application/json` | Yes |

## Authentication

| Header | Description | Required |
|--------|-------------|----------|
| Auth-Username | Username from `--username` flag | Yes |
| Auth-Password | Password from `--password` flag | Yes |

## SSL Configuration

| Flag | Description | Default |
|------|-------------|---------|
| `--insecure` | Skip SSL certificate verification | false |

## Request Format

All requests must:
1. Use HTTP PUT method
2. Set `Content-Type: application/json` header
3. Include authentication headers
4. Send JSON body

## Example Request Headers

```
PUT /executeAdpTask HTTP/1.1
Host: <configured-host>
Content-Type: application/json
Auth-Username: <username>
Auth-Password: <password>
```

## Client Options

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| baseUrl | string | required | ADP API base URL |
| username | string | required | Authentication username |
| password | string | required | Authentication password |
| insecure | bool | false | Skip SSL verification |
| timeout | duration | 30s | Request timeout |
