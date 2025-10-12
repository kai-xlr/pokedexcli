package commands

import (
	"fmt"

	"github.com/kai-xlr/pokedexcli/internal/errors"
	"github.com/kai-xlr/pokedexcli/pkg/repl"
)

// Explore returns the explore command implementation
func Explore() repl.Command {
	return repl.Command{
		Name:        "explore <location_name>",
		Description: "Explore a location",
		Callback: func(cfg *repl.Config, args ...string) error {
			if len(args) != 1 {
				return errors.NewValidationError("location_name", "you must provide a location name")
			}

			name := args[0]
			location, err := cfg.PokeapiClient.GetLocation(name)
			if err != nil {
				return errors.NewCommandError("explore", fmt.Sprintf("failed to explore location %s", name), err)
			}

			fmt.Printf("Exploring %s...\n", location.Name)
			fmt.Println("Found Pokemon:")
			for _, enc := range location.PokemonEncounters {
				fmt.Printf(" - %s\n", enc.Pokemon.Name)
			}
			return nil
		},
	}
}
