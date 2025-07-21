package arrays

func Sum(numbers []int) (sum int) {
	for _, val := range numbers {
		sum += val
	}
	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}
