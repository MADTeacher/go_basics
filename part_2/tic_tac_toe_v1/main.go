package main

import (
	"tic-tac-toe/game"
)

func main() {
	for {
		if game.InitBoard() {
			break
		}
	}

	game.Play()
}
