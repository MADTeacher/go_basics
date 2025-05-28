package main

import "fmt"

func main() {
	myMap := map[int]string{
		1:   "Alex",
		2:   "Maxim",
		200: "Jon",
	}

	for i, v := range myMap {
		// i - ключ
		// v – копия значения, хранящаяся по ключу i
		fmt.Printf("key = %d, value = %s\n", i, v)
	}
}
