package main

import "fmt"

type employee struct {
	name           string
	departmentName string
	age            uint8
	position       string
}

func main() {
	emp3 := employee{"Alex", "R&D", 25, "Assistant"} // ok
	// emp3 := employee{"Alex", "R&D"} – ошибка!!!

	fmt.Println(emp3) // {Alex R&D 25 Assistant}
}
