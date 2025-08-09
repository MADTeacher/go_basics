package main

import (
	"fmt"
	"slices"
)

func main() {
	base := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("Исходный срез:", base) // Исходный срез: [1 2 3 4 5 6]

	// Вставка элементов в срез в указанную позицию.
	// В нашем случае, начиная с 3-го индекса вставляем 99 и 100
	inserted := slices.Insert(base, 3, 99, 100)
	fmt.Println("Insert: ", inserted) // Insert:  [1 2 3 99 100 4 5 6]

	// Копирование среза
	cloned := slices.Clone(base)
	base[0] = 100
	fmt.Println("Оригинал:", base)   // Оригинал: [100 2 3 4 5 6]
	fmt.Println("Клон:    ", cloned) // Клон:     [1 2 3 4 5 6]

	// Объединение срезов
	parts := []int{7, 8, 9}
	concatenated := slices.Concat(base, parts)
	fmt.Println("Concat: ", concatenated) //Concat: [100 2 3 4 5 6 7 8 9]

	// Удаляем диапазон с i-го по j-й индекс
	deleted := slices.Clone(base)
	deleted = slices.Delete(deleted, 2, 5)
	fmt.Println("Delete: ", deleted) // Delete:  [100 2 6]

	// Удаляем по условию через лямбда-функцию.
	// В нашем случае удаляются элементы, у которых значение кратно 2
	filtered := slices.Clone(base)
	filtered = slices.DeleteFunc(filtered, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println("DeleteFunc: ", filtered) // DeleteFunc:  [3 5]

	// Повторение среза указанное количество раз (например, 2)
	repeated := slices.Repeat(base, 2)
	fmt.Println("Repeat: ", repeated) // [100 2 3 4 5 6 100 2 3 4 5 6]

	// С i-го до j-го индекса значения
	// элементов поменяются на 97,98,99
	replaced := slices.Replace(base, 1, 4, 97, 98, 99)
	fmt.Println("Replace: ", replaced) // Replace:  [100 97 98 99 5 6]

	// Инвертируем порядок элементов среза
	reversed := slices.Clone(base)
	slices.Reverse(reversed)
	fmt.Println("Reverse: ", reversed) // Reverse:  [6 5 99 98 97 100]
}
