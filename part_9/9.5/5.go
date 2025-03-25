package main

import (
	"fmt"
	"time"
)

func onlyRead(c <-chan int) { //только для чтения
	for i := 0; i <= 3; i++ {
		value := <-c
		fmt.Printf("%d ", value)
	}
}

func onlyWrite(c chan<- int) { //только для записи
	for i := 0; i <= 3; i++ {
		c <- i
	}
}

func main() {
	myChannel := make(chan int, 3) // объявление буферизированного канала
	defer close(myChannel)         // отложенное закрытие канала
	go onlyRead(myChannel)
	go onlyWrite(myChannel)
	time.Sleep(time.Second)
}
