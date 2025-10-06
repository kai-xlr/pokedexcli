package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/kai-xlr/pokedexcli/internal/pokeapi"
)

// cliCommand represents a command available in the REPL.
// It contains the command name, a description for help text, and a callback function to execute.
type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	pokeapiClient    *pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

// RunRepl starts the Read-Eval-Print Loop for the Pokédex CLI.
// It continuously reads user input, processes commands, and displays results until the program exits.
func RunRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		cleaned := cleanInput(scanner.Text())
		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

	}
}

// cleanInput normalizes user input by converting to lowercase and splitting on whitespace.
// It returns a slice of words with leading/trailing whitespace removed and multiple spaces collapsed.
func cleanInput(text string) []string {
	lowerCase := strings.ToLower(text)
	trimmed := strings.Fields(lowerCase)
	return trimmed
}

// getCommands returns a map of all available commands in the REPL.
// The map keys are command names and values are cliCommand structs containing command details.
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
