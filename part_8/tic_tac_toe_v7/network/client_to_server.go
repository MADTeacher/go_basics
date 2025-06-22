package network

// Сообщения от клиента к серверу

const (
	CmdNickname                Command = "nickname"
	CmdJoinRoomRequest         Command = "join_room"
	CmdLeaveRoomRequest        Command = "leave_room"
	CmdListRoomsRequest        Command = "list_rooms"
	CmdMakeMoveRequest         Command = "make_move"
	CmdFinishedGamesRequest    Command = "get_finished_games"
	CmdFinishedGameByIdRequest Command = "get_finished_game_by_id"
)

// Запрос от клиента для входа в систему
type NicknameRequest struct {
	Nickname string `json:"nickname"`
}

// Запрос от клиента для подключения к существующей комнате
type JoinRoomRequest struct {
	RoomName   string `json:"room_name"`
	PlayerName string `json:"player_name"`
}

// Запрос от клиента для выхода из текущей комнаты
type LeaveRoomRequest struct {
	RoomName   string `json:"room_name"`
	PlayerName string `json:"player_name"`
}

// Запрос от клиента для получения списка доступных комнат
// Обычно для этого запроса не требуется специальная полезная нагрузка
type ListRoomsRequest struct {
}

// Сообщение с данными о ходе игрока
type MakeMoveRequest struct {
	RoomName    string `json:"room_name"`
	PlayerName  string `json:"player_name"`
	PositionRow int    `json:"position_row"`
	PositionCol int    `json:"position_col"`
}

// Запрос от клиента для получения списка завершенных игр
type GetFinishedGamesRequest struct {
}

// Запрос от клиента для получения конкретной завершенной игры
type GetFinishedGameByIdRequest struct {
	GameID int `json:"game_id"`
}
