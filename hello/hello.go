package hello

import "fmt"

func Hello(name string) string {
	return "hello, " + name
}

func main() {
	fmt.Println(Hello("world"))
}
