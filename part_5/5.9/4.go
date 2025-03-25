package main

import "fmt"

type shape struct {
	name string
}

func (s *shape) getName() string {
	return s.name
}

func main() {
	var firstInterface interface{} = shape{name: "Cube"}
	value, ok := firstInterface.(string)

	if ok {
		fmt.Println(value) // работаем со значением
	} else {
		fmt.Println("Wrong type assertion!")
	}

	newValue, newOk := firstInterface.(shape)

	if newOk {
		fmt.Println(newValue.getName())
	} else {
		fmt.Println("Wrong type assertion!")
	}
}
