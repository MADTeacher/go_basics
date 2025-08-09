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

	myShape1 := shape.Shape{"Triangle", shape.Point{1, 9}}

	fmt.Printf("%+v\n", myShape)  // {Name:Cube Center:{X:4 Y:5}}
	fmt.Printf("%+v\n", myShape1) // {Name:Triangle Center:{X:1 Y:9}}
	myShape.Center.Y = 2
	myShape.Name = "Circle"
	fmt.Printf("%+v\n", myShape) // {Name:Circle Center:{X:4 Y:2}}
}
