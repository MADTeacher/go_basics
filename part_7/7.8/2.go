package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type AtomicCounter struct {
	value int64
}

func (a *AtomicCounter) Increment() {
	atomic.AddInt64(&a.value, 1)
}

func (a *AtomicCounter) Decrement() {
	atomic.AddInt64(&a.value, -1)
}

func (a *AtomicCounter) Value() int64 {
	return atomic.LoadInt64(&a.value)
}

func increment(waitGroup *sync.WaitGroup, counter *AtomicCounter) {
	counter.Increment()
	waitGroup.Done()
}

func main() {
	var waitGroup sync.WaitGroup
	counter := &AtomicCounter{0}
	for i := 0; i < 1000; i++ {
		waitGroup.Add(1)
		go increment(&waitGroup, counter)
	}
	waitGroup.Wait()
	fmt.Println("Value after 1k increment  = ", counter.Value())
}
