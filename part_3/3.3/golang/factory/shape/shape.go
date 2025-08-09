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
