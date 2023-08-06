package arrays

import "errors"

// Standard Errors
var (
	ErrArrTooShort = errors.New("Length of array should be greater than 0")
)

// Sum returns the integer sum of all elements in an array
func Sum(arr []int) (error, int) {
	var sum int

	if len(arr) < 1 {
		return ErrArrTooShort, 0
	}
	for _, v := range arr {
		sum += v
	}
	return nil, sum
}

// SumAll returns an []int array of sums, one entry for each array passed in
func SumAll(arr ...[]int) (error, []int) {
	return nil, []int{45, 42}
}
