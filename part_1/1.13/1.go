package main

import "fmt"

func main() {
	var val1 int8 = 5
	var val2 int16 = 500
	rez1 := int16(val1) + val2
	// rez2 := val1 + val2
	// invalid operation: val1 + val2 (mismatched types int8 and int16)

	fmt.Println(rez1) // 505
}
