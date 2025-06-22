package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func RunClient(port int) {
	// запускает реализацию клиента TCP и соединяет вас с нужным TCP-сервером
	connection, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", port))
	checkErr(err)

	fmt.Printf("The TCP server is %s\n", connection.RemoteAddr().String())
	defer connection.Close()
	count := 0
	for {
		// отправка сообщения на сервер
		data := []byte(fmt.Sprintf("%d)Hello!\n", count))
		fmt.Print("->", string(data))
		fmt.Fprint(connection, fmt.Sprintf("%d)Hello!\n", count))
		// или
		//_, err = connection.Write(data)
		// checkErr(err)

		// считываем ответ TCP-сервера
		message, err := bufio.NewReader(connection).ReadString('\n')
		// или
		// buffer := make([]byte, 1024)
		// _, err := connection.Read(buffer)
		// message := string(buffer)
		checkErr(err)

		fmt.Printf("<-: %s", message)
		count++
		if count >= 5 {
			data := []byte("STOP")
			fmt.Println("->", string(data))
			_, err = connection.Write(data)
			checkErr(err)

			fmt.Println("Finished!")
			return
		}
	}
}
