// Package pokeapi is a wrapper for interacting with the PokeAPI
package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations is a method function attached to Client object.
// It makes a request to PokeAPI to retrieve list of Pokemon locations.
// pageURL parameter allows to retrieve a specific page of the Pokemon locations list.
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	// building URL for the API request
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	// checking if the API response is in the cache
	if val, ok := c.cache.Get(url); ok {
		locationsResp := RespShallowLocations{}
		// unmarshalling the cached response into Go object
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			// returning empty response and error if unmarshalling failed
			return RespShallowLocations{}, err
		}

		// returning the response from the cache
		return locationsResp, nil
	}

	// creating new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// Return zero-value response struct and error
		return RespShallowLocations{}, err
	}

	// making the API call
	resp, err := c.httpClient.Do(req)
	if err != nil {
		// returning empty response and error if API call failed
		return RespShallowLocations{}, err
	}
	// Close response body once all operations finish
	defer resp.Body.Close()

	// Read all data from response body
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		// returning empty response and error if reading body failed
		return RespShallowLocations{}, err
	}

	// assigned the response to the locationsResp
	locationsResp := RespShallowLocations{}
	// Unmarshal JSON data into response struct
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		// returning empty response and error if unmarshalling failed
		return RespShallowLocations{}, err
	}

	// adding the API response to the cache
	c.cache.Add(url, dat)
	// returning the API response
	return locationsResp, nil
}
