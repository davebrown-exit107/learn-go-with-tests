package arrays

// Sum returns the integer sum of all elements in an array
func Sum(arr []int) int {
	var sum int

	for _, v := range arr {
		sum += v
	}
	return sum
}

// SumAll returns an []int array of sums, one entry for each array passed in
func SumAll(arr ...[]int) []int {
	var sums = []int{}
	for _, v := range arr {
		sums = append(sums, Sum(v))
	}
	return sums
}

// SumAllTails returns an []int array of sums of tails, one entry for each array passed in
func SumAllTails(arr ...[]int) []int {
	var sums = []int{}
	for _, v := range arr {
		if len(v) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(v[1:]))
		}
	}
	return sums
}
