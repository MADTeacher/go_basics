package main

import (
	"fmt"
	"sync"
)

func namedGorutine(number int, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done() // уменьшение счетчика sync.WaitGroup на 1
	fmt.Printf("%d ", number)
}

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(1) // приведет к ошибке!
	fmt.Printf("%#v\n", waitGroup)

	for i := 0; i <= 30; i++ {
		waitGroup.Add(1) // увеличение счетчика на 1
		go namedGorutine(i, &waitGroup)
	}

	fmt.Printf("\n%#v\n", waitGroup)
	waitGroup.Wait()
}
