package channel


import (
	"fmt"
)

func buffered_channel() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println("Buffered Channel2: ", <- ch)
	fmt.Println("Buffered Channel2: ", <- ch)
	fmt.Println("Bufferd Channel", <- ch)
}