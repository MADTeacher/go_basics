package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	numbers := []int{5, 2, 6, 3, 1, 4}
	words := []string{"banana", "apple", "cherry", "date"}

	// Сортировка среза
	slices.Sort(numbers)
	fmt.Println("Sort (int):", numbers) // Sort (int): [1 2 3 4 5 6]

	slices.Sort(words)
	fmt.Println("Sort (string):", words) // [apple banana cherry date]

	// Настраиваемая сортировка через лямбда-функцию
	mixed := []int{30, 15, 42, 7, 25}

	// Функция должна возвращать:
	// -1 если первый̆ аргумент меньше второго,
	// 0 если аргументы равны,
	// 1 если первый̆ аргумент больше второго
	// за нас это может рассчитать cmp.Compare
	slices.SortFunc(mixed, func(a, b int) int {
		// Сортировка по последней цифре
		return cmp.Compare(a%10, b%10)
	})
	fmt.Println("SortFunc: ", mixed) // SortFunc:  [30 42 15 25 7]

	// Проверка на то, что элементы среза отсортированы
	sortedCheck := []int{1, 3, 5, 7, 9}
	fmt.Printf("IsSorted(%v): %t\n", sortedCheck,
		slices.IsSorted(sortedCheck)) // IsSorted([1 3 5 7 9]): true

	unsortedCheck := []int{2, 1, 3, 4}
	fmt.Printf("IsSorted(%v): %t\n", unsortedCheck,
		slices.IsSorted(unsortedCheck)) // IsSorted([2 1 3 4]): false
}
