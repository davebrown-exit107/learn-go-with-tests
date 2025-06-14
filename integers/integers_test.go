package integers_test

import (
	"fmt"
	"testing"

	"github.com/davebrown-exit107/learn-go-with-tests/integers"
)

func assertValuesMatch(t testing.TB, got int, want int) {
	t.Helper()
	if got != want {
		fmt.Errorf("wanted %d, got %d", want, got)
	}
}

func TestAdder(t *testing.T) {
	t.Run("two positive numbers", func(t *testing.T) {
		got := integers.Add(2, 2)
		want := 4

		assertValuesMatch(t, got, want)
	})
	t.Run("one negative one positive number", func(t *testing.T) {
		got := integers.Add(-5, 10)
		want := 5

		assertValuesMatch(t, got, want)
	})
}

func ExampleAdd() {
	sum := integers.Add(2, 3)
	fmt.Println(sum)
	// Output: 5
}
