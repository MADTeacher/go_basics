package main

import (
	"errors"
	"fmt"
)

func main() {
	myError1 := errors.New("first error")
	myError2 := fmt.Errorf("second error")
	fmt.Println(myError1, myError2) // first error second error
}
