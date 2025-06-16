package storage

import m "tic-tac-toe/model"

type IGameLoader interface {
	LoadGame(path string) (*m.GameSnapshot, error)
}

type IGameSaver interface {
	SaveGame(path string, game *m.GameSnapshot) error
}

type IGameStorage interface {
	IGameLoader
	IGameSaver
}
