package main

import "fmt"

const (
	french  = "French"
	spanish = "Spanish"

	englishHelloPrefix = "Hello, "
	frenchHelloPrefix  = "Bonjour, "
	spanishHelloPrefix = "Hola, "
)

// Private function since it starts with a lowercase letter
func greetingPrefix(language string) (prefix string) { // Named return value
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func main() {
	fmt.Println(Hello("Dario", ""))
}
