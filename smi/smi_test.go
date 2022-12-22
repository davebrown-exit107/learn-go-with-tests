package smi

import (
	"fmt"
	"testing"
)

func TestPerimeter(t *testing.T) {
	testRectangle := Rectangle{W: 10, H: 10}
	got := testRectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("wanted: %.2f got: %.2f", want, got)
	}
}

func TestArea(t *testing.T) {
	testRectangle := Rectangle{W: 10, H: 10}
	got := testRectangle.Area()
	want := 100.0

	if got != want {
		t.Errorf("wanted: %.2f got: %.2f", want, got)
	}
}

func ExampleRectangle_Perimeter() {
	testRectangle := Rectangle{W: 10, H: 10}
	perimeter := testRectangle.Perimeter()
	fmt.Println(perimeter)
	// Output: 40
}

func ExampleRectangle_Area() {
	testRectangle := Rectangle{W: 10, H: 10}
	area := testRectangle.Area()
	fmt.Println(area)
	// Output: 100
}
