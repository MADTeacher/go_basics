package main

import "fmt"

func namedGorutine() {
	for i := 0; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
}

func main() {
	go namedGorutine() // объявление горутины
	go func() {
		for i := 6; i <= 10; i++ {
			fmt.Printf("%d", i)
		}
	}()
}
