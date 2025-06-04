package shape

type Point struct {
	X int
	Y int
}

type shape struct {
	name   string
	center Point
	color  string
}

func (s *shape) SetColor(color string) {
	s.color = color
}

func (s *shape) GetColor() string {
	return s.color
}

func (s *shape) GetName() string {
	return s.name
}

func (s *shape) MoveCenter(point Point) {
	s.center = point
}

func (s *shape) GetCenter() Point {
	return s.center
}
