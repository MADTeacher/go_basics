package main

import "fmt"

func main() {
	userInfo(3, 5, "Alex")
}

func userInfo(age, exp int, name string) {
	fmt.Printf("User name: %s, age: %d, exp years: %d", name, age, exp)
}

// User name: Alex, age: 3, exp years: 5
