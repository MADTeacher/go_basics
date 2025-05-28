package main

import "fmt"

func main() {
	value := 173.000473
	fmt.Printf("Value = %f\n", value)    // Value = 173.000473
	fmt.Printf("Value = %.f\n", value)   // Value = 173
	fmt.Printf("Value = %.5f\n", value)  // Value = 173.00047
	fmt.Printf("Value = %2.2f\n", value) // Value = 173.00
}
