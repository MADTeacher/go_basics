package main

import "fmt"

func main() {
	hello() // вызов функции hello
	fmt.Println("exit")
}

func hello() {
	fmt.Println("Hello World!")
}

// Hello World!
// exit
