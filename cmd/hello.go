package main

import "fmt"

const (
	spanish  = "Spanish"
	french   = "French"
	japanese = "Japanese"

	englishHelloPrefix  = "Hello, "
	spanishHelloPrefix  = "Hola, "
	frenchHelloPrefix   = "Bonjour, "
	japaneseHelloPrefix = "Kon'nichiwa, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case japanese:
		prefix = japaneseHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("Mario", "Spanish"))
}
