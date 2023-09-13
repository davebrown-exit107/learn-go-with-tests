package smi

// Perimeter takes two floats, one for each side, and returns the perimeter of the rectangle
func Perimeter(length, width float64) float64 {
	return (length * 2) + (width * 2)
}

// Area takes two floats, length and width, and returns the area of the rectangle
func Area(length, width float64) float64 {
	return length * width
}
