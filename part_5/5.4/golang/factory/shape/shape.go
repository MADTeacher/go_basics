package shape

type Point struct {
	X int
	Y int
}

type Shape struct {
	Name   string
	Center Point
	color  string
}

func (s *Shape) SetColor(color string) {
	s.color = color
}

func (s *Shape) GetColor() string {
	return s.color
}

func NewShape(name, color string, x, y int) Shape {
	return Shape{
		Name:   name,
		color:  color,
		Center: Point{X: x, Y: y},
	}
}

func NewShapeWithPoint(name, color string, center Point) *Shape {
	return &Shape{
		Name:   name,
		color:  color,
		Center: center,
	}
}
