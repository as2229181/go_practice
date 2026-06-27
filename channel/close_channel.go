package main

import (
	"fmt"
	// "time"
)

// receive data from close channel
// func main() {
// 	ch := make(chan int)
// 	close(ch)

// 	val, ok := <- ch
// 	if !ok {
// 		fmt.Println("channel is close")
// 		return
// 	}
// 	fmt.Println(val)
	
// }


// Receive data throgh range channel and close channel
// func main() {
// 	ch := make(chan int)
// 	go func() {
// 		for i := range 5 {
// 			ch <- i
// 		}
// 		close(ch)
// 	}()
// 	for val := range ch{
// 		fmt.Println(val)
// 	}
// }

//Close channel twice is ileagal
// func main() {
// 	ch := make(chan int)

// 	go func(){
// 		close(ch)
// 		close(ch)
// 	}()
// 	time.Sleep(time.Second)
// 	fmt.Println("end of program")
// }

func producer(ch chan <- int) {
	for i := range 5{
		ch <- i
	}
	close(ch)
}


func filter(in  <-chan int, out chan<- int){
	for val := range in{
		if val % 2 == 0{
			out <- val
		}
	}	
	close(out)
}


func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)
	go producer(chan1)
	go filter(chan1, chan2)
	
	for val := range chan2{
		fmt.Println(val)
	}
	fmt.Println("end of program")

}