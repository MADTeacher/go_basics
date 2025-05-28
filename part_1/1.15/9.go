package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "Hi, my dear friend!!!"
	str2 := "Hi, my dear friend!!!"
	str3 := "Hi, my dear friend!!"
	fmt.Println(strings.Compare(str1, str2)) // 0
	fmt.Println(strings.Compare(str1, str3)) // 1
	fmt.Println(strings.Compare(str3, str1)) // –1
}
