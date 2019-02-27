package shapes

// Square represents a square shape
type Square struct {
	Length int
}

// Area calculates the area of square
func (s Square) Area() int {
	return s.Length * s.Length
}

func (s Square) ScaleArea(num int) int {
	return s.Length * s.Length * num
}
