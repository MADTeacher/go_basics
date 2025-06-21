package player

import (
	"sync"
	b "tic-tac-toe/board"
)

// Метод запуска параллельной эвристической оценки ходов
func (p *ComputerPlayer) makeParallelHeuristicMove(board *b.Board) (int, int) {
	bestScore := -100000
	var bestMove []int
	emptyCells := p.getEmptyCells(board)

	if len(emptyCells) == 0 {
		return -1, -1
	}
	if len(emptyCells) == 1 {
		return emptyCells[0][0], emptyCells[0][1]
	}

	// Создаем канал для результатов
	type moveResult struct {
		move  []int
		score int
	}
	resultChan := make(chan moveResult, len(emptyCells))
	var wg sync.WaitGroup

	// Запускаем горутины для каждого возможного хода
	for _, cell := range emptyCells {
		wg.Add(1)
		go func(r, c int) {
			defer wg.Done()
			boardCopy := p.copyBoard(board)
			boardCopy.Board[r][c] = p.Figure
			score := p.evaluateBoardHeuristic(boardCopy, p.Figure)
			resultChan <- moveResult{move: []int{r, c}, score: score}
		}(cell[0], cell[1])
	}

	wg.Wait()
	close(resultChan)

	// Определяем лучший ход
	for result := range resultChan {
		if result.score > bestScore {
			bestScore = result.score
			bestMove = result.move
		}
	}

	if bestMove == nil {
		// Если по какой-то причине лучший ход не найден (маловероятно)
		// переходи на стратегию поведения среднего уровня сложности
		return p.makeMediumMove(board)
	}

	return bestMove[0], bestMove[1]
}

// Эвристическая оценка доски
// Количество рядов, столбцов или диагоналей, где у игрока есть N фигур
// и остальные клетки пусты. Также учитываем блокировку противника.
func (p *ComputerPlayer) evaluateBoardHeuristic(
	board *b.Board, player b.BoardField,
) int {
	score := 0
	opponent := b.Cross
	if player == b.Cross {
		opponent = b.Nought
	}

	// Оценка за почти выигрышные линии для игрока
	// Почти выигрыш
	score += p.countPotentialLines(board, player, board.Size-1) * 100
	// Две фигуры в ряд (для Size > 2)
	score += p.countPotentialLines(board, player, board.Size-2) * 10

	// Штраф за почти выигрышные линии для оппонента (блокировка)
	// Блокировка почти выигрыша оппонента
	score -= p.countPotentialLines(board, opponent, board.Size-1) * 90
	// Блокировка двух фигур оппонента
	score -= p.countPotentialLines(board, opponent, board.Size-2) * 5

	// Бонус за занятие центра (особенно на нечетных досках)
	if board.Size%2 == 1 {
		center := board.Size / 2
		if board.Board[center][center] == player {
			score += 5
		} else if board.Board[center][center] == opponent {
			score -= 5
		}
	}
	return score
}

// Вспомогательная функция для подсчета потенциальных линий
func (p *ComputerPlayer) countPotentialLines(
	board *b.Board, player b.BoardField, numPlayerSymbols int,
) int {
	count := 0
	lineSize := board.Size

	// Проверка строк
	for r := 0; r < lineSize; r++ {
		playerSymbols := 0
		emptySymbols := 0
		for c := 0; c < lineSize; c++ {
			if board.Board[r][c] == player {
				playerSymbols++
			} else if board.Board[r][c] == b.Empty {
				emptySymbols++
			}
		}
		if playerSymbols == numPlayerSymbols &&
			(playerSymbols+emptySymbols) == lineSize {
			count++
		}
	}

	// Проверка столбцов
	for c := 0; c < lineSize; c++ {
		playerSymbols := 0
		emptySymbols := 0
		for r := 0; r < lineSize; r++ {
			if board.Board[r][c] == player {
				playerSymbols++
			} else if board.Board[r][c] == b.Empty {
				emptySymbols++
			}
		}
		if playerSymbols == numPlayerSymbols &&
			(playerSymbols+emptySymbols) == lineSize {
			count++
		}
	}

	// Проверка главной диагонали
	playerSymbolsDiag1 := 0
	emptySymbolsDiag1 := 0
	for i := 0; i < lineSize; i++ {
		if board.Board[i][i] == player {
			playerSymbolsDiag1++
		} else if board.Board[i][i] == b.Empty {
			emptySymbolsDiag1++
		}
	}
	if playerSymbolsDiag1 == numPlayerSymbols &&
		(playerSymbolsDiag1+emptySymbolsDiag1) == lineSize {
		count++
	}

	// Проверка побочной диагонали
	playerSymbolsDiag2 := 0
	emptySymbolsDiag2 := 0
	for i := 0; i < lineSize; i++ {
		if board.Board[i][lineSize-1-i] == player {
			playerSymbolsDiag2++
		} else if board.Board[i][lineSize-1-i] == b.Empty {
			emptySymbolsDiag2++
		}
	}
	if playerSymbolsDiag2 == numPlayerSymbols &&
		(playerSymbolsDiag2+emptySymbolsDiag2) == lineSize {
		count++
	}

	return count
}
