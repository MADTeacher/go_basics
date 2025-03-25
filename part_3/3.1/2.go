package main

import "fmt"

func main() {
	add(3, 5)
}

func add(a int, b int) {
	fmt.Println(a + b) // 8
}

// func add(a, b int) {
// 	fmt.Println(a + b)
// }
