package model

import (
	"tic-tac-toe/board"
	"time"
)

type FinishGameSnapshot struct {
	Board          *board.Board     `json:"board"`
	PlayerFigure   board.BoardField `json:"player_figure"`
	WinnerName     string           `json:"winner_name"`
	PlayerNickName string           `json:"nick_name"`
	Time           time.Time        `json:"time"`
}
