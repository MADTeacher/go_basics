package main

import "fmt"

func main() {
	var val1 uint8 = 5 // от 0 до 255
	var val2 uint8 = 255
	fmt.Println(val2 + val1) // 4
	rez1 := uint16(val1) + uint16(val2)
	fmt.Println(rez1) // 260
}
