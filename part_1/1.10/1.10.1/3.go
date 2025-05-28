package main

import "fmt"

func main() {
	var value int = 15
	if value > 0 && value < 10 {
		fmt.Println("Число входит в диапазон от 0 до 10")
	} else if value > 10 && value < 20 {
		fmt.Println("Число входит в диапазон от 10 до 20")
	} else if value > 20 && value < 30 {
		fmt.Println("Число входит в диапазон от 20 до 30")
	} else {
		fmt.Println("Значение > 30")
	}
}

// Число входит в диапазон от 10 до 20
