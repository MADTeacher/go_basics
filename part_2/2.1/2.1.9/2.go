package main

import "fmt"

func createGenerator(start int, end int) chan int {
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
	generator := createGenerator(4, 8)
	for {
		value := <-generator // распаковка значения из канала в переменную
		fmt.Printf("%d || ", value)
		if len(generator) <= 0 { // проверка на выход из бесконечного цикла
			break
		}
	}
}
