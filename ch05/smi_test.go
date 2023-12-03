package smi_test

import (
	"testing"

	smi "github.com/davebrown-exit107/learn-go-with-tests/ch05"
)

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		name     string
		rectange smi.Rectangle
		want     float64
	}{
		{"square", smi.Rectangle{12, 12}, 48.0},
		{"rectange", smi.Rectangle{12, 6}, 36.0},
	}

	for _, tt := range perimeterTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.rectange.Perimeter()
			if got != tt.want {
				t.Errorf("got %.2f wanted %.2f", got, tt.want)
			}

		})
	}
}
func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape smi.Shape
		want  float64
	}{
		{"rectangle", smi.Rectangle{12, 6}, 72.0},
		{"circle", smi.Circle{10}, 314.1592653589793},
		{"triangle", smi.Triangle{5, 2}, 5},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("got %.2f wanted %.2f", got, tt.want)
			}
		})
	}
}
