package main

import (
	"fmt"
	"os"
)

func main() {
	var value1, value2 string
	fmt.Fscan(os.Stdin, &value1, &value2)
	fmt.Println(value1, value2)
}
