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
	var secondInterface interface{} = "Oo"

	fmt.Printf("%+v\n", firstInterface)  // {name:Cube}
	fmt.Printf("%+v\n", secondInterface) // Oo
}
