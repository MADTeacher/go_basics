package main

import (
	"fmt"
	calc "golang/testing/calculator"
)

func main() {
	fmt.Println(calc.Add(4, 2))   // 6
	fmt.Println(calc.Sub(23, 12)) // 11
	fmt.Println(calc.Mul(5, 6))   // 30
	fmt.Println(calc.Pow2(5))     // 25
}
