package main

import (
	"fmt"
	"sync"
)

var value int

func increment(waitGroup *sync.WaitGroup, c chan bool) {
	c <- true // блокируем критический раздел
	value = value + 1
	<-c //  разблокировка критического раздела
	waitGroup.Done()
}

func main() {
	var waitGroup sync.WaitGroup
	myChan := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		waitGroup.Add(1)
		go increment(&waitGroup, myChan)
	}
	waitGroup.Wait()
	fmt.Println("Value after 1k increment  = ", value)
}
