package game

type GameState int

// состояние игрового процесса
const (
	playing GameState = iota
	draw
	crossWin
	noughtWin
	quit
)
