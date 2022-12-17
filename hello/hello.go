package main

import "fmt"

const englishHelloPrefix = "Hello"

func Hello(name string) string {
	if name == "" {
		name = "world"
	}

	return fmt.Sprintf("%v, %v", englishHelloPrefix, name)
}

func main() {
	fmt.Println(Hello("world"))
}
