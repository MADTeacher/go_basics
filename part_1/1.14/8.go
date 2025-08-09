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
	var keys []int
	var values []string

	// Чтение ключей
	if !scanner.Scan() { // если произошла ошибка
		return // завершаем выполнение программы
	}
	keyParts := strings.Fields(scanner.Text())
	// преобразование в целочисленный срез
	for _, part := range keyParts {
		num, err := strconv.Atoi(part)
		if err != nil {
			fmt.Println("Ошибка в ключах:", part)
			return
		}
		keys = append(keys, num)
	}

	// Чтение значений
	if !scanner.Scan() {
		return
	}
	values = strings.Fields(scanner.Text())

	// Проверка на то, что количество ключей и значений совпадает
	if len(keys) != len(values) {
		fmt.Println("Количество ключей и значений не совпадает")
		return
	}

	// Чтение числа A
	if !scanner.Scan() {
		return
	}
	a, err := strconv.Atoi(scanner.Text())
	if err != nil || a == 0 {
		fmt.Println("Некорректное число A")
		return
	}

	// Создание и заполнение таблицы
	data := make(map[int]string)
	for i := range keys {
		data[keys[i]] = values[i]
	}

	// Поиск ключей для удаления
	var toDelete []int
	for k := range data {
		if k%a == 0 {
			toDelete = append(toDelete, k)
		}
	}

	// Удаление элементов
	for _, k := range toDelete {
		delete(data, k)
	}

	fmt.Println(data)
}
