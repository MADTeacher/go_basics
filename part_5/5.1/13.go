package main

import (
	"log"
	"os"
)

func main() {
	firstText := "Пятнадцать человек на сундук мертвеца,\n"
	secondText := "Йо-хо-хо, и бутылка рома,\n"
	file, err := os.OpenFile("pirates2.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.WriteString(firstText)
	file.WriteString(secondText)
}
