package database

import "time"

// Player представляет модель таблицы
// для хранения профилей игроков в БД
type Player struct {
	NickName string `gorm:"primary_key;not null"`
}

// PlayerFinishGame представляет модель таблицы
// для хранения завершенной игры в БД
type PlayerFinishGame struct {
	ID             int       `gorm:"primary_key;autoIncrement;not null"`
	WinnerName     string    `gorm:"not null"`
	BoardJSON      []byte    `gorm:"type:json;not null"`
	PlayerFigure   int       `gorm:"not null"`
	Time           time.Time `gorm:"not null"`
	PlayerNickName string    `gorm:"not null"`
	Player         *Player   `gorm:"foreignKey:PlayerNickName;references:NickName"`
}
