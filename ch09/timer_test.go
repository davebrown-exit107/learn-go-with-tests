package main_test

import (
	"bytes"
	"testing"
	"time"

	timer "github.com/davebrown-exit107/learn-go-with-tests/ch09"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep(d time.Duration) {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	t.Run("confirm all sleeps", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		sleeper := &SpySleeper{}
		timer.Countdown(buffer, sleeper)
		got := sleeper.Calls
		want := 3

		assertEqualInt(t, got, want)
	})
	t.Run("confirm all numbers", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		sleeper := &SpySleeper{}
		timer.Countdown(buffer, sleeper)
		got := buffer.String()
		want := "3\n2\n1\nGo!"

		assertEqualStr(t, got, want)
	})
}

func assertEqualInt(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func assertEqualStr(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
