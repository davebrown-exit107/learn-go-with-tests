package arrays

import "errors"

// Standard Errors
var (
	ArrTooShort = errors.New("Length of array should be greater than 0")
)

// Sum returns the integer sum of all elements in an array
func Sum(arr []int) (error, int) {
	var sum int

	if len(arr) < 1 {
		return ArrTooShort, 0
	}
	for _, v := range arr {
		sum += v
	}
	return nil, sum
}
