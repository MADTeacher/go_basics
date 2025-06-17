package game

import (
	"bufio"
	b "tic-tac-toe/board"
	db "tic-tac-toe/database"
	p "tic-tac-toe/player"
)

type Game struct {
	Board  *b.Board  `json:"board"`
	Player p.IPlayer `json:"player"`
	// Не сериализуется напрямую
	Player2 p.IPlayer `json:"-"`
	// Не сериализуется напрямую
	CurrentPlayer p.IPlayer      `json:"-"`
	Reader        *bufio.Reader  `json:"-"`
	State         GameState      `json:"state"`
	repository    db.IRepository `json:"-"`
	// Режим игры (PvP или PvC)
	Mode GameMode `json:"mode"`
	// Уровень сложности компьютера (только для PvC)
	Difficulty p.Difficulty `json:"difficulty,omitempty"`
	// Флаг для определения текущего игрока
	IsCurrentFirst bool `json:"is_current_first"`
}

// Создаем новую игру
func NewGame(board b.Board, reader *bufio.Reader, repository db.IRepository,
	mode GameMode, difficulty p.Difficulty) *Game {
	// Создаем первого игрока (всегда человек на X)
	player1 := p.NewHumanPlayer(b.Cross, reader)

	var player2 p.IPlayer
	if mode == PlayerVsPlayer {
		// Для режима игрок против игрока создаем второго человека-игрока
		player2 = p.NewHumanPlayer(b.Nought, reader)
	} else {
		// Для режима игрок против компьютера создаем компьютерного игрока
		player2 = p.NewComputerPlayer(b.Nought, difficulty)
	}

	return &Game{
		Board:          &board,
		Player:         player1,
		Player2:        player2,
		CurrentPlayer:  player1,
		Reader:         reader,
		State:          playing,
		repository:     repository,
		Mode:           mode,
		Difficulty:     difficulty,
		IsCurrentFirst: true,
	}
}

// Переключаем активного игрока
func (g *Game) switchCurrentPlayer() {
	if g.CurrentPlayer == g.Player {
		g.CurrentPlayer = g.Player2
	} else {
		g.CurrentPlayer = g.Player
	}
}

// Обновляем состояние игры
func (g *Game) updateState() {
	if g.Board.CheckWin(g.CurrentPlayer.GetFigure()) {
		if g.CurrentPlayer.GetFigure() == b.Cross {
			g.State = crossWin
		} else {
			g.State = noughtWin
		}
		return
	}

	if g.Board.CheckDraw() {
		g.State = draw
	}
}
