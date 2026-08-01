package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	numberGoroutines := 5
	wg.Add(numberGoroutines)
	increment := func() {
		defer wg.Done()
		for range 1000 {
			mu.Lock()
			counter++
			mu.Unlock()
		}
	}

	for range numberGoroutines {
		go increment()
	}
	wg.Wait()
	fmt.Println("Results:", counter)
}

// type counter struct {
// 	mu    sync.Mutex
// 	count int
// }

// func (c *counter) icreament() {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	c.count++
// }

// func (c *counter) getValue() int {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	return c.count
// }

// func main() {
// 	var wg sync.WaitGroup
// 	counter := &counter{}

// 	numberGrroutines := 10

// 	for range numberGrroutines {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			for range 1000 {
// 				// counter.icreament()
// 				counter.count++
// 			}
// 		}()
// 	}
// 	wg.Wait()
// 	fmt.Printf("Final count value: %d\n", counter.count)
// }
