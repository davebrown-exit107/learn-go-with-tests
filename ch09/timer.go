package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const startingNumber = 3
const finalWord = "Go!"

type DefaultSleeper struct{}

func (s *DefaultSleeper) Sleep(d time.Duration) {
	time.Sleep(d)
}

type Sleeper interface {
	Sleep(time.Duration)
}

func Countdown(w io.Writer, s Sleeper) {
	for i := startingNumber; i > 0; i-- {
		fmt.Fprintln(w, i)
		s.Sleep(time.Second * 1)
	}
	fmt.Fprint(w, finalWord)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
