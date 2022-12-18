package arrays

import "fmt"

func Sum(numbers [5]int) (total int) {
	for _, curNum := range numbers {
		total += curNum
	}
	return
}

func main() {
	fmt.Println("vim-go")
}
