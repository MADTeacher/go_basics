package player

import (
	"encoding/json"
	"net"
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
	return &HumanPlayer{
		Figure: b.Cross, Nickname: nickname, Conn: conn,
	}
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
// клиентской стороне
func (p *HumanPlayer) MakeMove(board *b.Board) (int, int, bool) {
	return -1, -1, false
}

func (p *HumanPlayer) IsComputer() bool {
	return false
}
