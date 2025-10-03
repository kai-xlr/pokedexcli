package repl

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "trims outer spaces and splits words",
			input:    "  hello  world ",
			expected: []string{"hello", "world"},
		},
		{
			name:     "empty string returns empty slice",
			input:    "",
			expected: []string{},
		},
		{
			name:     "only spaces returns empty slice",
			input:    "  ",
			expected: []string{},
		},
		{
			name:     "single space returns empty slice",
			input:    " ",
			expected: []string{},
		},
		{
			name:     "collapses multiple internal spaces",
			input:    "a    b    c",
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "handles tabs and newlines",
			input:    "foo\tbar\nbaz",
			expected: []string{"foo", "bar", "baz"},
		},
		{
			name:     "lowercases input",
			input:    "HeLLo   WoRLD",
			expected: []string{"hello", "world"},
		},
		{
			name:     "keeps punctuation but lowercases letters",
			input:    "Hello, world!",
			expected: []string{"hello,", "world!"},
		},
		{
			name:     "single word",
			input:    "pokemon",
			expected: []string{"pokemon"},
		},
		{
			name:     "leading/trailing whitespace with newlines",
			input:    "\t  Pikachu  \n ",
			expected: []string{"pikachu"},
		},
		{
			name:     "non-ascii characters are preserved and lowercased where applicable",
			input:    "¡Hola!\nMundo",
			expected: []string{"¡hola!", "mundo"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := CleanInput(c.input)
			if len(actual) != len(c.expected) {
				t.Fatalf("expected %d words, got %d (actual=%v)", len(c.expected), len(actual), actual)
			}
			for i := range actual {
				word := actual[i]
				expectedWord := c.expected[i]
				if word != expectedWord {
					t.Errorf("index %d: expected %q, got %q", i, expectedWord, word)
				}
			}
		})
	}
}
