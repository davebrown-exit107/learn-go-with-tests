package arrays_test

import (
	"testing"

	"github.com/davebrown-exit107/learn-go-with-tests/arrays"
)

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	got := arrays.Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d, given %v", got, want, numbers)
	}

}
