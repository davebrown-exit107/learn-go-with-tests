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
		slowServer := buildDelayedHttpServer(20 * time.Millisecond)
		fastServer := buildDelayedHttpServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		want := fastUrl
		got := racer.Racer(slowUrl, fastUrl)

		assertEqualStr(t, got, want)

	})

}

func buildDelayedHttpServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
func assertEqualStr(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
