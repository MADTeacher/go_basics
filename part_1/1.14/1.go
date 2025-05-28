package main

import (
	"fmt"
	"os"
)

func main() {
	var cityName string
	var age int
	fmt.Print("Введите название города, в котором проживаете: ")
	fmt.Fscan(os.Stdin, &cityName)

	fmt.Print("Введите возраст города: ")
	fmt.Fscan(os.Stdin, &age)

	fmt.Println(cityName, age)
}
