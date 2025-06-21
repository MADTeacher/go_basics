package board

type BoardField int

// фигуры в клетке поля
const (
	Empty BoardField = iota
	Cross
	Nought
)

func (bf BoardField) String() string {
	switch bf {
	case Empty:
		return "."
	case Cross:
		return "X"
	case Nought:
		return "O"
	default:
		return "?"
	}
}
