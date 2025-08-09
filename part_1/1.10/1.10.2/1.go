package main

import "fmt"

func main() {
	value := 2
	switch value {
	case 2:
		fmt.Println("2")
		fmt.Printf("%d + 2 = %d", value, value+2)
	case 4:
		fmt.Println("4")
	default:
		fmt.Println("8")
	}
}

// 2
// 2 + 2 = 4
