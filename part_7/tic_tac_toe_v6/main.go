package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tic-tac-toe/database"
	"tic-tac-toe/game"
	"tic-tac-toe/model"
	"time"
)

// Загрузка сохраненной игры
func loadGame(reader *bufio.Reader, repository database.IRepository) {
	loadedGame := &game.Game{}

	for {
		fmt.Print("Input your nickname: ")
		nickName, _ := reader.ReadString('\n')
		nickName = strings.TrimSpace(nickName)

		// Используем горутину для загрузки снапшотов
		type snapshotResult struct {
			snapshots *[]model.GameSnapshot
			err       error
		}
		snapshotChan := make(chan snapshotResult)

		go func() {
			snapshots, err := repository.GetSnapshots(nickName)
			snapshotChan <- snapshotResult{snapshots, err}
		}()

		// Пока загружаются снапшоты, можем показать индикатор загрузки
		fmt.Print("Loading saved games")
		for range 10 {
			time.Sleep(100 * time.Millisecond)
			fmt.Print(".")
		}
		fmt.Println()

		// Получаем результат
		result := <-snapshotChan
		if result.err != nil {
			fmt.Println("Error loading game: ", result.err)
			continue
		}
		snapshote := result.snapshots

		// Выводим все снапшоты игрока
		fmt.Println("\n═══════════════════ SAVED GAMES ════════════════════")
		fmt.Println("┌────────┬────────────────┬─────────┬─────────┬────────────┐")
		fmt.Println("│   ID   │      Name      │  Figure │   Mode  │ Difficulty │")
		fmt.Println("├────────┼────────────────┼─────────┼─────────┼────────────┤")

		if len(*snapshote) == 0 {
			fmt.Println("│        │  No saved games found                              │")
			fmt.Println("└────────┴───────────────────────────────────────────────────┘")
		} else {
			for ID, snapshot := range *snapshote {
				// Конвертируем режим игры (0=PvP, 1=PvC) в читаемый текст
				gameMode := "PvP"
				if snapshot.Mode == 1 {
					gameMode = "PvC"
				}

				// Конвертируем сложность (0=Easy, 1=Medium, 2=Hard) в читаемый текст
				difficulty := "-"
				if snapshot.Mode == 1 { // Только для режима PvC
					switch snapshot.Difficulty {
					case 0:
						difficulty = "Easy"
					case 1:
						difficulty = "Medium"
					case 2:
						difficulty = "Hard"
					}
				}

				// Форматированный вывод с выравниванием колонок
				figure := "X"
				if snapshot.PlayerFigure == 1 {
					figure = "O"
				}

				name := snapshot.SnapshotName
				if name == "" {
					name = "Game " + strconv.Itoa(ID)
				}

				fmt.Printf("│   %-4d │ %-14s │    %-4s │   %-5s │    %-7s │\n",
					ID, name, figure, gameMode, difficulty)
			}
			fmt.Println("└────────┴────────────────┴─────────┴─────────┴────────────┘")
		}

		// Запрашиваем номер снапшота
		snapID := -1
		for {
			fmt.Print("Enter snapshot number: ")
			num, _ := reader.ReadString('\n')
			num = strings.TrimSpace(num)

			if snapID, _ = strconv.Atoi(num); snapID < 0 || snapID >= len(*snapshote) {
				fmt.Println("Invalid snapshot number. Please try again.")
				continue
			}
			break
		}
		// Восстанавливаем все необходимые поля игры
		loadedGame.RestoreFromSnapshot(
			&(*snapshote)[snapID], reader,
			repository,
		)

		break
	}

	// Запускаем игру
	loadedGame.Play()
}

// Показать все завершенную игру
func showFinishedGames(
	reader *bufio.Reader, repository database.IRepository,
) {
	fmt.Print("Enter nickname: ")
	nickName, _ := reader.ReadString('\n')
	nickName = strings.TrimSpace(nickName)

	// Используем горутину для загрузки снапшотов
	type snapshotResult struct {
		snapshots *[]model.FinishGameSnapshot
		err       error
	}
	snapshotChan := make(chan snapshotResult)

	go func() {
		snapshots, err := repository.GetFinishedGames(nickName)
		snapshotChan <- snapshotResult{snapshots, err}
	}()

	// Пока загружаются снапшоты, можем показать индикатор загрузки
	fmt.Print("Loading finished games")
	for range 10 {
		time.Sleep(100 * time.Millisecond)
		fmt.Print(".")
	}
	fmt.Println()

	// Получаем результат
	result := <-snapshotChan
	if result.err != nil {
		fmt.Println("Error loading finished games: ", result.err)
		return
	}
	finishedGames := result.snapshots

	fmt.Println("\n═══════════════════ FINISHED GAMES ════════════════════")
	fmt.Println("┌────────┬─────────┬────────────┬────────────────┐")
	fmt.Println("│   ID   │  Figure │   Winner   │      Date      │")
	fmt.Println("├────────┼─────────┼────────────┼────────────────┤")

	if len(*finishedGames) == 0 {
		fmt.Println("│        │ No finished games found                      │")
		fmt.Println("└────────┴───────────────────────────────────────────────┘")
	} else {
		for ID, game := range *finishedGames {
			// Форматированный вывод с выравниванием колонок
			figure := "X"
			if game.PlayerFigure == 1 {
				figure = "O"
			}

			// Форматируем дату
			dateStr := game.Time.Format("02.01 15:04")

			fmt.Printf("│   %-4d │    %-4s │  %-8s  │  %-13s │\n",
				ID,
				figure,
				game.WinnerName,
				dateStr)
		}
		fmt.Println("└────────┴─────────┴────────────┴────────────────┘")
	}

	// Запрашиваем номер игры
	snapID := -1
	for {
		fmt.Print("Enter snapshot number: ")
		num, _ := reader.ReadString('\n')
		num = strings.TrimSpace(num)

		if snapID, _ = strconv.Atoi(num); snapID < 0 || snapID >= len(*finishedGames) {
			fmt.Println("Invalid snapshot number. Please try again.")
			continue
		}
		break
	}

	// Выводим выбранную игру
	chosenGame := (*finishedGames)[snapID]
	chosenGame.Board.PrintBoard()
	fmt.Println()
	fmt.Println("Winner: ", chosenGame.WinnerName)
	fmt.Println("Date: ", chosenGame.Time.Format("02.01.2006 15:04"))
	fmt.Println()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	repository, err := database.NewSQLiteRepository()
	if err != nil {
		fmt.Println("Error creating game storage: ", err)
		return
	}

	for {
		fmt.Println("Welcome to Tic-Tac-Toe!")
		fmt.Println("1 - Load game")
		fmt.Println("2 - New game")
		fmt.Println("3 - Show all finished games")
		fmt.Println("q - Exit")
		fmt.Print("Your choice: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1": // Загрузка сохраненной игры
			loadGame(reader, repository)

		case "2": // Создаем новую игру с помощью диалога настройки
			newGame := game.SetupGame(reader,
				repository)
			// Запускаем игру
			newGame.Play()

		case "3": // Показать все завершенные игры
			showFinishedGames(reader, repository)

		case "q":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
