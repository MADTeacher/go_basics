package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	myFunc := add             // var myFunc func(a int, b int) int = add
	fmt.Println(add(3, 5))    // 8
	fmt.Println(myFunc(3, 5)) // 8
}
