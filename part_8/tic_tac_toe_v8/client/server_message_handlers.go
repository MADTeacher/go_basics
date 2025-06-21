package client

import (
	"encoding/json"
	"fmt"
	"log"

	b "tic-tac-toe/board"
	g "tic-tac-toe/game"
	"tic-tac-toe/network"
)

// handleRoomJoinResponse processes the RoomJoinResponse message from the server.
func (c *Client) handleRoomJoinResponse(payload json.RawMessage) {
	var res network.RoomJoinResponse
	if err := json.Unmarshal(payload, &res); err == nil {
		c.mySymbol = res.PlayerSymbol
		c.roomName = res.RoomName
		if res.Board.Size > 0 { // Check if board is valid
			c.board = &res.Board
			fmt.Printf("\nSuccessfully joined room '%s' as %s.\n", res.RoomName, res.PlayerSymbol)
			c.board.PrintBoard()
		} else {
			fmt.Printf("\nSuccessfully joined room '%s' as %s. Waiting for game to start...\n", res.RoomName, res.PlayerSymbol)
		}
	} else {
		log.Printf("Error unmarshalling RoomJoinResponse: %v", err)
	}
	c.setState(waitingOpponentInRoom)
}

// handleInitGame processes the InitGameResponse message from the server.
func (c *Client) handleInitGame(payload json.RawMessage) {
	var res network.InitGameResponse
	if err := json.Unmarshal(payload, &res); err == nil {
		c.board = &res.Board
		c.currentPlayer = res.CurrentPlayer
		fmt.Println("\n--- Game Started ---")
		c.board.PrintBoard()
		c.printTurnInfo()
		if res.CurrentPlayer == c.mySymbol {
			c.setState(playerMove)
		} else {
			c.setState(opponentMove)
		}
	} else {
		log.Printf("Error unmarshalling InitGameResponse: %v", err)
	}
}

// handleUpdateState processes the GameStateUpdate message from the server.
func (c *Client) handleUpdateState(payload json.RawMessage) {
	var res network.GameStateUpdate
	if err := json.Unmarshal(payload, &res); err == nil {
		c.board = &res.Board
		c.currentPlayer = res.CurrentPlayer
		fmt.Println("\n--- Game State Update ---")
		c.board.PrintBoard()
		c.printTurnInfo()
		if res.CurrentPlayer == c.mySymbol {
			c.setState(playerMove)
		} else {
			c.setState(opponentMove)
		}
	} else {
		log.Printf("Error unmarshalling GameStateUpdate: %v", err)
	}
}

// handleEndGame processes the EndGameResponse message from the server.
func (c *Client) handleEndGame(payload json.RawMessage) {
	var res network.EndGameResponse
	if err := json.Unmarshal(payload, &res); err == nil {
		c.board = &res.Board
		fmt.Println("\n--- Game Over ---")
		c.board.PrintBoard()
		if res.CurrentPlayer == b.Empty {
			fmt.Println("It's a Draw!")
		} else {
			fmt.Printf("Player %s wins!\n", res.CurrentPlayer)
		}
		c.setState(endGame)
		fmt.Print("> ")
	} else {
		log.Printf("Error unmarshalling EndGameResponse: %v", err)
	}
}

// handleError processes the ErrorResponse message from the server.
func (c *Client) handleError(payload json.RawMessage) {
	var errPayload network.ErrorResponse
	if err := json.Unmarshal(payload, &errPayload); err == nil {
		fmt.Printf("\nServer Error: %s\n> ", errPayload.Message)
	} else {
		log.Printf("Error unmarshalling ErrorResponse: %v", err)
	}
	c.setState(mainMenu)
}

// gameModeToString converts GameMode to a string representation.
func gameModeToString(mode g.GameMode) string {
	switch mode {
	case g.PvP:
		return "PvP"
	case g.PvC:
		return "PvC"
	default:
		return "Unknown"
	}
}

func difficultyToString(difficulty g.Difficulty) string {
	switch difficulty {
	case g.Easy:
		return "Easy"
	case g.Medium:
		return "Medium"
	case g.Hard:
		return "Hard"
	default:
		return ""
	}
}

// handleFinishedGamesResponse processes the FinishedGamesResponse message from the server.
func (c *Client) handleFinishedGamesResponse(payload json.RawMessage) {
	var res network.FinishedGamesResponse
	if err := json.Unmarshal(payload, &res); err == nil {
		fmt.Println("\nFinished games:")
		if res.Games == nil || len(*res.Games) == 0 {
			fmt.Println("No finished games.")
		} else {
			for _, game := range *res.Games {
				fmt.Printf("- Game #%d: %s vs %s (Winner: %s) at %v\n",
					game.ID, game.WinnerName,
					game.AnotherPlayerName, game.WinnerName,
					game.Time.Format("2006-01-02 15:04:05"),
				)
			}
		}
	} else {
		log.Printf("Error unmarshalling FinishedGamesResponse: %v", err)
	}
	c.setState(mainMenu)
}

// handleRoomListResponse processes the RoomListResponse message from the server.
func (c *Client) handleRoomListResponse(payload json.RawMessage) {
	var roomList network.RoomListResponse
	if err := json.Unmarshal(payload, &roomList); err == nil {
		fmt.Println("\nAvailable rooms:")
		if len(roomList.Rooms) == 0 {
			fmt.Println("No rooms available.")
		} else {
			for _, room := range roomList.Rooms {
				fmt.Printf("- %s (Board Size: %dx%d, Full: %t, Mode: %s, Difficulty: %s)\n",
					room.Name,
					room.BoardSize, room.BoardSize,
					room.IsFull,
					gameModeToString(room.GameMode),
					difficultyToString(room.Difficult),
				)
			}
		}
	} else {
		log.Printf("Error unmarshalling RoomListResponse: %v", err)
	}
	c.setState(mainMenu)
}

// handleNickNameResponse processes the NickNameResponse message from the server.
func (c *Client) handleNickNameResponse(payload json.RawMessage) {
	var res network.NickNameResponse
	if err := json.Unmarshal(payload, &res); err == nil {
		fmt.Printf("\nWelcome, %s!\n> ", res.Nickname)
		c.setNickname(res.Nickname)
		c.setState(mainMenu)
	} else {
		log.Printf("Error unmarshalling NickNameResponse: %v", err)
	}
}

// handleOpponentLeft processes the OpponentLeft message from the server.
func (c *Client) handleOpponentLeft(payload json.RawMessage) {
	var res network.OpponentLeft
	if err := json.Unmarshal(payload, &res); err == nil {
		fmt.Printf("\nPlayer '%s' has left the game.\n> ", res.Nickname)
	} else {
		log.Printf("Error unmarshalling OpponentLeft: %v", err)
	}
	c.setState(waitingOpponentInRoom)
}

// handleFinishedGameResponse processes the FinishedGameResponse message from the server.
func (c *Client) handleFinishedGameResponse(payload json.RawMessage) {
	var res network.FinishedGameResponse
	if err := json.Unmarshal(payload, &res); err == nil {
		fmt.Println("\nFinished game:")
		fmt.Printf("- Game #%d: %s vs %s (Winner: %s) at %v\n",
			res.Game.ID, res.Game.WinnerName,
			res.Game.AnotherPlayerName, res.Game.WinnerName,
			res.Game.Time.Format("2006-01-02 15:04:05"),
		)
		c.board = res.Game.Board
		c.board.PrintBoard()
	} else {
		log.Printf("Error unmarshalling FinishedGameResponse: %v", err)
	}
	c.setState(mainMenu)
}
