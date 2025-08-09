package main

import (
	"fmt"
	"golang/factory/shape"
)

func main() {
	rectangle := shape.NewRectangle(10, 6, shape.Point{X: 6, Y: -8}, "Gold")

	fmt.Printf("%+v\n", *rectangle)
	fmt.Printf("%+v\n", rectangle.GetName())
	fmt.Printf("Perimeter = %+v\n", rectangle.GetPerimeter())
	fmt.Printf("Area = %+v\n", rectangle.GetArea())

	// меняем значения полей
	rectangle.SetLength(15)
	rectangle.SetWidth(7)
	rectangle.SetColor("Green")

	fmt.Printf("%+v\n", *rectangle)
	fmt.Printf("New perimeter = %+v\n", rectangle.GetPerimeter())
	fmt.Printf("New Area = %+v\n", rectangle.GetArea())
}
