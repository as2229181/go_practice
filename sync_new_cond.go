package main

import (
	"fmt"
	"sync"
	"time"
)

const bufferSize = 5

type buffer struct {
	item []int
	mu   sync.Mutex
	cond *sync.Cond
}

func newBuffer(size int) *buffer {
	b := &buffer{
		item: make([]int, 0, size),
		mu:   sync.Mutex{},
	}
	b.cond = sync.NewCond(&b.mu)
	return b
}

func (b *buffer) produce(item int) {
	b.mu.Lock()
	defer b.mu.Unlock()

	for len(b.item) == bufferSize {
		b.cond.Wait()
	}

	b.item = append(b.item, item)

	fmt.Println("Produce:", item)
	b.cond.Signal()
}

func (b *buffer) consume() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	for (len(b.item)) == 0 {
		b.cond.Wait()
		// This function stop do anything and waits for other function to append to the slice
	}

	item := b.item[0]
	b.item = b.item[1:]
	fmt.Println("Consume:", item)

	b.cond.Signal()
	return item
}

func producer(b *buffer, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range 10 {
		b.produce(i + 100)
		time.Sleep(100 * time.Millisecond)
	}
}

func consumer(b *buffer, wg *sync.WaitGroup) {
	defer wg.Done()
	for range 10 {
		b.consume()
		time.Sleep(200 * time.Microsecond)
	}
}

func main() {
	buffer := newBuffer(bufferSize)
	var wg sync.WaitGroup
	wg.Add(2)
	go producer(buffer, &wg)
	go consumer(buffer, &wg)

	wg.Wait()
}
