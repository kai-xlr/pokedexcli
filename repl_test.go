package main

import (
	"reflect"
	"testing"
	"time"

	"github.com/kai-xlr/pokedexcli/internal/pokeapi"
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
			expected: []string{"test"}, // strings.Fields removes unicode whitespace
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

func TestGetCommands(t *testing.T) {
	commands := getCommands()

	// Test that all expected commands exist
	expectedCommands := []string{"help", "explore", "map", "mapb", "exit"}
	for _, expectedCmd := range expectedCommands {
		if cmd, exists := commands[expectedCmd]; !exists {
			t.Errorf("expected command %q to exist", expectedCmd)
		} else {
			// Skip name comparison for explore command as it includes usage info
			if expectedCmd != "explore" && cmd.name != expectedCmd {
				t.Errorf("expected command name %q, got %q", expectedCmd, cmd.name)
			}
			if cmd.description == "" {
				t.Errorf("expected non-empty description for command %q", expectedCmd)
			}
			if cmd.callback == nil {
				t.Errorf("expected non-nil callback for command %q", expectedCmd)
			}
		}
	}

	// Test that no unexpected commands exist
	if len(commands) != len(expectedCommands) {
		t.Errorf("expected %d commands, got %d", len(expectedCommands), len(commands))
		for name := range commands {
			found := false
			for _, expected := range expectedCommands {
				if name == expected {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("unexpected command %q found", name)
			}
		}
	}
}

func TestConfigInitialization(t *testing.T) {
	// Test that a config can be properly initialized
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		pokeapiClient: &pokeClient,
	}

	if cfg.pokeapiClient == nil {
		t.Error("expected pokeapiClient to be initialized")
	}
	if cfg.nextLocationsURL != nil {
		t.Error("expected nextLocationsURL to be nil initially")
	}
	if cfg.prevLocationsURL != nil {
		t.Error("expected prevLocationsURL to be nil initially")
	}
}

func TestCliCommandStruct(t *testing.T) {
	// Test that cliCommand struct works as expected
	testCallback := func(cfg *config, args ...string) error {
		return nil
	}

	cmd := cliCommand{
		name:        "test",
		description: "test command",
		callback:    testCallback,
	}

	if cmd.name != "test" {
		t.Errorf("expected name %q, got %q", "test", cmd.name)
	}
	if cmd.description != "test command" {
		t.Errorf("expected description %q, got %q", "test command", cmd.description)
	}
	if cmd.callback == nil {
		t.Error("expected callback to not be nil")
	}

	// Test callback execution
	pokeClient := pokeapi.NewClient(1*time.Second, time.Minute)
	cfg := &config{
		pokeapiClient: &pokeClient,
	}
	err := cmd.callback(cfg)
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
