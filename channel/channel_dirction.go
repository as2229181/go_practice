package channel

import (
	"fmt"
)

func main() {

	ch := make(chan int)
	// go func(ch chan<- int){
	// 	defer close(ch)
	// 	for i := range 5{
	// 		ch <- i
	// 	}
	// }(ch)
	go producer(ch)
	// for value := range ch {
	// 	fmt.Println("recived:", value)
	// }
	comsumer(ch)
}

func producer(ch chan<- int){
	defer close(ch)
		for i := range 5{
			ch <- i
		}
}

// Recived only channel
func comsumer(ch <-chan int){
	for value := range ch {
		fmt.Println("Recived:", value)
	}
}