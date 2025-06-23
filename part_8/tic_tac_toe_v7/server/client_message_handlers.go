package server

import (
	"encoding/json"
	"log"
	"net"
	"strconv"
	"tic-tac-toe/network"
	p "tic-tac-toe/player"
)

// счетчик игроков на случай, когда игрок ввел никнейм
// который уже занят
var defaultPlayerCounts int = 0

// Обрабатываем входящие команды от клиента.
// В зависимости от типа команды вызываем соответствующий обработчик:
// регистрации никнейма, запроса на ход, списка комнат,
// входа/выхода из комнаты, работы с завершёнными играми.
// Если команда неизвестна, пишем об этом в лог.
func (s *Server) handleCommand(client net.Conn, msg *network.Message) {
	log.Printf(
		"Received command '%s' from %s",
		msg.Cmd, client.RemoteAddr(),
	)

	switch msg.Cmd {
	// Обрабатываем команду регистрации никнейма
	case network.CmdNickname:
		s.nickNameHandler(client, msg)
	// Обрабатываем команду хода игрока
	case network.CmdMakeMoveRequest:
		s.makeMoveHandler(client, msg)
	// Обрабатываем команду запроса списка комнат
	case network.CmdListRoomsRequest:
		s.listRoomsHandler(client, msg)
	// Обрабатываем команду присоединения к комнате
	case network.CmdJoinRoomRequest:
		s.joinRoomHandler(client, msg)
	// Обрабатываем команду выхода из комнаты
	case network.CmdLeaveRoomRequest:
		s.leaveRoomHandler(client, msg)
	// Обрабатываем команду запроса списка завершенных игр
	case network.CmdFinishedGamesRequest:
		s.getFinishedGamesHandler(client, msg)
	// Обрабатываем команду запроса завершенной игры по ID
	case network.CmdFinishedGameByIdRequest:
		s.getFinishedGameByIdHandler(client, msg)
	default:
		log.Printf("Unknown command: %s", msg.Cmd)
	}
}

// Обрабатываем запрос на регистрацию никнейма игрока.
// Проверяем уникальность никнейма, при необходимости
// добавляем суффикс, создаем нового игрока и отправляем
// клиенту подтверждение с итоговым никнеймом.
func (s *Server) nickNameHandler(client net.Conn, msg *network.Message) {
	// Десериализуем сообщение
	nicknameRequest := &network.NicknameRequest{}
	if err := json.Unmarshal(msg.Payload, nicknameRequest); err != nil {
		log.Printf("Error unmarshaling NicknameRequest: %v", err)
		return
	}
	s.mutex.Lock() // Защищаем доступ к данным
	if s.players[nicknameRequest.Nickname] != nil {
		// Если никнейм занят, добавляем суффикс
		nicknameRequest.Nickname = nicknameRequest.Nickname +
			"_" + strconv.Itoa(defaultPlayerCounts)
		defaultPlayerCounts++
	}
	// Создаем игрока и добавляем его в карту игроков
	s.players[nicknameRequest.Nickname] = p.NewHumanPlayer(
		nicknameRequest.Nickname, &client,
	)
	s.mutex.Unlock()
	// Формируем ответ клиенту
	response := &network.NickNameResponse{
		Nickname: nicknameRequest.Nickname,
	}
	msg.Payload, _ = json.Marshal(response)
	msg.Cmd = network.CmdNickNameResponse
	json.NewEncoder(client).Encode(msg) // Отправляем ответ клиенту
}

// Обрабатываем запрос на присоединение к комнате.
// Проверяем существование комнаты и игрока,
// добавляем игрока в комнату (если есть место),
// отправляем клиенту подтверждение и инициируем старт игры
func (s *Server) joinRoomHandler(client net.Conn, msg *network.Message) {
	// Десериализуем сообщение
	joinRoomRequest := &network.JoinRoomRequest{}
	if err := json.Unmarshal(msg.Payload, joinRoomRequest); err != nil {
		log.Printf("Error unmarshaling JoinRoomRequest: %v", err)
		return
	}
	// Получаем комнату и игрока
	s.mutex.RLock() // Защищаем доступ к данным
	room, okRoom := s.rooms[joinRoomRequest.RoomName]
	player, okPlayer := s.players[joinRoomRequest.PlayerName]
	s.mutex.RUnlock()
	// Проверяем существование комнаты и игрока
	if !okRoom || !okPlayer {
		// Если комнаты или игрока не существует
		response := &network.ErrorResponse{Message: "Room not found"}
		msg.Cmd = network.CmdError
		msg.Payload, _ = json.Marshal(response)
		json.NewEncoder(client).Encode(msg)
		return
	}
	// Защищаем доступ к данным
	s.mutex.Lock()
	if room.IsFull() { // Если комната заполнена
		s.mutex.Unlock()
		// Отправляем ответ клиенту
		response := &network.ErrorResponse{Message: "Room is full"}
		msg.Cmd = network.CmdError
		msg.Payload, _ = json.Marshal(response)
		json.NewEncoder(client).Encode(msg)
		return
	}
	room.AddPlayer(player) // Добавляем игрока в комнату
	s.mutex.Unlock()
	// Формируем ответ клиенту
	response := &network.RoomJoinResponse{
		RoomName:     joinRoomRequest.RoomName,
		PlayerSymbol: player.GetFigure(),
		Board:        *room.Board,
	}
	msg.Payload, _ = json.Marshal(response)
	msg.Cmd = network.CmdRoomJoinResponse
	json.NewEncoder(client).Encode(msg) // Отправляем ответ клиенту
	room.InitGame()                     // Инициируем старт игры
}

// Обрабатываем выход игрока из комнаты.
// Проверяем существование комнаты и игрока, удаляем игрока из комнаты.
func (s *Server) leaveRoomHandler(
	client net.Conn, msg *network.Message,
) {
	// Десериализуем сообщение
	leaveRoomRequest := &network.LeaveRoomRequest{}
	if err := json.Unmarshal(msg.Payload, leaveRoomRequest); err != nil {
		log.Printf("Error unmarshaling LeaveRoomRequest: %v", err)
		return
	}
	// Получаем комнату и игрока
	s.mutex.RLock() // Защищаем доступ к данным
	room, okRoom := s.rooms[leaveRoomRequest.RoomName]
	player, okPlayer := s.players[leaveRoomRequest.PlayerName]
	s.mutex.RUnlock()
	// Проверяем существование комнаты и игрока
	if !okRoom || !okPlayer {
		// Если комнаты или игрока не существует
		response := &network.ErrorResponse{Message: "Room not found"}
		msg.Cmd = network.CmdError
		msg.Payload, _ = json.Marshal(response)
		json.NewEncoder(client).Encode(msg)
		return
	}
	// Защищаем доступ к данным
	s.mutex.Lock()
	room.RemovePlayer(player) // Удаляем игрока из комнаты
	s.mutex.Unlock()
}

// Обрабатываем запрос на получение списка всех комнат
func (s *Server) listRoomsHandler(
	client net.Conn, msg *network.Message,
) {
	// Защищаем доступ к данным
	s.mutex.Lock()
	defer s.mutex.Unlock()
	// Создаем список комнат
	var roomInfos []network.RoomInfo
	for _, room := range s.rooms {
		roomInfos = append(roomInfos, network.RoomInfo{
			Name:      room.Name,
			BoardSize: room.BoardSize(),
			IsFull:    room.IsFull(),
			GameMode:  room.Mode,
			Difficult: room.Difficulty,
		})
	}
	// Формируем ответ клиенту
	response := &network.RoomListResponse{
		Rooms: roomInfos,
	}
	msg.Cmd = network.CmdRoomListResponse
	msg.Payload, _ = json.Marshal(response)
	json.NewEncoder(client).Encode(msg) // Отправляем сообщение
}

// Обрабатываем запрос на получение списка завершенных игр
func (s *Server) getFinishedGamesHandler(
	client net.Conn, msg *network.Message,
) {
	// Получаем данные из БД
	finishedGames, err := s.repository.GetAllFinishedGames()
	if err != nil {
		// Если ошибка, отправляем сообщение об ошибке
		response := &network.ErrorResponse{
			Message: "Error getting finished games",
		}
		msg.Cmd = network.CmdError
		msg.Payload, _ = json.Marshal(response)
		json.NewEncoder(client).Encode(msg)
		return
	}
	// Формируем ответ клиенту
	response := &network.FinishedGamesResponse{
		Games: finishedGames,
	}
	msg.Cmd = network.CmdFinishedGamesResponse
	msg.Payload, _ = json.Marshal(response)
	json.NewEncoder(client).Encode(msg)
}

// Обрабатываем запрос на получение завершенной игры по ID
func (s *Server) getFinishedGameByIdHandler(
	client net.Conn, msg *network.Message,
) {
	// Десериализуем сообщение
	getFinishedGameByIdRequest := &network.GetFinishedGameByIdRequest{}
	if err := json.Unmarshal(
		msg.Payload, getFinishedGameByIdRequest,
	); err != nil {
		log.Printf(
			"Error unmarshaling GetFinishedGameByIdRequest: %v",
			err,
		)
		return
	}
	// Получаем завершенную игру по ID
	finishedGame, err := s.repository.GetFinishedGameById(
		getFinishedGameByIdRequest.GameID,
	)
	if err != nil {
		// Если ошибка, отправляем сообщение об ошибке
		response := &network.ErrorResponse{
			Message: "Error getting finished game by id",
		}
		msg.Cmd = network.CmdError
		msg.Payload, _ = json.Marshal(response)
		json.NewEncoder(client).Encode(msg)
		return
	}
	// Формируем ответ клиенту
	response := &network.FinishedGameResponse{
		Game: finishedGame,
	}
	msg.Cmd = network.CmdFinishedGameResponse
	msg.Payload, _ = json.Marshal(response)
	json.NewEncoder(client).Encode(msg)
}

// Обрабатываем запрос на ход игрока
func (s *Server) makeMoveHandler(
	client net.Conn, msg *network.Message,
) {
	// Десериализуем сообщение
	makeMoveRequest := &network.MakeMoveRequest{}
	if err := json.Unmarshal(msg.Payload, makeMoveRequest); err != nil {
		log.Printf("Error unmarshaling MakeMoveRequest: %v", err)
		return
	}
	// Получаем комнату и игрока
	s.mutex.RLock() // Защищаем доступ к данным
	room, okRoom := s.rooms[makeMoveRequest.RoomName]
	player, okPlayer := s.players[makeMoveRequest.PlayerName]
	s.mutex.RUnlock()
	if !okRoom || !okPlayer { // Если комнаты или игрока не существует
		// Отправляем сообщение об ошибке
		response := &network.ErrorResponse{
			Message: "Room not found",
		}
		msg.Cmd = network.CmdError
		msg.Payload, _ = json.Marshal(response)
		json.NewEncoder(client).Encode(msg)
		return
	}
	// Выполняем ход игрока
	room.PlayerStep(
		player,
		makeMoveRequest.PositionRow,
		makeMoveRequest.PositionCol,
	)
}
