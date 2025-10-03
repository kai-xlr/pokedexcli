// Package main provides the entry point for the pokedexcli application.
// This is a command-line interface for interacting with Pokémon data.
package main

import (
	"github.com/kai-xlr/pokedexcli/internal/repl"
)

func main() {
	repl.RunRepl()
}
