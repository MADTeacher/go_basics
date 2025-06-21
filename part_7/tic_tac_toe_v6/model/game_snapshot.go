package model

import (
	b "tic-tac-toe/board"
	p "tic-tac-toe/player"
)

// Структура для сериализации/десериализации игры
type GameSnapshot struct {
	SnapshotName string       `json:"snapshot_name"`
	Board        *b.Board     `json:"board"`
	PlayerFigure b.BoardField `json:"player_figure"`
	State        int          `json:"state"`
	Mode         int          `json:"mode"`
	Difficulty   p.Difficulty `json:"difficulty,omitempty"`
}
