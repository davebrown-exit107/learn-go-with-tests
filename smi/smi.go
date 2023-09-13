package smi

// Rectangle represents a rectangle for calucations of area and perimeter
type Rectangle struct {
	Length float64
	Width  float64
}

// Perimeter takes two floats, one for each side, and returns the perimeter of the rectangle
func Perimeter(r Rectangle) float64 {
	return (r.Length * 2) + (r.Width * 2)
}

// Area takes two floats, length and width, and returns the area of the rectangle
func Area(r Rectangle) float64 {
	return r.Length * r.Width
}
