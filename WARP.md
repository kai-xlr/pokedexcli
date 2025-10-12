# WARP.md

This file provides guidance to WARP (warp.dev) when working with code in this repository.

Project summary
- Language/toolchain: Go (module mode)
- Module: github.com/kai-xlr/pokedexcli
- Minimum Go version (from go.mod): 1.21
- Entry-point: cmd/pokedexcli/main.go (follows Standard Go Project Layout)
- Architecture: Interactive REPL with modular command system, HTTP API client with caching
- Commands: help, exit, map, mapb, explore, catch, inspect, pokedex (extensible command architecture)
- External API: PokéAPI (https://pokeapi.co/api/v2/) for Pokémon data
- Build system: Makefile with development workflow automation
- Notes: Follows Standard Go Project Layout with pkg/ and internal/ separation

Common commands
- Build the CLI binary (recommended way):
```bash path=null start=null
make build
```
- Build the CLI binary (direct Go):
```bash path=null start=null
go build -o pokedexcli ./cmd/pokedexcli
```
- Run the built binary:
```bash path=null start=null
./pokedexcli
# or
make run
```
- Run all tests in all packages:
```bash path=null start=null
go test ./...
# or
make test
```
- Run tests verbosely for a specific package:
```bash path=null start=null
go test -v ./pkg/repl
go test -v ./pkg/pokecache
```
- Run a single test (or filtered subset):
```bash path=null start=null
go test -v -run 'TestCleanInput' ./pkg/repl
```
- Static checks and formatting:
```bash path=null start=null
# Vet for common issues
go vet ./...
# or
make vet

# Format code in-place
gofmt -s -w .
# or
make fmt
```
- Development setup:
```bash path=null start=null
make dev-setup  # Run mod tidy, fmt, and test
```

High-level architecture and structure

Follows the Standard Go Project Layout:

- cmd/pokedexcli/main.go
  - Program entry-point. Initializes PokéAPI client with HTTP caching.
  - Creates REPL configuration and registers all commands.
  - Bootstraps the application and starts the REPL.

- pkg/repl/ (Reusable REPL framework)
  - repl.go - Core REPL implementation with command registry.
  - types.go - Command, Config, and CommandRegistry type definitions.
  - repl_test.go - Tests for input processing and command registry.

- pkg/pokeapi/ (Reusable API client library)
  - client.go - HTTP client with integrated caching.
  - location_*.go - Location-related API operations.
  - pokemon_get.go - Pokémon data fetching.
  - types_*.go - API response type definitions.

- pkg/pokecache/ (Reusable caching library)
  - pokecache.go - TTL-based HTTP response cache.
  - pokecache_test.go - Cache functionality tests.

- internal/commands/ (Application-specific command implementations)
  - help.go - Help command with dynamic command listing.
  - exit.go - Application exit command.
  - map.go - Location navigation commands (map/mapb).
  - explore.go - Location exploration with Pokémon encounters.
  - catch.go - Pokémon catching with probability mechanics.
  - inspect.go - Detailed view of caught Pokémon.
  - pokedex.go - List all caught Pokémon.

- internal/errors/ (Application-specific error types)
  - errors.go - CommandError, ValidationError, APIError types.
  - Proper error wrapping and contextual error messages.

- Root files
  - go.mod declares the module and Go version (1.21).
  - Makefile provides development workflow automation.
  - A built binary (pokedexcli) may exist in the repo root after building.

Development notes
- Project follows the Standard Go Project Layout for better organization and maintainability.
- Makefile provides automated development workflows (build, test, fmt, vet, install).
- pkg/ packages are designed to be reusable by other projects.
- internal/ packages contain application-specific logic that cannot be imported.
- To add new commands: create a new file in internal/commands/ returning a repl.Command, then register it in cmd/pokedexcli/main.go.
- The REPL framework is generic and could be reused for other CLI applications.
- Proper error types provide better error handling and debugging.
- HTTP client includes automatic caching with configurable TTL to reduce API calls.
- Configuration is managed through the repl.Config struct with dependency injection.
- Caught Pokémon are stored in memory and reset on application restart.
- Use `make dev-setup` to initialize the development environment.
- Use `make help` to see all available Make targets.
