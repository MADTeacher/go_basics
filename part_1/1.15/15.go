package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	fmt.Println(strings.Join([]string{"a", "b", "c"}, " "))
	// a b c

	fmt.Println(strings.Join([]string{"a", "b", "c"}, "-"))
	// a-b-c

	// инвертирование строки
	str := "Hello!"
	splitStr := strings.Split(str, "")
	slices.Reverse(splitStr)
	fmt.Println(strings.Join(splitStr, "")) // !olleH
}
