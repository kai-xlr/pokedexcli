# Pokédex CLI

A command-line interface for interacting with Pokémon data using the PokéAPI, built in Go.

## Features

- Interactive REPL (Read-Eval-Print Loop) interface
- Location exploration and navigation
- Pokémon catching and collection system
- HTTP response caching for improved performance
- Built-in command system with extensible architecture
- Input normalization and command parsing

## Requirements

- Go 1.21 or later

## Installation (from source)

Clone the repository and build the binary:

```bash
git clone https://github.com/kai-xlr/pokedexcli.git
cd pokedexcli
go build -o pokedexcli .
```

## Usage

Run the CLI application:

```bash
./pokedexcli
```

The application will start an interactive prompt where you can enter commands:

- `help` - Display available commands and their descriptions
- `map` - Get the next page of locations to explore
- `mapb` - Get the previous page of locations
- `explore <location_name>` - Explore a specific location to find Pokémon
- `catch <pokemon_name>` - Attempt to catch a Pokémon you've encountered
- `exit` - Exit the Pokédex CLI

## Development

### Project structure

Following the [Standard Go Project Layout](https://github.com/golang-standards/project-layout):

```
.
├── cmd/
│   └── pokedexcli/           # Main application entry point
│       └── main.go           # Application bootstrap and command registration
├── pkg/                      # Public packages (can be imported by other projects)
│   ├── pokeapi/              # PokéAPI client library
│   │   ├── client.go         # HTTP client with caching integration
│   │   ├── location_get.go   # Individual location data fetching
│   │   ├── location_list.go  # Location listing and pagination
│   │   ├── pokemon_get.go    # Pokémon data retrieval
│   │   ├── types_locations.go # Location-related type definitions
│   │   ├── types_pokemon.go  # Pokémon-related type definitions
│   │   └── pokeapi.go        # Package constants and base URL
│   ├── pokecache/            # HTTP response caching library
│   │   ├── pokecache.go      # TTL-based cache with goroutine cleanup
│   │   └── pokecache_test.go # Cache functionality tests
│   └── repl/                 # REPL framework (reusable)
│       ├── repl.go           # REPL runner and command registry
│       ├── repl_test.go      # REPL and input processing tests
│       └── types.go          # Command, Config, and Registry types
├── internal/                 # Private packages (application-specific)
│   ├── commands/             # Command implementations
│   │   ├── catch.go          # Pokémon catching logic
│   │   ├── exit.go           # Application exit
│   │   ├── explore.go        # Location exploration
│   │   ├── help.go           # Help system
│   │   ├── inspect.go        # Pokémon inspection
│   │   ├── map.go            # Location navigation (map/mapb)
│   │   └── pokedex.go        # Caught Pokémon listing
│   └── errors/               # Application-specific error types
│       └── errors.go         # CommandError, ValidationError, APIError
├── Makefile                  # Build and development tasks
├── go.mod                    # Module definition (go 1.21)
├── WARP.md                   # Warp-specific guidance
└── README.md                 # This file
```

### Build

Using Go directly:
```bash
go build -o pokedexcli ./cmd/pokedexcli
```

Using Make (recommended):
```bash
make build
```

### Test

Run all tests:

```bash
go test ./...
# or
make test
```

Verbose tests for a specific package:

```bash
go test -v ./pkg/repl
go test -v ./pkg/pokecache
```

Run a single test (or subset):

```bash
go test -v -run 'TestCleanInput' ./pkg/repl
```

### Code quality

Format code:

```bash
gofmt -s -w .
# or
make fmt
```

Static analysis:

```bash
go vet ./...
# or
make vet
```

### Development workflow

Set up development environment:
```bash
make dev-setup
```

Build and run:
```bash
make run
```

Install globally:
```bash
make install
# or
go install ./cmd/pokedexcli
```

## Architecture

This project follows Go best practices and the Standard Go Project Layout:

- **`cmd/`** - Main applications (entry points)
- **`pkg/`** - Library code that can be imported by other projects
- **`internal/`** - Private application code that cannot be imported

The REPL framework in `pkg/repl/` is designed to be reusable, while command implementations in `internal/commands/` are application-specific.

## Contributing

- Open an issue or PR with proposed changes
- Include or update tests where applicable
- Run `make dev-setup` to set up your environment
- Ensure `make test` passes and code is formatted with `make fmt`
- Follow the existing project structure and Go idioms
