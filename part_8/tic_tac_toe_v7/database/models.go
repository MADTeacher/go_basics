package database

import (
	"encoding/json"
	b "tic-tac-toe/board"
	m "tic-tac-toe/model"
	"time"
)

// PlayerFinishGame представляет модель таблицы
// для хранения завершенной игры в БД
type PlayerFinishGame struct {
	// ID игры
	ID int `gorm:"primary_key;autoIncrement;not null"`
	// Имя победителя
	WinnerName string `gorm:"not null"`
	// Имя противника
	AnotherPlayerName string `gorm:"not null"`
	// JSON-представление доски
	BoardJSON []byte `gorm:"type:json;not null"`
	// Символ победившего игрока
	PlayerFigure int `gorm:"not null"`
	// Время завершения игры
	Time time.Time `gorm:"not null"`
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
		ID:                f.ID,
		Board:             &board,
		PlayerFigure:      b.BoardField(f.PlayerFigure),
		WinnerName:        f.WinnerName,
		AnotherPlayerName: f.AnotherPlayerName,
		Time:              f.Time,
	}, nil
}
