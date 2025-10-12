package commands

import (
	"fmt"

	"github.com/kai-xlr/pokedexcli/internal/errors"
	"github.com/kai-xlr/pokedexcli/pkg/repl"
)

// Inspect returns the inspect command implementation
func Inspect() repl.Command {
	return repl.Command{
		Name:        "inspect <pokemon_name>",
		Description: "View details about a caught Pokemon",
		Callback: func(cfg *repl.Config, args ...string) error {
			if len(args) != 1 {
				return errors.NewValidationError("pokemon_name", "you must provide a pokemon name")
			}

			name := args[0]
			pokemon, ok := cfg.CaughtPokemon[name]
			if !ok {
				return errors.NewCommandError("inspect", "you have not caught that pokemon", nil)
			}

			fmt.Println("Name:", pokemon.Name)
			fmt.Println("Height:", pokemon.Height)
			fmt.Println("Weight:", pokemon.Weight)
			fmt.Println("Stats:")
			for _, stat := range pokemon.Stats {
				fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
			}
			fmt.Println("Types:")
			for _, typeInfo := range pokemon.Types {
				fmt.Println("  -", typeInfo.Type.Name)
			}
			return nil
		},
	}
}
