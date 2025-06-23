package client

import (
	"fmt"
	b "tic-tac-toe/board"
	g "tic-tac-toe/game"
)

// Выводит информацию о ходе игрока
func (c *Client) printTurnInfo() {
	if c.board == nil {
		return
	}
	if c.currentPlayer == c.mySymbol { // Если ход игрока
		fmt.Println("It's your turn.")
	} else if c.currentPlayer != b.Empty { // Если ход оппонента
		fmt.Printf("It's player %s's turn.\n", c.currentPlayer)
	} else {
		// Игра может быть завершена или находиться
		// в промежуточном состоянии
	}
	fmt.Print("> ")
}

// Проверяем валидность хода игрока
func (c *Client) validateMove(row, col int) bool {
	if c.board == nil { // Если игра не начата
		fmt.Println("Game has not started yet.")
		return false
	}
	// Если ход вне поля
	if row < 1 || row > c.board.Size || col < 1 || col > c.board.Size {
		fmt.Printf(
			"Invalid move. Row and column must be between 1 and %d.\n",
			c.board.Size,
		)
		return false
	}
	// Преобразуем в 0-индексированный для доступа к полю
	if c.board.Board[row-1][col-1] != b.Empty {
		fmt.Println("Invalid move. Cell is already occupied.")
		return false
	}
	return true
}

// Конвертируем экземпляр типа GameMode в строку
func gameModeToString(mode g.GameMode) string {
	switch mode {
	case g.PvP:
		return "PvP"
	case g.PvC:
		return "PvC"
	default:
		return "Unknown"
	}
}

// Конвертируем экземпляр типа Difficulty в строку
func difficultyToString(difficulty g.Difficulty) string {
	switch difficulty {
	case g.Easy:
		return "Easy"
	case g.Medium:
		return "Medium"
	case g.Hard:
		return "Hard"
	default:
		return ""
	}
}
