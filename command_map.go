package main

import (
	"errors"
	"fmt"
)

// function commandMapf is responsible for getting a list of locations
// from the next page of the pokeapi and print them to the console.
func commandMapf(cfg *config) error {
	// Makes a call to the pokeapi client to list locations from the next URL
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)

	// If the API call returns an error, the function returns the error
	if err != nil {
		return err
	}

	// Updates 'nextLocationsURL' and 'prevLocationsURL' fields of the config
	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	// Loops through each location in the results and print its name
	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

// function commandeMapb does almost the same as function commandMapf,
// but it works in reverse, getting a list of locations from the
// previous page of the pokeapi and print them to the console.
func commandMapb(cfg *config) error {
	// If 'prevLocationsURL' is nil, an error will be returned as there isn't a previous page
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	// Makes a call to the pokeapi client to list locations from the previous URL
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)

	// If the API call returns an error, the function returns the error
	if err != nil {
		return err
	}

	// Updates 'nextLocationsURL' and 'prevLocationsURL' fields of the config
	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	// Loops through each location in the results and print its name
	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
