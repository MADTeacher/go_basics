package game

import (
	"bufio"
	"fmt"
	b "tic-tac-toe/board"
	m "tic-tac-toe/model"
	p "tic-tac-toe/player"
	s "tic-tac-toe/storage"
)

// Подготавливаем игру к сохранению
func (g *Game) prepareForSave() {
	// Устанавливаем флаг текущего игрока
	g.IsCurrentFirst = (g.CurrentPlayer == g.Player)
}

// Возвращаем снапшот игровой сессии
func (g *Game) gameSnapshot() *m.GameSnapshot {
	g.prepareForSave()
	return &m.GameSnapshot{
		Board:          g.Board,
		PlayerFigure:   g.Player.GetFigure(),
		State:          int(g.State),
		Mode:           int(g.Mode),
		Difficulty:     g.Difficulty,
		IsCurrentFirst: g.IsCurrentFirst,
	}
}

// Восстанавливаем игру из снапшота
func (g *Game) RestoreFromSnapshot(
	snapshot *m.GameSnapshot,
	reader *bufio.Reader,
	saver s.IGameSaver,
) {
	g.Board = snapshot.Board
	g.State = GameState(snapshot.State)
	g.Mode = GameMode(snapshot.Mode)
	g.Difficulty = snapshot.Difficulty
	g.IsCurrentFirst = snapshot.IsCurrentFirst

	// Создаем объекты игроков
	g.Player = &p.HumanPlayer{Figure: snapshot.PlayerFigure}

	g.Reader = reader
	g.Saver = saver

	g.recreatePlayersAfterLoad(reader)
}

// Восстанавливаем объекты игроков после загрузки из JSON
func (g *Game) recreatePlayersAfterLoad(reader *bufio.Reader) {
	// Создаем игроков в зависимости от режима игры
	if g.Player == nil {
		fmt.Println("Error: Player is nil")
		return
	}

	playerFigure := g.Player.GetFigure()
	g.Player = p.NewHumanPlayer(playerFigure, reader)

	// Получаем фигуру второго игрока
	var player2Figure b.BoardField
	if playerFigure == b.Cross {
		player2Figure = b.Nought
	} else {
		player2Figure = b.Cross
	}

	// Создаем второго игрока в зависимости от режима
	if g.Mode == PlayerVsPlayer {
		g.Player2 = p.NewHumanPlayer(player2Figure, reader)
	} else {
		g.Player2 = p.NewComputerPlayer(player2Figure, g.Difficulty)
	}

	// Восстанавливаем указатель на текущего игрока
	if g.IsCurrentFirst {
		g.CurrentPlayer = g.Player
	} else {
		g.CurrentPlayer = g.Player2
	}
}
