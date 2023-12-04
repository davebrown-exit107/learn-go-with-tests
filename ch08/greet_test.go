package main_test

import (
	"bytes"
	"testing"

	di "github.com/davebrown-exit107/learn-go-with-tests/ch08"
)

func TestGreet(t *testing.T) {
	t.Run("test greeting", func(t *testing.T) {
		buffer := bytes.Buffer{}
		di.Greet(&buffer, "Chris")

		got := buffer.String()
		want := "Hello, Chris"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})
}
