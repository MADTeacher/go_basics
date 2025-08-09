package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello mad world!"
	before, after, found := strings.Cut(str, " ")
	fmt.Printf("before: %s, after: %s, found: %t \n", before, after, found)
	// before: Hello, after: mad world!, found: true

	before, after, found = strings.Cut(str, "6")
	fmt.Printf("before: %s, after: %s, found: %t \n", before, after, found)
	// before: Hello mad world!, after: , found: false

	after, found = strings.CutPrefix(str, "He")
	fmt.Printf("after: %s, found: %t \n", after, found)
	// after: llo mad world!, found: true

	after, found = strings.CutPrefix(str, "yello")
	fmt.Printf("after: %s, found: %t \n", after, found)
	// after: Hello mad world!, found: false

	before, found = strings.CutSuffix(str, "!")
	fmt.Printf("before: %s, found: %t \n", before, found)
	// before: Hello mad world, found: true

	before, found = strings.CutSuffix(str, "?")
	fmt.Printf("before: %s, found: %t \n", before, found)
	// before: Hello mad world!, found: false
}
