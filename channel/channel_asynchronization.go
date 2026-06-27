package channel

import (
	"fmt"
	"time"
)

// func main() {
// 	done := make(chan struct{})

// 	go func() {
// 		fmt.Println("Start working.....")
// 		time.Sleep(3 * time.Second)
// 		done <- struct{}{}
// 	}()

// 	<- done
// 	fmt.Println("Job finished !!!!")
// }

// func main() {
// 	numGoroutines := 3
// 	done := make(chan int, 3)
// 	time.Sleep(time.Second)
// 	for i := range numGoroutines{
// 		go func(id int){
// 			fmt.Println("Goroutine %d working", id)
// 			done <- id
// 		}(i)
// 	}

// 	for range numGoroutines{
// 		<- done
// 	}
// 	fmt.Println("All goroutiines are complete")
// }

func main(){
	data := make(chan string)
	go func(){
		defer close(data)
		for i := range 5{
			data <- "hello" + string('0'+ i)
			time.Sleep(time.Second)
		}
	}()
	for recv := range data {
		fmt.Println("Recevied value:", recv, ":", time.Now())
	}
	fmt.Println("Channel colose")
}