package hello_test

import (
	"testing"

	"github.com/davebrown-exit107/learn-go-with-tests/hello"
)

func TestHello(t *testing.T) {
	t.Run("say hello with a name parameter", func(t *testing.T) {

		got := hello.Hello("Chris", "EN")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello with no name parameter", func(t *testing.T) {

		got := hello.Hello("", "EN")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)

	})

	t.Run("say hello in spanish with a name parameter", func(t *testing.T) {

		got := hello.Hello("Chris", "ES")
		want := "Hola, Chris"

		assertCorrectMessage(t, got, want)

	})

	t.Run("say hello in french with a name parameter", func(t *testing.T) {

		got := hello.Hello("Chris", "FR")
		want := "Bonjour, Chris"

		assertCorrectMessage(t, got, want)

	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
