package smi

import (
	"fmt"
	"testing"
)

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("wanted: %.2f got: %.2f", want, got)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{W: 10, H: 10}
		want := 100.0
		checkArea(t, rectangle, want)

	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{R: 5}
		want := 78.53981633974483
		checkArea(t, circle, want)
	})
}

func TestPerimeter(t *testing.T) {
	checkPerimeter := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()

		if got != want {
			t.Errorf("wanted: %.2f got: %.2f", want, got)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{W: 10, H: 10}
		want := 40.0
		checkPerimeter(t, rectangle, want)
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{R: 5}
		want := 31.41592653589793
		checkPerimeter(t, circle, want)
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
