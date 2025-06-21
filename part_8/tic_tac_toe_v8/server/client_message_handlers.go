package server

import (
	"encoding/json"
	"log"
	"net"
	"strconv"
	"tic-tac-toe/network"
	p "tic-tac-toe/player"
)

var defaultPlayerCounts int = 0

func (s *Server) handleCommand(client net.Conn, msg *network.Message) {
	log.Printf(
		"Received command '%s' from %s",
		msg.Cmd, client.RemoteAddr(),
	)

	switch msg.Cmd {
	case network.CmdNickname:
		s.nickNameHandler(client, msg)
	case network.CmdMakeMoveRequest:
		s.makeMoveHandler(client, msg)
	case network.CmdListRoomsRequest:
		s.listRoomsHandler(client, msg)
	case network.CmdJoinRoomRequest:
		s.joinRoomHandler(client, msg)
	case network.CmdLeaveRoomRequest:
		s.leaveRoomHandler(client, msg)
	case network.CmdFinishedGamesResponse:
		s.getFinishedGamesHandler(client, msg)
	case network.CmdFinishedGameByIdRequest:
		s.getFinishedGameByIdHandler(client, msg)
	default:
		log.Printf("Unknown command: %s", msg.Cmd)
	}
}

func (s *Server) nickNameHandler(client net.Conn, msg *network.Message) {
	nicknameRequest := &network.NicknameRequest{}
	if err := json.Unmarshal(msg.Payload, nicknameRequest); err != nil {
		log.Printf("Error unmarshaling NicknameRequest: %v", err)
		return
	}
	if s.players[nicknameRequest.Nickname] != nil {
		nicknameRequest.Nickname = nicknameRequest.Nickname +
			"_" + strconv.Itoa(defaultPlayerCounts)
		defaultPlayerCounts++
	}
	s.players[nicknameRequest.Nickname] = p.NewHumanPlayer(
		nicknameRequest.Nickname, &client,
	)
	response := &network.NickNameResponse{
		Nickname: nicknameRequest.Nickname,
	}
	msg.Payload, _ = json.Marshal(response)
	msg.Cmd = network.CmdNickNameResponse
	json.NewEncoder(client).Encode(msg)
}

func (s *Server) joinRoomHandler(client net.Conn, msg *network.Message) {
	joinRoomRequest := &network.JoinRoomRequest{}
	if err := json.Unmarshal(msg.Payload, joinRoomRequest); err != nil {
		log.Printf("Error unmarshaling JoinRoomRequest: %v", err)
		return
	}
	room, okRoom := s.rooms[joinRoomRequest.RoomName]
	player, okPlayer := s.players[joinRoomRequest.PlayerName]
	if !okRoom || !okPlayer {
		response := &network.ErrorResponse{Message: "Room not found"}
		msg.Cmd = network.CmdError
		msg.Payload, _ = json.Marshal(response)
		json.NewEncoder(client).Encode(msg)
		return
	}
	s.mutex.Lock()
	if room.IsFull() {
		s.mutex.Unlock()
		response := &network.ErrorResponse{Message: "Room is full"}
		msg.Cmd = network.CmdError
		msg.Payload, _ = json.Marshal(response)
		json.NewEncoder(client).Encode(msg)
		return
	}
	room.AddPlayer(player)
	s.mutex.Unlock()
	response := &network.RoomJoinResponse{
		RoomName:     joinRoomRequest.RoomName,
		PlayerSymbol: player.GetFigure(),
		Board:        *room.Board,
	}
	msg.Payload, _ = json.Marshal(response)
	msg.Cmd = network.CmdRoomJoinResponse
	json.NewEncoder(client).Encode(msg)
	room.InitGame()
}

func (s *Server) leaveRoomHandler(client net.Conn, msg *network.Message) {
	leaveRoomRequest := &network.LeaveRoomRequest{}
	if err := json.Unmarshal(msg.Payload, leaveRoomRequest); err != nil {
		log.Printf("Error unmarshaling LeaveRoomRequest: %v", err)
		return
	}
	room, okRoom := s.rooms[leaveRoomRequest.RoomName]
	player, okPlayer := s.players[leaveRoomRequest.PlayerName]
	if !okRoom || !okPlayer {
		response := &network.ErrorResponse{Message: "Room not found"}
		msg.Cmd = network.CmdError
		msg.Payload, _ = json.Marshal(response)
		json.NewEncoder(client).Encode(msg)
		return
	}
	s.mutex.Lock()
	room.RemovePlayer(player)
	s.mutex.Unlock()
}

func (s *Server) listRoomsHandler(client net.Conn, msg *network.Message) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
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

	response := &network.RoomListResponse{
		Rooms: roomInfos,
	}
	msg.Cmd = network.CmdRoomListResponse
	msg.Payload, _ = json.Marshal(response)
	json.NewEncoder(client).Encode(msg)
}

func (s *Server) getFinishedGamesHandler(client net.Conn, msg *network.Message) {
	// получаем данные из БД
	finishedGames, err := s.repository.GetAllFinishedGames()
	if err != nil {
		response := &network.ErrorResponse{Message: "Error getting finished games"}
		msg.Cmd = network.CmdError
		msg.Payload, _ = json.Marshal(response)
		json.NewEncoder(client).Encode(msg)
		return
	}
	response := &network.FinishedGamesResponse{
		Games: finishedGames,
	}
	msg.Cmd = network.CmdFinishedGamesResponse
	msg.Payload, _ = json.Marshal(response)
	json.NewEncoder(client).Encode(msg)
}

func (s *Server) getFinishedGameByIdHandler(client net.Conn, msg *network.Message) {
	getFinishedGameByIdRequest := &network.GetFinishedGameByIdRequest{}
	if err := json.Unmarshal(msg.Payload, getFinishedGameByIdRequest); err != nil {
		log.Printf("Error unmarshaling GetFinishedGameByIdRequest: %v", err)
		return
	}
	finishedGame, err := s.repository.GetFinishedGameById(getFinishedGameByIdRequest.GameID)
	if err != nil {
		response := &network.ErrorResponse{Message: "Error getting finished game by id"}
		json.NewEncoder(client).Encode(response)
		return
	}
	response := &network.FinishedGameResponse{
		Game: finishedGame,
	}
	json.NewEncoder(client).Encode(response)
}

func (s *Server) makeMoveHandler(client net.Conn, msg *network.Message) {
	makeMoveRequest := &network.MakeMoveRequest{}
	if err := json.Unmarshal(msg.Payload, makeMoveRequest); err != nil {
		log.Printf("Error unmarshaling MakeMoveRequest: %v", err)
		return
	}
	room, okRoom := s.rooms[makeMoveRequest.RoomName]
	player, okPlayer := s.players[makeMoveRequest.PlayerName]
	if !okRoom || !okPlayer {
		response := &network.ErrorResponse{Message: "Room not found"}
		msg.Cmd = network.CmdError
		msg.Payload, _ = json.Marshal(response)
		json.NewEncoder(client).Encode(msg)
		return
	}
	room.PlayerStep(
		player,
		makeMoveRequest.PositionRow,
		makeMoveRequest.PositionCol,
	)
}
