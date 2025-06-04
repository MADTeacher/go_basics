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

	fmt.Printf("%+v\n", secondInterface.(string)) // ok
	fmt.Printf("%+v\n", firstInterface.(string))  // error

}
