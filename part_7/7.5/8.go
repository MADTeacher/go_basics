package main

import (
	"fmt"
)

type Rectangle struct {
	Width  uint
	Length uint
}

func (r *Rectangle) GetPerimeter() uint {
	return (r.Length + r.Width) * 2
}

func (r *Rectangle) GetArea() uint {
	return r.Length * r.Width
}

func calculate(in <-chan Rectangle, out chan<- float64) {
	var sum float64
	for i := 0; i <= 3; i++ {
		reactangle := <-in
		sum += float64(reactangle.GetPerimeter())
	}
	out <- sum / float64(cap(in))
}

func main() {
	inChan := make(chan Rectangle, 3)
	outChan := make(chan float64)
	go calculate(inChan, outChan)

	inChan <- Rectangle{14, 6}
	inChan <- Rectangle{5, 21}
	inChan <- Rectangle{10, 33}
	inChan <- Rectangle{2, 5}

	result := <-outChan
	fmt.Printf("Result = %v", result)
}
