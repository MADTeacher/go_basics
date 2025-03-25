package main

import "fmt"

func main() {
	array := [4]float64{2.4, 5.6, 8.1, 9.22}
	myFunc(array)
	fmt.Printf("Array in main: %v\n", array)
}

func myFunc(array [4]float64) {
	for i := range array {
		array[i] += 2.33
	}
	fmt.Printf("Array in myFunc: %v\n", array)
}

// Array in myFunc: [4.73 7.93 10.43 11.55]
// Array in main: [2.4 5.6 8.1 9.22]
