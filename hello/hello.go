package main

import "fmt"

const englishHelloPrefix = "Hello"

func Hello(name string) string {
	output := fmt.Sprintf("%v, %v", englishHelloPrefix, name)
	return output
}

func main() {
	fmt.Println(Hello("world"))
}
