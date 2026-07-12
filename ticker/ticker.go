package main

import (
	"fmt"
	"time"
)

// func main() {
// 	ticker := time.NewTicker(time.Second)
// 	defer ticker.Stop()
// 	i := 0
// 	for range 5 {
// 		i += 1
// 		fmt.Println("Tick at:", i)
// 	}
// 	for tick := range ticker.C {
// 		fmt.Println("Tick at:", tick)
// 	}
// }

// Scheduling Logging, Periodic Task, Polling for U pdate
// func periodicTask() {
// 	fmt.Println("Performing periodic task:", time.Now())
// }

// func main() {
// 	ticker := time.NewTicker(5 * time.Second)
// 	defer ticker.Stop()
// 	for {
// 		select {
// 		case <-ticker.C:
// 			periodicTask()
// 		}
// 	}
// }

func main() {
	ticker := time.NewTicker(1 * time.Second)

	stop := time.After(5 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case tick := <-ticker.C:
			fmt.Println("Tick at:", tick)
		case <-stop:
			fmt.Println("Stoping Ticker")
			return
		}
	}
}
