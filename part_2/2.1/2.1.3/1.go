package main

import "fmt"

var global string = "Global value"

// global := "Global value" // так объявлять нельзя, будет ошибка

func myFunc() {
	fmt.Println(global)
}

func main() {
	fmt.Println(global)
	myFunc()
}

// Global value
// Global value
