package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message1, err1 := greetings.Hello("Dave")
	fmt.Println(message1)
	if err1 != nil {
		log.Fatal(err1)
	}

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Request greeting messages for the names.
	messages3, err3 := greetings.Hellos(names)
	if err3 != nil {
		log.Fatal(err3)
	}
	// If no error was returned, print the returned map of
	// messages to the console.
	fmt.Println(messages3)

	// Get a greeting message and print it.
	message2, err2 := greetings.Hello("")
	// If an error was returned, print it and exit the program.
	if err2 != nil {
		log.Fatal(err2)
	}

	// If no error was returned, print the greeting message.
	fmt.Println(message2)
}
