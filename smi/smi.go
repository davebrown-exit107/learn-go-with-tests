package smi

import "math"

// Rectangle represents a rectangle for calucations of area and perimeter
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

// Circle represents a circle for calucations of area and perimeter
type Circle struct {
	Radius float64
}

// Calculate the area of a circle
func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}
