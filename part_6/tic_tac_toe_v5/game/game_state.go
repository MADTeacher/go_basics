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

// Режим игры
type GameMode int

const (
	PlayerVsPlayer GameMode = iota
	PlayerVsComputer
)
