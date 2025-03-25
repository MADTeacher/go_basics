package main

import "fmt"

func main() {
	sum := func(values ...int) (sum int) { // объявление функции
		// переменную values рассматриваем как срез []int
		for _, value := range values {
			sum += value
		}
		return
	}(3, 5, 9, 25, -10) // вызов функции
	fmt.Printf("Sum = %d", sum) // Sum = 32
}
