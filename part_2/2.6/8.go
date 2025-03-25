package main

import "fmt"

func main() {
	array := [4]int{2, 5, 6, 0}
	for i := range array {
		array[i] += 3
	}
	fmt.Println(array) // [5 8 9 3]
}
