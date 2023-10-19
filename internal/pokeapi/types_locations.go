// Package pokeapi is a wrapper for interacting with the PokeAPI
package pokeapi

// RespShallowLocations struct is used to unmarshal the JSON response
// from the PokeAPI for a shallow listing of locations.
type RespShallowLocations struct {
	// Count refers to the total number of locations.
	Count int `json:"count"`

	// Next is a pointer to a string that represents the URL to the next page of results
	// It is a pointer, so it can be null if there is no next page.
	Next *string `json:"next"`

	// Previous is a pointer to a string that represents the URL to the previous page of results.
	// It is a pointer, so it can be null if there is no previous page.
	Previous *string `json:"previous"`

	// Results is a slice of anonymous structs that each contain the name and URL of a location.
	Results []struct {
		// Name is the name of the location.
		Name string `json:"name"`

		// URL is the URL of the specific location in the PokeAPI.
		URL string `json:"url"`
	} `json:"results"`
}
