package hello

import "fmt"

const enGreeting = "Hello"
const esGreeting = "Hola"
const frGreeting = "Bonjour"

func Hello(name string, lang string) string {
	if name == "" {
		name = "world"
	}

	var greeting string
	switch lang {
	case "EN":
		greeting = enGreeting
	case "ES":
		greeting = esGreeting
	case "FR":
		greeting = frGreeting
	}

	return fmt.Sprintf("%s, %s", greeting, name)
}

func main() {
	fmt.Println(Hello("", "EN"))
}
