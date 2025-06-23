package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tic-tac-toe/game"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	boardSize := 0
	for {
		fmt.Print("Enter the size of the board (3-9): ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input.")
			continue
		}
		input = strings.TrimSpace(input)
		boardSize, err = strconv.Atoi(input)
		if err != nil {
			// Использовать предыдущий размер по умолчанию
			boardSize = game.BoardDefaultSize
		}
		if boardSize < game.BoardMinSize ||
			boardSize > game.BoardMaxSize {
			fmt.Println("Invalid board size.")
		} else {
			break
		}
	}

	board := game.NewBoard(boardSize)
	player := game.NewPlayer()
	game := game.NewGame(*board, *player, reader)
	game.Play()
}
