package main

import (
	"fmt"
	"log"
	"os"
)

type Car struct {
	model  string
	age    uint8
	weight float64
	number string
}

func createCars() []Car {
	return []Car{
		{
			model:  "Oka",
			weight: 13.22,
			number: "ak1245j76",
			age:    13,
		},
		{
			model:  "Lada7",
			weight: 23.14,
			number: "ak2147j175",
			age:    1,
		},
	}
}

func main() {
	myFile, err := os.Create("car_fprintf.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer myFile.Close()

	myCar := createCars()
	for _, it := range myCar {
		n, _ := fmt.Fprintf(myFile, "Model: %s ||Number: %s ||Age: %d ||Weight: %f\n",
			it.model, it.number, it.age, it.weight)
		fmt.Println("Записано: ", n, "байт")
	}
}
