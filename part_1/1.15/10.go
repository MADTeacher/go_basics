package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "Hi, my dear friend!!!"
	fmt.Printf("%t\n", strings.Contains(str1, "i"))  // true
	fmt.Printf("%t\n", strings.Contains(str1, "my")) // true
	fmt.Printf("%t\n", strings.Contains(str1, "?"))  // false
}
