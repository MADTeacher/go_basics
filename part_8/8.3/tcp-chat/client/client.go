package client

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"sync"
)

var waitingGr sync.WaitGroup

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func sendMessage(connection net.Conn, waitingGr *sync.WaitGroup) {
	defer waitingGr.Done()
	for {
		reader := bufio.NewReader(os.Stdin) // считываем введенное сообщение
		text, _ := reader.ReadString('\n')

		if strings.TrimSuffix(strings.TrimSpace(text), "\r\n") == "/stop" {
			// завершение работы клиента
			break
		} else {
			_, err := connection.Write([]byte(text))
			checkErr(err)
		}
	}
}

func receiveMessage(connection net.Conn, waitingGr *sync.WaitGroup) {
	defer waitingGr.Done()
	for {
		// ждем сообщение от сервера и считываем его
		message, err := bufio.NewReader(connection).ReadString('\n')

		if err == io.EOF {
			fmt.Println("Connection close!")
			break
		} else if err != nil {
			fmt.Println(err.Error())
			break
		}

		message = message[:len(message)-1] // обрезаем символ перевода на следующую строку
		fmt.Println(string(message))
	}
}

func RunClient(port int) {
	// запускает реализацию клиента TCP и соединяет вас с нужным TCP-сервером
	connection, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", port))
	checkErr(err)

	fmt.Printf("The TCP server is %s\n", connection.RemoteAddr().String())
	defer connection.Close()

	waitingGr.Add(1)

	fmt.Println("Enter name: ")
	temp := bufio.NewReader(os.Stdin)
	userName, _ := temp.ReadString('\n')
	_, err = connection.Write([]byte(userName))
	checkErr(err)

	go sendMessage(connection, &waitingGr)
	go receiveMessage(connection, &waitingGr)

	waitingGr.Wait()
}
