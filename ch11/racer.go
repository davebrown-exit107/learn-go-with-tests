package racer

import (
	"net/http"
	"time"
)

func Racer(urlA, urlB string) (winner string) {
	durationA := measureResponseTime(urlA)
	durationB := measureResponseTime(urlB)

	if durationA < durationB {
		return urlA
	} else {
		return urlB
	}

}

func measureResponseTime(url string) time.Duration {
	startTime := time.Now()
	http.Get(url)
	return time.Since(startTime)
}
