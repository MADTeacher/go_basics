package database

import (
	"encoding/json"

	m "tic-tac-toe/model"
)

func (r *SQLiteRepository) SaveFinishedGame(
	snapshot *m.FinishGameSnapshot) error {
	boardJSON, err := json.Marshal(snapshot.Board)
	if err != nil {
		return err
	}

	return r.db.Create(&PlayerFinishGame{
		BoardJSON:         boardJSON,
		PlayerFigure:      int(snapshot.PlayerFigure),
		WinnerName:        snapshot.WinnerName,
		AnotherPlayerName: snapshot.AnotherPlayerName,
		Time:              snapshot.Time,
	}).Error
}

func (r *SQLiteRepository) GetAllFinishedGames() (*[]m.FinishGameSnapshot, error) {
	var playerFinishGames []PlayerFinishGame
	if err := r.db.Find(&playerFinishGames).Error; err != nil {
		return nil, err
	}
	var finishGameSnapshots []m.FinishGameSnapshot
	for _, playerFinishGame := range playerFinishGames {
		temp, err := playerFinishGame.ToModel()
		if err != nil {
			return nil, err
		}
		finishGameSnapshots = append(finishGameSnapshots, *temp)
	}
	return &finishGameSnapshots, nil
}

func (r *SQLiteRepository) GetFinishedGameById(id int) (*m.FinishGameSnapshot, error) {
	var playerFinishGame PlayerFinishGame
	if err := r.db.Where("id = ?", id).First(&playerFinishGame).Error; err != nil {
		return nil, err
	}
	return playerFinishGame.ToModel()
}
