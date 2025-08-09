package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := ""
	i := 0
	for i <= 5 {
		str += strconv.Itoa(i)
		i++
	}
	fmt.Println(str) // 012345
}
