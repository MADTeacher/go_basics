package main

import (
	"fmt"
	"golang/factory/shape"
)

func main() {

	myShape := shape.Shape{
		Name:   "Cube",
		Center: shape.Point{X: 4, Y: 5},
	}

	fmt.Printf("%+v\n", myShape) // {Name:Cube Center:{X:4 Y:5} color:}
	myShape.SetColor("Gold")
	fmt.Printf("%+v\n", myShape)                        // {Name:Cube Center:{X:4 Y:5} color:Gold}
	fmt.Printf("Shape color: %v\n", myShape.GetColor()) // Shape color: Gold
}
