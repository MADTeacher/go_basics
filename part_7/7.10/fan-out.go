package main

import (
	"fmt"
)

type TaskItem struct {
	id int
}

func main() {
	tasks := []TaskItem{
		{0}, {1}, {2},
		{3}, {4}, {5},
		{6}, {7}, {8},
		{9}, {10},
	}

	in := generateInput(&tasks)

	out1 := fanOut(in)
	out2 := fanOut(in)
	out3 := fanOut(in)

	for range tasks {
		select {
		case value := <-out1:
			fmt.Println("Task in output-1: ", value)
		case value := <-out2:
			fmt.Println("Task in output-2: ", value)
		case value := <-out3:
			fmt.Println("Task in output-3: ", value)
		}
	}
	fmt.Println("Finished")
}

func generateInput(work *[]TaskItem) (in1 <-chan TaskItem) {
	ch1 := make(chan TaskItem)
	go func() {
		defer close(ch1)
		for _, value := range *work {
			ch1 <- value
		}
	}()

	return ch1
}

func fanOut(in <-chan TaskItem) <-chan TaskItem {
	out := make(chan TaskItem)

	go func() {
		defer close(out)

		for data := range in {
			out <- data
		}
	}()

	return out
}
