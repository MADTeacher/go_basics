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
	// подключение к серверу
	conn net.Conn
	// игровое поле
	board *b.Board
	// фигура игрока
	mySymbol b.BoardField
	// фигура игрока, ход которой сейчас
	currentPlayer b.BoardField
	// никнейм игрока
	playerName string
	// имя комнаты
	roomName string
	// текущее состояние клиента
	state State
	// мьютекс для защиты доступа к данным
	mutex sync.RWMutex
	// время последнего сообщения
	lastMsgTime time.Time
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

// Устанавливаем никнейм игрока
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
	// Запускаем меню клиента для взаимодействия с пользователем
	c.menu()
}

// Читаем сообщения от сервера
func (c *Client) readFromServer() {
	decoder := json.NewDecoder(c.conn) // Создаем декодер
	for {                              // Бесконечный цикл
		var msg network.Message
		if err := decoder.Decode(&msg); err != nil {
			log.Printf("Disconnected from server: %v", err)
			return // если соединение потеряно, то выходим из горутины
		}

		switch msg.Cmd {
		// Обрабатываем ответ на запрос на присоединение к комнате
		case network.CmdRoomJoinResponse:
			c.handleRoomJoinResponse(msg.Payload)
		// Обрабатываем сообщение об инициализацию игры
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
