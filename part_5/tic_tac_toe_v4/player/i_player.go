package game

import b "tic-tac-toe/board"

// IPlayer представляет интерфейс для любого игрока (человека или компьютера)
type IPlayer interface {
	// Получение символа игрока (X или O)
	GetSymbol() string

	// Переключение хода на другого игрока
	SwitchPlayer()

	// Получение текущей фигуры игрока
	GetFigure() b.BoardField

	// Выполнение хода игрока
	// Возвращает координаты хода (x, y) и признак успешности
	MakeMove(board *b.Board) (int, int, bool)

	// Проверка, является ли игрок компьютером
	IsComputer() bool
}
