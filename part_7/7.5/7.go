package main

import (
	"fmt"
)

func infoChan(c chan int) {
	length, capacity := len(c), cap(c)
	fmt.Printf("Length = %d, capacity = %d\n", length, capacity)
}

func main() {
	var value int
	myChannel := make(chan int, 3) // объявление буферизированного канала
	defer close(myChannel)         // отложенное закрытие канала
	infoChan(myChannel)
	myChannel <- 1
	myChannel <- 2
	infoChan(myChannel)
	value = <-myChannel
	infoChan(myChannel)
	fmt.Println(value)
}
