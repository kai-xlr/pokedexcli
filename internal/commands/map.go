package commands

import (
	"fmt"

	"github.com/kai-xlr/pokedexcli/internal/errors"
	"github.com/kai-xlr/pokedexcli/pkg/repl"
)

// MapForward returns the map command implementation for forward navigation
func MapForward() repl.Command {
	return repl.Command{
		Name:        "map",
		Description: "Get the next page of locations",
		Callback: func(cfg *repl.Config, args ...string) error {
			locationsResp, err := cfg.PokeapiClient.ListLocations(cfg.NextLocationsURL)
			if err != nil {
				return errors.NewCommandError("map", "failed to fetch locations", err)
			}

			cfg.NextLocationsURL = locationsResp.Next
			cfg.PrevLocationsURL = locationsResp.Previous

			for _, loc := range locationsResp.Results {
				fmt.Println(loc.Name)
			}
			return nil
		},
	}
}

// MapBackward returns the mapb command implementation for backward navigation
func MapBackward() repl.Command {
	return repl.Command{
		Name:        "mapb",
		Description: "Get the previous page of locations",
		Callback: func(cfg *repl.Config, args ...string) error {
			if cfg.PrevLocationsURL == nil {
				return errors.NewCommandError("mapb", "you're on the first page", nil)
			}

			locationsResp, err := cfg.PokeapiClient.ListLocations(cfg.PrevLocationsURL)
			if err != nil {
				return errors.NewCommandError("mapb", "failed to fetch locations", err)
			}

			cfg.NextLocationsURL = locationsResp.Next
			cfg.PrevLocationsURL = locationsResp.Previous

			for _, loc := range locationsResp.Results {
				fmt.Println(loc.Name)
			}
			return nil
		},
	}
}
