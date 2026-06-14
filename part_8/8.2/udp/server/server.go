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
		log.Fatal(err)
	}
}

func RunServer(port int) {
	s, err := net.ResolveUDPAddr("udp4", fmt.Sprintf(":%d", port))
	checkErr(err)

	connection, err := net.ListenUDP("udp4", s)
	checkErr(err)

	defer connection.Close()
	buffer := make([]byte, 1024)
	fmt.Println("Oo")
	for {
		n, addr, err := connection.ReadFromUDP(buffer) // ждем подсоединение клиента
		checkErr(err)
		fmt.Printf("<- %s", string(buffer[0:n]))

		if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
			fmt.Println()
			fmt.Println("Exiting UDP server!")
			return
		}

		time.Sleep(5 * time.Second)

		fmt.Printf("->: %s", string(buffer))
		_, err = connection.WriteToUDP(buffer, addr)
		checkErr(err)
	}
}
