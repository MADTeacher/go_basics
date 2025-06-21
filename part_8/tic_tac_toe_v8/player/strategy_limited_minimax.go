package player

import (
	"sync"
	b "tic-tac-toe/board"
)

const maxDepth = 2 // Ограничение глубины для минимакса

// Метод запуска стратегии с ограничением глубины для минимакса
func (p *ComputerPlayer) makeLimitedDepthMinimax(board *b.Board) (int, int) {
	bestScore := -100000
	var bestMove []int
	emptyCells := p.getEmptyCells(board)

	if len(emptyCells) == 0 {
		return -1, -1 // Нет доступных ходов
	}
	if len(emptyCells) == 1 {
		return emptyCells[0][0], emptyCells[0][1] // Единственный возможный ход
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
			score := p.minimaxRecursive(boardCopy, 0, false, maxDepth)
			resultChan <- moveResult{move: []int{r, c}, score: score}
		}(cell[0], cell[1])
	}

	wg.Wait()         // Ждем завершения всех горутин
	close(resultChan) // Закрываем канал

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

// Рекурсивная часть минимакса с ограничением глубины
func (p *ComputerPlayer) minimaxRecursive(
	board *b.Board, depth int, isMaximizing bool,
	maxDepthLimit int,
) int {
	opponentFigure := b.Cross
	if p.Figure == b.Cross {
		opponentFigure = b.Nought
	}

	if board.CheckWin(p.Figure) {
		return 10 - depth // Выигрыш текущего игрока
	}
	if board.CheckWin(opponentFigure) {
		return depth - 10 // Проигрыш текущего игрока (выигрыш оппонента)
	}
	if board.CheckDraw() {
		return 0 // Ничья
	}

	if depth >= maxDepthLimit { // Ограничение глубины
		// Если достигнута максимальная глубина, используем эвристическую оценку
		return p.evaluateBoardHeuristic(board, p.Figure)
	}

	emptyCells := p.getEmptyCells(board)

	if isMaximizing {
		bestScore := -100000
		for _, cell := range emptyCells {
			boardCopy := p.copyBoard(board)
			boardCopy.Board[cell[0]][cell[1]] = p.Figure
			score := p.minimaxRecursive(
				boardCopy, depth+1, false, maxDepthLimit,
			)
			bestScore = max(bestScore, score)
		}
		return bestScore
	} else {
		bestScore := 100000
		// opponentFigure уже определен выше
		for _, cell := range emptyCells {
			boardCopy := p.copyBoard(board)
			boardCopy.Board[cell[0]][cell[1]] = opponentFigure
			score := p.minimaxRecursive(
				boardCopy, depth+1, true, maxDepthLimit,
			)
			bestScore = min(bestScore, score)
		}
		return bestScore
	}
}
