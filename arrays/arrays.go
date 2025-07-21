package arrays

func Sum(numbers []int) (sum int) {
	for _, val := range numbers {
		sum += val
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	numberOfArrays := len(numbersToSum)
	sums := make([]int, numberOfArrays)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return sums
}
