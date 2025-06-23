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

// Room — структура, которая описывает игровую комнату
type Room struct {
	Name    string     // Название комнаты
	Board   *b.Board   // Ссылка на игровую доску
	Player1 p.IPlayer  // Первый игрок (только человек)
	Mode    g.GameMode // Режим игры: PvP или PvC

	// Второй игрок (может быть человеком или компьютером)
	Player2 p.IPlayer
	// Текущий игрок, который должен сделать ход
	CurrentPlayer p.IPlayer
	// Текущее состояние игры (чей ход, победа, ничья и т.д.)
	State g.GameState
	// Уровень сложности компьютера (используется только в режиме PvC)
	Difficulty g.Difficulty
	// Интерфейс для сохранения завершенных игр в базе данных
	repository db.IRepository
}

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
	// Если режим игры — PvC, то создаем компьютерного игрока
	if gameMode == g.PvC {
		room.Player2 = p.NewComputerPlayer(b.Nought, difficulty)
	}
	return room
}

// Возвращает true, если в комнате есть два игрока
func (r *Room) IsFull() bool {
	return r.Player1 != nil && r.Player2 != nil // Если оба игрока не равны nil, значит комната полная
}

// Возвращает количество игроков в комнате
func (r *Room) PlayersAmount() int {
	if r.Player1 != nil && r.Player2 != nil {
		return 2 // Если оба игрока есть, возвращаем 2
	} else if r.Player1 != nil || r.Player2 != nil {
		return 1 // Если только один игрок, возвращаем 1
	}
	return 0 // Если ни один игрок не добавлен, возвращаем 0
}

// Возвращает размер доски
func (r *Room) BoardSize() int {
	return r.Board.Size
}

// Добавляем игрока в комнату
// Первый добавленный игрок становится Player1, второй — Player2
// Символы игроков автоматически корректируются: 1 — X, 2 — O
func (r *Room) AddPlayer(player p.IPlayer) {
	if r.Player1 == nil {
		r.Player1 = player // Первый игрок
		if r.Player1.GetSymbol() != "X" {
			r.Player1.SwitchPlayer() // Если символ не X, меняем на X
		}
	} else if r.Player2 == nil {
		r.Player2 = player // Второй игрок
		if r.Player2.GetSymbol() != "O" {
			r.Player2.SwitchPlayer() // Если символ не O, меняем на O
		}
	}
}

// Удаляем игрока из комнаты
// В случае выхода игрока его оппоненту (если присутствует в
// комнате и человек) отправляется сообщение о выходе соперника
func (r *Room) RemovePlayer(player p.IPlayer) {
	if r.Player1 == player {
		r.Player1 = nil // Удаляем первого игрока
		// Если в комнате есть второй игрок и он человек,
		// уведомляем его о выходе соперника
		if r.Player2 != nil && !r.Player2.IsComputer() {
			opponentLeft := &n.OpponentLeft{
				// Имя вышедшего игрока
				Nickname: player.GetNickname(),
			}
			payloadBytes, err := json.Marshal(opponentLeft)
			if err != nil {
				log.Printf("Error marshaling OpponentLeft: %v", err)
				return
			}
			msg := &n.Message{
				Cmd:     n.CmdOpponentLeft, // Команда "соперник вышел"
				Payload: payloadBytes,
			}
			// Отправляем сообщение второму игроку
			r.Player2.SendMessage(msg)
		}
	} else if r.Player2 == player {
		r.Player2 = nil // Удаляем второго игрока
		// Если в комнате есть первый игрок,
		// уведомляем его о выходе соперника
		if r.Player1 != nil {
			opponentLeft := &n.OpponentLeft{
				Nickname: player.GetNickname(),
			}
			payloadBytes, err := json.Marshal(opponentLeft)
			if err != nil {
				log.Printf("Error marshaling OpponentLeft: %v", err)
				return
			}
			msg := &n.Message{
				Cmd:     n.CmdOpponentLeft,
				Payload: payloadBytes,
			}
			// Отправляем сообщение первому игроку
			r.Player1.SendMessage(msg)
		}
	}
}

// Инициализируем новую игру в комнате.
// Здесь выбирается случайный игрок, который начинает первым,
// и отправляется сообщение обоим игрокам о начале игры.
// Если первый ход за компьютером, то он делает ход автоматически
func (r *Room) InitGame() {
	if !r.IsFull() {
		return // Если игроков меньше двух, игра не начинается
	}

	// Срез для определения игрока, ходящего первым
	randomPlayer := []b.BoardField{b.Cross, b.Nought}
	if !r.Board.IsEmpty() {
		// Если доска не пуста, создаем новую
		r.Board = b.NewBoard(r.Board.Size)
	}

	msg := &n.Message{Cmd: n.CmdInitGame} // Сообщение о начале игры
	initGamePayload := &n.InitGameResponse{
		Board: *r.Board, // Текущее состояние доски
	}
	// Выбираем случайным образом, кто ходит первым (X или O)
	starterSymbol := randomPlayer[rand.Intn(len(randomPlayer))]
	switch starterSymbol {
	case b.Cross:
		r.State = g.CrossStep
		initGamePayload.CurrentPlayer = b.Cross
	case b.Nought:
		r.State = g.NoughtStep
		initGamePayload.CurrentPlayer = b.Nought
	}

	// Устанавливаем активного игрока в зависимости
	// от режима игры и символа
	if r.Mode == g.PvC {
		// В режиме PvC человек всегда Player1
		if r.State == g.CrossStep {
			r.CurrentPlayer = r.Player1
		} else if r.State == g.NoughtStep {
			r.CurrentPlayer = r.Player2
		}
	} else {
		// В режиме PvP ищем, кто играет выбранным символом
		if (r.State == g.CrossStep &&
			r.Player1.GetFigure() == b.Cross) ||
			(r.State == g.NoughtStep &&
				r.Player1.GetFigure() == b.Nought) {
			r.CurrentPlayer = r.Player1
		} else {
			r.CurrentPlayer = r.Player2
		}
	}

	// Сериализуем данные для отправки игрокам
	payloadBytes, err := json.Marshal(initGamePayload)
	if err != nil {
		log.Printf(
			"Error marshaling InitGameResponse for Player1 "+
				"after Player2 left: %v", err)
		return
	}
	msg.Payload = payloadBytes
	r.Player1.SendMessage(msg) // Отправляем сообщение первому игроку
	r.Player2.SendMessage(msg) // Отправляем сообщение второму игроку

	// Если сейчас ход компьютера, он делает ход автоматически
	if r.CurrentPlayer.IsComputer() {
		row, col, _ := r.CurrentPlayer.MakeMove(r.Board)
		r.PlayerStep(r.CurrentPlayer, row, col)
	}
}

// Переключаем активного игрока
func (r *Room) switchCurrentPlayer() {
	if r.CurrentPlayer == r.Player1 {
		r.CurrentPlayer = r.Player2 // Если сейчас ходил первый, теперь ход второго
	} else {
		r.CurrentPlayer = r.Player1 // И наоборот
	}
}

// PlayerStep выполняет ход игрока и обновляет состояние игры
// Здесь проверяется правильность хода, обновляется доска,
// определяется победитель или ничья, и отправляются сообщения игрокам.
// Если после хода игра не закончена, ход переходит следующему игроку.
// Если ходит компьютер — он делает ход автоматически.
func (r *Room) PlayerStep(player p.IPlayer, row, col int) {
	msg := &n.Message{} // Создаем новое сообщение для игроков
	// Проверяем, что сейчас идет ход (игра не завершена)
	if r.State != g.CrossStep && r.State != g.NoughtStep {
		return // Если сейчас не ход, ничего не делаем
	}
	// Проверяем, что ход делает именно тот игрок, чей сейчас ход
	if player != r.CurrentPlayer {
		return
	}

	// Ставим символ игрока на выбранную клетку
	r.Board.SetSymbol(row, col, r.CurrentPlayer.GetFigure())
	// Проверяем, выиграл ли этот игрок
	if r.Board.CheckWin(r.CurrentPlayer.GetFigure()) {
		if r.CurrentPlayer.GetFigure() == b.Cross {
			r.State = g.CrossWin // Победа крестиков
		} else {
			r.State = g.NoughtWin // Победа ноликов
		}
		// Формируем сообщение о завершении игры
		msg.Cmd = n.CmdEndGame
		endGamePayload := &n.EndGameResponse{
			Board:         *r.Board,
			CurrentPlayer: r.CurrentPlayer.GetFigure(),
		}
		msg.Payload, _ = json.Marshal(endGamePayload)

		// Сохраняем информацию о завершенной игре в базе данных
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
		// Если доска заполнена, но победителя нет — ничья
		r.State = g.Draw
		msg.Cmd = n.CmdEndGame
		endGamePayload := &n.EndGameResponse{
			Board:         *r.Board,
			CurrentPlayer: b.Empty,
		}
		msg.Payload, _ = json.Marshal(endGamePayload)
	} else {
		// Если игра не закончена, меняем активного игрока и продолжаем
		if r.CurrentPlayer.GetFigure() == b.Cross {
			r.State = g.NoughtStep
		} else {
			r.State = g.CrossStep
		}
		r.switchCurrentPlayer()
		msg.Cmd = n.CmdUpdateState // Сообщаем о новом состоянии игры
		stateUpdatePayload := &n.GameStateUpdate{
			Board:         *r.Board,
			CurrentPlayer: r.CurrentPlayer.GetFigure(),
		}
		msg.Payload, _ = json.Marshal(stateUpdatePayload)
	}

	// Отправляем сообщение обоим игрокам
	r.Player1.SendMessage(msg)
	r.Player2.SendMessage(msg)

	// Если игра завершена (победа или ничья),
	// ждем 10 секунд и запускаем новую
	if r.State == g.CrossWin || r.State == g.NoughtWin ||
		r.State == g.Draw {
		time.Sleep(10 * time.Second)
		r.InitGame()
		return
	}

	// Если теперь ход компьютера, он делает ход автоматически
	if r.CurrentPlayer.IsComputer() {
		row, col, _ := r.CurrentPlayer.MakeMove(r.Board)
		r.PlayerStep(r.CurrentPlayer, row, col)
	}
}
