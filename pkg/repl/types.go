package repl

import (
	"github.com/kai-xlr/pokedexcli/pkg/pokeapi"
)

// Command represents a command available in the REPL.
// It contains the command name, a description for help text, and a callback function to execute.
type Command struct {
	Name        string
	Description string
	Callback    func(*Config, ...string) error
}

// Config holds the configuration and state for the REPL session
type Config struct {
	PokeapiClient    *pokeapi.Client
	NextLocationsURL *string
	PrevLocationsURL *string
	CaughtPokemon    map[string]pokeapi.Pokemon
}

// NewConfig creates a new REPL configuration
func NewConfig(client *pokeapi.Client) *Config {
	return &Config{
		PokeapiClient: client,
		CaughtPokemon: make(map[string]pokeapi.Pokemon),
	}
}

// CommandRegistry manages available commands
type CommandRegistry struct {
	commands map[string]Command
}

// NewCommandRegistry creates a new command registry
func NewCommandRegistry() *CommandRegistry {
	return &CommandRegistry{
		commands: make(map[string]Command),
	}
}

// Register adds a command to the registry
func (r *CommandRegistry) Register(name string, cmd Command) {
	r.commands[name] = cmd
}

// Get retrieves a command by name
func (r *CommandRegistry) Get(name string) (Command, bool) {
	cmd, exists := r.commands[name]
	return cmd, exists
}

// All returns all registered commands
func (r *CommandRegistry) All() map[string]Command {
	return r.commands
}
