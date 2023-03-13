package smi_test

import (
	"fmt"
	"testing"

	"github.com/davebrown-exit107/learn-go-with-tests/smi"
)

func TestArea(t *testing.T) {
	tests := map[string]struct {
		shape smi.Shape
		want  float64
	}{
		"rectangle": {
			shape: smi.Rectangle{10, 10},
			want:  100.0,
		},
		"circle": {
			shape: smi.Circle{5},
			want:  78.53981633974483,
		},
		"triangle": {
			shape: smi.Triangle{
				B: 1,
				H: 4},
			want: 2,
		},
	}

	for name, curTest := range tests {
		t.Run(name, func(t *testing.T) {
			got := curTest.shape.Area()
			if got != curTest.want {
				t.Errorf("wanted: %.2f got: %.2f", curTest.want, got)
			}
		})
	}

}

func TestPerimeter(t *testing.T) {
	tests := map[string]struct {
		shape smi.Shape
		want  float64
	}{
		"rectangle": {
			shape: smi.Rectangle{10, 10},
			want:  40.0,
		},
		"circle": {
			shape: smi.Circle{5},
			want:  31.41592653589793,
		},
		/* This hasn't been implemented yet so I'm leaving the test commented out
		"triangle": {
			shape: smi.Triangle{
				B: 1,
				H: 4},
			want:  31.41592653589793,
		},
		*/
	}

	for name, curTest := range tests {
		t.Run(name, func(t *testing.T) {
			got := curTest.shape.Perimeter()
			if got != curTest.want {
				t.Errorf("wanted: %.2f got: %.2f", curTest.want, got)
			}
		})
	}

}

func ExampleRectangle_Perimeter() {
	testRectangle := smi.Rectangle{W: 10, H: 10}
	perimeter := testRectangle.Perimeter()
	fmt.Println(perimeter)
	// Output: 40
}

func ExampleRectangle_Area() {
	testRectangle := smi.Rectangle{W: 10, H: 10}
	area := testRectangle.Area()
	fmt.Println(area)
	// Output: 100
}

func ExampleCircle_Perimeter() {
	circle := smi.Circle{R: 5}
	perimeter := circle.Perimeter()
	fmt.Println(perimeter)
	// Output: 31.41592653589793
}

func ExampleCircle_Area() {
	circle := smi.Circle{R: 5}
	area := circle.Area()
	fmt.Println(area)
	// Output: 78.53981633974483
}
