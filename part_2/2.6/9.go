package main

import "fmt"

func main() {
	myMap := map[int]string{
		1:   "Alex",
		2:   "Maxim",
		200: "Jon",
	}

	for i, v := range myMap {
		fmt.Printf("key = %d, value = %s\n", i, v)
	}
}
