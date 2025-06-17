package game

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	b "tic-tac-toe/board"
	db "tic-tac-toe/database"
	p "tic-tac-toe/player"
)

// Создаем новую игру с пользовательскими настройками
func SetupGame(reader *bufio.Reader, repository db.IRepository) *Game {
	// Запрашиваем размер игрового поля
	size := getBoardSize(reader)

	// Создаем доску
	board := *b.NewBoard(size)

	// Запрашиваем режим игры
	mode := getGameMode(reader)

	// Если выбран режим против компьютера, запрашиваем сложность
	var difficulty p.Difficulty
	if mode == PlayerVsComputer {
		difficulty = getDifficulty(reader)
	}

	// Создаем новую игру
	return NewGame(board, reader, repository, mode, difficulty)
}

// Запрашиваем у пользователя размер доски
func getBoardSize(reader *bufio.Reader) int {
	size := b.BoardDefaultSize
	var err error
	for {
		fmt.Printf("Choose board size (min: %d, max: %d, default: %d): ",
			b.BoardMinSize, b.BoardMaxSize, b.BoardDefaultSize)

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Если пользователь не ввел ничего, используем размер по умолчанию
		if input == "" {
			return b.BoardDefaultSize
		}

		// Пытаемся преобразовать ввод в число
		size, err = strconv.Atoi(input)
		if err != nil || size < b.BoardMinSize || size > b.BoardMaxSize {
			fmt.Println("Invalid input. Please try again!")
			continue
		}

		return size
	}
}

// Запрашиваем у пользователя режим игры
func getGameMode(reader *bufio.Reader) GameMode {
	for {
		fmt.Println("Choose game mode:")
		fmt.Println("1 - Player vs Player (PvP)")
		fmt.Println("2 - Player vs Computer (PvC)")
		fmt.Print("Your choice: ")

		input, err := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if err != nil {
			fmt.Println("Invalid input. Please try again!")
			continue
		}

		switch input {
		case "1":
			return PlayerVsPlayer
		case "2":
			return PlayerVsComputer
		default:
			fmt.Println("Invalid input. Please try again!")
		}
	}
}

// Запрашиваем у пользователя уровень сложности компьютера
func getDifficulty(reader *bufio.Reader) p.Difficulty {
	for {
		fmt.Println("Choose computer difficulty:")
		fmt.Println("1 - Easy (random moves)")
		fmt.Println("2 - Medium (block winning moves)")
		fmt.Println("3 - Hard (optimal strategy)")
		fmt.Print("Your choice: ")

		input, err := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if err != nil {
			fmt.Println("Invalid input. Please try again!")
			continue
		}

		switch input {
		case "1":
			return p.Easy
		case "2":
			return p.Medium
		case "3":
			return p.Hard
		default:
			fmt.Println("Invalid input. Please try again!")
		}
	}
}
