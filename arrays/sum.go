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
