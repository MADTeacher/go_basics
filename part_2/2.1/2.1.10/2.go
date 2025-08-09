package main

import "fmt"

func main() {
	defer fmt.Printf("%d ", 0)
	defer fmt.Printf("%d ", 1)
	defer fmt.Printf("%d ", 2)
	fmt.Println("Finish!")
}

// Finish!
// 2 1 0
