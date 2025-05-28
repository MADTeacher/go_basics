package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите целые числа, разделенные пробелами:")

	// Считываем одну строку ввода
	scanner.Scan()
	input := scanner.Text()

	// Разбиваем строку на срез строк,
	// используя в качестве разделителя символ пробела
	fields := strings.Fields(input)

	var mySlice []int
	// Преобразуем каждое элемент среза fields в целое число
	for _, field := range fields {
		n, err := strconv.Atoi(field)
		if err != nil {
			fmt.Println("Ошибка преобразования:", field)
			return
		}
		mySlice = append(mySlice, n)
	}

	// Удаляем дубликаты
	var result []int
	seen := make(map[int]bool)
	for _, num := range mySlice { // Перебираем элементы среза
		if !seen[num] { // Если элемент еще не встречался
			seen[num] = true
			// Добавляем элемент в результирующий срез
			result = append(result, num)
		}
	}

	fmt.Println("Срез без дубликатов:", result)
}
