package main

import "fmt"

func main() {
	if value := 15; value > 30 {
		fmt.Println("30 < value")
	} else {
		fmt.Println("30 >= value")
	}
}

// 30 >= value
