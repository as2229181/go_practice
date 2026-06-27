package main

import (
	"fmt"
	"time"
)

// func main() {
// 	ch1 := make(chan int)

// 	ch2 := make(chan int)
// 	go func() {
// 		ch1 <- 1
// 	}()
// 	// time.Sleep(time.Second)
// 	go func() {
// 		ch2 <- 1
// 	}()

// 	for range 2{
// 		select {
// 		case msg := <-ch1:
// 			fmt.Println("received from ch1:", msg)
// 		case msg := <-ch2:
// 			fmt.Println("received from ch2:", msg)
// 		// if default exist, select would be nonblocking
// 		// default:
// 		// 	fmt.Println("No message received")
// 		}
// 	}
// }

// func main(){
// 	ch := make(chan int)
// 	go func() {
// 		// time.Sleep(3 * time.Second)
// 		ch <- 1
// 		close(ch)
// 	}()
// 	for {
// 	select{
// 	case msg, ok := <- ch:
// 		if !ok{
// 			fmt.Println("Channel close")
// 			return
// 		}
// 		fmt.Println("Received:", msg)
// 	// case <-time.After(1 * time.Second):
// 	// 	fmt.Println("Timeout.")
// 	}
// }
// }

func main(){
	data := make(chan int)
	quit := make(chan bool)
	go func(){
		for{
			select {
			case d :=<- data:
				fmt.Println("data is ", d)
			case <- quit:
				fmt.Println("Stopping....")
			default:
				fmt.Println("waiting for data")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	for i := range 5 {
		data <- i
		time.Sleep(1 *time.Second)
	}
	quit <- true

}