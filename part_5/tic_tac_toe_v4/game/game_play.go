package game

import (
	"fmt"
	"strings"

	b "tic-tac-toe/board"
	p "tic-tac-toe/player"
)

// Play запускает игровой цикл
func (g *Game) Play() bool {
	fmt.Println("For saving the game enter: save filename")
	fmt.Println("For exiting the game enter : q")
	fmt.Println("For making a move enter: row col")

	for g.State == playing {
		g.Board.PrintBoard()

		// Определяем, кто делает ход: человек или компьютер
		if g.Mode == PlayerVsComputer && g.CurrentPlayer == g.Player2 {
			// Если ход компьютера, просто вызываем его MakeMove
			fmt.Println("Computer is making a move...")
			row, col, _ := g.CurrentPlayer.MakeMove(g.Board)

			// Применяем ход компьютера к доске
			g.Board.SetSymbol(row, col, g.CurrentPlayer.GetFigure())
		} else {
			// Если ход человека, запрашиваем ввод
			figure := g.CurrentPlayer.GetFigure()
			if figure == b.Cross {
				fmt.Print("X move: ")
			} else {
				fmt.Print("O move: ")
			}

			// Читаем ввод пользователя
			input, _ := g.Reader.ReadString('\n')
			input = strings.TrimSpace(input)

			// Проверка выхода из игры
			if input == "q" {
				g.State = quit
				break
			}

			// Проверка и выполнение сохранения игры
			if g.saveCheck(input) {
				continue
			}

			// Получаем ход человека-игрока через парсинг ввода
			hPlayer, ok := g.CurrentPlayer.(*p.HumanPlayer)
			if !ok {
				fmt.Println("Invalide data. Please try again!")
				continue
			}

			// Парсим ввод и получаем координаты хода
			row, col, validMove := hPlayer.ParseMove(input, g.Board)
			if !validMove {
				fmt.Println("Invalide data. Please try again!")
				continue
			}

			// Устанавливаем символ на доску
			if !g.Board.SetSymbol(row, col, hPlayer.GetFigure()) {
				fmt.Println("This cell is already occupied!")
				continue
			}
		}

		// Обновляем состояние игры
		g.updateState()

		// Если игра продолжается, меняем игрока
		if g.State == playing {
			g.switchCurrentPlayer()
		}
	}

	// Печатаем итоговую доску и результат
	g.Board.PrintBoard()
	fmt.Println()

	switch g.State {
	case crossWin:
		fmt.Println("X wins!")
	case noughtWin:
		fmt.Println("O wins!")
	case draw:
		fmt.Println("It's a draw!")
	}

	// Возвращаем true, если игра закончилась нормально (не выходом)
	return g.State != quit
}

// saveCheck проверяет, является ли ввод командой сохранения
func (g *Game) saveCheck(input string) bool {
	// Проверяем, если пользователь ввёл только "save" без имени файла
	if input == "save" {
		fmt.Println("Error: missing filename. " +
			"Please use the format: save filename")
		return true
	}

	// Проверяем команду сохранения с именем файла
	if len(input) > 5 && input[:5] == "save " {
		filename := input[5:]

		// Проверяем, что имя файла не пустое
		if len(strings.TrimSpace(filename)) == 0 {
			fmt.Println("Error: empty file name. " +
				"Please use the format: save filename")
			return true
		}

		fmt.Printf("Game saved to file: %s\n", filename)
		shapshot := g.gameSnapshot()
		if err := g.Saver.SaveGame(filename, shapshot); err != nil {
			fmt.Printf("Error saving game: %v\n", err)
		}
		return true
	}
	return false
}
