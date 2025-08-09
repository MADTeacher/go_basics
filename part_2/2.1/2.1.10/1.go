package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		defer fmt.Printf("%d ", i)
	}
}

// Finish!
// 2 1 0
