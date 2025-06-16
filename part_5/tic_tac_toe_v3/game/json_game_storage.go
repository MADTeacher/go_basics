package game

import (
	"encoding/json"
	"os"
	"strings"
)

func NewJsonGameStorage() IGameStorage {
	return &JsonGameStorage{}
}

type JsonGameStorage struct{}

func (j *JsonGameStorage) LoadGame(path string) (*Game, error) {
	if !strings.HasSuffix(path, ".json") {
		path += ".json"
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var game Game
	err = decoder.Decode(&game)
	if err != nil {
		return nil, err
	}
	return &game, nil
}

func (j *JsonGameStorage) SaveGame(path string, game *Game) error {
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
