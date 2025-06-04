package main

import "fmt"

type point struct {
	x int
	y int
}

func checkOutput(first, second int, check bool) {
	fmt.Printf("point%d == point%d - %t\n", first, second, check)
}

func main() {
	point1 := point{10, 20}
	point2 := point{10, 20}
	point3 := point{x: 2, y: 43}
	checkOutput(1, 2, point1 == point2)
	checkOutput(2, 3, point2 == point3)
}
