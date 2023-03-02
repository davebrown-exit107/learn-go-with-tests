package main

import (
	"fmt"
	"io"
	"os"
)

func Countdown(w io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintf(w, "%d\n", i)
	}
	fmt.Fprint(w, "Go!\n")
}

func main() {
	Countdown(os.Stdout)
}
