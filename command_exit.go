package main

import "os"

// The function commandExit is used to cleanly terminate the program.
// It receives the program's configuration as an argument.
// Although the function has an error return value, it actually always returns nil
// because the program exits before the return statement when the exit command is received.
func commandExit(config *config, args ...string) error {
	// os.Exit(0) is used to end the program
	// The argument 0 is a code that signifies the program has ended successfully
	os.Exit(0)

	// This is a placeholder return statement.
	// As the program has already exited before this point, it is never executed.
	return nil
}
