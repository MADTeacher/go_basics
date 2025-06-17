package database

import "tic-tac-toe/model"

// Интерфейс для работы с базой данных
type IRepository interface {
	// Сохраняет снапшот игры для указанного игрока
	SaveSnapshot(snapshot *model.GameSnapshot, playerNickName string) error
	// Получает все снапшоты игр для указанного игрока
	GetSnapshots(nickName string) (*[]model.GameSnapshot, error)
	// Проверяет существует ли снапшот с указанным именем для данного игрока
	IsSnapshotExist(snapshotName string, nickName string) (bool, error)
	// Сохраняет информацию о завершенной игре
	SaveFinishedGame(snapshot *model.FinishGameSnapshot) error
	// Получает все завершенные игры для указанного игрока
	GetFinishedGames(nickName string) (*[]model.FinishGameSnapshot, error)
}
