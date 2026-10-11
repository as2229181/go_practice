package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu1, mu2 sync.Mutex

	go func() {
		mu1.Lock()
		fmt.Println("Gorountine1 lock mutex 1")
		time.Sleep(time.Second)
		mu2.Lock()
		fmt.Println("Gorountine1 lock mutex 1")
		mu1.Unlock()
		mu2.Unlock()
	}()

	go func() {
		mu1.Lock()
		fmt.Println("Gorountine2 lock mutex 2")
		time.Sleep(time.Second)
		mu2.Lock()
		fmt.Println("Gorountine2 lock mutex 1")
		mu2.Unlock()
		mu1.Unlock()
	}()
	time.Sleep(time.Second * 8)
	// select {}
}
