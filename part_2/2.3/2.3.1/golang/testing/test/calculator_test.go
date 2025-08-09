package test_test

import (
	calc "golang/testing/calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	result := calc.Add(2, 5)
	expected := 7
	if result != expected {
		t.Errorf("result %d, expected %d", result, expected)
	}
}

func TestSub(t *testing.T) {
	result := calc.Sub(15, 5)
	expected := 10
	if result != expected {
		t.Errorf("result %d, expected %d", result, expected)
	}
}

func TestMul(t *testing.T) {
	result := calc.Mul(3, 6)
	expected := 18
	if result != expected {
		t.Errorf("result %d, expected %d", result, expected)
	}
}

func TestPow2(t *testing.T) {
	result := calc.Pow2(3)
	expected := 9
	if result != expected {
		t.Errorf("result %d, expected %d", result, expected)
	}
}
