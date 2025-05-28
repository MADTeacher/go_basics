package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hi, my dear friend!!!"
	fmt.Printf("Длина строки: %d\n", len(str)) // Длина строки: 21
	fmt.Printf("Символов 'i' – %d", strings.Count(str, "i"))
	// Символов 'i' – 2
}
