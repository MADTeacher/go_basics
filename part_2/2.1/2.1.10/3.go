package main

import "fmt"

func add(a, b int) int {
	defer fmt.Println("defer") // выполнится перед return
	fmt.Println("without defer")
	return a + b
}

func main() {
	fmt.Println(add(42, 13))
}
