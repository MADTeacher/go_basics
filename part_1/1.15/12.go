package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello mad world!"
	fmt.Println(strings.EqualFold(str, "hello mad world!")) // true
	fmt.Println(strings.EqualFold(str, "HELLO mad world!")) // true
	fmt.Println(strings.EqualFold(str, "Oo"))               // false
}
