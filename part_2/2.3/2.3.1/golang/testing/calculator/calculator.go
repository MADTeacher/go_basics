package calculator

import "math"

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Mul(a, b int) int {
	return a * b
}

func Pow2(a int) int {
	return int(math.Pow(float64(a), 2))
}
