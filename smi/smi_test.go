package smi

import (
	"fmt"
	"testing"
)

func TestRectangle(t *testing.T) {
	t.Run("perimeter", func(t *testing.T) {
		testRectangle := Rectangle{W: 10, H: 10}
		got := testRectangle.Perimeter()
		want := 40.0

		if got != want {
			t.Errorf("wanted: %.2f got: %.2f", want, got)
		}
	})

	t.Run("area", func(t *testing.T) {
		testRectangle := Rectangle{W: 10, H: 10}
		got := testRectangle.Area()
		want := 100.0

		if got != want {
			t.Errorf("wanted: %.2f got: %.2f", want, got)
		}
	})
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

func TestCircle(t *testing.T) {
	t.Run("perimeter", func(t *testing.T) {
		testCircle := Circle{R: 5}
		got := testCircle.Perimeter()
		want := 31.41592653589793

		if got != want {
			t.Errorf("wanted: %.2f got: %.2f", want, got)
		}
	})

	t.Run("area", func(t *testing.T) {
		testCircle := Circle{R: 5}
		got := testCircle.Area()
		want := 78.53981633974483

		if got != want {
			t.Errorf("wanted: %.2f got: %.2f", want, got)
		}
	})
}

func ExampleCircle_Perimeter() {
	circle := Circle{R: 5}
	perimeter := circle.Perimeter()
	fmt.Println(perimeter)
	// Output: 31.41592653589793
}

func ExampleCircle_Area() {
	circle := Circle{R: 5}
	area := circle.Area()
	fmt.Println(area)
	// Output: 78.53981633974483
}
