// Package pokeapi is a wrapper for interacting with the PokeAPI
package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations makes a GET request to the PokeAPI's location-area endpoint,
// retrieves a list of location data and unmarshals the received JSON into RespShallowLocations struct.
//
// It accepts a pointer to a string as an argument. If this pointer is not nil, it will use
// the passed string url to make the GET request. Else, it will use the default baseURL + "/location-area"
//
// It returns RespShallowLocations as a response or an error if any occurred during the process.
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	// Prepare new GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// Return zero-value response struct and error
		return RespShallowLocations{}, err
	}

	// Execute the request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		// Return zero-value response struct and error
		return RespShallowLocations{}, err
	}
	// Close response body once all operations finish
	defer resp.Body.Close()

	// Read all data from response body
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		// Return zero-value response struct and error
		return RespShallowLocations{}, err
	}

	// Prepare response struct
	locationsResp := RespShallowLocations{}
	// Unmarshal JSON data into response struct
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		// Return zero-value response struct and error
		return RespShallowLocations{}, err
	}

	// Return response and nil (indicating success)
	return locationsResp, nil
}
