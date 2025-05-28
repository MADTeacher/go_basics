package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "++Hello++ненаглядные++Гофферы!++"

	// Разбиваем строку на срезы по указанному разделителю
	splitRes := strings.Split(str, "++")
	fmt.Print("Split(s, \"++\"): ")
	fmt.Printf("%q\n", splitRes)

	// Разбиваем строку на срезы по указанному разделителю,
	// сохраняя сам разделитель в конце каждого элемента
	// (кроме последнего)
	splitAfterRes := strings.SplitAfter(str, "++")
	fmt.Print("\nSplitAfter(s, \"++\"): ")
	fmt.Printf("%q\n", splitAfterRes)

	// Разбиваем строку на срезы по указанному разделителю,
	// с ограничением на количество подстрок (n). n = 3 означает,
	// что результат будет содержать не более 3 элементов
	splitNRes := strings.SplitN(str, "++", 3)
	fmt.Print("\nSplitN(s, \"++\", 3): ")
	fmt.Printf("%q\n", splitNRes)

	// при n = 0 результат будет nil - []
	splitNRes = strings.SplitN(str, "++", 0)
	fmt.Print("\nSplitN(s, \"++\", 0): ")
	fmt.Printf("%q\n", splitNRes)

	// при n = -1 результат будет содержать все подстроки
	splitNRes = strings.SplitN(str, "++", -1)
	fmt.Print("\nSplitN(s, \"++\", -1): ")
	fmt.Printf("%q\n", splitNRes)
}
