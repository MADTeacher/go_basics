package main

import (
	"fmt"
	"golang/factory/shape"
)

func main() {
	myShape1 := shape.NewShape("Cube", "Gold", 4, 5)
	myShape2 := shape.NewShapeWithPoint("Circle", "Green", shape.Point{X: 5, Y: 7})
	fmt.Printf("%+v\n", myShape1)  // {Name:Cube Center:{X:4 Y:5} color:Gold}
	fmt.Printf("%+v\n", *myShape2) // {Name:Circle Center:{X:5 Y:7} color:Green}
}
