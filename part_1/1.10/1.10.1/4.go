package main

import "fmt"

func main() {
	var value int = 15
	if value > 30 {
		if value < 50 {
			fmt.Println("30 < value < 50")
		} else {
			fmt.Println("30 < value >= 50")
		}
	} else {
		if value > 10 {
			fmt.Println("30 >= value >= 10")
		} else {
			fmt.Println("30 >= value < 10")
		}
	}
}

// 30 >= value >= 10
