package main

import (
	"github.com/masteidel/pokedexcli/internal/pokeapi"
	"time"
)

func main() {
	// Initializing a new client for the PokeAPI with a 5-second timeout
	pokeClient := pokeapi.NewClient(5 * time.Second)

	// Setting up configuration where pokeapiClient is the initialized client
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	// Starting the REPL (Read-Eval-Print Loop) with the given configuration
	startRepl(cfg)
}
