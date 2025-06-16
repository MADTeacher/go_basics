package storage

import (
	"encoding/json"
	"os"
	"strings"
	m "tic-tac-toe/model"
)

func NewJsonGameStorage() IGameStorage {
	return &JsonGameStorage{}
}

type JsonGameStorage struct{}

func (j *JsonGameStorage) LoadGame(path string) (*m.GameSnapshot, error) {
	if !strings.HasSuffix(path, ".json") {
		path += ".json"
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var snapshot m.GameSnapshot
	err = decoder.Decode(&snapshot)
	if err != nil {
		return nil, err
	}

	return &snapshot, nil
}

func (j *JsonGameStorage) SaveGame(path string, game *m.GameSnapshot) error {
	if !strings.HasSuffix(path, ".json") {
		path += ".json"
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(game)
	if err != nil {
		return err
	}
	return nil
}
