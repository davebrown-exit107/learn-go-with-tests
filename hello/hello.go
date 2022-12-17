package main

import "fmt"

const englishHelloPrefix = "Hello"
const spanishHelloPrefix = "Hola"
const frenchHelloPrefix = "Bonjour"

func Hello(name, lang string) string {
	var prefix string
	if name == "" {
		name = "world"
	}

	switch lang {
	case "EN":
		prefix = englishHelloPrefix
	case "ES":
		prefix = spanishHelloPrefix
	case "FR":
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix

	}

	return fmt.Sprintf("%v, %v", prefix, name)
}

func main() {
	fmt.Println(Hello("world", "EN"))
}
