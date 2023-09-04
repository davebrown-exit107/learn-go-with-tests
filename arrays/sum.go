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
	sums := make([]int, len(arr))
	for i, v := range arr {
		curSum := Sum(v)
		sums[i] = curSum
	}
	return sums
}
