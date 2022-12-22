package arrays

func Sum(numbers []int) (total int) {
	for _, curNum := range numbers {
		total += curNum
	}
	return
}

func SumAll(arrays ...[]int) []int {
	numArrays := len(arrays) // How many arrays are we being handed
	sums := make([]int, numArrays)

	for i, array := range arrays {
		sums[i] = Sum(array)
	}

	return sums
}
