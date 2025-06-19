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
	myChannel := make(chan int, 3) // объявление буферизированного канала
	defer close(myChannel)         // отложенное закрытие канала
	go myPrint(myChannel)
	myChannel <- 3
	myChannel <- 10 // запись значения 10 в канал
	myChannel <- 77
	myChannel <- 105
	myChannel <- 104
	fmt.Printf("\nExit")
}
