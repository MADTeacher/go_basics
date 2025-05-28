package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// или
	for i := range 5 {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// или
	i := 0
	for range 5 {
		fmt.Printf("%d ", i)
		i++
	}
}
