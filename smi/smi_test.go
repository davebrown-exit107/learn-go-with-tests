package smi

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

func TestArea(t *testing.T) {
	got := Area(10.0, 10.0)
	want := 100.0

	if got != want {
		t.Errorf("wanted: %.2f got: %.2f", want, got)
	}
}

func ExamplePerimeter() {
	perimeter := Perimeter(10.0, 10.0)
	fmt.Println(perimeter)
	// Output: 40
}

func ExampleArea() {
	area := Area(10.0, 10.0)
	fmt.Println(area)
	// Output: 100
}
