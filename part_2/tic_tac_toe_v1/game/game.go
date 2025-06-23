package game

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type GameState int
type BoardField int

// фигуры в клетке поля
const (
	empty BoardField = iota
	cross
	nought
)

// состояние игрового процесса
const (
	playing GameState = iota
	draw
	crossWin
	noughtWin
	quit
)

// Объявление глобальных переменных
var (
	board [][]BoardField
	// размер игрового поля по умолчанию
	boardSize     int        = 3
	currentPlayer BoardField = cross // текущий игрок
	state         GameState  = playing
	reader                   = bufio.NewReader(os.Stdin)
)

// Инициализация игрового поля
// Имя функции начинается с большой буквы, что
// указывает на то, что к ней можно будет обратиться после
// импортирования пакета - game.InitBoard. Таким образом,
// InitBoard является публичной функцией
func InitBoard() bool {
	fmt.Print("Enter the size of the board (3-9): ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input.")
		return false
	}
	input = strings.TrimSpace(input)
	size, err := strconv.Atoi(input)
	if err != nil {
		size = boardSize // Использовать предыдущий размер по умолчанию
	}
	if size < 3 || size > 9 {
		fmt.Println("Invalid board size.")
		return false
	}
	boardSize = size
	board = make([][]BoardField, boardSize)
	for i := range board {
		board[i] = make([]BoardField, boardSize)
	}
	return true
}

// Отображение игрового поля
func printBoard() {
	fmt.Print("  ")
	for i := range boardSize {
		fmt.Printf("%d ", i+1)
	}
	fmt.Println()
	for i := range boardSize {
		fmt.Printf("%d ", i+1)
		for j := range boardSize {
			switch board[i][j] {
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
func makeMove(x, y int) bool {
	if x < 0 || x >= boardSize || y < 0 || y >= boardSize {
		return false
	}
	if board[x][y] != empty {
		return false
	}
	board[x][y] = currentPlayer
	return true
}

// Проверка выигрыша
func checkWin(player BoardField) bool {
	// Проверка строк и столбцов
	for i := range boardSize {
		rowWin, colWin := true, true
		for j := range boardSize {
			if board[i][j] != player {
				rowWin = false
			}
			if board[j][i] != player {
				colWin = false
			}
		}
		if rowWin || colWin {
			return true
		}
	}

	// Главная диагональ
	mainDiag := true
	for i := range boardSize {
		if board[i][i] != player {
			mainDiag = false
			break
		}
	}
	if mainDiag {
		return true
	}

	// Побочная диагональ
	antiDiag := true
	for i := range boardSize {
		if board[i][boardSize-i-1] != player {
			antiDiag = false
			break
		}
	}
	return antiDiag
}

// Проверка на ничью
func checkDraw() bool {
	for i := range boardSize {
		for j := range boardSize {
			if board[i][j] == empty {
				return false
			}
		}
	}
	return true
}

// Смена игрока
func switchPlayer() {
	if currentPlayer == cross {
		currentPlayer = nought
	} else {
		currentPlayer = cross
	}
}

// Обновление состояния игры
func updateState() {
	if checkWin(currentPlayer) {
		if currentPlayer == cross {
			state = crossWin
		} else {
			state = noughtWin
		}
	} else if checkDraw() {
		state = draw
	}
}

// Игровой цикл
func Play() {
	for state == playing {
		playerSymbol := "X"
		if currentPlayer == nought {
			playerSymbol = "O"
		}
		fmt.Printf(
			"%s's turn. Enter row and column (e.g. 1 2): ",
			playerSymbol)

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Invalid input. Please try again.")
			continue
		}

		input = strings.TrimSpace(input)
		if input == "q" {
			state = quit
			break
		}

		parts := strings.Fields(input)
		if len(parts) != 2 {
			fmt.Println("Invalid input. Please try again.")
			continue
		}

		row, err1 := strconv.Atoi(parts[0])
		col, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil || !makeMove(row-1, col-1) {
			fmt.Println("Invalid move. Please try again.")
			continue
		}

		updateState()
		printBoard()

		if state == crossWin {
			fmt.Println("X wins!")
		} else if state == noughtWin {
			fmt.Println("O wins!")
		} else if state == draw {
			fmt.Println("It's a draw!")
		} else {
			switchPlayer()
		}
	}
}
