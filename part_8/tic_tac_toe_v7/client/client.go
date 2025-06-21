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

// Client represents the client-side application.
type Client struct {
	conn          net.Conn
	board         *b.Board
	mySymbol      b.BoardField
	currentPlayer b.BoardField
	playerName    string
	roomName      string
	state         State
	mutex         sync.RWMutex
	lastMsgTime   time.Time
}

// NewClient creates a new client and connects to the server.
func NewClient(addr string) (*Client, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	return &Client{
		conn:     conn,
		state:    waitNickNameConfirm,
		mySymbol: b.Empty, // Will be set upon joining a room
	}, nil
}

func (c *Client) setNickname(nickname string) {
	c.playerName = nickname
}

func (c *Client) getState() State {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.state
}

func (c *Client) setState(state State) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// Display a message only when transitioning to opponentMove
	if state == opponentMove && c.state != opponentMove {
		fmt.Println("\nWaiting for opponent's move...")
	} else if state == waitingOpponentInRoom && c.state != waitingOpponentInRoom {
		fmt.Println("\nWaiting for opponent to join...")
	}

	c.state = state
}

// Start begins the client's main loop for sending and receiving messages.
func (c *Client) Start() {
	defer c.conn.Close()

	fmt.Println("Connected to server. ")
	go c.readFromServer()
	c.menu()
}

// readFromServer continuously reads messages from the server and handles them.
func (c *Client) readFromServer() {
	decoder := json.NewDecoder(c.conn)
	for {
		var msg network.Message
		if err := decoder.Decode(&msg); err != nil {
			log.Printf("Disconnected from server: %v", err)
			return // Exit goroutine if connection is lost
		}

		switch msg.Cmd {
		case network.CmdRoomJoinResponse:
			c.handleRoomJoinResponse(msg.Payload)
		case network.CmdInitGame:
			c.handleInitGame(msg.Payload)
		case network.CmdUpdateState:
			c.handleUpdateState(msg.Payload)
		case network.CmdEndGame:
			c.handleEndGame(msg.Payload)
		case network.CmdError:
			c.handleError(msg.Payload)
		case network.CmdRoomListResponse:
			c.handleRoomListResponse(msg.Payload)
		case network.CmdNickNameResponse:
			c.handleNickNameResponse(msg.Payload)
		case network.CmdOpponentLeft:
			c.handleOpponentLeft(msg.Payload)
		case network.CmdFinishedGamesResponse:
			c.handleFinishedGamesResponse(msg.Payload)
		case network.CmdFinishedGameResponse:
			c.handleFinishedGameResponse(msg.Payload)
		default:
			log.Printf(
				"Received unhandled message type '%s' "+
					"from server. Payload: %s\n> ",
				msg.Cmd, string(msg.Payload),
			)
		}
	}
}
