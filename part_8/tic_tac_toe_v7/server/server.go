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

type Server struct {
	// Слушатель для подключения клиентов
	listener net.Listener
	// Интерфейс для сохранения завершенных игр в базе данных
	repository db.IRepository
	// Карта комнат
	rooms map[string]*room.Room
	// Карта игроков
	players map[string]player.IPlayer
	// Мьютекс для защиты доступа к данным
	mutex sync.RWMutex
}

func NewServer(addr string, repository db.IRepository) (*Server, error) {
	listener, err := net.Listen("tcp", addr) // Создаем слушатель
	if err != nil {
		return nil, err
	}

	// Создаем экземпляр структуры Server
	server := &Server{
		listener:   listener,
		repository: repository,
		rooms:      make(map[string]*room.Room),
		players:    make(map[string]player.IPlayer),
	}

	// Создаем комнаты и добавляем их в карту rooms по ключу,
	// который является именем комнаты
	server.rooms["room1"] = room.NewRoom(
		"room1", server.repository, 3, g.PvP, g.None,
	)
	server.rooms["room2"] = room.NewRoom(
		"room2", server.repository, 3, g.PvC, g.Easy,
	)
	server.rooms["room3"] = room.NewRoom(
		"room3", server.repository, 5, g.PvC, g.Medium,
	)
	server.rooms["room4"] = room.NewRoom(
		"room4", server.repository, 6, g.PvC, g.Hard,
	)

	// Возвращаем экземпляр созданного сервера
	return server, nil
}

// Запускаем прослушивание подключений и обработку сообщений от клиентов
func (s *Server) Start() {
	log.Printf("Server started, listening on %s", s.listener.Addr())
	defer s.listener.Close() // Закрываем слушателя при завершении

	// Запускаем бесконечный цикл обработки подключений
	for {
		// Принимаем подключение
		conn, err := s.listener.Accept()
		if err != nil { // Если ошибка, то пропускаем итерацию
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		// Чтобы не блокировать основной поток
		// передаем сокетное подключение в отдельную горутину, где и
		// будем обрабатывать сообщения от клиента по данному подключению
		go s.handleConnection(conn)
	}
}

// Обрабатываем подключение клиента
func (s *Server) handleConnection(conn net.Conn) {
	log.Printf("New client connected: %s", conn.RemoteAddr())
	defer conn.Close() // Закрываем подключение при завершении

	// Создаем декодер для чтения сообщений от клиента
	decoder := json.NewDecoder(conn)
	for { // Бесконечный цикл обработки сообщений
		// Создаем переменную для хранения сообщения
		var msg network.Message
		// Декодируем сообщение от клиента
		if err := decoder.Decode(&msg); err != nil { // Если ошибка
			log.Printf(
				"Client %s disconnected: %v", conn.RemoteAddr(),
				err,
			)
			// Обрабатываем отключение клиента
			s.disconnectedClientHandler(conn)
			return
		}

		// Обрабатываем сообщение от клиента
		s.handleCommand(conn, &msg)
	}
}

// Обрабатываем отключение клиента
func (s *Server) disconnectedClientHandler(conn net.Conn) {
	var player player.IPlayer // Создаем переменную для хранения игрока
	// Проходим по всем комнатам
	for _, room := range s.rooms {
		// Если игрок 1 в комнате
		if room.Player1 != nil {
			// Если игрок 1 подключился по этому сокету
			if room.Player1.CheckSocket(conn) {
				player = room.Player1
				// Удаляем игрока из комнаты
				room.RemovePlayer(room.Player1)
				break
			}
		}
		// Если игрок 2 в комнате
		if room.Player2 != nil {
			// Если игрок 2 подключился по этому сокету
			if room.Player2.CheckSocket(conn) {
				player = room.Player2
				// Удаляем игрока из комнаты
				room.RemovePlayer(room.Player2)
				break
			}
		}
	}
	// Если игрок не найден
	if player == nil {
		log.Printf(
			"Client %s disconnected: player not found",
			conn.RemoteAddr(),
		)
		return
	}
	// Удаляем игрока из карты, предварительно защищая доступ к ней
	s.mutex.Lock()
	delete(s.players, player.GetNickname())
	s.mutex.Unlock()
}
