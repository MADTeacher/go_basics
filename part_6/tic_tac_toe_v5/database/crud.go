package database

import (
	"encoding/json"

	m "tic-tac-toe/model"
)

func (r *SQLiteRepository) CreatePlayer(nickName string) (*Player, error) {
	player := &Player{NickName: nickName}
	if err := r.db.Create(player).Error; err != nil {
		return nil, err
	}
	return player, nil
}

func (r *SQLiteRepository) GetPlayer(nickName string) (*Player, error) {
	var player Player
	if err := r.db.Where(
		"nick_name = ?", nickName,
	).First(&player).Error; err != nil {
		return nil, err
	}
	return &player, nil
}

func (r *SQLiteRepository) SaveSnapshot(
	snapshot *m.GameSnapshot,
	playerNickName string,
) error {
	player, _ := r.GetPlayer(playerNickName)
	if player == nil {
		player, _ = r.CreatePlayer(playerNickName)
	}

	boardJSON, err := json.Marshal(snapshot.Board)
	if err != nil {
		return err
	}

	return r.db.Create(&GameSnapshot{
		BoardJSON:      boardJSON,
		PlayerFigure:   int(snapshot.PlayerFigure),
		State:          int(snapshot.State),
		Mode:           int(snapshot.Mode),
		Difficulty:     int(snapshot.Difficulty),
		IsCurrentFirst: snapshot.IsCurrentFirst,
		PlayerNickName: player.NickName,
	}).Error
}

func (r *SQLiteRepository) GetSnapshots(
	nickName string) (*[]m.GameSnapshot, error) {
	var snapshots []GameSnapshot
	// ищем игрока по никнейму
	player, err := r.GetPlayer(nickName)
	if err != nil {
		return nil, err
	}

	// находим все снапшоты игрока
	if err := r.db.Where(
		"player_nick_name = ?", player.NickName,
	).Find(&snapshots).Error; err != nil {
		return nil, err
	}

	var gameSnapshots []m.GameSnapshot
	for _, snapshot := range snapshots {
		temp, err := snapshot.ToModel()
		if err != nil {
			return nil, err
		}
		gameSnapshots = append(gameSnapshots, *temp)
	}
	return &gameSnapshots, nil
}

func (r *SQLiteRepository) IsSnapshotExist(snapshotName string, nickName string) (bool, error) {
	var snapshot GameSnapshot
	if err := r.db.Where(
		"snapshot_name = ? AND player_nick_name = ?", snapshotName, nickName,
	).First(&snapshot).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *SQLiteRepository) SaveFinishedGame(
	snapshot *m.FinishGameSnapshot) error {
	boardJSON, err := json.Marshal(snapshot.Board)
	if err != nil {
		return err
	}

	player, _ := r.GetPlayer(snapshot.PlayerNickName)
	if player == nil {
		player, _ = r.CreatePlayer(snapshot.PlayerNickName)
	}

	return r.db.Create(&PlayerFinishGame{
		BoardJSON:      boardJSON,
		PlayerFigure:   int(snapshot.PlayerFigure),
		WinnerName:     snapshot.WinnerName,
		PlayerNickName: player.NickName,
		Time:           snapshot.Time,
	}).Error
}

func (r *SQLiteRepository) GetFinishedGames(nickName string) (*[]m.FinishGameSnapshot, error) {
	var playerFinishGames []PlayerFinishGame
	if err := r.db.Where(
		"player_nick_name = ?", nickName,
	).Find(&playerFinishGames).Error; err != nil {
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
