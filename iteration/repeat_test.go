package iteration_test

import (
	"fmt"
	"testing"

	"github.com/davebrown-exit107/learn-go-with-tests/iteration"
)

func TestRepeat(t *testing.T) {
	got := iteration.Repeat("a")
	want := "aaaaa"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func ExampleRepeat() {
	repeated := iteration.Repeat("a")
	fmt.Println(repeated)
	// Output: aaaaa
}
