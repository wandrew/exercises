package main

import "fmt"

// This says: "Hi.  If you are a type that implements a method called 'getGreetings' which returns a string, you are
// allowed to use any function that accepts a bot as a parameter."

// interface type - interfaces are implicit.  Types do not need to be declared as type 'bot'
type bot interface {
	// an interface definition is composed of a list of one or more function signatures
	// functionName(argtype1,argtype2,...) (returntype1,returntype2,...)
	getGreeting() string
}

// To be "in the club" as a function, you MUST meet every requirement be be considered of this type and use a function
// that requires this interface as a parameter

// Concrete types (map, struct, int, string, etc)
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

// If you meet the requirements of the 'bot' interface, you can use me
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	// we assume getGreeting for englishBot has very specific logic that cannot really apply to spanishBot
	return "Hello!"
}
func (sb spanishBot) getGreeting() string {
	return "¡Hola!"
}
