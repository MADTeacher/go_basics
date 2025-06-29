package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	b "tic-tac-toe/board"
	"tic-tac-toe/network"
)

// Объявление структуры клиента
type Client struct {
	conn          net.Conn     // подключение к серверу
	board         *b.Board     // игровое поле
	mySymbol      b.BoardField // фигура игрока
	currentPlayer b.BoardField // фигура игрока, чей сейчас ход
	playerName    string       // имя комнаты
	roomName      string       // имя комнаты
	state         State        // текущее состояние клиента
	mutex         sync.RWMutex // мьютекс для защиты доступа к данным
	lastMsgTime   time.Time    // время последнего сообщения
}

// Констукторная функция для создания клиента
func NewClient(addr string) (*Client, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	return &Client{
		// подключение к серверу
		conn: conn,
		// начальное состояние клиента
		state: waitNickNameConfirm,
		// mySymbol будет установлен при подключении к комнате
		mySymbol: b.Empty,
	}, nil
}

// Устанавливаем никнейма игрока
func (c *Client) setNickname(nickname string) {
	c.playerName = nickname
}

// Получаем текущее состояние клиента
func (c *Client) getState() State {
	c.mutex.RLock() // защищаем доступ к данным
	defer c.mutex.RUnlock()
	return c.state
}

// Устанавливаем текущее состояние клиента
func (c *Client) setState(state State) {
	c.mutex.Lock() // защищаем доступ к данным
	defer c.mutex.Unlock()

	// Если переходим в состояние opponentMove
	if state == opponentMove && c.state != opponentMove {
		fmt.Println("\nWaiting for opponent's move...")
	} else if state == waitingOpponentInRoom &&
		c.state != waitingOpponentInRoom {
		// Если переходим в состояние waitingOpponentInRoom
		fmt.Println("\nWaiting for opponent to join...")
	}

	c.state = state // устанавливаем новое состояние
}

// Запускаем клиента
func (c *Client) Start() {
	defer c.conn.Close()

	fmt.Println("Connected to server. ")
	// Запускаем горутину для чтения сообщений от сервера
	go c.readFromServer()
	// Запускаем управление пользовательским потоком
	c.handleUserFlow()
}

// Читаем сообщения от сервера
func (c *Client) readFromServer() {
	decoder := json.NewDecoder(c.conn) // Создаем декодер
	for {                              // Бесконечный цикл
		// Создаем буфер и десериализуем в него сообщение от сервера
		var msg network.Message
		if err := decoder.Decode(&msg); err != nil {
			log.Printf("Disconnected from server: %v", err)
			return // если соединение потеряно, то выходим из горутины
		}

		switch msg.Cmd {
		// Обрабатываем ответ на запрос о присоединении к комнате
		case network.CmdRoomJoinResponse:
			c.handleRoomJoinResponse(msg.Payload)
		// Обрабатываем сообщение об инициализации игры
		case network.CmdInitGame:
			c.handleInitGame(msg.Payload)
		// Обрабатываем сообщение об обновлении состояния игры
		case network.CmdUpdateState:
			c.handleUpdateState(msg.Payload)
		// Обрабатываем сообщение об окончании игры
		case network.CmdEndGame:
			c.handleEndGame(msg.Payload)
		// Обрабатываем сообщение об ошибке
		case network.CmdError:
			c.handleError(msg.Payload)
		// Обрабатываем сообщение о списке комнат
		case network.CmdRoomListResponse:
			c.handleRoomListResponse(msg.Payload)
		// Обрабатываем сообщение о подтверждении никнейма
		case network.CmdNickNameResponse:
			c.handleNickNameResponse(msg.Payload)
		// Обрабатываем сообщение об отключении оппонента
		case network.CmdOpponentLeft:
			c.handleOpponentLeft(msg.Payload)
		// Обрабатываем сообщение о списке завершенных игр
		case network.CmdFinishedGamesResponse:
			c.handleFinishedGamesResponse(msg.Payload)
		// Обрабатываем сообщение с данными по
		// запрошенной завершенной игре
		case network.CmdFinishedGameResponse:
			c.handleFinishedGameResponse(msg.Payload)
		default: // Если пришло неизвестное сообщение
			log.Printf(
				"Received unhandled message type '%s' "+
					"from server. Payload: %s\n> ",
				msg.Cmd, string(msg.Payload),
			)
		}
	}
}
