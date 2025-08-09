package main

import "fmt"

func sumElem(slice []int) int {
	fmt.Println(slice)
	if len(slice) <= 1 {
		return slice[0]
	}
	return anotherFunction(slice)
}

func anotherFunction(slice []int) int {
	return slice[0] + sumElem(slice[1:])
}

func main() {
	mySlice := []int{2, 5, 22, 9, 0}
	fmt.Println(sumElem(mySlice))
}
