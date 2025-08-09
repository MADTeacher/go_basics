package main

import "fmt"

type employee struct {
	name           string
	departmentName string
	age            uint8
	position       string
}

func main() {
	emp1 := employee{
		name:           "Alex",
		age:            25,
		departmentName: "R&D",
		position:       "Assistant",
	}
	emp2 := employee{
		name:     "Tom",
		position: "Intern",
	}

	fmt.Println(emp1) // {Alex R&D 25 Assistant}
	fmt.Println(emp2) // {Tom  0 Intern}
}
