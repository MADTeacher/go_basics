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
	loader := game.NewJsonGameLoader()
	boardSize := 0

	for {
		fmt.Println("1 - load game")
		fmt.Println("2 - new game")
		fmt.Println("q - quit")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			var loadedGame *game.Game
			var err error
			for {
				fmt.Println("Enter file name: ")
				fileName, _ := reader.ReadString('\n')
				fileName = strings.TrimSpace(fileName)
				loadedGame, err = loader.LoadGame(fileName)
				if err != nil {
					fmt.Println("Error loading game.")
					continue
				}
				break
			}
			loadedGame.Reader = reader
			loadedGame.Saver = loader.(game.IGameSaver)
			loadedGame.Play()
		case "2":
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
			game := game.NewGame(*board, *player, reader,
				loader.(game.IGameSaver))
			game.Play()
		case "q":
			return
		default:
			fmt.Println("Invalid input. Please try again.")
			return
		}
	}
}
