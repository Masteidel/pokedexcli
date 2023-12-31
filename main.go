package main

import (
	"time"

	"github.com/masteidel/pokedexcli/internal/pokeapi"
)

func main() {
	// Initializing a new client for the PokeAPI with a 5-second timeout
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)

	// Setting up configuration where pokeapiClient is the initialized client
	cfg := &config{
		caughtPokemon: map[string]pokeapi.Pokemon{},
		pokeapiClient: pokeClient,
	}

	// Starting the REPL (Read-Eval-Print Loop) with the given configuration
	startRepl(cfg)
}
