package main

import "fmt"

func main() {
	str1 := "Hi, my маленький friend!!!"
	runes := []rune(str1)
	substr := string(runes[4:16])
	fmt.Println(substr) // my маленький
}
