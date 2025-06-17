package database

import (
	"encoding/json"

	b "tic-tac-toe/board"
	m "tic-tac-toe/model"
	p "tic-tac-toe/player"
)

func (p *Player) TableName() string {
	return "players"
}

func (pfg *PlayerFinishGame) TableName() string {
	return "player_finish_games"
}

func (f *PlayerFinishGame) ToModel() (*m.FinishGameSnapshot, error) {
	var board b.Board
	if err := json.Unmarshal(f.BoardJSON, &board); err != nil {
		return nil, err
	}

	return &m.FinishGameSnapshot{
		Board:          &board,
		PlayerFigure:   b.BoardField(f.PlayerFigure),
		WinnerName:     f.WinnerName,
		PlayerNickName: f.PlayerNickName,
		Time:           f.Time,
	}, nil
}

func (g *GameSnapshot) TableName() string {
	return "game_snapshots"
}

func (gs *GameSnapshot) ToModel() (*m.GameSnapshot, error) {
	// Десериализуем BoardJSON в структуру Board
	var board b.Board
	if err := json.Unmarshal(gs.BoardJSON, &board); err != nil {
		return nil, err
	}

	return &m.GameSnapshot{
		Board:          &board,
		PlayerFigure:   b.BoardField(gs.PlayerFigure),
		State:          gs.State,
		Mode:           gs.Mode,
		Difficulty:     p.Difficulty(gs.Difficulty),
		IsCurrentFirst: gs.IsCurrentFirst,
		SnapshotName:   gs.SnapshotName,
	}, nil
}
