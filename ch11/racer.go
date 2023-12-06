package racer

import (
	"net/http"
	"time"
)

func Racer(urlA, urlB string) (winner string) {
	startA := time.Now()
	http.Get(urlA)
	durationA := time.Since(startA)

	startB := time.Now()
	http.Get(urlB)
	durationB := time.Since(startB)

	if durationA < durationB {
		return urlA
	} else {
		return urlB
	}

}
