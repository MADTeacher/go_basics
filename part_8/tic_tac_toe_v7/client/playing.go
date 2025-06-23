package client

import (
	"bufio"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"tic-tac-toe/network"
)

// Обрабатываем ход игрока
func (c *Client) playing(reader *bufio.Reader, encoder *json.Encoder) {
	fmt.Printf(
		"\nEnter command: <row> <col> or q for exit to main menu\n> ",
	)
	// Считываем ввод игрока
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "q" { // Если игрок хочет выйти в меню
		var msg network.Message               // Создаем сообщение
		msg.Cmd = network.CmdLeaveRoomRequest // Устанавливаем команду
		payload := network.LeaveRoomRequest{
			RoomName:   c.roomName,
			PlayerName: c.playerName,
		}
		jsonPayload, _ := json.Marshal(payload)
		msg.Payload = jsonPayload
		encoder.Encode(msg)  // Отправляем сообщение
		c.setState(mainMenu) // Переходим в главное меню
		return
	}

	// Разделяем ввод игрока на строки
	parts := strings.Fields(input)
	if len(parts) != 2 {
		fmt.Println("Usage: <row> <col>")
		return
	}

	var msg network.Message // Создаем сообщение
	// Преобразуем ввод игрока в числа
	row, err1 := strconv.Atoi(parts[0])
	col, err2 := strconv.Atoi(parts[1])
	if err1 != nil || err2 != nil {
		fmt.Println("Row and column must be numbers.")
		return
	}

	// Валидируем ввод игрока
	if !c.validateMove(row, col) {
		return // validateMove prints the error
	}

	// Создаем сообщение о ходе игрока
	msg.Cmd = network.CmdMakeMoveRequest // Устанавливаем команду
	payload := network.MakeMoveRequest{
		RoomName:    c.roomName,   // Устанавливаем имя комнаты
		PlayerName:  c.playerName, // Устанавливаем никнейм игрока
		PositionRow: row - 1,      // Устанавливаем строку
		PositionCol: col - 1,      // Устанавливаем столбец
	}
	jsonPayload, _ := json.Marshal(payload)
	msg.Payload = jsonPayload
	encoder.Encode(msg) // Отправляем сообщение
	// Переходим в состояние ожидания ответа от сервера
	c.setState(waitResponseFromServer)
}
