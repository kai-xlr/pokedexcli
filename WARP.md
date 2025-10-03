# WARP.md

This file provides guidance to WARP (warp.dev) when working with code in this repository.

Project summary
- Language/toolchain: Go (module mode)
- Module: github.com/kai-xlr/pokedexcli
- Minimum Go version (from go.mod): 1.25.1
- Entry-point: cmd/main/main.go
- Packages: internal/repl (with unit tests)
- Notes: No Makefile, CI, or linter config detected. README.md is minimal.

Common commands
- Build the CLI binary to the repo root:
```bash path=null start=null
go build -o pokedexcli ./cmd/main
```
- Run the built binary:
```bash path=null start=null
./pokedexcli
```
- Run all tests in all packages:
```bash path=null start=null
go test ./...
```
- Run tests verbosely for the repl package:
```bash path=null start=null
go test -v ./internal/repl
```
- Run a single test (or filtered subset) in the repl package:
```bash path=null start=null
go test -v -run 'TestCleanInput' ./internal/repl
```
- Static checks and formatting (no external linter configured):
```bash path=null start=null
# Vet for common issues
go vet ./...

# Format code in-place
gofmt -s -w .
```

High-level architecture and structure
- cmd/main/main.go
  - Program entry-point. Currently prints a placeholder message ("Hello, World!").
  - This is the only main package; building from ./cmd/main produces the pokedexcli binary.

- internal/repl
  - Purpose: Shared library code for the CLI (currently only input normalization).
  - cleanInput(text string) []string
    - Lowercases the input string and splits on whitespace (collapses multiple spaces, handles tabs/newlines).
    - Leaves punctuation intact while normalizing case.
  - Unit tests (internal/repl/repl_test.go) cover whitespace handling, lowercasing, punctuation retention, and non-ASCII behavior.

- Root
  - go.mod declares the module and Go version.
  - A built binary (pokedexcli) may exist in the repo root after running the build command.

Development notes
- There is no Makefile or task runner; prefer direct go commands noted above.
- No repository-level linter configuration is present; use go vet and gofmt as shown.
- If you add new packages intended for internal use by the CLI, prefer placing them under internal/ to keep APIs constrained to this module.
