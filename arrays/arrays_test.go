package arrays_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/davebrown-exit107/learn-go-with-tests/arrays"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}

	got := arrays.Sum(numbers)
	want := 36

	if got != want {
		t.Errorf("got %d want %d, given %v", got, want, numbers)
	}
}

func TestSumAll(t *testing.T) {
	numbersA := []int{1, 2, 3, 4, 5}
	numbersB := []int{6, 7, 8, 9, 10}

	got := arrays.SumAll(numbersA, numbersB)
	want := []int{15, 40}

	if !slices.Equal(got, want) {
		t.Errorf("got %d want %d", got, want)
	}
}

func ExampleSum() {
	numbers := []int{1, 2, 3, 4, 5}
	sum := arrays.Sum(numbers)
	fmt.Println(sum)
	// Output: 15
}
