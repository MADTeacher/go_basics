package main

import "fmt"

func main() {
	i := 13
	for i > 0 {
		i--
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
}
