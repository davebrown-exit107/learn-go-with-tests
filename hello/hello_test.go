package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello with a name parameter", func(t *testing.T) {

		got := Hello("Chris", "EN")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello with no name parameter", func(t *testing.T) {

		got := Hello("", "EN")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)

	})

	t.Run("say hello in spanish with a name parameter", func(t *testing.T) {

		got := Hello("Chris", "ES")
		want := "Hola, Chris"

		assertCorrectMessage(t, got, want)

	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
