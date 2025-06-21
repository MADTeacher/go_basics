package client

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"tic-tac-toe/network"
	"time"
)

func (c *Client) menu() {
	reader := bufio.NewReader(os.Stdin)
	encoder := json.NewEncoder(c.conn)

	fmt.Print("Enter your nickname: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	c.playerName = input
	var msg network.Message

	msg.Cmd = network.CmdNickname
	payloadData := network.NicknameRequest{Nickname: c.playerName}
	jsonPayload, err := json.Marshal(payloadData)
	if err != nil {
		log.Printf("Error marshalling payload for command %s: %v", msg.Cmd, err)
		return
	}
	msg.Payload = jsonPayload
	if err := encoder.Encode(msg); err != nil {
		log.Printf("Failed to send message to server: %v. Disconnecting.", err)
		return // Exit if we can't send to server
	}

	for {
		switch c.getState() {
		case waitNickNameConfirm:
			time.Sleep(100 * time.Millisecond)
			continue
		case mainMenu:
			c.mainMenu(reader, encoder)
		case playerMove:
			c.playing(reader, encoder)
		case opponentMove:
			// Just wait silently for opponent's move
			time.Sleep(1000 * time.Millisecond)
			continue
		case endGame:
			fmt.Println("\nGame has ended. Restarting in 10 seconds...")
			time.Sleep(10 * time.Second)
			continue
		case waitResponseFromServer:
			time.Sleep(100 * time.Millisecond)
			continue
		case waitingOpponentInRoom:
			// Rate-limit messages to once every 3 seconds
			now := time.Now()
			if now.Sub(c.lastMsgTime) > 3*time.Second {
				c.lastMsgTime = now
				fmt.Println("\nWaiting for opponent to join...")
				fmt.Println("Press 'q' and Enter to return to main menu")
				fmt.Print("> ")
			}

			// Poll for input every cycle but don't block
			var buffer [1]byte
			n, _ := os.Stdin.Read(buffer[:])
			if n > 0 && (buffer[0] == 'q' || buffer[0] == 'Q') {
				fmt.Println("Leaving room...")
				var msg network.Message
				msg.Cmd = network.CmdLeaveRoomRequest
				payload := network.LeaveRoomRequest{
					RoomName:   c.roomName,
					PlayerName: c.playerName,
				}
				jsonPayload, _ := json.Marshal(payload)
				msg.Payload = jsonPayload
				encoder.Encode(msg)
				c.setState(mainMenu)
				continue
			}

			// Short sleep to avoid CPU spinning
			time.Sleep(100 * time.Millisecond)
			continue
		}
	}
}

func (c *Client) mainMenu(reader *bufio.Reader, encoder *json.Encoder) {
	var msg network.Message

	fmt.Println("Enter command:")
	fmt.Println("1 - Get room list")
	fmt.Println("2 - Join room")
	fmt.Println("3 - Get finished games")
	fmt.Println("4 - Get finished game by id")
	fmt.Println("5 - Exit")
	fmt.Print("> ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	command, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid command.")
		return
	}

	switch command {
	case 1:
		msg.Cmd = network.CmdListRoomsRequest
		encoder.Encode(msg)
		c.setState(waitResponseFromServer)
	case 2:
		fmt.Print("Enter room name: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		c.roomName = input
		msg.Cmd = network.CmdJoinRoomRequest
		payload := network.JoinRoomRequest{
			RoomName:   c.roomName,
			PlayerName: c.playerName,
		}
		jsonPayload, _ := json.Marshal(payload)
		msg.Payload = jsonPayload
		encoder.Encode(msg)
		c.setState(waitResponseFromServer)
	case 3:
		msg.Cmd = network.CmdFinishedGamesRequest
		encoder.Encode(msg)
		c.setState(waitResponseFromServer)
	case 4:
		fmt.Print("Enter game id: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		gameId, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid game id.")
			return
		}

		msg.Cmd = network.CmdFinishedGameByIdRequest
		payload := network.GetFinishedGameByIdRequest{GameID: gameId}
		jsonPayload, _ := json.Marshal(payload)
		msg.Payload = jsonPayload
		encoder.Encode(msg)
		c.setState(waitResponseFromServer)
	case 5:
		os.Exit(0)
	default:
		fmt.Println("Unknown command.")
		return
	}
}
