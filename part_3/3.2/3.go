package main

import "fmt"

func main() {
	array := []float64{2.4, 5.6, 8.1, 9.22}
	myFunc(&array)
	fmt.Printf("Array in main: %v\n", array)
}

func myFunc(array *[]float64) {
	for i := range *array {
		(*array)[i] += 1.5
	}
	fmt.Printf("Array in myFunc: %v\n", *array)
}

// Array in myFunc: [3.9 7.1 9.6 10.72]
// Array in main: [3.9 7.1 9.6 10.72]
