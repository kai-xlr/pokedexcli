# Pokédex CLI

A command-line interface for interacting with Pokémon data, built in Go.

## Features

- Command-line interface for Pokémon workflows (scaffolded)
- Input normalization utilities in `internal/repl`
- Modular structure for future expansion

## Requirements

- Go 1.25.1 or later

## Installation (from source)

Clone the repository and build the binary:

```bash
git clone https://github.com/kai-xlr/pokedexcli.git
cd pokedexcli
go build -o pokedexcli ./cmd/main
```

## Usage

Run the CLI application:

```bash
./pokedexcli
```

## Development

### Project structure

```
.
├── cmd/main/           # Application entry point
│   └── main.go        # Main executable (currently prints "Hello, World!")
├── internal/repl/     # Input normalization utilities with tests
│   ├── repl.go       # cleanInput: lowercase + whitespace tokenization
│   └── repl_test.go  # Unit tests for cleanInput
├── go.mod            # Module definition (go 1.25.1)
├── WARP.md           # Warp-specific guidance
└── README.md         # This file
```

### Build

```bash
go build -o pokedexcli ./cmd/main
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
