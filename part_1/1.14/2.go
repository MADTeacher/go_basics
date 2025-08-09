package main

import (
	"fmt"
	"os"
)

func main() {
	var cityName string
	var age int
	fmt.Println("Введите название города, в котором проживаете и его возраст: ")
	fmt.Fscan(os.Stdin, &cityName, &age)

	fmt.Println(cityName, age)
}
