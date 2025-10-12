// Package main provides a command-line Pokédex application with an interactive REPL interface.
package main

import (
	"time"

	"github.com/kai-xlr/pokedexcli/internal/pokeapi"
)

// main is the entry point of the Pokédex CLI application.
// It starts the interactive REPL (Read-Eval-Print Loop) that allows users to interact with Pokémon data.
func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		caughtPokemon: map[string]pokeapi.Pokemon{},
		pokeapiClient: &pokeClient,
	}
	RunRepl(cfg)
}
