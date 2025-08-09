package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func sub(c, a, b int, addFunc func(a int, b int) int) int {
	return c - addFunc(a, b)
}

func main() {
	myFunc := add                      // var myFunc func(a int, b int) int = add
	fmt.Println(sub(4, 3, 10, add))    // -9
	fmt.Println(sub(4, 3, 10, myFunc)) // -9
}
