package iteration_test

import (
	"fmt"
	"testing"

	iteration "github.com/davebrown-exit107/learn-go-with-tests/ch03"
)

func TestRepeat(t *testing.T) {
	t.Run("test string repetition", func(t *testing.T) {
		err, got := iteration.Repeat("a", 10)
		want := "aaaaaaaaaa"

		assertNoError(t, err)
		assertEqual(t, got, want)
	})
	t.Run("numRepeat is zero", func(t *testing.T) {
		err, got := iteration.Repeat("b", 0)
		assertError(t, err)
		assertEqual(t, got, "")
	})
	t.Run("numRepeat less than zero", func(t *testing.T) {
		err, got := iteration.Repeat("b", -5)
		assertError(t, err)
		assertEqual(t, got, "")
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iteration.Repeat("a", 1)
	}

}

func ExampleRepeat() {
	_, repeated := iteration.Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("expected no error, got %q", err)
	}
}

func assertEqual(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func assertError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Error("Expected an error and got nil")
	}
}
