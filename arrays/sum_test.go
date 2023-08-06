package arrays_test

import (
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
		wantErr := arrays.ArrTooShort

		assertError(t, err, wantErr)
	})

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
