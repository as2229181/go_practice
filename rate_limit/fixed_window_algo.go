package main

import (
	"fmt"
	"sync"
	"time"
)

type RateLimiter struct {
	mu        sync.Mutex
	count     int
	limit     int
	window    time.Duration
	resetTime time.Time
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		limit:  limit,
		window: window,
	}
}

func (rl *RateLimiter) allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	if now.After(rl.resetTime) {
		rl.resetTime = now.Add(rl.window)
		rl.count = 0
	}
	if rl.count < rl.limit {
		rl.count++
		return true
	}
	return false
}

func main() {
	var wg sync.WaitGroup
	rl := NewRateLimiter(5, time.Second)
	for range 10 {
		wg.Add(1)
		go func() {
			if rl.allow() {
				fmt.Println("Request allowed")
			} else {
				fmt.Println("Reqeust denied")
			}
			wg.Done()
		}()
		// time.Sleep(200 * time.Millisecond)
	}
	wg.Wait()
}
