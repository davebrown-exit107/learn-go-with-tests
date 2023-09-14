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
	t.Run("rectangle", func(t *testing.T) {
		rect := smi.Rectangle{12, 6}
		got := rect.Area()
		want := 72.0

		if got != want {
			t.Errorf("got %.2f wanted %.2f", got, want)
		}
	})

	t.Run("circle", func(t *testing.T) {
		circle := smi.Circle{10.0}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %.2f wanted %.2f", got, want)
		}
	})
}
