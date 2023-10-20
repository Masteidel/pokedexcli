// Package pokeapi is a wrapper for interacting with the PokeAPI
package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// GetLocation is a method attached to the Client struct.
// This method retrieves a location object from the Pokemon API by its name.
// The location data is first searched in the client cache. If it is found,
// the cached version is returned. Otherwise, the API is queried directly.
// The API response is then cached for subsequent calls.
func (c *Client) GetLocation(locationName string) (Location, error) {
	// Construct the URL for the API call
	url := baseURL + "/location-area/" + locationName

	// Check if the location data is in the cache
	if val, ok := c.cache.Get(url); ok {
		// Create a new Location object
		locationResp := Location{}

		// Unmarshal the cached JSON data into the Location object
		err := json.Unmarshal(val, &locationResp)

		// Check for unmarshalling errors
		if err != nil {
			// If there's an error return an empty Location and the error
			return Location{}, err
		}

		// Return the cached Location if no unmarshalling error occured
		return locationResp, nil
	}

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// If there's an error creating the request, return an empty Location and the error
		return Location{}, err
	}

	// Execute the HTTP request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		// If there's an error doing the request, return an empty Location and the error
		return Location{}, err
	}
	defer resp.Body.Close()

	// Read all the content of the response body
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		// If there's an error reading the body, return an empty Location and the error
		return Location{}, err
	}

	// Create a Location object to hold the API response
	locationResp := Location{}

	// Unmarshal the response data into the Location object
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		// If there's an error unmarshalling the response, return an empty Location and the error
		return Location{}, err
	}

	// If everything went well, add the API response to the cache for future use
	c.cache.Add(url, dat)

	// Return the fetched Location and nil as there's no error
	return locationResp, nil
}
