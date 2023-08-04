package hello

import "fmt"

const (
	EN                 = "english"
	ES                 = "spanish"
	FR                 = "french"
	englishHelloPrefix = "hello, "
	spanishHelloPrefix = "hola, "
	frenchHelloPrefix  = "bonjour, "
)

func Hello(name, language string) string {
	var helloPrefix string

	switch language {
	case EN:
		helloPrefix = englishHelloPrefix
	case ES:
		helloPrefix = spanishHelloPrefix
	case FR:
		helloPrefix = frenchHelloPrefix
	}

	if name == "" {
		name = "world"
	}
	return helloPrefix + name
}

func main() {
	fmt.Println(Hello("world", EN))
}
