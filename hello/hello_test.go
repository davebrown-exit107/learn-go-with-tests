package hello_test

import (
	"testing"

	"github.com/davebrown-exit107/learn-go-with-tests/hello"
)

func TestHello(t *testing.T) {
	t.Run("say hello", func(t *testing.T) {
		got := hello.Hello()
		want := "hello, world"

		if got != want {
			t.Errorf("wanted %q got %q", got, want)
		}
	})
}
