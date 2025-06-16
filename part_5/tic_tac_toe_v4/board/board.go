package board

import (
	"fmt"
)

const (
	BoardDefaultSize int = 3
	BoardMinSize     int = 3
	BoardMaxSize     int = 9
)

type Board struct {
	Board [][]BoardField `json:"board"`
	Size  int            `json:"size"`
}

func NewBoard(size int) *Board {
	board := make([][]BoardField, size)
	for i := range board {
		board[i] = make([]BoardField, size)
	}
	return &Board{Board: board, Size: size}
}

// Отображение игрового поля
func (b *Board) PrintBoard() {
	fmt.Print("  ")
	for i := range b.Size {
		fmt.Printf("%d ", i+1)
	}
	fmt.Println()
	for i := range b.Size {
		fmt.Printf("%d ", i+1)
		for j := range b.Size {
			switch b.Board[i][j] {
			case Empty:
				fmt.Print(". ")
			case Cross:
				fmt.Print("X ")
			case Nought:
				fmt.Print("O ")
			}
		}
		fmt.Println()
	}
}

// Проверка возможности и выполнения хода
func (b *Board) makeMove(x, y int) bool {
	return b.Board[x][y] == Empty
}

func (b *Board) SetSymbol(x, y int, player BoardField) bool {
	if b.makeMove(x, y) {
		b.Board[x][y] = player
		return true
	}
	return false
}

// Проверка выигрыша
func (b *Board) CheckWin(player BoardField) bool {
	// Проверка строк и столбцов
	for i := range b.Size {
		rowWin, colWin := true, true
		for j := range b.Size {
			if b.Board[i][j] != player {
				rowWin = false
			}
			if b.Board[j][i] != player {
				colWin = false
			}
		}
		if rowWin || colWin {
			return true
		}
	}

	// Главная диагональ
	mainDiag := true
	for i := range b.Size {
		if b.Board[i][i] != player {
			mainDiag = false
			break
		}
	}
	if mainDiag {
		return true
	}

	// Побочная диагональ
	antiDiag := true
	for i := range b.Size {
		if b.Board[i][b.Size-i-1] != player {
			antiDiag = false
			break
		}
	}
	return antiDiag
}

// Проверка на ничью
func (b *Board) CheckDraw() bool {
	for i := range b.Size {
		for j := range b.Size {
			if b.Board[i][j] == Empty {
				return false
			}
		}
	}
	return true
}
