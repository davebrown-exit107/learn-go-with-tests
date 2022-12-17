package main

import "fmt"

const englishHelloPrefix = "Hello"
const spanishHelloPrefix = "Hola"

func Hello(name, lang string) string {
	var prefix string
	if name == "" {
		name = "world"
	}

	if lang == "EN" || lang == "" {
		prefix = englishHelloPrefix
	}
	if lang == "ES" {
		prefix = spanishHelloPrefix
	}

	return fmt.Sprintf("%v, %v", prefix, name)
}

func main() {
	fmt.Println(Hello("world", "EN"))
}
