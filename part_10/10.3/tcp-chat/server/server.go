package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

var connections []net.Conn

func checkErr(err error) {
	if err != nil {
		log.Print(err)
	}
}

func handleConnection(connection net.Conn) {
	connections = append(connections, connection)
	userName, _ := bufio.NewReader(connection).ReadString('\n')
	userName = userName[:len(userName)-2]
	_, err := connection.Write([]byte("Hi " + userName + "\n"))
	checkErr(err)

	for {
		text, err := bufio.NewReader(connection).ReadString('\n')
		if err != nil {
			connection.Close()
			removeConnection(connection)
			broadCastMessage(userName+" is offline\n", connection)
			break
		}

		broadCastMessage(userName+":"+text, connection)
	}
}

func removeConnection(connection net.Conn) {
	var i int

	for i = range connections {
		if connections[i] == connection {
			break
		}
	}

	if len(connections) > 1 {
		connections = append(connections[:i], connections[i+1:]...)
	} else {
		connections = nil
	}
}

func broadCastMessage(msg string, connection net.Conn) {
	// отправка сообщения всем клиентам
	for _, c := range connections {
		if connection != c {
			_, err := c.Write([]byte(msg))
			checkErr(err)
		}
	}

	msg = msg[:len(msg)-1]
	fmt.Println(msg)
}

func RunServer(port int) {
	// net.Listen возвращает Listener переменную,
	// которая является общим сетевым прослушивателем для потоковых протоколов
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	checkErr(err)

	defer listener.Close()

	for {
		connection, err := listener.Accept() // ожидание подключения клиента к серверу
		// Только после успешного вызова Accept()TCP-сервер может начать
		// взаимодействовать с TCP-клиентами
		if err != nil {
			log.Print(err)
			return
		}
		go handleConnection(connection)
	}
}
