package main

import (
	"fmt"
	"sync"
	"time"
)

type LeakyBucket struct {
	capacity int
	leakRate time.Duration
	tokens   int
	lastLeak time.Time
	mu       sync.Mutex
}

func NewLeakyBucket(capacity int, leakRate time.Duration) *LeakyBucket {
	return &LeakyBucket{
		capacity: capacity,
		leakRate: leakRate,
		tokens:   capacity,
		lastLeak: time.Now(),
	}
}

func (lb *LeakyBucket) allow() bool {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	now := time.Now()
	elapsedTime := now.Sub(lb.lastLeak)
	tokensToAdd := int(elapsedTime / lb.leakRate)
	fmt.Println(tokensToAdd)
	lb.tokens += tokensToAdd

	if lb.tokens > lb.capacity {
		lb.tokens = lb.capacity
		lb.lastLeak = now
	} else {
		lb.lastLeak = lb.lastLeak.Add(time.Duration(tokensToAdd) * lb.leakRate) // not just add elpasedTime
	}
	// elpasedTime = 1.3 seconds
	// initial lastLeak = 0
	// tokensToAdd = int(1.3/ 0.5) = 2.6 => 2 tokens
	// lb.lastLeak = lb.lastLeak.Add(time.Duration(tokensToAdd) * lb.leakRate) => 0 + 2 * 0.5 = 1
	// lb.lastLeak = lb.lastLeak + elapsed time 0 + 1.3 => 0.3 wasted

	if lb.tokens > 0 {
		fmt.Println(lb.tokens)
		lb.tokens--
		return true
	}
	return false
}

func main() {
	leakyBucket := NewLeakyBucket(5, 500*time.Millisecond)

	for range 10 {
		if leakyBucket.allow() {
			fmt.Println("Request accept")
		} else {
			fmt.Println("Request denied")
		}
		time.Sleep(200 * time.Millisecond)
	}
}
