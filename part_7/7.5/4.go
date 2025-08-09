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
	for i := 0; i < 10; i++ {
		myChannel <- i
	}
	fmt.Printf("\nExit")
}
