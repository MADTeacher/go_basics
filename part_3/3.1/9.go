package main

import "fmt"

type employee struct {
	name string
	age  uint8
}

func valueFunc(emp employee) {
	emp.name = "O_O"
	fmt.Printf("Copy value: %+v\n", emp)
}

func pointFunc(emp *employee) {
	emp.name = "^_^"
	fmt.Printf("Point to value: %+v\n", emp)
}

func main() {
	emp := employee{
		name: "Tom",
		age:  45,
	}
	newEmployee := emp
	newEmployee.age = 22
	fmt.Printf("newEmployee = %+v\n", newEmployee)
	fmt.Printf("emp = %+v\n", emp)

	valueFunc(emp) // передача в функцию по значению
	fmt.Printf("emp = %+v\n", emp)
	pointFunc(&emp) // передача в функцию по указателю
	fmt.Printf("emp = %+v\n", emp)
}
