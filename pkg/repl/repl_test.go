package repl

import (
	"reflect"
	"testing"
)

// TestCleanInput tests the cleanInput function with various input scenarios.
// It verifies that the function properly normalizes input by lowercasing text,
// trimming whitespace, and splitting words while preserving punctuation.
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
		{
			name:     "mixed case command",
			input:    "HELP",
			expected: []string{"help"},
		},
		{
			name:     "command with arguments",
			input:    "map forward 5",
			expected: []string{"map", "forward", "5"},
		},
		{
			name:     "unicode whitespace",
			input:    "\u00A0test\u00A0", // non-breaking spaces
			expected: []string{"test"},   // strings.Fields removes unicode whitespace
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := cleanInput(c.input)
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

func TestCommandRegistry(t *testing.T) {
	registry := NewCommandRegistry()

	// Test registering a command
	cmd := Command{
		Name:        "test",
		Description: "A test command",
		Callback:    func(cfg *Config, args ...string) error { return nil },
	}

	registry.Register("test", cmd)

	// Test retrieving the command
	retrieved, exists := registry.Get("test")
	if !exists {
		t.Error("expected command to exist")
	}

	if retrieved.Name != "test" {
		t.Errorf("expected name 'test', got '%s'", retrieved.Name)
	}

	// Test non-existent command
	_, exists = registry.Get("nonexistent")
	if exists {
		t.Error("expected command not to exist")
	}
}

func TestNewConfig(t *testing.T) {
	// Test that NewConfig creates a proper config
	cfg := NewConfig(nil)

	if cfg == nil {
		t.Error("expected config to be created")
	}

	if cfg.CaughtPokemon == nil {
		t.Error("expected CaughtPokemon map to be initialized")
	}

	if len(cfg.CaughtPokemon) != 0 {
		t.Error("expected CaughtPokemon map to be empty initially")
	}
}

func TestCommandStruct(t *testing.T) {
	// Test that Command struct works as expected
	testCallback := func(cfg *Config, args ...string) error {
		return nil
	}

	cmd := Command{
		Name:        "test",
		Description: "test command",
		Callback:    testCallback,
	}

	if cmd.Name != "test" {
		t.Errorf("expected name %q, got %q", "test", cmd.Name)
	}
	if cmd.Description != "test command" {
		t.Errorf("expected description %q, got %q", "test command", cmd.Description)
	}
	if cmd.Callback == nil {
		t.Error("expected callback to not be nil")
	}

	// Test callback execution
	cfg := NewConfig(nil)
	err := cmd.Callback(cfg)
	if err != nil {
		t.Errorf("expected no error from test callback, got %v", err)
	}
}

func TestCleanInputEdgeCases(t *testing.T) {
	// Additional edge cases for cleanInput
	cases := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "only tabs",
			input:    "\t\t\t",
			expected: []string{},
		},
		{
			name:     "only newlines",
			input:    "\n\n\n",
			expected: []string{},
		},
		{
			name:     "mixed whitespace only",
			input:    " \t\n \r ",
			expected: []string{},
		},
		{
			name:     "consecutive whitespace types",
			input:    "word1 \t\n word2",
			expected: []string{"word1", "word2"},
		},
		{
			name:     "very long input",
			input:    "This Is A Very Long Command With Many Words That Tests Input Processing",
			expected: []string{"this", "is", "a", "very", "long", "command", "with", "many", "words", "that", "tests", "input", "processing"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := cleanInput(c.input)
			if !reflect.DeepEqual(actual, c.expected) {
				t.Errorf("expected %v, got %v", c.expected, actual)
			}
		})
	}
}
