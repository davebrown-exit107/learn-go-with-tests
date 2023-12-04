package main_test

import (
	"bytes"
	"testing"

	timer "github.com/davebrown-exit107/learn-go-with-tests/ch09"
)

func TestCountdown(t *testing.T) {
	t.Run("confirm all numbers", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		timer.Countdown(buffer)
		got := buffer.String()
		want := `3
		2
		1
		Go!`

		assertEqual(t, got, want)
	})
}

func assertEqual(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("wanted %q, got %q", got, want)
	}
}
