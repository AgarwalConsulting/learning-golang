package shapes

// Rectangle represents a rectangular shape
type Rectangle struct {
	Width  int
	Height int
}

// Area calculates the area of rectangle
func (r Rectangle) Area() int {
	return r.Width * r.Height
}

// Square represents a square shape
type Square struct {
	Length int
}

// Area calculates the area of square
func (s Square) Area() int {
	return s.Length * s.Length
}
