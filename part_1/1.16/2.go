package main

import (
	"fmt"
	"slices"
)

func main() {
	mySlice := []int{1, 2, 3, 4, 5, 6, 7}

	// Разбиваем срез на подсрезы размером 3
	for subSlice := range slices.Chunk(mySlice, 3) {
		fmt.Println(subSlice)
	}
}
