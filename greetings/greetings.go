package greetings

import "fmt"

// Returns a greeting for a named person
func Hello(name string) string {
	// Return a greeting that embeds the name in the message
	message := fmt.Sprintf("Hi, %v. Greetings!", name)
	return message
}
