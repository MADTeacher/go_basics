package game

import (
	"fmt"
)

const (
	BoardDefaultSize int = 3
	BoardMinSize     int = 3
	BoardMaxSize     int = 9
)

type Board struct {
	board [][]BoardField
	size  int
}

func NewBoard(size int) *Board {
	board := make([][]BoardField, size)
	for i := range board {
		board[i] = make([]BoardField, size)
	}
	return &Board{board: board, size: size}
}

// Отображение игрового поля
func (b *Board) printBoard() {
	fmt.Print("  ")
	for i := range b.size {
		fmt.Printf("%d ", i+1)
	}
	fmt.Println()
	for i := range b.size {
		fmt.Printf("%d ", i+1)
		for j := range b.size {
			switch b.board[i][j] {
			case empty:
				fmt.Print(". ")
			case cross:
				fmt.Print("X ")
			case nought:
				fmt.Print("O ")
			}
		}
		fmt.Println()
	}
}

// Проверка возможности и выполнения хода
func (b *Board) makeMove(x, y int) bool {
	return b.board[x][y] == empty
}

func (b *Board) setSymbol(x, y int, player BoardField) bool {
	if b.makeMove(x, y) {
		b.board[x][y] = player
		return true
	}
	return false
}

// Проверка выигрыша
func (b *Board) checkWin(player BoardField) bool {
	// Проверка строк и столбцов
	for i := range b.size {
		rowWin, colWin := true, true
		for j := range b.size {
			if b.board[i][j] != player {
				rowWin = false
			}
			if b.board[j][i] != player {
				colWin = false
			}
		}
		if rowWin || colWin {
			return true
		}
	}

	// Главная диагональ
	mainDiag := true
	for i := range b.size {
		if b.board[i][i] != player {
			mainDiag = false
			break
		}
	}
	if mainDiag {
		return true
	}

	// Побочная диагональ
	antiDiag := true
	for i := range b.size {
		if b.board[i][b.size-i-1] != player {
			antiDiag = false
			break
		}
	}
	return antiDiag
}

// Проверка на ничью
func (b *Board) checkDraw() bool {
	for i := range b.size {
		for j := range b.size {
			if b.board[i][j] == empty {
				return false
			}
		}
	}
	return true
}
