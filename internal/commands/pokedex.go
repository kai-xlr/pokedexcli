package commands

import (
	"fmt"

	"github.com/kai-xlr/pokedexcli/pkg/repl"
)

// Pokedex returns the pokedex command implementation
func Pokedex() repl.Command {
	return repl.Command{
		Name:        "pokedex",
		Description: "See all the pokemon you've caught",
		Callback: func(cfg *repl.Config, args ...string) error {
			if len(cfg.CaughtPokemon) == 0 {
				fmt.Println("You haven't caught any pokemon yet!")
				return nil
			}

			fmt.Println("Your Pokedex:")
			for _, p := range cfg.CaughtPokemon {
				fmt.Printf(" - %s\n", p.Name)
			}
			return nil
		},
	}
}
