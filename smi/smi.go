// Testing structs, methods, and interfaces
package smi

import "math"

// A type to hold Rectangle data and methods
type Rectangle struct {
	W float64 // Width of Rectangle
	H float64 // Height of Rectangle
}

// Calculate the perimeter of a Rectangle
func (r *Rectangle) Perimeter() (perimeter float64) {
	return (r.W + r.H) * 2
}

// Calculate the area of a Rectangle
func (r *Rectangle) Area() (area float64) {
	return r.W * r.H
}

// A type to hold Circle data and methods
type Circle struct {
	R float64 // Radius
}

// Calculate the perimeter of a Rectangle
func (c *Circle) Perimeter() (perimeter float64) {
	return 2 * math.Pi * c.R
}

// Calculate the area of a Rectangle
func (c *Circle) Area() (area float64) {
	return math.Pi * math.Pow(c.R, 2)
}
