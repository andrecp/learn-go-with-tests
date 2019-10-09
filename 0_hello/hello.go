package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func greetingPrefix(language string) string {
	switch language {
	case "English":
		return englishHelloPrefix
	case "Spanish":
		return spanishHelloPrefix
	case "French":
		return frenchHelloPrefix
	default:
		return "N/A"
	}
}

func hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "" {
		language = "English"
	}

	prefix := greetingPrefix(language)
	return prefix + name

}

func main() {
	fmt.Println(hello("world", "English"))
}
