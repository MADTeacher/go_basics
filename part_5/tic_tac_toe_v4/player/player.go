package game

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	b "tic-tac-toe/board"
)

// HumanPlayer представляет игрока-человека
type HumanPlayer struct {
	Figure b.BoardField  `json:"figure"`
	Reader *bufio.Reader `json:"-"`
}

// NewHumanPlayer создает нового игрока-человека
func NewHumanPlayer(figure b.BoardField, reader *bufio.Reader) *HumanPlayer {
	return &HumanPlayer{Figure: figure, Reader: reader}
}

// GetSymbol возвращает символ игрока
func (p *HumanPlayer) GetSymbol() string {
	if p.Figure == b.Cross {
		return "X"
	}
	return "O"
}

// SwitchPlayer изменяет фигуру текущего игрока
func (p *HumanPlayer) SwitchPlayer() {
	if p.Figure == b.Cross {
		p.Figure = b.Nought
	} else {
		p.Figure = b.Cross
	}
}

// GetFigure возвращает текущую фигуру игрока
func (p *HumanPlayer) GetFigure() b.BoardField {
	return p.Figure
}

// MakeMove обрабатывает строку ввода от человека и преобразует её в координаты хода
// input - строка ввода в формате "1 2"
func (p *HumanPlayer) MakeMove(board *b.Board) (int, int, bool) {
	fmt.Printf(
		"%s's turn. Enter row and column (e.g. 1 2): ",
		p.GetSymbol(),
	)

	input, err := p.Reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if err != nil {
		fmt.Println("Invalid input. Please try again.")
		return -1, -1, false
	}

	return p.ParseMove(input, board)
}

// ParseMove обрабатывает строку ввода от человека и преобразует её в координаты хода
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

	// Преобразуем введенные координаты (начиная с 1) в индексы массива (начиная с 0)
	return row - 1, col - 1, true
}

// IsComputer возвращает false для человека-игрока
func (p *HumanPlayer) IsComputer() bool {
	return false
}

// Для обратной совместимости
type Player HumanPlayer

func NewPlayer() *Player {
	return (*Player)(NewHumanPlayer(b.Cross, nil))
}

func (p *Player) SwitchPlayer() {
	(*HumanPlayer)(p).SwitchPlayer()
}

func (p *Player) GetSymbol() string {
	return (*HumanPlayer)(p).GetSymbol()
}
