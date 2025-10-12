package commands

import (
	"fmt"
	"os"

	"github.com/kai-xlr/pokedexcli/pkg/repl"
)

// Exit returns the exit command implementation
func Exit() repl.Command {
	return repl.Command{
		Name:        "exit",
		Description: "Exit the Pokedex",
		Callback: func(cfg *repl.Config, args ...string) error {
			fmt.Println("Closing the Pokedex... Goodbye!")
			os.Exit(0)
			return nil
		},
	}
}
