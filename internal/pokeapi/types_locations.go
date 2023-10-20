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

// Location defines the structure for the game location data
type Location struct {
	// EncounterMethodRates holds the data of encounter method rates
	EncounterMethodRates []struct {
		// EncounterMethod defines the method of encounter
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		// VersionDetails holds the details of the game version
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	// GameIndex represents the game index
	GameIndex int `json:"game_index"`
	// ID is the Identifier for the location
	ID int `json:"id"`
	// Location is the name of the location
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name string `json:"name"`
	// Names holds a list of names of the location in multiple languages
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	// PokemonEncounters holds the data related to encounters with Pokemon's in this location
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		// VersionDetails holds the details of the game version this encounter relates to
		VersionDetails []struct {
			// EncounterDetails holds the details related to this encounter
			EncounterDetails []struct {
				Chance          int           `json:"chance"`
				ConditionValues []interface{} `json:"condition_values"`
				MaxLevel        int           `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}
