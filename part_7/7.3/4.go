package main

import (
	"fmt"
	"time"
)

func namedGorutine(number int) int{
	fmt.Printf("%d ", number)
	return number
}

func main() {
	number := go namedGorutine(1)
	time.Sleep(time.Second)
}
