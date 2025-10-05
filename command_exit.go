package main

import (
	"fmt"
	"os"
)

// commandExit displays a goodbye message and terminates the application.
// This function is called when the user enters the "exit" command in the REPL.
func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
