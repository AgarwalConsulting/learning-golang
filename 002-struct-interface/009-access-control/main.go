package main

import (
	"fmt"

	"github.com/AgarwalConsulting/learning-golang/002-struct-interface/009-access-control/shapes"
)

func printGeometry(g shapes.Geometry) {
	fmt.Printf("Area of %T: %d\n", g, g.Area())
	fmt.Printf("Scale Area of %T by 5: %d\n", g, g.ScaleArea(5))
}

func main() {
	rect := shapes.Rectangle{Width: 20, Height: 10}

	printGeometry(rect)
	// fmt.Println(rect.Area())

	// var square shapes.Square
	square := shapes.Square{Length: 10}

	printGeometry(square)
	// fmt.Println(square.Area())

	fmt.Println(square.ScaleArea(5))
}
