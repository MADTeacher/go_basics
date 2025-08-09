package main

import (
	"fmt"
	"time"
)

func main() {
	ping := make(chan int, 1)
	pong := make(chan int, 1)

	// ping <- 1

	go func(ping, pong chan int) {
		var ball int
		for {
			select {
			case ball = <-ping:
				fmt.Println("Ping", ball)
				time.Sleep(500 * time.Millisecond)
				pong <- ball + 1
			case ball = <-pong:
				fmt.Println("Pong", ball)
				time.Sleep(500 * time.Millisecond)
				ping <- ball + 1
			default:
				fmt.Println("Time out!!!")
				return
			}
		}
	}(ping, pong)

	time.Sleep(2 * time.Second)
	fmt.Println("Exit")
}
