package main

import "fmt"

type shape struct{}
type car struct{}

func isEqual(i interface{}, j interface{}) {
	if i == j { // одинаковое значение и базовый тип?
		fmt.Println("Equal")
	} else {
		fmt.Println("Inequal")
	}
}

func main() {
	isEqual(shape{}, shape{}) // Equal
	isEqual(shape{}, car{})   // Inequal

	var firstInterface interface{} // по умолчанию - nil
	var secondInterface interface{}

	isEqual(firstInterface, secondInterface) // Equal
}
