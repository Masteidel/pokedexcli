package main

import (
	"testing"
)

// TestCleanInput is a test function that verifies the behavior of the cleanInput function.
// cleanInput function is assumed to process the input string and return the expected output.
func TestCleanInput(t *testing.T) {

	// Declaring an array of struct for test case scenarios
	cases := []struct {
		input    string
		expected []string
	}{
		// Test case scenario: empty string input should return empty []string
		{
			input:    "  ",
			expected: []string{},
		},
		// Test case scenario: string with a single word should return the single word
		{
			input:    "  hello  ",
			expected: []string{"hello"},
		},
		// Test case scenario: string with multiple words should return the words
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		// Test case scenario: string with capitalized word should return the words in lower case
		{
			input:    "  HellO  World  ",
			expected: []string{"hello", "world"},
		},
	}

	// Iterating over the declared test cases
	for _, c := range cases {
		// Executing the cleanInput function with each input
		actual := cleanInput(c.input)
		// Testing if length of expected output and actual output matches
		if len(actual) != len(c.expected) {
			t.Errorf("lengths don't match: '%v' vs '%v'", actual, c.expected)
			continue
		}
		// Testing if value of expected output and actual output matches
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("cleanInput(%v) == %v, expected %v", c.input, actual, c.expected)
			}
		}
	}
}
