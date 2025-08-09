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
	myFile, err := os.Open("car_fprintf.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer myFile.Close()

	myCars := []Car{}

	for {
		var car Car
		n, err := fmt.Fscanf(myFile, "Model: %s ||Number: %s ||Age: %d ||Weight: %f\n",
			&car.model, &car.number, &car.age, &car.weight)
		if err != nil {
			if err == io.EOF || err.Error() == "unexpected EOF" {
				break
			}
			log.Fatal(err)
		}
		fmt.Println("Прочитано: ", n, "элемента(-ов)")
		myCars = append(myCars, car)
	}
	fmt.Println(myCars)
}
