package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"tic-tac-toe/game"
	"tic-tac-toe/storage"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	gameStorage := storage.NewJsonGameStorage()

	for {
		fmt.Println("Welcome to Tic-Tac-Toe!")
		fmt.Println("1 - Load game")
		fmt.Println("2 - New game")
		fmt.Println("q - Exit")
		fmt.Print("Your choice: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1": // Загрузка сохраненной игры
			loadedGame := &game.Game{}

			for {
				fmt.Println("Input file name: ")
				fileName, _ := reader.ReadString('\n')
				fileName = strings.TrimSpace(fileName)

				snapshote, err := gameStorage.LoadGame(fileName)
				if err != nil {
					fmt.Println("Error loading game: ", err)
					continue
				}

				// Восстанавливаем все необходимые поля игры
				loadedGame.RestoreFromSnapshot(
					snapshote, reader,
					gameStorage.(storage.IGameSaver),
				)

				break
			}

			// Запускаем игру
			loadedGame.Play()

		case "2": // Создаем новую игру с помощью диалога настройки
			newGame := game.SetupGame(reader,
				gameStorage.(storage.IGameSaver))

			// Запускаем игру
			newGame.Play()

		case "q":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
