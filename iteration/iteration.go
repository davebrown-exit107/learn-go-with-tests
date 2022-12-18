package iteration

import "fmt"

func Repeat(input string) (repeated string) {
	for i := 0; i < 5; i++ {
		repeated += input
	}
	return
}

func main() {

	fmt.Println("vim-go")
}
