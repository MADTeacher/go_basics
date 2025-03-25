package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	firstText := "Пятнадцать человек на сундук мертвеца,\n"
	secondText := "Йо-хо-хо, и бутылка рома,"
	file, err := os.Create("pirates.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// функция WriteString преобразует string в []byte, после чего записывает в файл
	file.WriteString(firstText)
	// ручное приведение типа string в []byte и запись данных в файл
	file.Write([]byte(secondText))
	fmt.Println("Text was wrote to file")
}
