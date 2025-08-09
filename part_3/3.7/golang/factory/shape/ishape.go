package shape

type IShape interface {
	GetArea() float64
	GetPerimeter() uint
	MoveCenter(point Point)
	GetCenter() Point
}

type IShapeArea interface {
	GetArea() float64
}

type IShapePerimeter interface {
	GetPerimeter() uint
}

type IShapeName interface {
	GetName() string
}
