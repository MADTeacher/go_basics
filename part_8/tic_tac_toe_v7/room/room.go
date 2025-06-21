package room

import (
	"encoding/json"
	"log"
	"math/rand"
	b "tic-tac-toe/board"
	db "tic-tac-toe/database"
	g "tic-tac-toe/game"
	"tic-tac-toe/model"
	n "tic-tac-toe/network"
	p "tic-tac-toe/player"
	"time"
)

// Room manages the state of a single game room.
type Room struct {
	Name          string
	Board         *b.Board
	Player1       p.IPlayer
	Player2       p.IPlayer
	CurrentPlayer p.IPlayer
	State         g.GameState
	repository    db.IRepository
	Mode          g.GameMode
	// Уровень сложности компьютера (только для PvC)
	Difficulty g.Difficulty
}

// NewRoom creates a new game room.
func NewRoom(
	name string, repository db.IRepository, boardSize int,
	gameMode g.GameMode, difficulty g.Difficulty,
) *Room {
	room := &Room{
		Name:       name,
		repository: repository,
		Mode:       gameMode,
		Difficulty: difficulty,
		Board:      b.NewBoard(boardSize),
		State:      g.WaitingOpponent,
	}
	if gameMode == g.PvC {
		room.Player2 = p.NewComputerPlayer(b.Nought, difficulty)
	}
	return room
}

func (r *Room) IsFull() bool {
	return r.Player1 != nil && r.Player2 != nil
}

func (r *Room) PlayersAmount() int {
	if r.Player1 != nil && r.Player2 != nil {
		return 2
	}
	return 1
}

func (r *Room) BoardSize() int {
	return r.Board.Size
}

func (r *Room) AddPlayer(player p.IPlayer) {
	if r.Player1 == nil {
		r.Player1 = player
		if r.Player1.GetSymbol() != "X" {
			r.Player1.SwitchPlayer()
		}
	} else if r.Player2 == nil {
		r.Player2 = player
		if r.Player2.GetSymbol() != "O" {
			r.Player2.SwitchPlayer()
		}
	}
}

func (r *Room) RemovePlayer(player p.IPlayer) {
	if r.Player1 == player {
		r.Player1 = nil
		if r.Player2 != nil && !r.Player2.IsComputer() {
			opponentLeft := &n.OpponentLeft{Nickname: player.GetNickname()}
			payloadBytes, err := json.Marshal(opponentLeft)
			if err != nil {
				log.Printf("Error marshaling OpponentLeft: %v", err)
				return
			}
			msg := &n.Message{
				Cmd:     n.CmdOpponentLeft,
				Payload: payloadBytes,
			}
			r.Player2.SendMessage(msg)
		}
	} else if r.Player2 == player {
		r.Player2 = nil
		if r.Player1 != nil && !r.Player1.IsComputer() {
			opponentLeft := &n.OpponentLeft{Nickname: player.GetNickname()}
			payloadBytes, err := json.Marshal(opponentLeft)
			if err != nil {
				log.Printf("Error marshaling OpponentLeft: %v", err)
				return
			}
			msg := &n.Message{
				Cmd:     n.CmdOpponentLeft,
				Payload: payloadBytes,
			}
			r.Player1.SendMessage(msg)
		}
	}
}

func (r *Room) InitGame() {
	if !r.IsFull() {
		return
	}

	randomPlayer := []b.BoardField{b.Cross, b.Nought}
	if !r.Board.IsEmpty() {
		r.Board = b.NewBoard(r.Board.Size)
	}

	msg := &n.Message{Cmd: n.CmdInitGame}
	initGamePayload := &n.InitGameResponse{
		Board: *r.Board,
	}
	// Select a random starting symbol
	starterSymbol := randomPlayer[rand.Intn(len(randomPlayer))]
	switch starterSymbol {
	case b.Cross:
		r.State = g.CrossStep
		initGamePayload.CurrentPlayer = b.Cross
	case b.Nought:
		r.State = g.NoughtStep
		initGamePayload.CurrentPlayer = b.Nought
	}

	// Set the current player based on game mode and starter symbol
	if r.Mode == g.PvC {
		// In PvC mode, Player1 is always the human player
		if r.State == g.CrossStep {
			r.CurrentPlayer = r.Player1
		} else if r.State == g.NoughtStep {
			r.CurrentPlayer = r.Player2
		}
	} else {
		// In PvP mode, set the current player based on who has the starter symbol
		if (r.State == g.CrossStep && r.Player1.GetFigure() == b.Cross) ||
			(r.State == g.NoughtStep && r.Player1.GetFigure() == b.Nought) {
			r.CurrentPlayer = r.Player1
		} else {
			r.CurrentPlayer = r.Player2
		}
	}

	payloadBytes, err := json.Marshal(initGamePayload)
	if err != nil {
		log.Printf("Error marshaling InitGameResponse for Player1 after Player2 left: %v", err)
		return
	}
	msg.Payload = payloadBytes
	r.Player1.SendMessage(msg)
	r.Player2.SendMessage(msg)

	if r.CurrentPlayer.IsComputer() {
		row, col, _ := r.CurrentPlayer.MakeMove(r.Board)
		r.PlayerStep(r.CurrentPlayer, row, col)
	}
}

// Переключаем активного игрока
func (r *Room) switchCurrentPlayer() {
	if r.CurrentPlayer == r.Player1 {
		r.CurrentPlayer = r.Player2
	} else {
		r.CurrentPlayer = r.Player1
	}
}

func (r *Room) PlayerStep(player p.IPlayer, row, col int) {
	msg := &n.Message{}
	if r.State != g.CrossStep && r.State != g.NoughtStep {
		return
	}
	// проверяем, что ход делает текущий игрок
	if player != r.CurrentPlayer {
		return
	}

	r.Board.SetSymbol(row, col, r.CurrentPlayer.GetFigure())
	if r.Board.CheckWin(r.CurrentPlayer.GetFigure()) {
		if r.CurrentPlayer.GetFigure() == b.Cross {
			r.State = g.CrossWin
		} else {
			r.State = g.NoughtWin
		}
		msg.Cmd = n.CmdEndGame
		endGamePayload := &n.EndGameResponse{
			Board:         *r.Board,
			CurrentPlayer: r.CurrentPlayer.GetFigure(),
		}
		msg.Payload, _ = json.Marshal(endGamePayload)

		figureWinner := r.CurrentPlayer.GetFigure()
		winnerNickName := r.CurrentPlayer.GetNickname()

		var anotherPlayerNickName string
		if r.CurrentPlayer == r.Player1 {
			anotherPlayerNickName = r.Player2.GetNickname()
		} else {
			anotherPlayerNickName = r.Player1.GetNickname()
		}
		r.repository.SaveFinishedGame(&model.FinishGameSnapshot{
			Board:             r.Board,
			PlayerFigure:      figureWinner,
			WinnerName:        winnerNickName,
			AnotherPlayerName: anotherPlayerNickName,
			Time:              time.Now(),
		})
	} else if r.Board.CheckDraw() {
		r.State = g.Draw
		msg.Cmd = n.CmdEndGame
		endGamePayload := &n.EndGameResponse{
			Board:         *r.Board,
			CurrentPlayer: b.Empty,
		}
		msg.Payload, _ = json.Marshal(endGamePayload)
	} else {
		if r.CurrentPlayer.GetFigure() == b.Cross {
			r.State = g.NoughtStep
		} else {
			r.State = g.CrossStep
		}
		r.switchCurrentPlayer()
		msg.Cmd = n.CmdUpdateState
		stateUpdatePayload := &n.GameStateUpdate{
			Board:         *r.Board,
			CurrentPlayer: r.CurrentPlayer.GetFigure(),
		}
		msg.Payload, _ = json.Marshal(stateUpdatePayload)
	}

	r.Player1.SendMessage(msg)
	r.Player2.SendMessage(msg)

	if r.State == g.CrossWin || r.State == g.NoughtWin || r.State == g.Draw {
		time.Sleep(10 * time.Second)
		r.InitGame()
		return
	}

	if r.CurrentPlayer.IsComputer() {
		row, col, _ := r.CurrentPlayer.MakeMove(r.Board)
		r.PlayerStep(r.CurrentPlayer, row, col)
	}
}
