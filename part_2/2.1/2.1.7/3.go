package main

import "fmt"

type InputFunc func(int, int) int
type OutputFunc func(int) int

func myClosure(a, b int, foo InputFunc) OutputFunc {
	return func(value int) int {
		return value - foo(a, b)
	}
}

func main() {
	globalValue := 99
	mainFunc := func(a, b int) int {
		globalValue--
		fmt.Printf("globalValue: %d, ", globalValue)
		if a < b {
			return globalValue - a + b
		}
		return -globalValue + b*a
	}

	calculation := myClosure(3, 5, mainFunc)
	fmt.Println(calculation(3))
	fmt.Println(calculation(2))
	calculation = myClosure(6, -2, mainFunc)
	fmt.Println(calculation(3))
	fmt.Println(calculation(7))

	fmt.Println("Final globalValue:", globalValue)
}
