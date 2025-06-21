package game

type GameState int

// состояние игрового процесса
const (
	WaitingOpponent GameState = iota
	Draw
	CrossWin
	NoughtWin
	CrossStep
	NoughtStep
)

// Режим игры
type GameMode int

const (
	PvP GameMode = iota
	PvC
)

// Уровни сложности компьютера
type Difficulty int

const (
	None Difficulty = iota
	Easy
	Medium
	Hard
)
