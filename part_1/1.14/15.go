package main

import "fmt"

func main() {
	value := 173.000473
	str := fmt.Sprintf("Value = %x", value)
	fmt.Printf("str = %q\n", str)
}
