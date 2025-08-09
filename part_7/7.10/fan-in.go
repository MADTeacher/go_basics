package main

import (
	"fmt"
	"sync"
)

type TaskItem struct {
	id int
}

func generateInputs(work *[]TaskItem) (in1, in2 <-chan TaskItem) {
	ch1 := make(chan TaskItem)
	ch2 := make(chan TaskItem)

	go func() {
		defer close(ch1)
		defer close(ch2)

		for it, value := range *work {
			if it%2 == 0 {
				ch1 <- value
			} else {
				ch2 <- value
			}
		}
	}()

	return ch1, ch2
}

func main() {
	tasks := []TaskItem{
		{0}, {1}, {2},
		{3}, {4}, {5},
		{6}, {7}, {8},
		{9}, {10},
	}

	in1, in2 := generateInputs(&tasks)

	out := fanIn(in1, in2)

	for value := range out {
		fmt.Println("Value:", value)
	}
	fmt.Println("Finished")
}

func fanIn(inputs ...<-chan TaskItem) <-chan TaskItem {
	var wg sync.WaitGroup
	out := make(chan TaskItem)

	wg.Add(len(inputs))

	for _, in := range inputs {
		go func(ch <-chan TaskItem) {
			for {
				value, ok := <-ch

				if !ok {
					wg.Done()
					break
				}

				out <- value
			}
		}(in)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
