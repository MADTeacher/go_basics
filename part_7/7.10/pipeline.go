package main

import (
	"fmt"
	"math"
)

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	out := source(slice) // первый блок конвейера
	out = square(out)    // блок возведения в квадрат
	out = decriment(out) // блок декремента значения

	for value := range out { // вывод результата работы конвейера в терминал
		fmt.Printf("%d  ", value)
	}
}

func source(in []int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, it := range in {
			if it%3 != 0 {
				out <- it
			}
		}
	}()

	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for i := range in {
			value := math.Pow(float64(i), 2)
			out <- int(value)
		}
	}()

	return out
}

func decriment(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for i := range in {
			out <- i - 1
		}
	}()

	return out
}
