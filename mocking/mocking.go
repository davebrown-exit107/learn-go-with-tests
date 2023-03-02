package main

import (
	"fmt"
	"io"
	"os"
)

func Countdown(w io.Writer) {
	fmt.Fprintf(w, "3\n")
}

func main() {
	Countdown(os.Stdout)
}
