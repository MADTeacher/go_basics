package main

import (
	"fmt"
	"sync"
)

const totalElements = 6
const totalWorkers = 3

type TaskItem struct {
	id   int
	data int
}

func main() {
	input := make(chan TaskItem, totalElements)
	output := make(chan TaskItem, totalElements)

	for w := 1; w <= totalWorkers; w++ {
		go worker(w, input, output)
	}

	// Отправка данных для обработки
	for j := 0; j < totalElements; j++ {
		input <- TaskItem{j, j}
	}

	close(input)

	// Работа с результатами
	for a := 0; a < totalElements; a++ {
		task := <-output
		fmt.Printf("Task-%d finished!\n", task.id)
	}

	close(output)
}

func worker(id int, input <-chan TaskItem, output chan<- TaskItem) {
	var wg sync.WaitGroup

	for j := range input {
		wg.Add(1)

		go func(task TaskItem) {
			defer wg.Done()

			fmt.Printf("Worker %d started task %+v\n", id, task)
			task.data *= task.data
			output <- task

			fmt.Printf("Worker %d finished task %+v\n", id, task)
		}(j)
	}

	wg.Wait()
}
