package main

import (
	"fmt"
	"golang/factory/shape"
)

func printShapeName(name shape.IShapeName) {
	fmt.Printf("IShapeName: %v\n", name.GetName())
}

func getIShapeFromRectangle(rectangle *shape.Rectangle) shape.IShape {
	return rectangle // неявное приведение
	// return shape.IShape(rectangle) // явное приведение
}

func main() {
	rectangle := shape.NewRectangle(10, 6, shape.Point{X: 6, Y: -8}, "Gold")

	ishape := shape.IShape(rectangle) // приведение Rectangle к интерфейсу
	// приведение интерфейса IShape к интерфейсу IShapeArea
	iShapeArea := shape.IShapeArea(ishape)
	fmt.Printf("IShapeArea: %+v\n", iShapeArea.GetArea())

	// явное обратное приведение интерфейса IShapePerimeter к интерфейсу IShape невозможно
	// ishape = shape.IShape(iShapePerimeter) //error:
	// cannot convert iShapePerimeter (variable of type shape.IShapePerimeter) to shape.IShape
	// (shape.IShapePerimeter does not implement shape.IShape (missing method GetArea))

	// правильное приведение интерфейса IShapeArea к интерфейсу IShape
	ishape = iShapeArea.(*shape.Rectangle)
	// аналогично
	// newIshape := shape.IShape(iShapeArea.(*shape.Rectangle))
	fmt.Printf("IShape center: %+v\n", ishape.GetCenter())

	// правильное приведение интерфейса IShapeArea к структуре Rectangle
	rectangle = iShapeArea.(*shape.Rectangle)
	fmt.Printf("%s: %+v\n", rectangle.GetName(), *rectangle)
}

// func main() {
// 	rectangle1 := shape.NewRectangle(10, 6, shape.Point{X: 6, Y: -8}, "Gold")
// 	triangle1 := shape.NewTriangle(6, 10, 8, shape.Point{X: -85, Y: 15}, "Green")
// 	rectangle2 := shape.NewRectangle(3, 5, shape.Point{X: 46, Y: -48}, "Yellow")
// 	triangle2 := shape.NewTriangle(4, 8, 6, shape.Point{X: 61, Y: 98}, "Orange")

// 	// объявляем и инициализируем срез интерфейсного типа
// 	ishapeSlice := []shape.IShape{
// 		rectangle1,
// 		triangle1,
// 		triangle2,
// 		rectangle2,
// 	}

// 	for _, it := range ishapeSlice {
// 		fmt.Printf("Shape center coordinate: %+v\n", it.GetCenter())
// 		newCenter := shape.Point{
// 			X: it.GetCenter().X + 33,
// 			Y: it.GetCenter().Y - 20,
// 		}
// 		it.MoveCenter(newCenter)
// 		fmt.Printf("Shape new center coordinate: %+v\n", it.GetCenter())
// 		fmt.Println()
// 	}
// }

// func main() {
// 	rectangle := shape.NewRectangle(10, 6, shape.Point{X: 6, Y: -8}, "Gold")
// 	triangle := shape.NewTriangle(6, 10, 8, shape.Point{X: -85, Y: 15}, "Green")
// 	printShapeName(rectangle)                   // неявное приведение
// 	//printShapeName(shape.IShapeName(rectangle)) // явное приведение
// 	printShapeName(triangle)
// 	ishape := getIShapeFromRectangle(rectangle)
// 	fmt.Printf("Shape center: %v\n", ishape.GetCenter())
// 	fmt.Printf("Shape perimeter: %v\n", ishape.GetPerimeter())
// }

// func main() {
// 	rectangle := shape.NewRectangle(10, 6, shape.Point{X: 6, Y: -8}, "Gold")
// 	triangle := shape.NewTriangle(6, 10, 8, shape.Point{X: -85, Y: 15}, "Green")
// 	iShapeName := shape.IShapeName(triangle) // приведение Triangle к интерфейсу
// 	fmt.Printf("IShapeName: %v\n", iShapeName.GetName())
// 	iShapeName = shape.IShapeName(rectangle) // приведение Rectangle к интерфейсу
// 	fmt.Printf("IShapeName: %v\n", iShapeName.GetName())
// }
