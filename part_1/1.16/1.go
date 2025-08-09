package main

import (
	"fmt"
	"slices"
)

func main() {
	mySlice := []int{1, 2, 3, 4, 5, 6, 7}

	// Поэлементно обходим срез в формате индекс:значение
	fmt.Printf("All: ")
	for i, v := range slices.All(mySlice) {
		fmt.Printf(" %d:%d |", i, v)
	}
	fmt.Println()

	// Поэлементно обходим срез в обратном порядке
	fmt.Printf("Backward: ")
	for i, v := range slices.Backward(mySlice) {
		fmt.Printf(" %d:%d |", i, v)
	}
	fmt.Println()
}
