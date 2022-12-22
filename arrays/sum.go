package arrays

func Sum(numbers []int) (total int) {
	for _, curNum := range numbers {
		total += curNum
	}
	return
}

func SumAll(arrays ...[]int) (sums []int) {
	for _, array := range arrays {
		sums = append(sums, Sum(array))
	}
	return
}

func SumAllTails(arrays ...[]int) (sums []int) {
	for _, array := range arrays {
		sums = append(sums, Sum(array[1:]))
	}
	return
}
