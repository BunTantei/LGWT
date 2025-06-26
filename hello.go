package main

import "fmt"

const (
	englishHelloPrefix  = "Hello, "
	spanishHelloPrefix  = "Hola, "
	spanishLanguage     = "Spanish"
	frenchHelloPrefix   = "Bonjour, "
	frenchLanguage      = "French"
	japaneseLanguage    = "Japanese"
	japaneseHelloPrefix = "こんにちは, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	prefix := englishHelloPrefix

	switch language {
	case spanishLanguage:
		prefix = spanishHelloPrefix
	case frenchLanguage:
		prefix = frenchHelloPrefix
	case japaneseLanguage:
		prefix = japaneseHelloPrefix
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("World", ""))
}
