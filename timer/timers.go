package main

// Basic Timer
// func main() {
// 	timer := time.NewTimer(2 * time.Second)
// 	stop := timer.Stop()
// 	if stop {
// 		fmt.Println("Timer stopped")
// 	}
// 	timer.Reset(1 * time.Second)
// 	fmt.Println("Timer reset")
// 	<-timer.C
// 	fmt.Println("Timer expired")
// }

// func longRunningOperation() {
// 	for i := range 1 {
// 		fmt.Println(i)
// 		time.Sleep(time.Second)
// 	}
// }

// Time out
// func main() {
// 	timeout := time.After(2 * time.Second)
// 	done := make(chan bool)
// 	go func() {
// 		longRunningOperation()
// 		done <- true
// 	}()
// 	select {
// 	case <-timeout:
// 		fmt.Println("Operation time out")
// 	case <-done:
// 		fmt.Println("Operation Done")
// 	}

// }

// ================== Scheduling Delayed Operations
// func main() {
// 	timer := time.NewTimer(2 * time.Second) // non-blocking timers starts

// 	go func() {
// 		<-timer.C
// 		fmt.Println("Delayed operation executed")
// 	}()
// 	fmt.Println("Waiting....")
// 	time.Sleep(5 * time.Second)
// 	fmt.Println("End of program")

// }
