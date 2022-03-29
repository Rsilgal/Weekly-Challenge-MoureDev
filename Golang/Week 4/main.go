package main

import (
	"challenge/Shapes"
	"fmt"
)

func main() {
	rectangle := shapes.Rectangle{Width: 5, Height: 6}
	triangle := shapes.Triangle{Base: 3, Height: 5}
	square := shapes.Square{Side: 25}

	fmt.Println(shapes.GetArea(rectangle))
	fmt.Println(shapes.GetArea(triangle))
	fmt.Println(shapes.GetArea(square))
}
