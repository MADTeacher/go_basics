package model

import (
	"tic-tac-toe/board"
	"time"
)

type FinishGameSnapshot struct {
	ID                int              `json:"id"`
	Board             *board.Board     `json:"board"`
	PlayerFigure      board.BoardField `json:"player_figure"`
	WinnerName        string           `json:"winner_name"`
	AnotherPlayerName string           `json:"another_player_name"`
	Time              time.Time        `json:"time"`
}
