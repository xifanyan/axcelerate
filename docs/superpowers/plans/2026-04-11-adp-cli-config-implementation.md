# ADP CLI Config Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Add universal `adp_config.json` support to the ADP CLI contract and implement the feature in the existing Go CLI so global flags may come from config and explicit command-line values always override config values.

**Architecture:** Update the authoritative language-agnostic CLI specs first, then add a small config-loading and value-resolution layer to the Go CLI entrypoint before client construction. Keep the implementation narrow: one config struct, one loader, explicit per-field precedence helpers, and focused CLI tests proving config-only resolution, overrides, and failure handling.

**Tech Stack:** Markdown specs, Go, `urfave/cli/v3`, Go standard library JSON/filepath/os packages, `go test`

---

## File Map

- Modify: `projects/adp/specs/cli.md`
  - Add the universal `--path` global flag.
  - Define `adp_config.json`, config keys, precedence rules, and examples.
- Modify: `projects/adp/specs/index.md`
  - Keep the global CLI summary consistent with `cli.md`.
- Modify: `projects/adp/src/go/cmd/adpgo/main.go`
  - Add config-file loading and global-value resolution before client construction.
  - Move `host`, `user`, and `password` from parse-time required flags to post-resolution validation.
- Modify: `projects/adp/src/go/cmd/adpgo/main_test.go`
  - Add CLI-level tests for config-only required values, CLI overrides, invalid config, and `--path` resolution.

Prefer keeping the feature inside `cmd/adpgo/main.go` unless that file becomes unreasonably crowded during implementation.

---

### Task 1: Update The Authoritative CLI Specs

**Files:**
- Modify: `projects/adp/specs/cli.md`
- Modify: `projects/adp/specs/index.md`

- [ ] **Step 1: Write the target spec checklist in your notes**

Use this checklist while editing:

```text
- cli.md global flags include --path with default /adp/rest/api/task
- cli.md defines adp_config.json as universal for all language CLIs
- cli.md says config may provide host, port, path, user, password, insecure, debug
- cli.md says explicit CLI flags override config values for all global flags
- cli.md says host/user/password are required after resolution, not necessarily as CLI-only inputs
- index.md summary includes --path and config precedence note
```

- [ ] **Step 2: Update `projects/adp/specs/cli.md`**

Edit the global flag table to this shape:

```md
| `--host` | string | — | ADP server host (required after config resolution) |
| `--port` | integer | 8443 | ADP server port |
| `--path` | string | `/adp/rest/api/task` | ADP task API path |
| `--user` | string | — | ADP username (required after config resolution) |
| `--password` | string | — | ADP password (required after config resolution) |
| `--insecure` | boolean | false | Skip TLS certificate verification |
| `--debug`, `-d` | boolean | false | Enable debug logging — traces request/response payloads |
```

Add a new section after global flags:

```md
## Config File

All language CLIs must support a config file named exactly `adp_config.json`.

Supported keys:

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

Add a precedence section:

```md
## Precedence Rules

Global values must resolve in this order:

1. Explicit command-line flag
2. `adp_config.json`
3. Built-in default, where one exists

Built-in defaults exist for `port`, `path`, `insecure`, and `debug`.
`host`, `user`, and `password` must be present after resolution or the CLI must fail with a user-facing error.
```

Update at least one example to show `--path` explicitly or show a config-backed invocation in prose.

- [ ] **Step 3: Update `projects/adp/specs/index.md`**

Replace the CLI summary block with this content:

```md
### CLI Interface

- Must support subcommands for each task
- Global flags: `--host`, `--port`, `--path`, `--user`, `--password`, `--insecure`, `--debug`
- Shared CLI config file: `adp_config.json`
- Explicit command-line flags override config-file values for all global flags
- `--port` default: 8443
- `--path` default: `/adp/rest/api/task`
- Example: `adpgo --host example.com --user adp --password adp list-entities --type singleMindServer`
- CLI naming: `[project][lang]` (e.g., `adpgo` for Go, `adppy` for Python, `adprs` for Rust)
```

- [ ] **Step 4: Review only the spec files you changed**

Check these exact conditions:

```text
- cli.md and index.md use the same global flag list
- both mention --path
- both describe config precedence consistently
- cli.md says all global flags can be overridden, not just debug
```

- [ ] **Step 5: Commit the spec changes**

```bash
git add projects/adp/specs/cli.md projects/adp/specs/index.md
git commit -m "docs: define ADP CLI config precedence"
```

---

### Task 2: Add Failing Go CLI Tests For Config Resolution

**Files:**
- Modify: `projects/adp/src/go/cmd/adpgo/main_test.go`

- [ ] **Step 1: Write the failing test for config-only required values**

Add this test:

```go
func TestListEntitiesCommandReadsGlobalConfigFile(t *testing.T) {
	tempDir := t.TempDir()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("Auth-Username"); got != "config-user" {
			t.Fatalf("Auth-Username = %q", got)
		}
		if got := r.Header.Get("Auth-Password"); got != "config-pass" {
			t.Fatalf("Auth-Password = %q", got)
		}
		_, _ = io.WriteString(w, `{"executionId":"1","taskType":"List Entities","loggingEnabled":"true","progressMax":1,"executionStatus":"success","executionRootDir":"root","contextId":"ctx","executionPersistent":"true","progressCurrent":1,"progressPercentage":1.0,"taskDisplayName":"List entities","executionMetaData":{"adp_entities_output_file_name":"output.json","adp_entities_json_output":"[]"}}`)
	}))
	defer server.Close()

	configPath := filepath.Join(tempDir, "adp_config.json")
	if err := os.WriteFile(configPath, []byte(`{"host":`+strconv.Quote(server.URL)+`,"user":"config-user","password":"config-pass","path":""}`), 0o600); err != nil {
		t.Fatalf("WriteFile error: %v", err)
	}

	oldWD, err := os.Getwd()
	if err != nil {
		t.Fatalf("Getwd error: %v", err)
	}
	defer func() { _ = os.Chdir(oldWD) }()
	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("Chdir error: %v", err)
	}

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd := newApp(stdout, stderr)

	err = cmd.Run(context.Background(), []string{"adpgo", "list-entities", "--type", "singleMindServer"})
	if err != nil {
		t.Fatalf("Run error: %v", err)
	}
}
```

- [ ] **Step 2: Run the test and verify it fails**

Run:

```bash
go test ./cmd/adpgo -run TestListEntitiesCommandReadsGlobalConfigFile -count=1
```

Expected: FAIL because `host`, `user`, and `password` are still parse-time required flags.

- [ ] **Step 3: Add failing precedence and config-error tests**

Add these test names and cover the listed behavior:

```go
func TestListEntitiesCommandCLIFlagsOverrideConfigFile(t *testing.T) {}

func TestRunReportsInvalidConfigFile(t *testing.T) {}

func TestRunReportsMissingRequiredGlobalsAfterConfigResolution(t *testing.T) {}
```

Required assertions:

- config file provides one set of `host`, `user`, and `password`
- CLI flags provide different values
- request uses CLI values, not config values
- malformed `adp_config.json` returns exit code `1` and prints a user-facing config error
- missing `host`, `user`, or `password` after resolution returns exit code `1` and prints a user-facing error

Import `path/filepath`, `os`, and `strconv` in the test file if they are not already present.

- [ ] **Step 4: Add the failing test for `--path` override behavior**

Add this test name:

```go
func TestListEntitiesCommandCLIPathOverridesConfigFile(t *testing.T) {}
```

Required assertions:

- config file sets `path` to one value
- CLI passes `--path` with a different value
- the request uses the CLI-selected path, not the config-selected path

- [ ] **Step 5: Run the focused CLI test set and confirm it fails**

Run:

```bash
go test ./cmd/adpgo -run "Test(ListEntitiesCommandReadsGlobalConfigFile|ListEntitiesCommandCLIFlagsOverrideConfigFile|RunReportsInvalidConfigFile|RunReportsMissingRequiredGlobalsAfterConfigResolution|ListEntitiesCommandCLIPathOverridesConfigFile)$" -count=1
```

Expected: FAIL on the new tests before implementation.

- [ ] **Step 6: Commit the failing tests**

```bash
git add projects/adp/src/go/cmd/adpgo/main_test.go
git commit -m "test: cover ADP CLI config resolution"
```

---

### Task 3: Implement Config Loading And Global Value Resolution In Go CLI

**Files:**
- Modify: `projects/adp/src/go/cmd/adpgo/main.go`

- [ ] **Step 1: Add a minimal config type and loader**

Add this code near the helper section in `main.go`:

```go
type cliConfigFile struct {
	Host     string `json:"host"`
	Port     *int   `json:"port"`
	Path     string `json:"path"`
	User     string `json:"user"`
	Password string `json:"password"`
	Insecure *bool  `json:"insecure"`
	Debug    *bool  `json:"debug"`
}

func loadCLIConfigFile() (cliConfigFile, error) {
	body, err := os.ReadFile("adp_config.json")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return cliConfigFile{}, nil
		}
		return cliConfigFile{}, err
	}

	var cfg cliConfigFile
	if err := json.Unmarshal(body, &cfg); err != nil {
		return cliConfigFile{}, fmt.Errorf("invalid adp_config.json: %w", err)
	}
	return cfg, nil
}
```

- [ ] **Step 2: Add explicit resolution helpers for string, int, and bool globals**

Add helpers with this shape:

```go
func resolvedString(cmd *cli.Command, name string, configValue string) string {
	if cmd.IsSet(name) {
		return cmd.String(name)
	}
	return configValue
}

func resolvedInt(cmd *cli.Command, name string, configValue *int) int {
	if cmd.IsSet(name) {
		return cmd.Int(name)
	}
	if configValue != nil {
		return *configValue
	}
	return cmd.Int(name)
}

func resolvedBool(cmd *cli.Command, name string, configValue *bool) bool {
	if cmd.IsSet(name) {
		return cmd.Bool(name)
	}
	if configValue != nil {
		return *configValue
	}
	return cmd.Bool(name)
}
```

- [ ] **Step 3: Remove parse-time requiredness from global flags only**

Change the global flags in `newApp` so these are no longer `Required: true`:

```go
&cli.StringFlag{Name: "host", Usage: "ADP server host"}
&cli.StringFlag{Name: "user", Usage: "ADP username"}
&cli.StringFlag{Name: "password", Usage: "ADP password"}
```

Do not change task-specific required flags.

- [ ] **Step 4: Make `newClient` config-aware**

Refactor `newClient` to this structure:

```go
func newClient(cmd *cli.Command) (*adp.Client, error) {
	cfg, err := loadCLIConfigFile()
	if err != nil {
		return nil, err
	}

	host := strings.TrimSpace(resolvedString(cmd, "host", cfg.Host))
	user := strings.TrimSpace(resolvedString(cmd, "user", cfg.User))
	password := strings.TrimSpace(resolvedString(cmd, "password", cfg.Password))
	if host == "" {
		return nil, errors.New("missing required global setting: host")
	}
	if user == "" {
		return nil, errors.New("missing required global setting: user")
	}
	if password == "" {
		return nil, errors.New("missing required global setting: password")
	}

	port := resolvedInt(cmd, "port", cfg.Port)
	path := resolvedString(cmd, "path", cfg.Path)
	insecure := resolvedBool(cmd, "insecure", cfg.Insecure)
	debug := resolvedBool(cmd, "debug", cfg.Debug)

	return adp.NewClient(adp.ClientConfig{
		BaseURL:  adp.MustBaseURL(host, port, path),
		Username: user,
		Password: password,
		Insecure: insecure,
		Debug:    debug,
		DebugOut: cmd.ErrWriter,
	})
}
```

- [ ] **Step 5: Keep task-specific parse-time validation unchanged**

Leave these patterns intact:

```go
&cli.StringFlag{Name: "engineName", Required: true}
&cli.StringFlag{Name: "csvFile", Required: true}
&cli.StringFlag{Name: "batchScriptPath", Required: true}
```

Only global `host`, `user`, and `password` move to post-resolution validation.

- [ ] **Step 6: Run the focused CLI config tests and verify they pass**

Run:

```bash
go test ./cmd/adpgo -run "Test(ListEntitiesCommandReadsGlobalConfigFile|ListEntitiesCommandCLIFlagsOverrideConfigFile|RunReportsInvalidConfigFile|RunReportsMissingRequiredGlobalsAfterConfigResolution|ListEntitiesCommandCLIPathOverridesConfigFile)$" -count=1
```

Expected: PASS.

- [ ] **Step 7: Run existing CLI regression tests to verify no behavior drift**

Run:

```bash
go test ./cmd/adpgo -run "Test(ListEntitiesCommandPrintsDecodedEntities|QueryEngineCommandParsesTaxonomyFlags|RunDoesNotDuplicateTaskExecutionErrors|CreateOcrJobCommandStartsWithoutWaitingByDefault|CreateOcrJobCommandWaitsWhenRequested|CSVMergeCommandRequiresCSVFile|CLICommandRequiresBatchScriptPath)$" -count=1
```

Expected: PASS.

- [ ] **Step 8: Commit the Go CLI implementation**

```bash
git add projects/adp/src/go/cmd/adpgo/main.go projects/adp/src/go/cmd/adpgo/main_test.go
git commit -m "feat: add ADP CLI global config resolution"
```

---

### Task 4: Final Verification

**Files:**
- Modify: none expected

- [ ] **Step 1: Run the full Go test suite**

Run:

```bash
go test ./... -count=1
```

Expected: PASS.

- [ ] **Step 2: Verify CLI help still shows the global options**

Run:

```bash
go run ./cmd/adpgo --help
```

Expected output includes:

```text
--host
--port
--path
--user
--password
--insecure
--debug, -d
```

- [ ] **Step 3: Re-run the config-focused CLI tests as a smoke check**

Run:

```bash
go test ./cmd/adpgo -run "Test(ListEntitiesCommandReadsGlobalConfigFile|ListEntitiesCommandCLIFlagsOverrideConfigFile|RunReportsInvalidConfigFile|RunReportsMissingRequiredGlobalsAfterConfigResolution|ListEntitiesCommandCLIPathOverridesConfigFile)$" -count=1
```

Expected: PASS.

- [ ] **Step 4: Record verification state**

If files changed during final verification, commit them:

```bash
git add projects/adp/src/go/cmd/adpgo/main.go projects/adp/src/go/cmd/adpgo/main_test.go projects/adp/specs/cli.md projects/adp/specs/index.md
git commit -m "test: verify ADP CLI config support"
```

If no files changed, do not create an empty commit. Report the verification commands in the final handoff instead.

---

## Spec Coverage Check

- Universal `--path` flag: Task 1
- Universal `adp_config.json` contract: Task 1
- CLI-overrides-config for all global flags: Tasks 1 and 3
- Go CLI support for config-backed `host` / `user` / `password`: Task 3
- Go CLI error handling for invalid config and missing required resolved values: Tasks 2 and 3
- Regression safety for existing CLI behavior: Tasks 2, 3, and 4

## Notes For The Implementer

- Keep the config feature scoped to global CLI settings only.
- Do not add task-specific config-file entries.
- Prefer direct, explicit resolution code over clever mutation of `urfave/cli` flag state.
- Keep user-facing config errors short and actionable.
- Do not reintroduce parse-time `Required: true` for global `host`, `user`, or `password` once config-backed resolution exists.
