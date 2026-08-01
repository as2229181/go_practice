package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type AtomicCounter struct {
	count int64
}

func (a *AtomicCounter) increment() {
	atomic.AddInt64(&a.count, 1)
}

func (a *AtomicCounter) getValue() int64 {
	return atomic.LoadInt64(&a.count)
}

func main() {
	var wg sync.WaitGroup
	numberGoroutines := 10
	counter := &AtomicCounter{}
	for range numberGoroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				counter.increment()
			}
		}()
	}
	wg.Wait()
	fmt.Printf("Results: %d\n", counter.getValue())
}
