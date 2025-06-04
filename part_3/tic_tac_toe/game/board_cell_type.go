package game

type BoardField int

// фигуры в клетке поля
const (
	empty BoardField = iota
	cross
	nought
)
