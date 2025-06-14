package game

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Game struct {
	Board  *Board        `json:"board"`
	Player *Player       `json:"player"`
	Reader *bufio.Reader `json:"-"`
	State  GameState     `json:"state"`
	Saver  IGameSaver    `json:"-"`
}

func NewGame(board Board, player Player,
	reader *bufio.Reader, saver IGameSaver) *Game {
	return &Game{
		Board:  &board,
		Player: &player,
		Reader: reader,
		State:  playing,
		Saver:  saver,
	}
}

func (g *Game) updateState() {
	if g.Board.checkWin(g.Player.Figure) {
		if g.Player.Figure == cross {
			g.State = crossWin
		} else {
			g.State = noughtWin
		}
	} else if g.Board.checkDraw() {
		g.State = draw
	}
}

func (g *Game) saveCheck(input string) bool {
	if input == "save" {
		fmt.Println("Enter file name: ")
		fileName, err := g.Reader.ReadString('\n')
		if err != nil {
			fmt.Println("Invalid input. Please try again.")
			return false
		}

		fileName = strings.TrimSpace(fileName)
		err = g.Saver.SaveGame(fileName, g)
		if err != nil {
			fmt.Println("Error saving game.")
			return false
		}
		fmt.Println("Game saved successfully!!!")
		return true
	}

	return false
}

// Игровой цикл
func (g *Game) Play() {
	for g.State == playing {
		g.Board.printBoard()
		fmt.Printf(
			"%s's turn. Enter row and column (e.g. 1 2): ",
			g.Player.getSymbol())

		input, err := g.Reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if err != nil {
			fmt.Println("Invalid input. Please try again.")
			continue
		}

		input = strings.TrimSpace(input)
		if input == "q" {
			g.State = quit
			break
		}

		if g.saveCheck(input) {
			continue
		}

		parts := strings.Fields(input)
		if len(parts) != 2 {
			fmt.Println("Invalid input. Please try again.")
			continue
		}

		row, err1 := strconv.Atoi(parts[0])
		col, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil ||
			row < 1 || col < 1 || row > g.Board.Size ||
			col > g.Board.Size {
			fmt.Println("Invalid input. Please try again.")
			continue
		}
		if g.Board.setSymbol(row-1, col-1, g.Player.Figure) {
			g.updateState()
			g.Player.switchPlayer()

		} else {
			fmt.Println("This cell is already occupied!")
		}
	}

	g.Board.printBoard()

	if g.State == crossWin {
		fmt.Println("X wins!")
	} else if g.State == noughtWin {
		fmt.Println("O wins!")
	} else if g.State == draw {
		fmt.Println("It's a draw!")
	} else {
		fmt.Println("Game over!")
	}
}
