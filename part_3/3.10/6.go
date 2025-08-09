package main

import "fmt"

type MyInt int

type MyInterface interface {
	~int | float64 | string | uint16 | uint32
}

func Sum[T MyInterface](a, b T) T {
	return a + b
}

func main() {
	var value1, value2 MyInt = 34, 22
	fmt.Println(Sum(value1, value2)) // 56
	fmt.Println(Sum(10.3, 45.1))     // 55.4
	fmt.Println(Sum(10, 45))         // 55
	fmt.Println(Sum("^_", "^"))      // ^_^
}
