package arrays_test

import (
	"fmt"
	"testing"

	"github.com/davebrown-exit107/learn-go-with-tests/arrays"
)

func TestSum(t *testing.T) {
	t.Run("sum an array of ints", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		err, got := arrays.Sum(numbers)
		want := 15

		assertEqual(t, got, want)
		assertNoError(t, err)
	})

	t.Run("sum an empty array of ints", func(t *testing.T) {
		numbers := []int{}

		err, _ := arrays.Sum(numbers)
		wantErr := arrays.ErrArrTooShort

		assertError(t, err, wantErr)
	})

}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numbers := []int{1, 2, 3, 4, 5}
		_, _ = arrays.Sum(numbers)
	}
}

func ExampleSum() {
	numbers := []int{5, 6, 7, 8, 9, 10}
	err, sum := arrays.Sum(numbers)
	if err != nil {
		// Error encountered
	}
	fmt.Println(sum)
	// Output: 45
}

func TestSumAll(t *testing.T) {
	t.Run("sum the contents of two arrays", func(t *testing.T) {
		arrOne := []int{5, 6, 7, 8}
		arrTwo := []int{7, 6, 5, 4}

		err, got := arrays.SumAll(arrOne, arrTwo)
		want := []int{26, 22}

		assertNoError(t, err)
		assertArraysEqual(t, got, want)
	})
}

func assertArraysEqual(t *testing.T, got, want []int) {
	t.Helper()
	if len(got) != len(want) {
		t.Errorf("arrays not of equal length")
	}
	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("wanted %q, got %q", want[i], got[i])
		}
	}
}

func assertEqual(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("wanted %q, got %q", want, got)
	}
}

func assertNoError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("Expected no error, got %q", err)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("Expected error %q, got %q", want, got)
	}
}
