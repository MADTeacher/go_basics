package main

import (
	"errors"
	"fmt"
)

type DivError struct {
	Numerator   int
	Denominator int
	Msg         string
}

func (d *DivError) Error() string {
	return d.Msg
}

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, &DivError{
			Numerator:   a,
			Denominator: b,
			Msg:         fmt.Sprintf("Сan't divide '%v' by zero", a),
		}
	}
	if b == 1 {
		return a, &DivError{
			Numerator:   a,
			Denominator: b,
			Msg:         "Divide by one",
		}
	}
	return a / b, nil
}

func printResult(result int, err error) {
	if err != nil {
		var myDivError *DivError
		switch {
		case errors.As(err, &myDivError):
			fmt.Printf("Numerator: %d, Denominator: %d, Error: %s\n",
				myDivError.Numerator, myDivError.Denominator, myDivError.Msg)
			if myDivError.Denominator == 1 {
				fmt.Printf("result = %d, but was %v\n",
					myDivError.Numerator, myDivError.Msg)
			}
		default:
			fmt.Printf("unexpected error: %v\n", err)
		}

	} else {
		fmt.Println(result)
	}
}

func main() {
	printResult(div(10, 2)) // 5
	printResult(div(4, 0))  // Numerator: 4, Denominator: 0, Error: Сan't divide '4' by zero
	printResult(div(12, 1))
	// Numerator: 12, Denominator: 1, Error: Divide by one
	// result = 12, but was divide by one
}
