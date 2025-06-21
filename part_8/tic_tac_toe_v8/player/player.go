package player

import (
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"strings"
	b "tic-tac-toe/board"
	"tic-tac-toe/network"
)

// Структура для представления игрока-человека
type HumanPlayer struct {
	Figure   b.BoardField `json:"figure"`
	Nickname string       `json:"nickname"`
	Conn     *net.Conn    `json:"-"`
}

func NewHumanPlayer(
	nickname string, conn *net.Conn,
) *HumanPlayer {
	return &HumanPlayer{Figure: b.Cross, Nickname: nickname, Conn: conn}
}

func (p *HumanPlayer) CheckSocket(conn net.Conn) bool {
	return *p.Conn == conn
}

// Возвращаем символ игрока
func (p *HumanPlayer) GetSymbol() string {
	if p.Figure == b.Cross {
		return "X"
	}
	return "O"
}

func (p *HumanPlayer) SendMessage(msg *network.Message) {
	json.NewEncoder(*p.Conn).Encode(msg)
}

func (p *HumanPlayer) GetNickname() string {
	return p.Nickname
}

// Изменяем фигуру текущего игрока
func (p *HumanPlayer) SwitchPlayer() {
	if p.Figure == b.Cross {
		p.Figure = b.Nought
	} else {
		p.Figure = b.Cross
	}
}

// Возвращаем текущую фигуру игрока
func (p *HumanPlayer) GetFigure() b.BoardField {
	return p.Figure
}

// Метод-заглушка, т.к. ввод игрока осуществляется на
// уровне пакета game, где нужно еще отрабатывать
// команду на выход и сохранение игровой сессии
func (p *HumanPlayer) MakeMove(board *b.Board) (int, int, bool) {
	return -1, -1, false
}

// Обрабатываем строку ввода и
// преобразуем ее в координаты хода
func (p *HumanPlayer) ParseMove(
	input string,
	board *b.Board,
) (int, int, bool) {
	parts := strings.Fields(input)
	if len(parts) != 2 {
		fmt.Println("Invalid input. Please try again.")
		return -1, -1, false
	}

	row, err1 := strconv.Atoi(parts[0])
	col, err2 := strconv.Atoi(parts[1])
	if err1 != nil || err2 != nil ||
		row < 1 || col < 1 || row > board.Size ||
		col > board.Size {
		fmt.Println("Invalid input. Please try again.")
		return -1, -1, false
	}

	// Преобразуем введенные координаты (начиная с 1)
	// в индексы массива (начиная с 0)
	return row - 1, col - 1, true
}

func (p *HumanPlayer) IsComputer() bool {
	return false
}
