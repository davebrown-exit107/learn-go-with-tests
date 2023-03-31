package racer_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/davebrown-exit107/learn-go-with-tests/racer"
)

func TestRacer(t *testing.T) {
	t.Run("return faster server", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, _ := racer.Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("return error on timeout", func(t *testing.T) {
		server := makeDelayedServer(25 * time.Millisecond)

		defer server.Close()

		_, err := racer.ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Errorf("expected an error but did not get one")
		}

	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
