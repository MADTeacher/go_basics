package client

import (
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
	s, _ := net.ResolveUDPAddr("udp4", fmt.Sprintf("localhost:%d", port))
	connection, err := net.DialUDP("udp4", nil, s)
	checkErr(err)

	fmt.Printf("The UDP server is %s\n", connection.RemoteAddr().String())
	defer connection.Close()
	count := 0
	for {
		data := []byte(fmt.Sprintf("%d)Hello!\n", count))
		fmt.Print("->", string(data))
		_, err = connection.Write(data)

		checkErr(err)

		buffer := make([]byte, 1024)
		n, _, err := connection.ReadFromUDP(buffer)
		checkErr(err)

		fmt.Printf("<-: %s", string(buffer[0:n]))
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
