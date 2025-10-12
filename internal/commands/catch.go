package commands

import (
	"fmt"
	"math/rand"

	"github.com/kai-xlr/pokedexcli/internal/errors"
	"github.com/kai-xlr/pokedexcli/pkg/repl"
)

// Catch returns the catch command implementation
func Catch() repl.Command {
	return repl.Command{
		Name:        "catch <pokemon_name>",
		Description: "Attempt to catch a pokemon",
		Callback: func(cfg *repl.Config, args ...string) error {
			if len(args) != 1 {
				return errors.NewValidationError("pokemon_name", "you must provide a pokemon name")
			}

			name := args[0]
			pokemon, err := cfg.PokeapiClient.GetPokemon(name)
			if err != nil {
				return errors.NewCommandError("catch", fmt.Sprintf("failed to fetch pokemon %s", name), err)
			}

			fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

			// Simple catch probability based on base experience
			// Higher base experience = harder to catch
			threshold := 50
			if pokemon.BaseExperience > 100 {
				threshold = 30
			} else if pokemon.BaseExperience > 200 {
				threshold = 15
			}

			if rand.Intn(100) > threshold {
				fmt.Printf("%s escaped!\n", pokemon.Name)
				return nil
			}

			fmt.Printf("%s was caught!\n", pokemon.Name)
			fmt.Println("You may now inspect it with the inspect command.")
			cfg.CaughtPokemon[pokemon.Name] = pokemon
			return nil
		},
	}
}
