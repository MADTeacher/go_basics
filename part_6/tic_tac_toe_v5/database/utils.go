package database

import (
	"encoding/json"

	b "tic-tac-toe/board"
	m "tic-tac-toe/model"
	p "tic-tac-toe/player"
)

// Задаем имя таблицы для структуры Player
func (p *Player) TableName() string {
	return "players"
}

// Задаем имя таблицы для структуры PlayerFinishGame
func (pfg *PlayerFinishGame) TableName() string {
	return "player_finish_games"
}

// Преобразуем таблицу PlayerFinishGame в модель PlayerFinishGame
// из пакета model
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

// Задаем имя таблицы для структуры GameSnapshot
func (g *GameSnapshot) TableName() string {
	return "game_snapshots"
}

// Преобразуем таблицу GameSnapshot в модель GameSnapshot
// из пакета model
func (gs *GameSnapshot) ToModel() (*m.GameSnapshot, error) {
	var board b.Board
	if err := json.Unmarshal(gs.BoardJSON, &board); err != nil {
		return nil, err
	}

	return &m.GameSnapshot{
		Board:        &board,
		PlayerFigure: b.BoardField(gs.PlayerFigure),
		State:        gs.State,
		Mode:         gs.Mode,
		Difficulty:   p.Difficulty(gs.Difficulty),
		SnapshotName: gs.SnapshotName,
	}, nil
}
