package main

import "fmt"

func main() {
	str1 := "hello"
	fmt.Println(str1) // hello

	byteArray := []byte(str1)
	fmt.Println(byteArray) // [104 101 108 108 111]

	str2 := string(byteArray)
	fmt.Println(str2) // hello
}
