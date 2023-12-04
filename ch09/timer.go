package main

import (
	"fmt"
	"io"
	"os"
)

const startingNumber = 3
const finalWord = "Go!"

func Countdown(w io.Writer) {
	for i := startingNumber; i > 0; i-- {
		fmt.Fprintln(w, i)
	}
	fmt.Fprint(w, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
