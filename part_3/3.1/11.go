package main

import "fmt"

type point struct {
	x int
	y int
}

type shape struct {
	name   string
	center point
}

func main() {
	myShape := shape{
		name:   "Cube",
		center: point{x: 10, y: 6},
	}
	fmt.Printf("%+v\n", myShape) // {name:Cube center:{x:10 y:6}}
}
