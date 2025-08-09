package main

import (
	"fmt"
	"time"
)

func namedGorutine(number int) {
	fmt.Printf("%d ", number)
}

func main() {
	for i := 0; i <= 30; i++ {
		go namedGorutine(i)
	}
	time.Sleep(time.Second)
}
