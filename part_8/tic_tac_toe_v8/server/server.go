package server

import (
	"encoding/json"
	"log"
	"net"
	"sync"

	db "tic-tac-toe/database"
	g "tic-tac-toe/game"
	"tic-tac-toe/network"
	"tic-tac-toe/player"
	"tic-tac-toe/room"
)

// Server manages client connections and game rooms.
type Server struct {
	listener   net.Listener
	repository db.IRepository
	rooms      map[string]*room.Room
	players    map[string]player.IPlayer
	mutex      sync.RWMutex
}

// NewServer creates and returns a new server instance.
func NewServer(addr string, repository db.IRepository) (*Server, error) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	server := &Server{
		listener:   listener,
		repository: repository,
		rooms:      make(map[string]*room.Room),
		players:    make(map[string]player.IPlayer),
	}

	server.rooms["room1"] = room.NewRoom(
		"room1", server.repository, 3, g.PvP, g.None,
	)
	server.rooms["room2"] = room.NewRoom(
		"room2", server.repository, 3, g.PvC, g.Easy,
	)
	server.rooms["room3"] = room.NewRoom(
		"room3", server.repository, 3, g.PvC, g.Medium,
	)
	server.rooms["room4"] = room.NewRoom(
		"room4", server.repository, 3, g.PvC, g.Hard,
	)

	return server, nil
}

// Start begins listening for and handling client connections.
func (s *Server) Start() {
	log.Printf("Server started, listening on %s", s.listener.Addr())
	defer s.listener.Close()

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		go s.handleConnection(conn)
	}
}

// handleConnection manages a single client connection.
func (s *Server) handleConnection(conn net.Conn) {
	log.Printf("New client connected: %s", conn.RemoteAddr())
	defer conn.Close()

	decoder := json.NewDecoder(conn)
	for {
		var msg network.Message
		if err := decoder.Decode(&msg); err != nil {
			log.Printf("Client %s disconnected: %v", conn.RemoteAddr(), err)
			s.disconnectedClientHandler(conn)
			return
		}

		s.handleCommand(conn, &msg)
	}
}

func (s *Server) disconnectedClientHandler(conn net.Conn) {
	var player player.IPlayer
	for _, room := range s.rooms {
		if room.Player1 != nil {
			if room.Player1.CheckSocket(conn) {
				player = room.Player1
				room.RemovePlayer(room.Player1)
				break
			}
		}
		if room.Player2 != nil {
			if room.Player2.CheckSocket(conn) {
				player = room.Player2
				room.RemovePlayer(room.Player2)
				break
			}
		}
	}
	if player == nil {
		log.Printf(
			"Client %s disconnected: player not found",
			conn.RemoteAddr(),
		)
		return
	}
	s.mutex.Lock()
	delete(s.players, player.GetNickname())
	s.mutex.Unlock()
}
