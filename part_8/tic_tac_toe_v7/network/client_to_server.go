package network

const (
	// Client to Server Commands
	CmdNickname                Command = "nickname"
	CmdJoinRoomRequest         Command = "join_room"
	CmdLeaveRoomRequest        Command = "leave_room"
	CmdListRoomsRequest        Command = "list_rooms"
	CmdMakeMoveRequest         Command = "make_move"
	CmdFinishedGamesRequest    Command = "get_finished_games"
	CmdFinishedGameByIdRequest Command = "get_finished_game_by_id"
)

// LoginRequest отправляется клиентом для входа в систему.
type NicknameRequest struct {
	Nickname string `json:"nickname"`
}

// JoinRoomRequest отправляется клиентом для подключения к существующей комнате.
type JoinRoomRequest struct {
	RoomName   string `json:"room_name"`
	PlayerName string `json:"player_name"`
}

// LeaveRoomRequest отправляется клиентом для выхода из текущей комнаты.
type LeaveRoomRequest struct {
	RoomName   string `json:"room_name"`
	PlayerName string `json:"player_name"`
}

// ListRoomsRequest отправляется клиентом для получения списка доступных комнат.
// Обычно для этого запроса не требуется специальная полезная нагрузка.
type ListRoomsRequest struct {
}

// MakeMoveRequest отправляется клиентом для совершения хода в игре.
type MakeMoveRequest struct {
	RoomName    string `json:"room_name"`
	PlayerName  string `json:"player_name"`
	PositionRow int    `json:"position_row"`
	PositionCol int    `json:"position_col"`
}

// GetFinishedGamesRequest отправляется клиентом для получения списка завершенных игр.
// Обычно для этого запроса не требуется специальная полезная нагрузка, если запрашиваются все игры для пользователя.
type GetFinishedGamesRequest struct {
}

// GetFinishedGameByIdRequest отправляется клиентом для получения конкретной завершенной игры.
type GetFinishedGameByIdRequest struct {
	GameID int `json:"game_id"`
}
