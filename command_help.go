package main

import "fmt"

// The function 'commandHelp' is responsible for printing out
// the names and descriptions of all the possible commands
// that a user can use in the application.
//
// It needs a pointer as parameter (*config), which is a data structure holding
// all the user's current settings and preferences.
//
// It returns an error, allowing the caller to handle situations
// where the commands cannot be correctly printed to the console.
func commandHelp(config *config) error {
	// Blank print to add a new line for neatness
	fmt.Println()

	// Welcome message
	fmt.Println("Welcome to the Pokedex!")

	// General usage message
	fmt.Println("Usage:")

	// Blank print to neatly separate the usage from the commands
	fmt.Println()

	// Iterate over the commands obtained by calling 'getCommands'
	for _, cmd := range getCommands() {
		// Print the name of the command and its description
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	// Blank print to add a new line at the end
	fmt.Println()

	// Return nil as no error has occurred.
	return nil
}
