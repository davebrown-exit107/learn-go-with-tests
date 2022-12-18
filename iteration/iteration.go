package iteration

import "fmt"

func Repeat(character string, repeatCount int) (repeated string) {
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return
}

func main() {

	fmt.Println("vim-go")
}
