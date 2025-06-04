package game

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Game struct {
	board  *Board
	player *Player
	reader *bufio.Reader
	state  GameState
}

func NewGame(board Board, player Player, reader *bufio.Reader) *Game {
	return &Game{
		board:  &board,
		player: &player,
		reader: reader,
		state:  playing,
	}
}

func (g *Game) updateState() {
	if g.board.checkWin(g.player.figure) {
		if g.player.figure == cross {
			g.state = crossWin
		} else {
			g.state = noughtWin
		}
	} else if g.board.checkDraw() {
		g.state = draw
	}
}

// Игровой цикл
func (g *Game) Play() {
	for g.state == playing {
		g.board.printBoard()
		fmt.Printf(
			"%s's turn. Enter row and column (e.g. 1 2): ",
			g.player.getSymbol())

		input, err := g.reader.ReadString('\n')
		if err != nil {
			fmt.Println("Invalid input. Please try again.")
			continue
		}

		input = strings.TrimSpace(input)
		if input == "q" {
			g.state = quit
			break
		}

		parts := strings.Fields(input)
		if len(parts) != 2 {
			fmt.Println("Invalid input. Please try again.")
			continue
		}

		row, err1 := strconv.Atoi(parts[0])
		col, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil ||
			row < 1 || col < 1 || row > g.board.size ||
			col > g.board.size {
			fmt.Println("Invalid input. Please try again.")
			continue
		}
		if g.board.setSymbol(row-1, col-1, g.player.figure) {
			g.updateState()
			g.player.switchPlayer()

		} else {
			fmt.Println("This cell is already occupied!")
		}
	}

	g.board.printBoard()

	if g.state == crossWin {
		fmt.Println("X wins!")
	} else if g.state == noughtWin {
		fmt.Println("O wins!")
	} else if g.state == draw {
		fmt.Println("It's a draw!")
	} else {
		fmt.Println("Game over!")
	}
}
