package player

import (
	b "tic-tac-toe/board"
)

// Параллельный анализ на основе зон
func (p *ComputerPlayer) makeZoneBasedMove(board *b.Board) (int, int) {
	// Если доска не очень большая, используем эвристику
	if board.Size <= 5 { // Пороговое значение, можно настроить
		return p.makeParallelHeuristicMove(board)
	}

	bestScore := -100000
	var bestMove []int
	emptyCells := p.getEmptyCells(board)
	if len(emptyCells) == 0 {
		return -1, -1
	}
	if len(emptyCells) == 1 {
		return emptyCells[0][0], emptyCells[0][1]
	}

	// Определяем размер зоны (например, 3x3)
	zoneSize := 3
	if board.Size < zoneSize {
		zoneSize = board.Size // Если доска меньше зоны, зона равна доске
	}

	type moveResult struct {
		move  []int
		score int
	}
	// Используем буферизированный канал, чтобы не блокировать горутины,
	// если основная горутина не успевает обрабатывать результаты

	// Размер канала равен количеству пустых клеток,
	// т.к. для каждой может быть запущена горутина
	resultChan := make(chan moveResult, len(emptyCells))
	numZonesToProcess := 0 // Счетчик для корректного ожидания

	for _, cell := range emptyCells {
		numZonesToProcess++
		// Запускаем горутину для каждой пустой клетки
		go func(centerCell []int) {
			localBestScore := -100000
			var localBestMove []int

			// Определяем границы зоны вокруг centerCell
			minRow := max(0, centerCell[0]-zoneSize/2)
			maxRow := min(board.Size-1, centerCell[0]+zoneSize/2)
			minCol := max(0, centerCell[1]-zoneSize/2)
			maxCol := min(board.Size-1, centerCell[1]+zoneSize/2)

			// Ищем ходы в зоне
			foundMoveInZone := false
			for r := minRow; r <= maxRow; r++ {
				for c := minCol; c <= maxCol; c++ {
					// Если найден пустая клетка в зоне
					if board.Board[r][c] == b.Empty {
						foundMoveInZone = true
						boardCopy := p.copyBoard(board)
						boardCopy.Board[r][c] = p.Figure
						// Оцениваем ход испо
						score := p.evaluateBoardHeuristic(boardCopy, p.Figure)

						// Если найден лучший ход
						if score > localBestScore {
							localBestScore = score
							localBestMove = []int{r, c}
						}
					}
				}
			}

			// Если найден лучший ход в зоне
			if foundMoveInZone && localBestMove != nil {
				resultChan <- moveResult{
					move: localBestMove, score: localBestScore,
				}
			} else if !foundMoveInZone &&
				board.Board[centerCell[0]][centerCell[1]] == b.Empty {
				// Если зона вокруг centerCell не содержит других
				// пустых клеток, но сама centerCell пуста –
				// оцениваем ход в centerCell
				boardCopy := p.copyBoard(board)
				boardCopy.Board[centerCell[0]][centerCell[1]] = p.Figure
				score := p.evaluateBoardHeuristic(boardCopy, p.Figure)
				resultChan <- moveResult{move: centerCell, score: score}
			} else {
				// Если не найдено ходов в зоне или centerCell не пуста
				// (не должно случиться, если итерируем по emptyCells),
				// отправляем фиктивный результат,
				// чтобы не блокировать ожидание.
				// Этого не должно происходить в нормальном потоке.
				resultChan <- moveResult{move: nil, score: -200000}
			}
		}(cell)
	}

	// Ожидаем завершения всех горутин
	processedGoroutines := 0 // Счетчик для корректного ожидания
	for processedGoroutines < numZonesToProcess {
		result := <-resultChan
		processedGoroutines++
		// Если найден лучший ход
		if result.move != nil && result.score > bestScore {
			bestScore = result.score
			bestMove = result.move
		}
	}

	if bestMove == nil {
		// Если по какой-то причине лучший ход не найден (маловероятно)
		// переходи на стратегию поведения среднего уровня сложности
		return p.makeMediumMove(board)
	}

	// Возвращаем лучший ход
	return bestMove[0], bestMove[1]
}
