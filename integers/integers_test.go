package integers_test

import (
	"fmt"
	"testing"

	"github.com/davebrown-exit107/learn-go-with-tests/integers"
)

func TestAdder(t *testing.T) {
	got := integers.Add(2, 2)
	want := 4

	if got != want {
		fmt.Errorf("wanted %q, got %q", want, got)
	}
}
