// Testing structs, methods, and interfaces
package smi

import "math"

// Common interface to all shapes
type Shape interface {
	Area() float64      // Should return the area of a shape
	Perimeter() float64 // Should return the perimeter of a shape
}

// A type to hold Rectangle data and methods
type Rectangle struct {
	W float64 // Width of Rectangle
	H float64 // Height of Rectangle
}

// Calculate the perimeter of a Rectangle
func (r Rectangle) Perimeter() (perimeter float64) {
	return (r.W + r.H) * 2
}

// Calculate the area of a Rectangle
func (r Rectangle) Area() (area float64) {
	return r.W * r.H
}

// A type to hold Circle data and methods
type Circle struct {
	R float64 // Radius
}

// Calculate the perimeter of a Rectangle
func (c Circle) Perimeter() (perimeter float64) {
	return 2 * math.Pi * c.R
}

// Calculate the area of a Rectangle
func (c Circle) Area() (area float64) {
	return math.Pi * math.Pow(c.R, 2)
}

// Triangle shape
type Triangle struct {
	B float64 // Base
	H float64 // Height
}

// Calculate the perimeter of a Rectangle
func (t Triangle) Perimeter() (perimeter float64) {
	return 0.0
}

// Calculate the area of a Rectangle
func (t Triangle) Area() (area float64) {
	return (t.H * t.B) / 2.0
}
