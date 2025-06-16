package board

type BoardField int

// фигуры в клетке поля
const (
	Empty BoardField = iota
	Cross
	Nought
)
