package client

import (
	"bufio"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"tic-tac-toe/network"
)

func (c *Client) playing(reader *bufio.Reader, encoder *json.Encoder) {
	fmt.Printf("\nEnter command: <row> <col> or q for exit to main menu\n> ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "q" {
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
		return
	}

	parts := strings.Fields(input)
	if len(parts) != 2 {
		fmt.Println("Usage: <row> <col>")
		return
	}

	var msg network.Message
	row, err1 := strconv.Atoi(parts[0])
	col, err2 := strconv.Atoi(parts[1])
	if err1 != nil || err2 != nil {
		fmt.Println("Row and column must be numbers.")
		return
	}

	if !c.validateMove(row, col) {
		return // validateMove prints the error
	}

	msg.Cmd = network.CmdMakeMoveRequest
	payload := network.MakeMoveRequest{
		RoomName:    c.roomName,
		PlayerName:  c.playerName,
		PositionRow: row - 1,
		PositionCol: col - 1,
	}
	jsonPayload, _ := json.Marshal(payload)
	msg.Payload = jsonPayload
	encoder.Encode(msg)
	c.setState(waitResponseFromServer)
}
