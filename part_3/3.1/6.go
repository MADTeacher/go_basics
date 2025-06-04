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
	fmt.Printf("%+v\n", emp2.age) // 22
	employeeAge := emp2.age
	fmt.Println(employeeAge) // 22
}
