package main

import "fmt"

func Hello(name string) string {
	output := fmt.Sprintf("Hello, %v", name)
	return output
}

func main() {
	fmt.Println(Hello("world"))
}
