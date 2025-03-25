package main

import "fmt"

func main() {
	array := [4]int{2, 5, 6, 0}
	for i, v := range array {
		v += 3
		fmt.Printf("%d) %d || ", i, v)
	}
	// 0) 5 || 1) 8 || 2) 9 || 3) 3 ||
	fmt.Println()
	fmt.Println(array) // [2 5 6 0]
}
