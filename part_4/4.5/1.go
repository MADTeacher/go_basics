package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string = "10"
	value, err := strconv.Atoi(str) // преобразование строки в целочисленное значение
	if err != nil {                 // если err != nil, то обрабатываем ошибку
		value = 22
	}
	fmt.Println(value) // 10

	str = "2we"
	value, err = strconv.Atoi(str)
	if err != nil { // если err != nil, то обрабатываем ошибку
		value = 22
	}
	fmt.Println(value) // 22
}
