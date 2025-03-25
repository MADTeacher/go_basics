package main

import (
	"errors"
	"fmt"
)

// Объявление пользовательского варианта ошибки деления на ноль
var ErrDivByZero = errors.New("divide by zero")

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivByZero
	}
	return a / b, nil
}

func printResult(result int, err error) {
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func main() {
	a, b := 10, 2
	printResult(div(a, b)) // 5
	a, b = 4, 0
	printResult(div(a, b)) // divide by zero
}
