package main

import "fmt"

type employee struct {
	name           string
	departmentName string
	age            uint8
	position       string
}

func main() {
	emp2 := employee{
		name:     "Tom",
		position: "Intern",
	}
	emp2.age = 22
	emp2.departmentName = "R&D"
	fmt.Println(emp2) // {Tom R&D 22 Intern}
	emp2.position = "Engineer"
	fmt.Println(emp2) // {Tom R&D 22 Engineer}
}
