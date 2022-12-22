// Testing structs, methods, and interfaces
package smi

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
