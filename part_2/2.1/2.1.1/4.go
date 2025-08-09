package main

import "fmt"

func main() {
	a, b, rez := 10, 13, 0
	add(a, b, &rez)
	fmt.Printf("%d + %d = %d\n", a, b, rez)
}

func add(a, b int, rez *int) {
	*rez = a + b
}

// 10 + 13 = 23
