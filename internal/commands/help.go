package commands

import (
	"fmt"

	"github.com/kai-xlr/pokedexcli/pkg/repl"
)

// Help returns the help command implementation
func Help(replInstance *repl.REPL) repl.Command {
	return repl.Command{
		Name:        "help",
		Description: "Displays a help message",
		Callback: func(cfg *repl.Config, args ...string) error {
			fmt.Println()
			fmt.Println("Welcome to the Pokedex!")
			fmt.Println("Usage:")
			fmt.Println()

			availableCommands := replInstance.GetCommands()
			for _, cmd := range availableCommands {
				fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
			}
			fmt.Println()
			return nil
		},
	}
}
