package smi_test

import (
	"testing"

	"github.com/davebrown-exit107/learn-go-with-tests/smi"
)

func assertEqual(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
func TestPerimeter(t *testing.T) {
	t.Run("rectangle", func(t *testing.T) {
		rect := smi.Rectangle{Height: 10.0, Width: 10.0}
		got := rect.Perimeter()
		want := 40.0

		assertEqual(t, got, want)
	})
	t.Run("circle", func(t *testing.T) {
		circle := smi.Circle{Radius: 10.0}
		got := circle.Perimeter()
		want := 62.83185307179586

		assertEqual(t, got, want)
	})
}

func TestArea(t *testing.T) {
	t.Run("rectangle", func(t *testing.T) {
		rect := smi.Rectangle{Height: 10.0, Width: 10.0}
		got := rect.Area()
		want := 100.0

		assertEqual(t, got, want)
	})
	t.Run("circle", func(t *testing.T) {
		circle := smi.Circle{Radius: 10.0}
		got := circle.Area()
		want := 314.1592653589793

		assertEqual(t, got, want)
	})
}
