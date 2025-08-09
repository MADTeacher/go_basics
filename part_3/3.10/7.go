package main

import "fmt"

func reverseSlice[T any](slice []T) []T {
	length := len(slice)
	newSlice := make([]T, length)

	for i, elem := range slice {
		newSlice[length-i-1] = elem
	}
	return newSlice
}

func printMySlice[T any](slice []T) {
	fmt.Printf("Slice values before reverse: %v\n", slice)
	fmt.Printf("Slice values after reverse: %v\n", reverseSlice(slice))
	fmt.Println()
}

func main() {
	intSlice := []int{1, 3, 4, 6, 7}
	floatSlice := []float64{3.12, 45.6, 21.6, 5.11}
	stringSlice := []string{"Oo", "^_^", "-_-"}
	printMySlice(intSlice)
	printMySlice(floatSlice)
	printMySlice(stringSlice)
}
