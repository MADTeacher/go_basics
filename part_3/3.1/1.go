package main

import "fmt"

type employee struct {
	name           string
	departmentName string
	age            uint8
	position       string
}

func main() {
	var emp employee
	emp2 := employee{}
	fmt.Println(emp, emp2) // {  0 } {  0 }
}
