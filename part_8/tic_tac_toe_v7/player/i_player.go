package player

import (
	"net"
	b "tic-tac-toe/board"
	"tic-tac-toe/network"
)

// Интерфейс для любого игрока, будь то человек или компьютер
type IPlayer interface {
	// Получение символа игрока (X или O)
	GetSymbol() string

	// Переключение хода на другого игрока
	SwitchPlayer()

	SendMessage(msg *network.Message)

	GetNickname() string

	// Получение текущей фигуры игрока
	GetFigure() b.BoardField

	// Выполнение хода игрока
	// Возвращает координаты хода (x, y) и признак успешности
	MakeMove(board *b.Board) (int, int, bool)

	// Проверка, является ли игрок компьютером
	IsComputer() bool

	CheckSocket(conn net.Conn) bool
}
