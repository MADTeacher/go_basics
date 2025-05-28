package main

import (
	"fmt"
	"maps"
)

func main() {
	original := map[string]int{
		"1": 5,
		"2": 3,
		"3": 7,
	}

	// Создание копии таблицы
	cloned := maps.Clone(original)

	// Копирование элементов из одной таблицы в другую
	dest := map[string]int{"x": 1}
	maps.Copy(dest, original)   // копируем элементы original в dest
	fmt.Println("Copy: ", dest) // Copy:  map[1:5 2:3 3:7 x:1]

	// Удаление элементов таблицы по условию
	filtered := maps.Clone(original)
	maps.DeleteFunc(filtered, func(key string, value int) bool {
		return value < 5 // удаляем элементы, где значение <5
	})
	fmt.Println("DeleteFunc: ", filtered) // DeleteFunc:  map[1:5 3:7]

	// Cравнение таблиц
	fmt.Println(maps.Equal(original, cloned))   // false
	fmt.Println(maps.Equal(original, original)) // true

	// Итерация по ключам
	fmt.Print("Keys: ")
	for k := range maps.Keys(original) {
		fmt.Print(" ", k) // Keys:  1 2 3
	}
	fmt.Println()

	// Итерация по значениям
	fmt.Print("Values: ")
	for v := range maps.Values(original) {
		fmt.Print(" ", v) // Values:  3 7 5
	}
}
