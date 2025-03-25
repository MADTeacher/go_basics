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
	fmt.Printf("%+v\n", emp2) // {name:Tom departmentName:R&D age:22 position:Intern}
	emp2.position = "Engineer"
	fmt.Printf("%+v\n", emp2) // {name:Tom departmentName:R&D age:22 position:Engineer}
}
