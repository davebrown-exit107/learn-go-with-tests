package racer_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	racer "github.com/davebrown-exit107/learn-go-with-tests/ch11"
)

func TestRacer(t *testing.T) {
	t.Run("compare two sites", func(t *testing.T) {
		slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(20 * time.Millisecond)
			w.WriteHeader(http.StatusOK)
		}))
		fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		want := fastUrl
		got := racer.Racer(slowUrl, fastUrl)

		assertEqualStr(t, got, want)

		slowServer.Close()
		fastServer.Close()
	})

}

func assertEqualStr(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
