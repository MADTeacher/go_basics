package main

import "fmt"

func main() {
	var value int = 5
	switch {
	case value > 30:
		switch {
		case value < 50:
			fmt.Println("30 < value < 50")
		default:
			fmt.Println("30 < value >= 50")
		}
	case value < 30:
		switch {
		case value > 10:
			fmt.Println("30 >= value >= 10")
		default:
			fmt.Println("30 >= value < 10")
		}
	}
}

// 30 >= value < 10
