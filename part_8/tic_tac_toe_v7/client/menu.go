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

// Метод управления всем пользовательским потоком
func (c *Client) handleUserFlow() {
	// Создаем буфер для чтения ввода
	reader := bufio.NewReader(os.Stdin)
	// Создаем энкодер для отправки сообщений
	encoder := json.NewEncoder(c.conn)

	// Запрашиваем никнейм у пользователя
	fmt.Print("Enter your nickname: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	c.playerName = input
	var msg network.Message

	// Формируем сообщение и отправляем никнейм на сервер
	msg.Cmd = network.CmdNickname
	payloadData := network.NicknameRequest{Nickname: c.playerName}
	jsonPayload, err := json.Marshal(payloadData)
	if err != nil {
		log.Printf(
			"Error marshalling payload for command %s: %v",
			msg.Cmd, err,
		)
		return
	}
	msg.Payload = jsonPayload
	// Отправляем сообщение на сервер
	if err := encoder.Encode(msg); err != nil {
		log.Printf(
			"Failed to send message to server: %v. Disconnecting.",
			err,
		)
		// Выходим из программы, если не удалось отправить сообщение
		return
	}

	for { // Бесконечный цикл
		switch c.getState() { // Переключение состояний
		case waitNickNameConfirm: // Ожидание подтверждения никнейма
			// Ожидаем подтверждения никнейма от сервера
			time.Sleep(100 * time.Millisecond)
			continue
		case mainMenu: // Главное меню
			// Переходим в главное меню
			c.mainMenu(reader, encoder)
		case playerMove: // Ход игрока
			// Отрабатываем ход игрока
			c.playing(reader, encoder)
		case opponentMove: // Ход противника
			// Ожидаем данные по ходу противника
			time.Sleep(1000 * time.Millisecond)
			continue
		case endGame: // Конец игры
			// Игра завершена. Ждем ее перезапуск от сервера
			fmt.Println("\nGame has ended. Restarting in 10 seconds...")
			time.Sleep(10 * time.Second)
			continue
		case waitResponseFromServer: // Ожидание ответа от сервера
			time.Sleep(100 * time.Millisecond)
			continue
		case waitingOpponentInRoom: // Ожидание противника в комнате
			// Здесь нам надо учесть ситуацию, сто противник может
			// так и не подключиться к комнате. Поэтому, чтобы не
			// заставлять игрока страдать в бесконечном цикле ожидания
			// мы ограничиваем сообщения 1 раз в 3 секунды и считываем
			// ввод пользователя посредством неблокирующего чтения
			now := time.Now()
			// Если прошло более 3 секунд с момента последнего сообщения
			if now.Sub(c.lastMsgTime) > 3*time.Second {
				c.lastMsgTime = now
				fmt.Println("\nWaiting for opponent to join...")
				fmt.Println("Press 'q' and Enter to return to main menu")
				fmt.Print("> ")
			}

			// Проверяем ввод пользователя
			var buffer [1]byte
			n, _ := os.Stdin.Read(buffer[:])
			// Если пользователь нажал 'q' или 'Q',
			// то выходим в главное меню
			if n > 0 && (buffer[0] == 'q' || buffer[0] == 'Q') {
				fmt.Println("Leaving room...")
				// Формируем сообщение о выходе из комнаты
				// и отправляем на сервер
				var msg network.Message
				msg.Cmd = network.CmdLeaveRoomRequest
				payload := network.LeaveRoomRequest{
					RoomName:   c.roomName,
					PlayerName: c.playerName,
				}
				jsonPayload, _ := json.Marshal(payload)
				msg.Payload = jsonPayload
				encoder.Encode(msg)  // Отправляем сообщение на сервер
				c.setState(mainMenu) // Переходим в главное меню
				continue
			}

			// Небольшое время сна для избежания загрузки процессора
			time.Sleep(100 * time.Millisecond)
			continue
		}
	}
}

// Главное меню клиента для взаимодействия с пользователем
func (c *Client) mainMenu(reader *bufio.Reader, encoder *json.Encoder) {
	var msg network.Message // Создаем буфер для сообщения

	// Выводим меню
	fmt.Println("Enter command:")
	fmt.Println("1 - Get room list")
	fmt.Println("2 - Join room")
	fmt.Println("3 - Get finished games")
	fmt.Println("4 - Get finished game by id")
	fmt.Println("5 - Exit")
	fmt.Print("> ")
	// Считываем ввод пользователя
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Преобразуем ввод пользователя в число
	command, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid command.")
		return
	}

	// Обрабатываем ввод пользователя
	switch command {
	case 1: // Получаем список комнат
		// Формируем сообщение и отправляем на сервер
		msg.Cmd = network.CmdListRoomsRequest
		encoder.Encode(msg)
		// Переходим в состояние ожидания ответа от сервера
		c.setState(waitResponseFromServer)
	case 2: // Присоединяемся к комнате
		// Запрашиваем имя комнаты у пользователя
		fmt.Print("Enter room name: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		// Формируем сообщение и отправляем на сервер
		c.roomName = input
		msg.Cmd = network.CmdJoinRoomRequest
		payload := network.JoinRoomRequest{
			RoomName:   c.roomName,
			PlayerName: c.playerName,
		}
		jsonPayload, _ := json.Marshal(payload)
		msg.Payload = jsonPayload
		encoder.Encode(msg) // Отправляем сообщение на сервер
		// Переходим в состояние ожидания ответа от сервера
		c.setState(waitResponseFromServer)
	case 3: // Получаем список завершенных игр
		// Формируем сообщение и отправляем на сервер
		msg.Cmd = network.CmdFinishedGamesRequest
		encoder.Encode(msg)
		// Переходим в состояние ожидания ответа от сервера
		c.setState(waitResponseFromServer)
	case 4: // Получаем завершенную игру по id
		// Запрашиваем id игры у пользователя
		fmt.Print("Enter game id: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Преобразуем ввод пользователя в число
		gameId, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid game id.")
			return
		}

		// Формируем сообщение и отправляем на сервер
		msg.Cmd = network.CmdFinishedGameByIdRequest
		payload := network.GetFinishedGameByIdRequest{GameID: gameId}
		jsonPayload, _ := json.Marshal(payload)
		msg.Payload = jsonPayload
		encoder.Encode(msg) // Отправляем сообщение на сервер
		// Переходим в состояние ожидания ответа от сервера
		c.setState(waitResponseFromServer)
	case 5: // Выходим из программы
		os.Exit(0)
	default: // Неизвестная команда
		fmt.Println("Unknown command.")
		return
	}
}
