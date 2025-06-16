package game

type IGameLoader interface {
	LoadGame(path string) (*Game, error)
}

type IGameSaver interface {
	SaveGame(path string, game *Game) error
}
