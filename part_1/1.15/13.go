package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello mad world!"
	fmt.Println(strings.HasPrefix(str, "He"))  // true
	fmt.Println(strings.HasPrefix(str, "he"))  // false
	fmt.Println(strings.HasPrefix(str, "mad")) // false

	fmt.Println(strings.HasSuffix(str, "!"))   // true
	fmt.Println(strings.HasSuffix(str, "mad")) // false
}
