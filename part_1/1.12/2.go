package main

import "fmt"

func main() {
	test()
	fmt.Println(1)
	goto myLable
	fmt.Println(2)
myLable:
	fmt.Println(3)
}

func test() {
	goto myLable // error: label myLable not defined
}
