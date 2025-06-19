package main

import (
	"fmt"
	"sync"
)

var value = 0

func increment(waitGroup *sync.WaitGroup) {
	value = value + 1
	waitGroup.Done()
}

func main() {
	var waitGroup sync.WaitGroup
	for i := 0; i < 1000; i++ {
		waitGroup.Add(1)
		go increment(&waitGroup)
	}
	waitGroup.Wait()
	fmt.Println("Value after 1k increment  = ", value)
}
