package shape

import "math"

type Triangle struct {
	shape
	abLength uint
	acLength uint
	bcLength uint
}

func (t *Triangle) GetABLength() uint {
	return t.abLength
}

func (t *Triangle) GetACLength() uint {
	return t.acLength
}

func (t *Triangle) GetBCLength() uint {
	return t.bcLength
}

func (t *Triangle) SetABLength(ab uint) {
	t.abLength = ab
}

func (t *Triangle) SetACLength(ac uint) {
	t.abLength = ac
}

func (t *Triangle) SetBCLength(bc uint) {
	t.abLength = bc
}

func (t *Triangle) GetPerimeter() uint {
	return t.abLength + t.acLength + t.bcLength
}

func (t *Triangle) GetArea() float64 {
	// расчет площади треугольника
	// высота рассчитывается по формуле разностороннего треугольника
	// через длины всех сторон
	p := float64(t.GetPerimeter() / 2) // полупериметр
	numerator := p * (p - float64(t.abLength)) *
		(p - float64(t.acLength)) *
		(p - float64(t.acLength))
	numerator = 2 * math.Sqrt(numerator)
	height := numerator / 2
	return float64(t.acLength) / 2 * height
}

func NewTriangle(ab, ac, bc uint, center Point, color string) *Triangle {
	return &Triangle{
		shape: shape{
			name:   "Triangle",
			color:  color,
			center: center,
		},
		abLength: ab,
		bcLength: bc,
		acLength: ac,
	}
}
