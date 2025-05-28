package main

import "fmt"

func main() {
	s := "Привет"

	for i, r := range s {
		// i - номер индекса, где располагается первый байт символа
		// v – копия символа приведенная к типу Rune
		fmt.Printf("Rune index: %d, Rune: %c (Hex: %x)\n", i, r, r)
	}
}
