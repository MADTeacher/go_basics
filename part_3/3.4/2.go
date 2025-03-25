package main

import "fmt"

var global string = "Global value"

func myFunc() {
	global := 10
	fmt.Println(global)
}

func main() {
	fmt.Println(global)
	myFunc()
}

// Global value
// 10
