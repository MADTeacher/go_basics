package main

import (
	"fmt"
	"slices"
)

func main() {
	// Удаление последовательных дубликатов
	values := []int{1, 1, 2, 3, 3, 3, 4, 5, 5, 6, 1}
	compacted := slices.Compact(values)
	fmt.Println("Compact:", compacted) // Compact: [1 2 3 4 5 6 1]

	// Поиск максимума
	values = []int{5, 2, 9, 1, 7}
	fmt.Println("Max:", slices.Max(values)) // Max: 9

	// Поиск минимума
	fmt.Println("Min:", slices.Min(values)) // Min: 1
}
