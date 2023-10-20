// Package pokeapi is a wrapper for interacting with the PokeAPI
package pokeapi

import (
	"net/http"
	"time"

	"github.com/masteidel/pokedexcli/internal/pokecache"
)

// Client is a type that represents a custom HTTP client for making requests.
type Client struct {
	cache      pokecache.Cache // cache is  used for storing and retrieving data to reduce network requests
	httpClient http.Client     // httpClient makes the http requests to PokeAPI
}

// NewClient is a function that creates and returns a new Client object.
// It accepts two parameters: timeout and cacheInterval, both of type time.Duration.
func NewClient(timeout, cacheInterval time.Duration) Client {
	// The function returns a new Client instance.
	return Client{
		// cache is initialized with a new Cache object, with cacheInterval as the argument.
		// The pokecache.NewCache() function is assumed to return a new Cache instance.
		cache: pokecache.NewCache(cacheInterval),

		// httpClient is assigned a value of a new http.Client object.
		// The Timeout of this http.Client is set to the given timeout parameter.
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
