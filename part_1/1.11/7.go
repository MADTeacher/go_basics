package main

import "fmt"

func main() {
	array := [4]int{2, 5, 6, 0}
	for i, v := range array {
		// i - номер индекса элемента
		// v – копия значения, хранящаяся по индексу i
		v += 3
		fmt.Printf("%d) %d || ", i, v)
	}
	// 0) 5 || 1) 8 || 2) 9 || 3) 3 ||
	fmt.Println()
	fmt.Println(array) // [2 5 6 0]

	// изменение значения элемента коллекции
	for i := range array {
		array[i] += 3
	}
	fmt.Println(array) // [5 8 9 3]
}
