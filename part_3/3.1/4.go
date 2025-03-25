package main

import "fmt"

func main() {
	fmt.Println(userInfo(3, 5, "Alex")) // User name: Alex, age: 3, exp years: 5
}

func userInfo(age, exp int, name string) string {
	return fmt.Sprintf("User name: %s, age: %d, exp year: %d", name, age, exp)
}
