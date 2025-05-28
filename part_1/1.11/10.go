package main

import "fmt"

func main() {
	s := "Hello"

	for i := 0; i < len(s); i++ {
		fmt.Printf("Byte index: %d, Symbol: %c\n", i, s[i])
	}
}
