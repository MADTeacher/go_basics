package main

import "fmt"

func main() {
	rez, check := add(10, 8)
	fmt.Printf("a + b = %d, a > b = %t", rez, check)
}

func add(a, b int) (int, bool) {
	return a + b, a > b
}

// a + b = 18, a > b = true

// func main() {
// 	rez := add(10, 8) // assignment mismatch: 1 variable but add returns 2 values
// 	fmt.Printf("a + b = %d", rez)
// }

// func main() {
// 	rez, _ := add(10, 8)
// 	fmt.Printf("a + b = %d", rez) // a + b = 18
// }
