package main

import (
	"fmt"
	"math/rand"
	"time"
)

// генератор случайных чисел
var generator = rand.New(rand.NewSource(time.Now().UnixNano()))

func sliceCreator1(size int16) (slice *[]int16) {
	slice = new([]int16)
	*slice = make([]int16, size, size+10)
	for i := range *slice {
		(*slice)[i] += int16(generator.Int())
	}
	return
}

func sliceCreator2(size int16) *[]int16 {
	slice := make([]int16, size, size+10)
	for i := range slice {
		slice[i] += int16(generator.Int())
	}
	return &slice
}

func main() {
	slice := sliceCreator1(7)
	fmt.Println(*slice) // [-22201 -30245 6682 28346 23159 -28589 16762]
	slice = sliceCreator2(5)
	fmt.Println(*slice) // [3896 28732 -30834 1722 -14725]
}
