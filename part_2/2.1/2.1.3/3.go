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
	if true {
		local := 20
		fmt.Println(local)
	}
	local = 5 // undefined: local
}
