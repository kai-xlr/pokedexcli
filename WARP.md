# WARP.md

This file provides guidance to WARP (warp.dev) when working with code in this repository.

Project summary
- Language/toolchain: Go (module mode)
- Module: github.com/kai-xlr/pokedexcli
- Minimum Go version (from go.mod): 1.21
- Entry-point: main.go (root directory)
- Architecture: Interactive REPL with modular command system
- Commands: help, exit (with extensible command architecture)
- Notes: No Makefile, CI, or linter config detected. Simple flat structure.

Common commands
- Build the CLI binary to the repo root:
```bash path=null start=null
go build -o pokedexcli .
```
- Run the built binary:
```bash path=null start=null
./pokedexcli
```
- Run all tests in all packages:
```bash path=null start=null
go test ./...
```
- Run tests verbosely for the current package:
```bash path=null start=null
go test -v .
```
- Run a single test (or filtered subset) in the current package:
```bash path=null start=null
go test -v -run 'TestCleanInput' .
```
- Static checks and formatting (no external linter configured):
```bash path=null start=null
# Vet for common issues
go vet ./...

# Format code in-place
gofmt -s -w .
```

High-level architecture and structure
- main.go
  - Program entry-point. Calls RunRepl() to start the interactive REPL.
  - Simple main function that delegates to the REPL implementation.

- repl.go
  - Core REPL implementation with command system architecture.
  - RunRepl() - Main REPL loop that reads input, processes commands, and handles responses.
  - cleanInput(text string) []string - Input normalization (lowercase, whitespace splitting).
  - getCommands() - Returns map of available commands with their callbacks.
  - cliCommand struct - Defines command structure (name, description, callback).

- command_help.go
  - Implementation of the 'help' command.
  - Lists available commands with descriptions.

- command_exit.go
  - Implementation of the 'exit' command.
  - Cleanly exits the application.

- repl_test.go
  - Unit tests for the cleanInput function.
  - Tests cover whitespace handling, lowercasing, punctuation retention, and edge cases.

- Root
  - go.mod declares the module and Go version (1.21).
  - A built binary (pokedexcli) may exist in the repo root after running the build command.

Development notes
- There is no Makefile or task runner; prefer direct go commands noted above.
- No repository-level linter configuration is present; use go vet and gofmt as shown.
- The current architecture uses a flat structure with all code in the main package.
- To add new commands: create a new command_*.go file with a function matching the callback signature, then add it to the getCommands() map in repl.go.
- For future expansion, consider moving to internal/ packages if the codebase grows significantly.
