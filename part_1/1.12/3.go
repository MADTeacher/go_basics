package main

import "fmt"

func main() {
myLable: // бесконечный цикл
	fmt.Println(1)
	goto myLable
	fmt.Println(2)
}
