package database

import "time"

type Player struct {
	NickName string `gorm:"primary_key;not null"`
}

type PlayerFinishGame struct {
	ID             int       `gorm:"primary_key;autoIncrement;not null"`
	WinnerName     string    `gorm:"not null"`
	BoardJSON      []byte    `gorm:"type:json;not null"`
	PlayerFigure   int       `gorm:"not null"`
	Time           time.Time `gorm:"not null"`
	PlayerNickName string    `gorm:"not null"`
	Player         *Player   `gorm:"foreignKey:PlayerNickName;references:NickName"`
}

// GameSnapshot представляет модель для хранения снапшота игры в БД
type GameSnapshot struct {
	ID             int     `gorm:"primaryKey;autoIncrement;not null"`
	SnapshotName   string  `gorm:"not null"`
	BoardJSON      []byte  `gorm:"type:json;not null"`
	PlayerFigure   int     `gorm:"not null"`
	State          int     `gorm:"not null"`
	Mode           int     `gorm:"not null"`
	Difficulty     int     `gorm:"not null"`
	IsCurrentFirst bool    `gorm:"not null"`
	PlayerNickName string  `gorm:"not null"`
	Player         *Player `gorm:"foreignKey:PlayerNickName;references:NickName"`
}
