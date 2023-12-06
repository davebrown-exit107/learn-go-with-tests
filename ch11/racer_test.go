package racer_test

import (
	"testing"

	racer "github.com/davebrown-exit107/learn-go-with-tests/ch11"
)

func TestRacer(t *testing.T) {
	t.Run("compare two sites", func(t *testing.T) {
		slowUrl := "http://www.facebook.com"
		fastUrl := "http://www.quii.dev"

		want := fastUrl
		got := racer.Racer(slowUrl, fastUrl)

		assertEqualStr(t, got, want)
	})

}

func assertEqualStr(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
