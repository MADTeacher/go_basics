package player

import (
	"fmt"
	"math/rand"
	"net"
	b "tic-tac-toe/board"
	g "tic-tac-toe/game"
	"tic-tac-toe/network"
	"time"
)

// Структура для представления игрока-компьютера
type ComputerPlayer struct {
	Figure     b.BoardField `json:"figure"`
	Difficulty g.Difficulty `json:"difficulty"`
	rand       *rand.Rand
}

// Создаем нового игрока-компьютера с заданным уровнем сложности
func NewComputerPlayer(
	figure b.BoardField,
	difficulty g.Difficulty,
) *ComputerPlayer {
	source := rand.NewSource(time.Now().UnixNano())
	return &ComputerPlayer{
		Figure:     figure,
		Difficulty: difficulty,
		rand:       rand.New(source),
	}
}

func (p *ComputerPlayer) GetSymbol() string {
	if p.Figure == b.Cross {
		return "X"
	}
	return "O"
}

func (p *ComputerPlayer) SendMessage(msg *network.Message) {

}

func (p *ComputerPlayer) GetNickname() string {
	return "Computer"
}

func (p *ComputerPlayer) SwitchPlayer() {
	if p.Figure == b.Cross {
		p.Figure = b.Nought
	} else {
		p.Figure = b.Cross
	}
}

func (p *ComputerPlayer) GetFigure() b.BoardField {
	return p.Figure
}

func (p *ComputerPlayer) IsComputer() bool {
	return true
}

// Реализуем ход компьютера в зависимости от выбранной сложности
func (p *ComputerPlayer) MakeMove(board *b.Board) (int, int, bool) {
	fmt.Printf("%s (Computer) making move... ", p.GetSymbol())

	var row, col int
	switch p.Difficulty {
	case g.Easy:
		row, col = p.makeEasyMove(board)
	case g.Medium:
		row, col = p.makeMediumMove(board)
	case g.Hard:
		row, col = p.makeHardMove(board)
	}

	fmt.Printf("Move made (%d, %d)\n", row+1, col+1)
	return row, col, true
}

// Легкий уровень: случайный ход на свободную клетку
func (p *ComputerPlayer) makeEasyMove(board *b.Board) (int, int) {
	emptyCells := p.getEmptyCells(board)
	if len(emptyCells) == 0 {
		return -1, -1
	}

	// Выбираем случайную свободную клетку
	randomIndex := p.rand.Intn(len(emptyCells))
	return emptyCells[randomIndex][0], emptyCells[randomIndex][1]
}

func (p *ComputerPlayer) CheckSocket(conn net.Conn) bool {
	return false
}

// Средний уровень: проверяет возможность выигрыша
// или блокировки выигрыша противника
func (p *ComputerPlayer) makeMediumMove(board *b.Board) (int, int) {
	// Проверяем, можем ли мы выиграть за один ход
	if move := p.findWinningMove(board, p.Figure); move != nil {
		return move[0], move[1]
	}

	// Проверяем, нужно ли блокировать победу противника
	opponentFigure := b.Cross
	if p.Figure == b.Cross {
		opponentFigure = b.Nought
	}

	if move := p.findWinningMove(board, opponentFigure); move != nil {
		return move[0], move[1]
	}

	// Занимаем центр, если свободен (хорошая стратегия)
	center := board.Size / 2
	if board.Board[center][center] == b.Empty {
		return center, center
	}

	// Занимаем угол, если свободен
	corners := [][]int{
		{0, 0},
		{0, board.Size - 1},
		{board.Size - 1, 0},
		{board.Size - 1, board.Size - 1},
	}

	for _, corner := range corners {
		if board.Board[corner[0]][corner[1]] == b.Empty {
			return corner[0], corner[1]
		}
	}

	// Если нет лучшего хода, делаем случайный ход
	return p.makeEasyMove(board)
}

// Сложный уровень: использует алгоритм минимакс для оптимального хода
func (p *ComputerPlayer) makeHardMove(board *b.Board) (int, int) {
	// Если доска пустая, ходим в центр или угол (оптимальный первый ход)
	emptyCells := p.getEmptyCells(board)
	if len(emptyCells) == board.Size*board.Size {
		// Первый ход - центр или угол
		center := board.Size / 2
		return center, center
	}

	// Используем минимакс для доски 3x3
	// Для больших досок это слишком ресурсоемко
	if board.Size <= 3 {
		bestScore := -1000
		bestMove := []int{-1, -1}

		// Создаем канал для результатов
		type moveResult struct {
			move  []int
			score int
		}
		resultChan := make(chan moveResult, len(emptyCells))

		// Запускаем горутину для каждого возможного хода
		for _, cell := range emptyCells {
			go func(cell []int) {
				row, col := cell[0], cell[1]
				// Копируем доску чтобы избежать гонок данных
				boardCopy := p.copyBoard(board)

				// Пробуем сделать ход
				boardCopy.Board[row][col] = p.Figure

				// Вычисляем оценку хода через минимакс
				score := p.minimax(boardCopy, 0, false)

				// Отправляем результат в канал
				resultChan <- moveResult{
					move:  []int{row, col},
					score: score,
				}
			}(cell)
		}

		// Собираем результаты всех горутин
		for i := 0; i < len(emptyCells); i++ {
			result := <-resultChan
			if result.score > bestScore {
				bestScore = result.score
				bestMove = result.move
			}
		}

		return bestMove[0], bestMove[1]
	}

	// Для больших досок выбираем случайно одну из трех параллельных стратегий
	strategyChoice := p.rand.Intn(3)
	switch strategyChoice {
	case 0:
		fmt.Println("Using limited-depth parallel minimax strategy")
		return p.makeLimitedDepthMinimax(board)
	case 1:
		fmt.Println("Using parallel heuristic evaluation strategy")
		return p.makeParallelHeuristicMove(board)
	case 2:
		fmt.Println("Using zone-based parallel analysis strategy")
		return p.makeZoneBasedMove(board)
	default:
		//В случае ошибки используем стратегию среднего уровня
		return p.makeMediumMove(board)
	}
}

// Алгоритм минимакс для определения оптимального хода
func (p *ComputerPlayer) minimax(
	board *b.Board,
	depth int, isMaximizing bool,
) int {
	opponentFigure := b.Cross
	if p.Figure == b.Cross {
		opponentFigure = b.Nought
	}

	// Проверяем терминальное состояние
	if board.CheckWin(p.Figure) {
		return 10 - depth // Выигрыш, чем быстрее, тем лучше
	} else if board.CheckWin(opponentFigure) {
		return depth - 10 // Проигрыш, чем дольше, тем лучше
	} else if board.CheckDraw() {
		return 0 // Ничья
	}

	emptyCells := p.getEmptyCells(board)

	if isMaximizing {
		bestScore := -1000

		// Проходим по всем свободным клеткам
		for _, cell := range emptyCells {
			row, col := cell[0], cell[1]

			// Делаем ход
			board.Board[row][col] = p.Figure

			// Рекурсивно оцениваем ход
			score := p.minimax(board, depth+1, false)

			// Отменяем ход
			board.Board[row][col] = b.Empty

			bestScore = max(score, bestScore)
		}

		return bestScore
	} else {
		bestScore := 1000

		// Проходим по всем свободным клеткам
		for _, cell := range emptyCells {
			row, col := cell[0], cell[1]

			// Делаем ход противника
			board.Board[row][col] = opponentFigure

			// Рекурсивно оцениваем ход
			score := p.minimax(board, depth+1, true)

			// Отменяем ход
			board.Board[row][col] = b.Empty

			bestScore = min(score, bestScore)
		}

		return bestScore
	}
}

// Вспомогательная функция для поиска хода, приводящего к выигрышу
func (p *ComputerPlayer) findWinningMove(
	board *b.Board,
	figure b.BoardField,
) []int {
	for _, cell := range p.getEmptyCells(board) {
		row, col := cell[0], cell[1]

		// Пробуем сделать ход
		board.Board[row][col] = figure

		// Проверяем, приведет ли этот ход к выигрышу
		if board.CheckWin(figure) {
			// Отменяем ход и возвращаем координаты
			board.Board[row][col] = b.Empty
			return []int{row, col}
		}

		// Отменяем ход
		board.Board[row][col] = b.Empty
	}

	return nil // Нет выигрышного хода
}

// Получение списка пустых клеток
func (p *ComputerPlayer) getEmptyCells(board *b.Board) [][]int {
	var emptyCells [][]int

	for i := 0; i < board.Size; i++ {
		for j := 0; j < board.Size; j++ {
			if board.Board[i][j] == b.Empty {
				emptyCells = append(emptyCells, []int{i, j})
			}
		}
	}

	return emptyCells
}

// Вспомогательные функции max и min
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Копирование доски для избежания гонок данных при параллельном вычислении
func (p *ComputerPlayer) copyBoard(board *b.Board) *b.Board {
	newBoard := b.NewBoard(board.Size)
	for i := 0; i < board.Size; i++ {
		for j := 0; j < board.Size; j++ {
			newBoard.Board[i][j] = board.Board[i][j]
		}
	}
	return newBoard
}
