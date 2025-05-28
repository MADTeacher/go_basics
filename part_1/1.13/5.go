package main

import "fmt"

func main() {
	var val1 uint8 = 255 // 1111 1111 в двоичной системе счисления
	val2 := int8(val1)
	fmt.Println(val2) //–1, что =1111 1111 в двоичной системе счисления
}
