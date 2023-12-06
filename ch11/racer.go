package racer

import (
	"errors"
	"net/http"
	"time"
)

var ErrTimeOut = errors.New("both urls timed out")

func Racer(urlA, urlB string) (winner string, err error) {
	select {
	case <-ping(urlA):
		return urlA, nil
	case <-ping(urlB):
		return urlB, nil
	case <-time.After(10 * time.Second):
		return "", ErrTimeOut
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
