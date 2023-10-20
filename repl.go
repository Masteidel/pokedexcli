package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/masteidel/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client // Client for Pokeapi
	nextLocationsURL *string        // URL of next page of locations
	prevLocationsURL *string        // URL of previous page of locations
}

// Function to start the REPL
func startRepl(cfg *config) {
	// create a new scanner for reading from the standard input
	reader := bufio.NewScanner(os.Stdin)

	// infinite for loop to keep the REPL running until forced exit
	for {
		// prompt for the user
		fmt.Print("Pokedex > ")

		// scan the next line from the standard input
		reader.Scan()

		// clean the input and split it into words
		words := cleanInput(reader.Text())
		// if there are no words, continue with the next iteration
		if len(words) == 0 {
			continue
		}

		// get the command name (first word)
		commandName := words[0]

		// Initialize an empty string slice for arguments
		args := []string{}

		// Check if the length of words is greater than 1
		if len(words) > 1 {
			// If true, slice the words array from index 1 to end, and assign it to args
			args = words[1:]
		}

		// get the corresponding command struct if exists
		command, exists := getCommands()[commandName]
		if exists {
			// execute the command callback and check for errors
			err := command.callback(cfg, args...)
			if err != nil {
				// if there is an error, print it and continue with the next iteration
				fmt.Println(err)
			}
			continue
		} else {
			// if command does not exist, print an error message and continue with the next iteration
			fmt.Println("Unknown command")
			continue
		}
	}
}

// Function to clean the input, convert it to lower case and split it into words
func cleanInput(text string) []string {
	output := strings.ToLower(text) // convert the input to lower case
	words := strings.Fields(output) // split the lower cased input text into words
	return words                    // return the slice of words
}

// struct to hold the details of a CLI command
type cliCommand struct {
	name        string                         // name of the command
	description string                         // description of the command
	callback    func(*config, ...string) error // callback function to be executed when the command is called
}

// Function to get a map of available commands
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": { // Help command details
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": { // Map command details
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": { // Mapb command details
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"explore": { // Explore command details
			name:        "explore <location_name>",
			description: "Explore a location",
			callback:    commandExplore,
		},
		"exit": { // Exit command details
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
