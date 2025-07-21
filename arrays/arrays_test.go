package arrays_test

import (
	"fmt"
	"testing"

	"github.com/davebrown-exit107/learn-go-with-tests/arrays"
)

func TestSum(t *testing.T) {
	t.Run("collection of five numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := arrays.Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d, given %v", got, want, numbers)
		}
	})

	t.Run("collection of indeterminant size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}

		got := arrays.Sum(numbers)
		want := 36

		if got != want {
			t.Errorf("got %d want %d, given %v", got, want, numbers)
		}
	})

}
