package integers_test

import (
	"testing"

	"github.com/davebrown-exit107/learn-go-with-tests/integers"
)

func TestAdder(t *testing.T) {
	got := integers.Adder(3, 3)
	want := 6

	if want != got {
		t.Errorf("wanted %q got %q", want, got)
	}

}
