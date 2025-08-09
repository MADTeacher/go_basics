package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Создаем сканер для чтения из стандартного потока ввода
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() { // ожидаем ввод данных
		// если данные прочитаны успешно, выполнится код
		// тела оператора if
		input := scanner.Text() // Считываем строку
		input = strings.TrimSpace(input)
		fmt.Println(input)
	}
}
