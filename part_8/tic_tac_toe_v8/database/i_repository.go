package database

import "tic-tac-toe/model"

// Интерфейс для работы с базой данных
type IRepository interface {
	// Сохраняет информацию о завершенной игре
	SaveFinishedGame(snapshot *model.FinishGameSnapshot) error
	// Получает все завершенные игры для указанного игрока
	GetAllFinishedGames() (*[]model.FinishGameSnapshot, error)
	// Получает конкретную завершенную игру по ID
	GetFinishedGameById(id int) (*model.FinishGameSnapshot, error)
}
