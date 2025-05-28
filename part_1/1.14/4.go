package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Создаем буфер для чтения из стандартного потока ввода
	reader := bufio.NewReader(os.Stdin)

	// Считываем первую строку до указанного символа,
	// где '\n' - символ конца строки + нажатый Enter
	line1, err := reader.ReadString('\n')
	if err != nil { // считали строку?
		return
	}
	// Удаляем лишние пробелы
	line1 = strings.TrimSpace(line1)
	fmt.Println(line1) // 3 7 8 9 8 7 6 7
}
