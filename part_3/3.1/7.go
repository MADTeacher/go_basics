package main

import "fmt"

func main() {
	fmt.Printf("Sum = %d", myFunc(3, 5, 9, 25, -10))
}

func myFunc(values ...int) (sum int) {
	// переменную values рассматриваем как срез []int
	for _, value := range values {
		sum += value
	}
	return
}

// Sum = 32
