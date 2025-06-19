package board

import (
	"fmt"
	"sync"
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
	if b.Size <= 4 {
		// Для маленьких досок используем обычную проверку
		return b.checkWinSequential(player)
	}

	// Для больших досок используем параллельную проверку

	// 3 направления проверок: строки/столбцы, 2 диагонали
	resultChan := make(chan bool, 3)
	var wg sync.WaitGroup
	wg.Add(3)

	// Параллельная проверка строк и столбцов
	go func() {
		defer wg.Done()
		for i := range b.Size {
			rowWin, colWin := true, true
			for j := 0; j < b.Size; j++ {
				if b.Board[i][j] != player {
					rowWin = false
				}
				if b.Board[j][i] != player {
					colWin = false
				}
			}
			if rowWin || colWin {
				resultChan <- true
				return // Нашли выигрыш, выходим из горутины
			}
		}
		resultChan <- false
	}()

	// Параллельная проверка главной диагонали
	go func() {
		defer wg.Done()
		mainDiag := true
		for i := range b.Size {
			if b.Board[i][i] != player {
				mainDiag = false
				break
			}
		}
		resultChan <- mainDiag
	}()

	// Параллельная проверка побочной диагонали
	go func() {
		defer wg.Done()
		antiDiag := true
		for i := range b.Size {
			if b.Board[i][b.Size-i-1] != player {
				antiDiag = false
				break
			}
		}
		resultChan <- antiDiag
	}()

	// Запускаем горутину, которая закроет канал после завершения всех проверок
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Получаем результаты проверок с помощью for range.
	// Этот цикл будет ждать, пока канал не будет закрыт.
	for result := range resultChan {
		if result {
			return true // Найден выигрыш.
		}
	}

	return false
}

// Оригинальный алгоритм проверки выигрыша для малых досок
func (b *Board) checkWinSequential(player BoardField) bool {
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
