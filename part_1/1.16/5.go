package main

import (
	"fmt"
	"slices"
)

func main() {
	mySlice := []int{1, 2, 3, 4, 5, 6}

	// BinarySearch принимает на вход срез и значение,
	// которое нужно найти. Работает только с отсортированными срезами
	index, found := slices.BinarySearch(mySlice, 4)
	fmt.Printf("BinarySearch(4): index=%d, found=%t\n", index, found)
	// BinarySearch(4): index=3, found=true

	index, found = slices.BinarySearch(mySlice, 7)
	fmt.Printf("BinarySearch(7): index=%d, found=%t\n", index, found)
	// BinarySearch(7): index=6, found=false

	// Ищем индекс первого вхождения элемента
	data := []int{5, 2, 5, 3, 5}
	fmt.Println(slices.Index(data, 5)) // 0
	fmt.Println(slices.Index(data, 7)) // -1
}
