package main

import "fmt"

func Sum(numbers [5]int) (total int) {
	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}
	return
}

func main() {
	fmt.Println("vim-go")
}
