package hello

import "fmt"

const englishHelloPrefix = "Hello"
const spanishHelloPrefix = "Hola"
const frenchHelloPrefix = "Bonjour"

func greetingPrefix(lang string) (prefix string) {
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
	return
}

func Hello(name, lang string) string {
	if name == "" {
		name = "world"
	}

	prefix := greetingPrefix(lang)

	return fmt.Sprintf("%v, %v", prefix, name)
}

func main() {
	fmt.Println(Hello("world", "EN"))
}
