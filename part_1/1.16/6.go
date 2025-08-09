package main

import (
	"fmt"
	"slices"
)

func main() {
	// Проверка наличия элемента с таким значением в срезе
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(slices.Contains(numbers, 3)) // true
	fmt.Println(slices.Contains(numbers, 6)) // false

	// Лексикографическое сравнение срезов
	a := []int{1, 2, 3}
	b := []int{1, 2, 4}

	fmt.Println(slices.Compare(a, b)) // -1 (a < b)
	fmt.Println(slices.Compare(b, a)) // 1 (b > a)
	fmt.Println(slices.Compare(a, a)) // 0 (equal)

	// Проверка срезов на равенство
	x := []int{1, 2, 3}
	y := []int{1, 2, 3}
	z := []int{3, 2, 1}

	fmt.Println(slices.Equal(x, y)) // true
	fmt.Println(slices.Equal(x, z)) // false
}
