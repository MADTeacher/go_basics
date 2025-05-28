package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "Hi, my dear friend!!!"
	str2 := strings.ToUpper(str1)
	fmt.Println(str2) // HI, MY DEAR FRIEND!!!
	fmt.Println(str1) // Hi, my dear friend!!!

	fmt.Println(strings.ToLower(str2)) // hi, my dear friend!!!
	fmt.Println(strings.ToTitle(str1)) // HI, MY DEAR FRIEND!!!
}
