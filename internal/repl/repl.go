package repl

import "strings"

func CleanInput(text string) []string {
	lowerCase := strings.ToLower(text)
	trimmed := strings.Fields(lowerCase)
	return trimmed
}
