package arrays_test

import (
	"fmt"
	"testing"

	"github.com/davebrown-exit107/learn-go-with-tests/arrays"
	"github.com/google/go-cmp/cmp"
)

func TestSum(t *testing.T) {
	t.Run("sum an array of ints", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := arrays.Sum(numbers)
		want := 15

		assertEqual(t, got, want)
	})

}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numbers := []int{1, 2, 3, 4, 5}
		_ = arrays.Sum(numbers)
	}
}

func ExampleSum() {
	numbers := []int{5, 6, 7, 8, 9, 10}
	sum := arrays.Sum(numbers)
	fmt.Println(sum)

	// Output: 45
}

func TestSumAll(t *testing.T) {
	t.Run("sum the contents of two arrays", func(t *testing.T) {
		arrOne := []int{5, 6, 7, 8}
		arrTwo := []int{7, 6, 5, 4}

		got := arrays.SumAll(arrOne, arrTwo)
		want := []int{26, 22}

		assertArraysEqual(t, got, want)
	})
	t.Run("ensure empty arrays are handled correctly", func(t *testing.T) {
		arrOne := []int{}
		arrTwo := []int{5, 6, 7, 8}

		got := arrays.SumAll(arrOne, arrTwo)
		want := []int{0, 26}

		assertArraysEqual(t, got, want)
	})
}

func assertArraysEqual(t *testing.T, got, want []int) {
	t.Helper()

	if !cmp.Equal(got, want) {
		t.Errorf("wanted %v, got %v", want, got)
	}
}

func assertEqual(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("wanted %q, got %q", want, got)
	}
}
