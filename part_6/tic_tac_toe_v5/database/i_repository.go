package database

import "tic-tac-toe/model"

type IRepository interface {
	CreatePlayer(nickName string) (*Player, error)
	GetPlayer(nickName string) (*Player, error)
	SaveSnapshot(snapshot *model.GameSnapshot, playerNickName string) error
	GetSnapshots(nickName string) (*[]model.GameSnapshot, error)
	IsSnapshotExist(snapshotName string, nickName string) (bool, error)
	SaveFinishedGame(snapshot *model.FinishGameSnapshot) error
	GetFinishedGames(nickName string) (*[]model.FinishGameSnapshot, error)
}
