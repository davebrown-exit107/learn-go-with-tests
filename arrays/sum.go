package arrays

func Sum(numbers []int) (total int) {
	for _, curNum := range numbers {
		total += curNum
	}
	return
}

func SumAll(arrays ...[]int) (sums []int) {
	for _, array := range arrays {
		if len(array) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(array))
		}
	}
	return
}

func SumAllTails(arrays ...[]int) (sums []int) {
	for _, array := range arrays {
		if len(array) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(array[1:]))
		}
	}
	return
}
