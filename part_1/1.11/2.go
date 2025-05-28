package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := ""
	i := 0
	for {
		str += strconv.Itoa(i)
		if i >= 5 {
			break
		}
		i++
	}
	fmt.Println(str) // 012345
}
