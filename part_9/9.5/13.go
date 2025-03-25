package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	ping := make(chan int, 1)
	pong := make(chan int, 1)

	ping <- 1

	go func(ping, pong chan int) {
		timeChan := make(chan time.Time)
		go func(timeChan chan time.Time) {
			timeChan <- (<-time.After(time.Second * 2))
		}(timeChan)

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
			case <-timeChan:
				fmt.Println("Exit")
				os.Exit(0)
			}
		}
	}(ping, pong)

	select {} // вечное ожидание
}
