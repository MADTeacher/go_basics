package main

import "fmt"

func main() {
	value := 2
	for i := 0; i < 10; i++ {
		fmt.Printf("%d * %d = %d\n", value, i, value*i)
	}
}
