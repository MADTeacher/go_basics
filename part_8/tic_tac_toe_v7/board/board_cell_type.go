package board

type BoardField int

// фигуры в клетке поля
const (
	Empty BoardField = iota
	Cross
	Nought
)

// Возвращаем строковое представление фигуры
func (f BoardField) String() string {
	switch f {
	case Cross:
		return "X"
	case Nought:
		return "O"
	default:
		return " "
	}
}
