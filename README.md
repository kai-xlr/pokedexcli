# Pokédex CLI

A command-line interface for interacting with Pokémon data, built in Go.

## Features

- Interactive REPL (Read-Eval-Print Loop) interface
- Built-in command system with help and exit commands
- Input normalization and command parsing
- Modular command architecture for easy extension

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
- `exit` - Exit the Pokédex CLI

## Development

### Project structure

```
.
├── main.go           # Application entry point - starts the REPL
├── repl.go           # REPL implementation with command system
├── repl_test.go      # Unit tests for input processing
├── command_help.go   # Help command implementation
├── command_exit.go   # Exit command implementation
├── go.mod            # Module definition (go 1.21)
├── WARP.md           # Warp-specific guidance
└── README.md         # This file
```

### Build

```bash
go build -o pokedexcli .
```

### Test

Run all tests:

```bash
go test ./...
```

Verbose tests for a package:

```bash
go test -v ./internal/repl
```

Run a single test (or subset):

```bash
go test -v -run 'TestCleanInput' ./internal/repl
```

### Code quality

Format code:

```bash
gofmt -s -w .
```

Static analysis:

```bash
go vet ./...
```

## Contributing

- Open an issue or PR with proposed changes.
- Include or update tests where applicable.
- Ensure `go test ./...` passes and code is formatted.
