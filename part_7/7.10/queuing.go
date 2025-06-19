package main

import (
	"fmt"
	"sync"
	"time"
)

type TaskItem struct {
	id int
}

func main() {
	var wg sync.WaitGroup
	limit := make(chan interface{}, 2)
	workers := func(l chan<- interface{}, wg *sync.WaitGroup, tasks *[]TaskItem) {
		for _, value := range *tasks {
			value := value // удалите и посмотрите на вывод в терминал
			// не используйте в замыкании передачу задачи по ссылке на переменную value
			// объявленную в цикле (_, value := range *tasks)

			limit <- struct{}{}
			wg.Add(1)

			go func(workItem *TaskItem, w *sync.WaitGroup) {
				defer w.Done()
				fmt.Printf("Task %d processing\n", workItem.id)
				time.Sleep(1 * time.Second)
				<-limit
			}(&value, wg)
		}
	}

	tasks := []TaskItem{
		{0}, {1}, {2},
		{3}, {4}, {5},
		{6}, {7}, {8},
		{9},
	}
	workers(limit, &wg, &tasks)
	wg.Wait()

	fmt.Println("Finished")
}
