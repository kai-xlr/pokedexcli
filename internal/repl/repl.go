package repl

import "strings"

func cleanInput(text string) []string {
	lowerCase := strings.ToLower(text)
	trimmed := strings.Fields(lowerCase)
	return trimmed
}
