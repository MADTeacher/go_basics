package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	val1 := 33
	var str string = strconv.Itoa(val1)
	fmt.Println(str) // 33
	val2, err := strconv.Atoi(str)
	if err == nil { // если получилось преобразовать, то err == nil
		fmt.Println(val2)                 // 33
		fmt.Println(reflect.TypeOf(val2)) // int
	}
	val2, err = strconv.Atoi("ee")
	if err != nil {
		fmt.Println(val2) // 0
	}

	val3 := 100
	fmt.Println("Rezult: " + strconv.Itoa(val3)) // Rezult: 100
}
