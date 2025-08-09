package main

import (
	"fmt"
	"sync"
)

var value int

func increment(waitGroup *sync.WaitGroup, mutex *sync.Mutex) {
	mutex.Lock()         // блокируем критический раздел
	defer mutex.Unlock() // отложенная разблокировка
	value = value + 1
	waitGroup.Done()
}

func main() {
	var waitGroup sync.WaitGroup
	mutex := new(sync.Mutex) // либо var mutex sync.Mutex
	for i := 0; i < 1000; i++ {
		waitGroup.Add(1)
		go increment(&waitGroup, mutex)
	}
	waitGroup.Wait()
	fmt.Println("Value after 1k increment  = ", value)
}
