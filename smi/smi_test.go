package smi_test

import (
	"testing"

	"github.com/davebrown-exit107/learn-go-with-tests/smi"
)

func TestPerimeter(t *testing.T) {
	rect := smi.Rectangle{10.0, 10.0}
	got := smi.Perimeter(rect)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f wanted %.2f", got, want)
	}

}

func TestArea(t *testing.T) {
	rect := smi.Rectangle{12.0, 6.0}
	got := smi.Area(rect)
	want := 72.0

	if got != want {
		t.Errorf("got %.2f wanted %.2f", got, want)
	}
}
