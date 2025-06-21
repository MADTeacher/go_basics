package database

import "time"

// PlayerFinishGame представляет модель таблицы
// для хранения завершенной игры в БД
type PlayerFinishGame struct {
	ID                int       `gorm:"primary_key;autoIncrement;not null"`
	WinnerName        string    `gorm:"not null"`
	AnotherPlayerName string    `gorm:"not null"`
	BoardJSON         []byte    `gorm:"type:json;not null"`
	PlayerFigure      int       `gorm:"not null"`
	Time              time.Time `gorm:"not null"`
}
