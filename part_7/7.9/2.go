package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
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

func getGOMAXPROCS() int {
	return runtime.GOMAXPROCS(0)
}

func main() {
	var waitGroup sync.WaitGroup
	counter := &AtomicCounter{0}
	myTime := time.Now()
	for i := 0; i < 100_000; i++ {
		waitGroup.Add(1)
		go increment(&waitGroup, counter)
	}
	waitGroup.Wait()
	fmt.Println("Value after 1k increment  = ", counter.Value())
	fmt.Printf("Time work: %v", time.Since(myTime))
}
