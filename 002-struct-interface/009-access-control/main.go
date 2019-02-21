package main

import (
	"fmt"

	"github.com/Chennai-Golang/learning-golang/002-struct-interface/009-access-control/shapes"
)

func main() {
	rect := shapes.Rectangle{20, 10}

	fmt.Println(rect.Area())

	square := shapes.Square{10}

	fmt.Println(square.Area())
}
