package smi_test

import (
	"testing"

	"github.com/davebrown-exit107/learn-go-with-tests/smi"
)

func TestPerimeter(t *testing.T) {
	input := smi.Rectangle{Height: 10.0, Width: 10.0}
	got := smi.Perimeter(input)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	input := smi.Rectangle{Height: 10.0, Width: 10.0}
	got := smi.Area(input)
	want := 100.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
