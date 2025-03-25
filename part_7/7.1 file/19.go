package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type Car struct {
	model  string
	age    uint8
	weight float64
	number string
}

func main() {
	myFile, err := os.Open("car_fprintln.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer myFile.Close()

	myCars := []Car{}

	for {
		var car Car
		n, err := fmt.Fscanln(myFile, &car.model, &car.number, &car.age, &car.weight)
		if err != nil {
			if err == io.EOF { // Читаем до конца файла
				break
			}
			log.Fatal(err)
		}
		fmt.Println("Прочитано: ", n, "элемента(-ов)")
		myCars = append(myCars, car)
	}
	fmt.Println(myCars)
}
