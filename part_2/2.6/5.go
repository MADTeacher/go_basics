package main

import "fmt"

func main() {
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			if i == j {
				continue
			}
			fmt.Printf("%d %d || ", i, j)
		}
	}
	// 0 1 || 0 2 || 1 0 || 1 2 ||

	fmt.Println()
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			if i == j {
				break
			}
			fmt.Printf("%d %d || ", i, j)
		}
	}
	fmt.Println("Oo")
	// 1 0 || Oo
}
