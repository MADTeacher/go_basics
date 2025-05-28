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
	for range 2 { // код в теле цикла for будет выполнен 2 раза
		line1, err := reader.ReadString('|')
		if err != nil {
			return
		}
		// Удаляем лишние пробелы
		line1 = strings.TrimSpace(line1)
		fmt.Println(line1)
	}
}
