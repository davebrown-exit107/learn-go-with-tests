package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello with a parameter", func(t *testing.T) {

		got := Hello("Chris")
		want := "Hello, Chris"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("say hello with no parameter", func(t *testing.T) {

		got := Hello("")
		want := "Hello, world"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
