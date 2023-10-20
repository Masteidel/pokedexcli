package main

import (
	"errors"
	"fmt"
)

// commandMapf retrieves a list of `locations` from the next page of
// the pokeapi and outputs these locations to the console.
func commandMapf(cfg *config, args ...string) error {
	// Outgoing API request to the pokeapi client to receive a list of locations
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)

	// Error handling for the API request
	if err != nil {
		return err
	}

	// Update the 'nextLocationsURL' and 'prevLocationsURL' fields in the config with the received response
	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	// Iterate over each location in the response and display its name
	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

// commandMapb retrieves a list of `locations` from the previous page
// of the pokeapi and outputs these locations to the console.
func commandMapb(cfg *config, args ...string) error {
	// If 'prevLocationsURL' is nil, it implies there are no more previous pages, and hence an error is returned
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	// Outgoing API request to the pokeapi client to receive a list of locations from the previous page
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)

	// Error handling for the API request
	if err != nil {
		return err
	}

	// Update the 'nextLocationsURL' and 'prevLocationsURL' fields in the config with the received response
	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	// Iterate over each location in the response and display its name
	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
