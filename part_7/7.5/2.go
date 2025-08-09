package main

import (
	"fmt"
)

func myPrint(c chan int) {
	for i := 0; i < 3; i++ {
		value := <-c
		fmt.Printf("%d ", value)
	}

}

func main() {
	// объявление буферизированного канала
	myChannel := make(chan int, 3)
	defer close(myChannel) // отложенное закрытие канала
	go myPrint(myChannel)
	myChannel <- 3
	myChannel <- 10 // запись значения 10 в канал
	myChannel <- 77
	fmt.Printf("\nExit ")
}
