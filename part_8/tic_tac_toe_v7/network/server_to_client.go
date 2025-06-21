package network

import (
	b "tic-tac-toe/board"
	g "tic-tac-toe/game"
	"tic-tac-toe/model"
)

const (
	// Server to Client Commands
	CmdUpdateState           Command = "update_state"
	CmdError                 Command = "error"
	CmdNickNameResponse      Command = "nick_name_response"
	CmdRoomCreated           Command = "room_created"
	CmdRoomJoinResponse      Command = "room_join_response"
	CmdRoomListResponse      Command = "room_list_response"
	CmdInitGame              Command = "init_game"
	CmdOpponentLeft          Command = "opponent_left"
	CmdEndGame               Command = "end_game"
	CmdFinishedGamesResponse Command = "finished_games_response"
	CmdFinishedGameResponse  Command = "finished_game_response"
)

// сообщение о том, что противник покинул игру
// инициализирующее сообщение в начале партии

// InitGameResponse отправляется сервером при инициализации игры.
type InitGameResponse struct {
	Board         b.Board      `json:"board"`
	CurrentPlayer b.BoardField `json:"current_player"`
}

// EndGameResponse отправляется сервером при завершении игры.
type EndGameResponse struct {
	Board         b.Board      `json:"board"`
	CurrentPlayer b.BoardField `json:"current_player"`
}

type OpponentLeft struct {
	Nickname string `json:"nickname"`
}

// RoomInfo содержит информацию о комнате.
type RoomInfo struct {
	Name      string       `json:"name"`
	BoardSize int          `json:"board_size"`
	IsFull    bool         `json:"is_full"`
	GameMode  g.GameMode   `json:"game_mode"`
	Difficult g.Difficulty `json:"difficult"`
}

// RoomListсодержит список доступных комнат.
type RoomListResponse struct {
	Rooms []RoomInfo `json:"rooms"`
}

// GameStateUpdate содержит информацию об обновлении состояния игры.
type GameStateUpdate struct {
	Board         b.Board      `json:"board"`
	CurrentPlayer b.BoardField `json:"current_player"`
}

// ErrorResponse отправляется сервером при возникновении ошибки.
type ErrorResponse struct {
	Message string `json:"message"`
}

// NickNameResponse отправляется сервером при успешном входе клиента.
type NickNameResponse struct {
	Nickname string `json:"nickname"`
}

// RoomCreatedResponse отправляется сервером после успешного создания комнаты.
type RoomCreatedResponse struct {
	RoomID   string `json:"room_id"`
	RoomName string `json:"room_name"`
}

// RoomJoinResponse отправляется сервером, когда клиент успешно присоединился к комнате.
type RoomJoinResponse struct {
	RoomName     string       `json:"room_name"`
	PlayerSymbol b.BoardField `json:"player_symbol"`
	Board        b.Board      `json:"board"`
}

// FinishedGamesResponse отправляется сервером со списком завершенных игр.
type FinishedGamesResponse struct {
	Games *[]model.FinishGameSnapshot `json:"games"`
}

// FinishedGameResponse отправляется сервером с информацией о конкретной завершенной игре.
type FinishedGameResponse struct {
	Game *model.FinishGameSnapshot `json:"game"`
}
