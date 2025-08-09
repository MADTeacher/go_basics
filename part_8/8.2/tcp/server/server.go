package server

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func checkErr(err error) {
	if err != nil {
		log.Print(err)
	}
}

func handleConnection(connection net.Conn) {
	defer connection.Close()

	buffer := make([]byte, 1024)
	for {
		n, err := connection.Read(buffer) // считывание данных
		checkErr(err)
		fmt.Printf("Receive: %s from %s\n",
			string(buffer[0:n-1]), connection.RemoteAddr().String())

		if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
			// завершение работы сервера
			fmt.Println()
			fmt.Printf("Close connection with client: %s\n",
				connection.RemoteAddr().String())
			break
		}

		time.Sleep(2 * time.Second)

		fmt.Printf("Send : %s to %s\n",
			string(buffer[0:n-1]), connection.RemoteAddr().String())
		_, err = connection.Write(buffer) // отправка сообщения клиенту
		checkErr(err)
	}
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
