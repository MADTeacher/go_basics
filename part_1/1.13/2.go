package main

import "fmt"

func main() {
	var val1 int8 = 5
	var val2 int16 = 500
	rez1 := val1 + int8(val2)
	fmt.Println(rez1)       // –7
	fmt.Println(int8(val2)) // –12
}
