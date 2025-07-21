package smi_test

import (
	"testing"

	"github.com/davebrown-exit107/learn-go-with-tests/smi"
)

func TestPerimeter(t *testing.T) {
	t.Run("rectangle", func(t *testing.T) {
		rect := smi.Rectangle{Height: 10.0, Width: 10.0}
		got := rect.Perimeter()
		want := 40.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
	t.Run("circle", func(t *testing.T) {
		circle := smi.Circle{Radius: 10.0}
		got := circle.Perimeter()
		want := 62.83185307179586

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestArea(t *testing.T) {
	t.Run("rectangle", func(t *testing.T) {
		rect := smi.Rectangle{Height: 10.0, Width: 10.0}
		got := rect.Area()
		want := 100.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
	t.Run("circle", func(t *testing.T) {
		circle := smi.Circle{Radius: 10.0}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}
