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
	checkPerimeter := func(t testing.TB, shape smi.Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()
		assertEqual(t, got, want)
	}
	t.Run("rectangle", func(t *testing.T) {
		rect := smi.Rectangle{Height: 10.0, Width: 10.0}
		want := 40.0
		checkPerimeter(t, rect, want)
	})
	t.Run("circle", func(t *testing.T) {
		circle := smi.Circle{Radius: 10.0}
		want := 62.83185307179586
		checkPerimeter(t, circle, want)
	})
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape smi.Shape, want float64) {
		t.Helper()
		got := shape.Area()
		assertEqual(t, got, want)
	}
	t.Run("rectangle", func(t *testing.T) {
		rect := smi.Rectangle{Height: 10.0, Width: 10.0}
		want := 100.0
		checkArea(t, rect, want)
	})
	t.Run("circle", func(t *testing.T) {
		circle := smi.Circle{Radius: 10.0}
		want := 314.1592653589793
		checkArea(t, circle, want)
	})
}
