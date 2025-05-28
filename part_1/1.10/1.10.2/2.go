package main

import "fmt"

func main() {
	value := 2
	switch value := 4; value {
	case 2:
		fmt.Println("2")
		fmt.Printf("%d + 2 = %d", value, value+2)
	case 4:
		fmt.Println("4")
	default:
		fmt.Println("8")
	}
	fmt.Printf("value = %d", value)
}

// 4
// value = 2
