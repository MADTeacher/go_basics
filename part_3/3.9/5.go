package main

import "fmt"

type shape struct{}
type car struct{}

func checkType(i interface{}) {
	switch i.(type) {
	case shape:
		fmt.Println("It is shape")
	case car:
		fmt.Println("It is car")
	default:
		fmt.Println("What is this pokemon?")
	}
}

func main() {
	checkType(shape{})
	checkType(car{})
	checkType(6)
}
