package client

import (
	"encoding/json"
	"fmt"
	"log"

	b "tic-tac-toe/board"
	"tic-tac-toe/network"
)

// Обрабатываем ответ сервера, содержащий никнейм игрока
func (c *Client) handleNickNameResponse(payload json.RawMessage) {
	var res network.NickNameResponse // Десериализуем ответ
	if err := json.Unmarshal(payload, &res); err == nil {
		fmt.Printf("\nWelcome, %s!\n> ", res.Nickname)
		c.setNickname(res.Nickname) // Устанавливаем никнейм игрока
		c.setState(mainMenu)        // Переходим в главное меню
	} else {
		log.Printf("Error unmarshalling NickNameResponse: %v", err)
	}
}

// Обрабатываем сообщение о списке комнат
func (c *Client) handleRoomListResponse(payload json.RawMessage) {
	var roomList network.RoomListResponse // Десериализуем ответ
	if err := json.Unmarshal(payload, &roomList); err == nil {
		fmt.Println("\nAvailable rooms:")
		// Если список комнат пуст
		if len(roomList.Rooms) == 0 {
			fmt.Println("No rooms available.")
		} else { // Иначе выводим список комнат
			for _, room := range roomList.Rooms {
				fmt.Printf("- %s (Board Size: %dx%d, Full: %t, "+
					"Mode: %s, Difficulty: %s)\n",
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
	c.setState(mainMenu) // Переходим в главное меню
}

// Обрабатываем ответ на запрос на присоединение к комнате
func (c *Client) handleRoomJoinResponse(payload json.RawMessage) {
	var res network.RoomJoinResponse // Десериализуем ответ
	if err := json.Unmarshal(payload, &res); err == nil {
		c.mySymbol = res.PlayerSymbol // Устанавливаем фигуру игрока
		c.roomName = res.RoomName     // Устанавливаем имя комнаты
		if res.Board.Size > 0 {       // Проверяем размер поля
			c.board = &res.Board // Устанавливаем игровое поле
			fmt.Printf(
				"\nSuccessfully joined room '%s' as %s.\n",
				res.RoomName, res.PlayerSymbol,
			)
			c.board.PrintBoard()
		} else {
			fmt.Printf(
				"\nSuccessfully joined room '%s' as %s. "+
					"Waiting for game to start...\n",
				res.RoomName, res.PlayerSymbol,
			)
		}
	} else {
		log.Printf("Error unmarshalling RoomJoinResponse: %v", err)
	}
	// переходим в состояние ожидания присоединения оппонента к комнате
	c.setState(waitingOpponentInRoom)
}

// Обрабатываем ответ на запрос на получение списка завершенных игр
func (c *Client) handleFinishedGamesResponse(payload json.RawMessage) {
	var res network.FinishedGamesResponse // Десериализуем ответ
	if err := json.Unmarshal(payload, &res); err == nil {
		fmt.Println("\nFinished games:")
		// Если список завершенных игр пуст
		if res.Games == nil || len(*res.Games) == 0 {
			fmt.Println("No finished games.")
		} else { // Иначе выводим список завершенных игр
			for _, game := range *res.Games {
				fmt.Printf("- Game #%d: %s vs %s (Winner: %s) at %v\n",
					game.ID, game.WinnerName,
					game.AnotherPlayerName, game.WinnerName,
					game.Time.Format("2006-01-02 15:04:05"),
				)
			}
		}
	} else {
		log.Printf(
			"Error unmarshalling FinishedGamesResponse: %v",
			err,
		)
	}
	c.setState(mainMenu) // Переходим в главное меню
}

// Обрабатываем ответ на запрос на получение данных
// о конкретной завершенной игре
func (c *Client) handleFinishedGameResponse(payload json.RawMessage) {
	var res network.FinishedGameResponse // Десериализуем ответ
	if err := json.Unmarshal(payload, &res); err == nil {
		// Выводим информацию о завершенной игре
		fmt.Println("\nFinished game:")
		fmt.Printf("- Game #%d: %s vs %s (Winner: %s) at %v\n",
			res.Game.ID, res.Game.WinnerName,
			res.Game.AnotherPlayerName, res.Game.WinnerName,
			res.Game.Time.Format("2006-01-02 15:04:05"),
		)
		c.board = res.Game.Board
		c.board.PrintBoard()
		fmt.Println()
	} else {
		log.Printf(
			"Error unmarshalling FinishedGameResponse: %v",
			err,
		)
	}
	c.setState(mainMenu) // Переходим в главное меню
}

// Обрабатываем ответ на запрос на инициализацию игры
func (c *Client) handleInitGame(payload json.RawMessage) {
	var res network.InitGameResponse // Десериализуем ответ
	if err := json.Unmarshal(payload, &res); err == nil {
		c.board = &res.Board                // Устанавливаем игровое поле
		c.currentPlayer = res.CurrentPlayer // Устанавливаем фигуру игрока
		fmt.Println("\n--- Game Started ---")
		c.board.PrintBoard() // Выводим игровое поле
		c.printTurnInfo()    // Выводим информацию о ходе игрока
		if res.CurrentPlayer == c.mySymbol {
			// Устанавливаем состояние, что ходит игрок
			c.setState(playerMove)
		} else {
			// Устанавливаем состояние, что ходит оппонент
			c.setState(opponentMove)
		}
	} else {
		log.Printf(
			"Error unmarshalling InitGameResponse: %v",
			err,
		)
	}
}

// Обрабатываем сообщение об обновлении состояния игры
func (c *Client) handleUpdateState(payload json.RawMessage) {
	var res network.GameStateUpdate // Десериализуем ответ
	if err := json.Unmarshal(payload, &res); err == nil {
		c.board = &res.Board                // Устанавливаем игровое поле
		c.currentPlayer = res.CurrentPlayer // Устанавливаем фигуру игрока
		fmt.Println("\n--- Game State Update ---")
		c.board.PrintBoard() // Выводим игровое поле
		c.printTurnInfo()    // Выводим информацию о ходе игрока
		if res.CurrentPlayer == c.mySymbol {
			// Устанавливаем состояние, что ходит игрок
			c.setState(playerMove)
		} else {
			// Устанавливаем состояние, что ходит оппонент
			c.setState(opponentMove)
		}
	} else {
		log.Printf("Error unmarshalling GameStateUpdate: %v", err)
	}
}

// Обрабатываем сообщение об окончании игры
func (c *Client) handleEndGame(payload json.RawMessage) {
	var res network.EndGameResponse // Десериализуем ответ
	if err := json.Unmarshal(payload, &res); err == nil {
		c.board = &res.Board // Устанавливаем игровое поле
		fmt.Println("\n--- Game Over ---")
		c.board.PrintBoard() // Выводим игровое поле
		// Выводим информацию о победителе
		if res.CurrentPlayer == b.Empty {
			fmt.Println("It's a Draw!")
		} else {
			fmt.Printf("Player %s wins!\n", res.CurrentPlayer)
		}
		c.setState(endGame) /// Переходи состояние завершения игры
		fmt.Print("> ")
	} else {
		log.Printf("Error unmarshalling EndGameResponse: %v", err)
	}
}

// Обрабатываем сообщение об ошибке
func (c *Client) handleError(payload json.RawMessage) {
	var errPayload network.ErrorResponse // Десериализуем ответ
	if err := json.Unmarshal(payload, &errPayload); err == nil {
		fmt.Printf("\nServer Error: %s\n> ", errPayload.Message)
	} else {
		log.Printf("Error unmarshalling ErrorResponse: %v", err)
	}
	c.setState(mainMenu) // Переходим в главное меню
}

// Обрабатываем сообщение об отключении оппонента
func (c *Client) handleOpponentLeft(payload json.RawMessage) {
	var res network.OpponentLeft // Десериализуем ответ
	if err := json.Unmarshal(payload, &res); err == nil {
		fmt.Printf(
			"\nPlayer '%s' has left the game.\n> ",
			res.Nickname,
		)
	} else {
		log.Printf("Error unmarshalling OpponentLeft: %v", err)
	}
	// Переходим в состояние ожидания присоединения оппонента к комнате
	c.setState(waitingOpponentInRoom)
}
