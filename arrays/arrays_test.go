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
	checkSlices := func(t testing.TB, got, want []int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("more than one slice", func(t *testing.T) {
		numbersA := []int{1, 2, 3, 4, 5}
		numbersB := []int{6, 7, 8, 9, 10}

		got := arrays.SumAll(numbersA, numbersB)
		want := []int{15, 40}

		checkSlices(t, got, want)
	})
	t.Run("empty slices", func(t *testing.T) {
		numbersA := []int{}
		numbersB := []int{6, 7, 8, 9, 10}

		got := arrays.SumAll(numbersA, numbersB)
		want := []int{0, 40}

		checkSlices(t, got, want)
	})
}

func ExampleSumAll() {
	numbersA := []int{1, 2, 3, 4, 5}
	numbersB := []int{6, 7, 8, 9, 10}
	sum := arrays.SumAll(numbersA, numbersB)
	fmt.Println(sum)
	// Output: [15 40]
}
func ExampleSum() {
	numbers := []int{1, 2, 3, 4, 5}
	sum := arrays.Sum(numbers)
	fmt.Println(sum)
	// Output: 15
}
