package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello mad world!"
	fmt.Println(strings.Index(str, "l")) // 2
	fmt.Println(strings.Index(str, "W")) // -1

	fmt.Println(strings.LastIndex(str, "l")) // 13
	fmt.Println(strings.LastIndex(str, "W")) // -1
}
