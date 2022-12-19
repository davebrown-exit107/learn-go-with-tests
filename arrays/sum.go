package arrays

func Sum(numbers []int) (total int) {
	for _, curNum := range numbers {
		total += curNum
	}
	return
}
