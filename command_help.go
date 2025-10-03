package main

import (
	"fmt"
)

// commandHelp displays a welcome message and lists all available commands with their descriptions.
// This function is called when the user enters the "help" command in the REPL.
func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}
