package hello

const (
	EN                 = "english"
	ES                 = "spanish"
	FR                 = "french"
	englishHelloPrefix = "hello, "
	spanishHelloPrefix = "hola, "
	frenchHelloPrefix  = "bonjour, "
)

// Takes a name and language and returns the equivalent "hello, name" in that language
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
