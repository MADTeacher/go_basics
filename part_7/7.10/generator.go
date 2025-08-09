package main

import "fmt"

func Generator(start int, end int) <-chan int {
	ch := make(chan int, end-start)
	go func(ch chan int) {
		for i := start; i <= end; i++ {
			ch <- i // помещение значения в канал
		}
		close(ch)
	}(ch)
	return ch
}

func main() {
	for it := range Generator(1, 10) {
		fmt.Printf("%d || ", it)
	}
}
