package smi_test

import (
	"testing"

	"github.com/davebrown-exit107/learn-go-with-tests/smi"
)

func assertEqual(t testing.TB, got, want float64) {
	t.Helper()
}
func TestPerimeter(t *testing.T) {
	testCases := map[string]struct {
		input smi.Shape
		want  float64
	}{
		"rectangle": {
			input: smi.Rectangle{10, 10},
			want:  40.0,
		},
		"circle": {
			input: smi.Circle{10.0},
			want:  62.83185307179586,
		},
		"triangle": {
			// not implemented
			input: smi.Triangle{12, 6},
			want:  0.0,
		},
	}
	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			got := test.input.Perimeter()
			if got != test.want {
				t.Errorf("%#v got %v want %v", test.input, got, test.want)
			}
		})
	}
}

func TestArea(t *testing.T) {
	testCases := map[string]struct {
		input smi.Shape
		want  float64
	}{
		"rectangle": {
			input: smi.Rectangle{10, 10},
			want:  100.0,
		},
		"circle": {
			input: smi.Circle{10.0},
			want:  314.1592653589793,
		},
		"triangle": {
			input: smi.Triangle{12, 6},
			want:  36.0,
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			got := test.input.Area()
			if got != test.want {
				t.Errorf("%#v got %v want %v", test.input, got, test.want)
			}
		})
	}
}
