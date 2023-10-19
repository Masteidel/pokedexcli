// Package pokeapi is a wrapper for interacting with the PokeAPI
package pokeapi

import (
	"net/http"
	"time"
)

// Client is a type that represents a custom HTTP client for making requests.
type Client struct {
	httpClient http.Client // httpClient makes the http requests to PokeAPI
}

// NewClient creates a new instance of Client with a specified timeout duration.
// It takes a 'timeout' parameter of type time.Duration and sets the http request timeout with it.
// It returns a new Client instance.
func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout, // setting timeout for http requests
		},
	}
}
