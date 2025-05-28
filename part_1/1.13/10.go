package main

import "fmt"

func main() {
	myBool := true
	var value int
	if myBool {
		value = 1
	} else {
		value = 0
	}
	fmt.Println("Value = ", value) // Value =  1

	myBool2 := value == 1
	fmt.Println("myBool2 value = ", myBool2) // myBool2 value = true
}
