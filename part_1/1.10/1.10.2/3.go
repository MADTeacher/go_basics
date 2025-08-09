package main

import "fmt"

func main() {
	name := "Alex"
	switch name {
	case "Stanislav":
		fmt.Println("Admin")
	case "Maxim", "Alex", "Bill":
		fmt.Println("Employee")
	}
}

// Employee
