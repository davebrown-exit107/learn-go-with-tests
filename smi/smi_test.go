package main

import (
	"fmt"
	"testing"
)

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("wanted: %.2f got: %.2f", want, got)
	}
}

func ExamplePerimeter() {
	perimeter := Perimeter(10.0, 10.0)
	fmt.Println(perimeter)
	// Output: 40
}
