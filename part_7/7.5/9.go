package main

import (
	"fmt"
	"time"
)

func main() {
	pingChan := make(chan int, 1)
	pongChan := make(chan int, 1)

	go ping(pingChan, pongChan)
	go pong(pongChan, pingChan)

	pingChan <- 1

	time.Sleep(2 * time.Second)
	fmt.Println("Exit")
}

func ping(pingChan <-chan int, pongChan chan<- int) {
	for {
		ball := <-pingChan

		fmt.Println("Ping", ball)
		time.Sleep(500 * time.Millisecond)

		pongChan <- ball + 1
	}
}

func pong(pongChan <-chan int, pingChan chan<- int) {
	for {
		ball := <-pongChan

		fmt.Println("Pong", ball)
		time.Sleep(500 * time.Millisecond)

		pingChan <- ball + 1
	}
}
