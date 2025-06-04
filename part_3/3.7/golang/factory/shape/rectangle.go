package shape

type Rectangle struct {
	shape  // композиция
	width  uint
	length uint
}

func (r *Rectangle) GetWidth() uint {
	return r.width
}

func (r *Rectangle) GetLength() uint {
	return r.length
}

func (r *Rectangle) SetWidth(width uint) {
	r.width = width
}

func (r *Rectangle) SetLength(length uint) {
	r.length = length
}

func (r *Rectangle) GetPerimeter() uint {
	return (r.length + r.width) * 2
}

func (r *Rectangle) GetArea() float64 {
	return float64(r.length * r.width)
}

func (r *Rectangle) GetName() string {
	return "Super-puper " + r.shape.GetName()
}

func NewRectangle(width, length uint, center Point, color string) *Rectangle {
	return &Rectangle{
		shape: shape{
			name:   "Rectangle",
			color:  color,
			center: center,
		},
		width:  width,
		length: length,
	}
}
