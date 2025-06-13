package hello

import "fmt"

const enGreeting = "Hello"
const esGreeting = "Hola"

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
	}
	if lang == "EN" {
		greeting = enGreeting
	}
	return fmt.Sprintf("%s, %s", greeting, name)
}

func main() {
	fmt.Println(Hello("", "EN"))
}
