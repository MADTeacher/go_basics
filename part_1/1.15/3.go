package main

import "fmt"

func main() {
	str1 := "Hi, my dear friend!!!"
	fmt.Println("Символы с 1 по 5: " + str1[:5])
	fmt.Println("Символы с 4: " + str1[4:])
}
