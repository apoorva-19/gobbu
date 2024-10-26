package main

import (
	"fmt"
)

const (
	spanish = "Spanish"
	french  = "French"
	italian = "Italian"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	italianHelloPrefix = "Bonjourno, "
)

/*
Hello generates a greeting based on the language

Parameters:
  - `name` : the name of the person
  - `language` : the language in which the greeting should be generated

Returns:
  - Greeting in the requested language
*/
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

/*
greetingPrefix returns the greeting prefix based on the requested language

Parameters:
  - `language` : the language in which the greeting should be generated

Returns:
  - `prefix` : greeting prefix for the requested language
*/
func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case italian:
		prefix = italianHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Apoorva", "Italian"))
}
