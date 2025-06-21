package client

import (
	"fmt"
	b "tic-tac-toe/board"
)

func (c *Client) printTurnInfo() {
	if c.board == nil {
		return
	}
	if c.currentPlayer == c.mySymbol {
		fmt.Println("It's your turn.")
	} else if c.currentPlayer != b.Empty {
		fmt.Printf("It's player %s's turn.\n", c.currentPlayer)
	} else {
		// Game might be over or in an intermediate state
	}
	fmt.Print("> ")
}

// validateMove checks if a move is valid based on the local board state.
func (c *Client) validateMove(row, col int) bool {
	if c.board == nil {
		fmt.Println("Game has not started yet.")
		return false
	}
	if row < 1 || row > c.board.Size || col < 1 || col > c.board.Size {
		fmt.Printf("Invalid move. Row and column must be between 1 and %d.\n", c.board.Size)
		return false
	}
	// Convert to 0-indexed for board access
	if c.board.Board[row-1][col-1] != b.Empty {
		fmt.Println("Invalid move. Cell is already occupied.")
		return false
	}
	return true
}
