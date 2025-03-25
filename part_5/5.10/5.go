package main

import "fmt"

type MyInt int

func Sum[T ~int | float64 | string](a, b T) T {
	return a + b
}

func main() {
	var value1, value2 MyInt = 34, 22
	fmt.Println(Sum[MyInt](value1, value2)) // 56
	fmt.Println(Sum[float64](10.3, 45.1))   // 55.4
	fmt.Println(Sum[int](10, 45))           // 55
	fmt.Println(Sum[string]("^_", "^"))     // ^_^
}
