package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func RunRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		cleaned := CleanInput(scanner.Text())
		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]
		fmt.Printf("Your command was: %s\n", commandName)
	}
}

func CleanInput(text string) []string {
	lowerCase := strings.ToLower(text)
	trimmed := strings.Fields(lowerCase)
	return trimmed
}
