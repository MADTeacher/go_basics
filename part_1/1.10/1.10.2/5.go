package main

import "fmt"

func main() {
	name := "Stanislav"
	switch name {
	case "Stanislav":
		fmt.Println("Admin")
		break
		fallthrough
	case "Maxim", "Alex", "Bill":
		fmt.Println("Employee")
	}
}

// Admin
