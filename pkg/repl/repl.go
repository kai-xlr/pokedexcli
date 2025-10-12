package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// REPL represents the Read-Eval-Print Loop
type REPL struct {
	registry *CommandRegistry
	config   *Config
}

// New creates a new REPL instance
func New(config *Config) *REPL {
	return &REPL{
		registry: NewCommandRegistry(),
		config:   config,
	}
}

// RegisterCommand registers a new command with the REPL
func (r *REPL) RegisterCommand(name string, cmd Command) {
	r.registry.Register(name, cmd)
}

// Run starts the Read-Eval-Print Loop for the Pokédex CLI.
// It continuously reads user input, processes commands, and displays results until the program exits.
func (r *REPL) Run() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		cleaned := cleanInput(scanner.Text())
		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		command, exists := r.registry.Get(commandName)
		if exists {
			err := command.Callback(r.config, args...)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
			continue
		} else {
			fmt.Println("Unknown command. Type 'help' for available commands.")
			continue
		}
	}
}

// GetCommands returns all registered commands
func (r *REPL) GetCommands() map[string]Command {
	return r.registry.All()
}

// cleanInput normalizes user input by converting to lowercase and splitting on whitespace.
// It returns a slice of words with leading/trailing whitespace removed and multiple spaces collapsed.
func cleanInput(text string) []string {
	lowerCase := strings.ToLower(text)
	trimmed := strings.Fields(lowerCase)
	return trimmed
}
