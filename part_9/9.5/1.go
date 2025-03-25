package main

import (
	"fmt"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

func myPrint(c chan int) {
	value, ok := <-c // чтение из канала. Горутина блокируется до тех пор,
	// пока не появятся данные для чтения
	if ok {
		fmt.Println("Channel is open!")
	} else {
		fmt.Println("Channel is closed!")
	}
	fmt.Println(value, ok)
	waitGroup.Done()
}

func main() {
	myChannel := make(chan int) // объявление канала из 1-го элемента
	defer close(myChannel)      // отложенное закрытие канала
	go myPrint(myChannel)
	waitGroup.Add(1)
	time.Sleep(time.Second)
	myChannel <- 10 // запись значения 10 в канал
	waitGroup.Wait()
}
