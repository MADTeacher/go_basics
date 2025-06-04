package main

import (
	"errors"
	"fmt"
)

var ErrDivByZero = errors.New("divide by zero")
var ErrDivByOne = errors.New("divide by one")

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivByZero
	}
	if b == 1 {
		return a, ErrDivByOne
	}
	return a / b, nil
}

func printResult(result int, err error) {
	if err != nil {
		switch {
		case errors.Is(err, ErrDivByZero):
			fmt.Println(err)
		case errors.Is(err, ErrDivByOne):
			fmt.Printf("result = %d, but was %v\n", result, err)
		default:
			fmt.Printf("unexpected error: %v\n", err)
		}

	} else {
		fmt.Println(result)
	}
}

func main() {
	printResult(div(10, 2)) // 5
	printResult(div(4, 0))  // divide by zero
	printResult(div(12, 1)) // result = 12, but was divide by one
}
