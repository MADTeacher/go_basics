package main

import "fmt"

func main() {
	fmt.Println(1) // 1
	goto myLable
	fmt.Println(2)
myLable: // метка для безусловного перехода
	fmt.Println(3) // 3
}
