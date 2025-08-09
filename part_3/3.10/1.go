package main

import "fmt"

func SumInt(a, b int) int {
	return a + b
}

func SumFloat(a, b float64) float64 {
	return a + b
}

func SumString(a, b string) string {
	return a + b
}

func main() {
	fmt.Println(SumFloat(10.3, 45.1)) // 55.4
	fmt.Println(SumInt(10, 45))       // 55
	fmt.Println(SumString("^_", "^")) // ^_^
}
