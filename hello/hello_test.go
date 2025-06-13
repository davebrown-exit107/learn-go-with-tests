package hello_test

import (
	"testing"

	"github.com/davebrown-exit107/learn-go-with-tests/hello"
)

func assertStringsMatch(t testing.TB, want string, got string) {
	t.Helper()
	if want != got {
		t.Errorf("wanted %q, got %q", want, got)
	}
}
func TestHelloEn(t *testing.T) {
	t.Run("say hello with name", func(t *testing.T) {
		got := hello.Hello("Dave", "EN")
		want := "Hello, Dave"

		assertStringsMatch(t, want, got)
	})
	t.Run("say hello no name", func(t *testing.T) {
		got := hello.Hello("", "EN")
		want := "Hello, world"

		assertStringsMatch(t, want, got)
	})
}

func TestHelloEs(t *testing.T) {
	t.Run("say hello with name", func(t *testing.T) {
		got := hello.Hello("Dave", "ES")
		want := "Hola, Dave"

		assertStringsMatch(t, want, got)
	})
	t.Run("say hello no name", func(t *testing.T) {
		got := hello.Hello("", "ES")
		want := "Hola, world"

		assertStringsMatch(t, want, got)
	})
}

func TestHelloFr(t *testing.T) {
	t.Run("say hello with name", func(t *testing.T) {
		got := hello.Hello("Dave", "FR")
		want := "Bonjour, Dave"

		assertStringsMatch(t, want, got)
	})
	t.Run("say hello no name", func(t *testing.T) {
		got := hello.Hello("", "FR")
		want := "Bonjour, world"

		assertStringsMatch(t, want, got)
	})
}
