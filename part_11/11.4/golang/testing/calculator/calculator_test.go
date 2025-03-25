package calculator

import (
	"testing"
)

type testData struct {
	arg1, arg2, expected int
}

type testPowData struct {
	arg, expected int
}

func TestSub(t *testing.T) {
	cases := []testData{
		{2, 3, -1},
		{10, 5, 5},
		{-8, -3, -5},
	}
	for _, it := range cases {
		result := Sub(it.arg1, it.arg2)
		if result != it.expected {
			t.Errorf("result %d, expected %d", result, it.expected)
		}
	}
}

func TestPow2(t *testing.T) {
	cases := []testPowData{
		{2, 4},
		{-2, 4},
		{3, 9},
	}
	for _, it := range cases {
		result := Pow2(it.arg)
		if result != it.expected {
			t.Errorf("result %d, expected %d", result, it.expected)
		}
	}
}
