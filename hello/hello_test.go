package hello_test

import (
	"testing"

	"github.com/davebrown-exit107/learn-go-with-tests/hello"
)

func TestHelloEn(t *testing.T) {
	t.Run("say hello - with name", func(t *testing.T) {
		got := hello.Hello("Dave", "EN")
		want := "Hello, Dave"

		if got != want {
			t.Errorf("wanted %q got %q", want, got)
		}
	})
	t.Run("say hello - no name", func(t *testing.T) {
		got := hello.Hello("", "EN")
		want := "Hello, world"

		if got != want {
			t.Errorf("wanted %q got %q", want, got)
		}
	})
}

func TestHelloEs(t *testing.T) {
	t.Run("say hello - with name", func(t *testing.T) {
		got := hello.Hello("Dave", "ES")
		want := "Hola, Dave"

		if got != want {
			t.Errorf("wanted %q got %q", want, got)
		}
	})
	t.Run("say hello - no name", func(t *testing.T) {
		got := hello.Hello("", "ES")
		want := "Hola, world"

		if got != want {
			t.Errorf("wanted %q got %q", want, got)
		}
	})
}
