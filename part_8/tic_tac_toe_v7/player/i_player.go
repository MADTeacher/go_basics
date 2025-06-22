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

	// Переключение хода на другой символ
	SwitchPlayer()

	// Отправка сообщения игроку-клиенту
	SendMessage(msg *network.Message)

	// Получение никнейма игрока
	GetNickname() string

	// Получение текущей фигуры игрока
	GetFigure() b.BoardField

	// Метод для выполнения хода компьютером
	// Возвращает координаты хода (x, y) и признак успешности
	MakeMove(board *b.Board) (int, int, bool)

	// Проверка, является ли игрок компьютером
	IsComputer() bool

	// Проверка владения сокетом
	CheckSocket(conn net.Conn) bool
}
