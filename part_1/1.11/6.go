package main

import "fmt"

func main() {
	array := [4]int{2, 5, 6, 0}
	for i := 0; i < len(array); i++ {
		array[i] += 3
		fmt.Printf("%d ", array[i])
	}
}
