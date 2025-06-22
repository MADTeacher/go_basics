package model

import (
	"tic-tac-toe/board"
	"time"
)

type FinishGameSnapshot struct {
	// Идентификатор игры
	ID int `json:"id"`
	// Состояние доски в момент завершения игры
	Board *board.Board `json:"board"`
	// Символ (Х/О) победившего игрока
	PlayerFigure board.BoardField `json:"player_figure"`
	// Имя победителя
	WinnerName string `json:"winner_name"`
	// Имя противника
	AnotherPlayerName string `json:"another_player_name"`
	// Время завершения игры
	Time time.Time `json:"time"`
}
