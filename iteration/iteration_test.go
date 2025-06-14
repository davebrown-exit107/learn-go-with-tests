package iteration_test

import (
	"fmt"
	"testing"

	"github.com/davebrown-exit107/learn-go-with-tests/iteration"
)

func TestRepeat(t *testing.T) {
	t.Run("first test", func(t *testing.T) {
		got := iteration.Repeat("a", 4)
		want := "aaaa"

		if got != want {
			fmt.Errorf("wanted %q, got %q", want, got)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iteration.Repeat("a", 4)
	}
}
