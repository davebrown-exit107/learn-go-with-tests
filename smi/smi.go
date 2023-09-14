package smi

import "math"

// Shape is an interface to encompass shape types
type Shape interface {
	Area() float64
}

type Rectangle struct {
	Length float64
	Height float64
}

// Calculate the perimeter of a rectangle
func (r Rectangle) Perimeter() float64 {
	return (r.Length * 2) + (r.Height * 2)
}

// Calculate the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Length * r.Height
}

type Circle struct {
	Radius float64
}

// Calculate the area of a circle
func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

type Triangle struct {
	Height float64
	Base   float64
}

// Calculate the area of a triangle
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}
