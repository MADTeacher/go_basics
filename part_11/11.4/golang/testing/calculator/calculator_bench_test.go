package calculator

import "testing"

func BenchmarkMul(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Mul(5, 10)
	}
}

func BenchmarkPow2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Pow2(5)
	}
}
