package smi_test

import (
	"testing"

	"github.com/davebrown-exit107/learn-go-with-tests/smi"
)

func TestPerimeter(t *testing.T) {
	t.Run("rectangle", func(t *testing.T) {
		rect := smi.Rectangle{10, 10}
		got := rect.Perimeter()
		want := 40.0

		if got != want {
			t.Errorf("got %.2f wanted %.2f", got, want)
		}
	})
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
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %.2f wanted %.2f", got, tt.want)
		}
	}
}
