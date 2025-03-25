package main

import "fmt"

func main() {
	rez, check := add(10, 8)
	fmt.Printf("a + b = %d, a > b = %t", rez, check)
}

func add(a, b int) (rez int, check bool) {
	rez = a + b
	check = a > b
	return // аналогично return rez, check

}

// a + b = 18, a > b = true
