package main

import (
	"log"
	"os"
)

func main() {
	firstText := "Пей, и дьявол тебя доведет до конца,\n"
	secondText := "Йо-хо-хо, и бутылка рома!"
	file, err := os.OpenFile("pirates2.txt", os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.WriteString(firstText)
	file.WriteString(secondText)
}
