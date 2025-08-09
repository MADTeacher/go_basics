package main

import "fmt"

func Sum[T int | float64 | string](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(Sum(10.3, 45.1)) // 55.4
	fmt.Println(Sum(10, 45))     // 55
	fmt.Println(Sum("^_", "^"))  // ^_^
}
